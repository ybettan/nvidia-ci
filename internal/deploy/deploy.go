package deploy

import (
	"fmt"
	"log"

	"github.com/golang/glog"
	"github.com/kelseyhightower/envconfig"
	"github.com/openshift-kni/eco-goinfra/pkg/namespace"
	"github.com/openshift-kni/eco-goinfra/pkg/olm"
	"github.com/operator-framework/api/pkg/operators/v1alpha1"
	"github.com/rh-ecosystem-edge/nvidia-ci/internal/inittools"
)

type MachineSetConfig struct {
	InstanceType string `required:"true" envconfig:"GPU_MACHINESET_INSTANCE_TYPE"`
}

type SubscriptionConfig struct {
	CatalogSource       string `required:"true" envconfig:"GPU_CATALOGSOURCE"`
	SubscriptionChannel string `required:"true" envconfig:"GPU_SUBSCRIPTION_CHANNEL"`
}

type BundleConfig struct {
	BundleImage string `required:"true" envconfig:"GPU_BUNDLE_IMAGE"`
}

type deployFromSubscription interface {
	GetSubscriptionConfig(logLevel glog.Level) (*SubscriptionConfig, error)
	DeployFromSubscription(logLevel glog.Level, targetNs, subName, operatorGroup, pkg, catalogSrc string,
		installPlanApproval v1alpha1.Approval) (*olm.SubscriptionBuilder, *olm.OperatorGroupBuilder, error)
}

type deployFromBundle interface {
	GetBundleConfig() (*BundleConfig, error)
}

type Deploy interface {
	deployFromSubscription
	deployFromBundle
	GetMachineSetConfig() (*MachineSetConfig, error)
	CreateAndLabelNamespaceIfNeeded(logLevel glog.Level, targetNs string, labels map[string]string) (*namespace.Builder, error)
}

type deploy struct{}

func NewDeploy() Deploy {
	return deploy{}
}

func (d deploy) GetSubscriptionConfig(logLevel glog.Level) (*SubscriptionConfig, error) {

	glog.V(logLevel).Infof("Getting the GPU subscription config")

	var subConfig SubscriptionConfig

	if err := envconfig.Process("gpu", &subConfig); err != nil {
		return nil, fmt.Errorf("failed to instantiate subscription config: %v", err)
	}

	return &subConfig, nil
}

// FIXME: get a SubscriptionConfig instead of getting all those parameters
func (d deploy) DeployFromSubscription(logLevel glog.Level, targetNs, subName, operatorGroup, pkg,
	catalogSrc string, installPlanApproval v1alpha1.Approval) (*olm.SubscriptionBuilder, *olm.OperatorGroupBuilder, error) {

	// Get the packagemanifest
	glog.V(logLevel).Infof("Using catalogsource '%s'", catalogSrc)
	pkgManifest, err := olm.PullPackageManifestByCatalog(inittools.APIClient, pkg, targetNs, catalogSrc)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get packagemanifest %s from catalog %s: %v", pkg, catalogSrc, err)
	}
	glog.V(logLevel).Infof("The packagemanifest name returned: %s", pkgManifest.Object.Name)
	pkgDefaultChannel := pkgManifest.Object.Status.DefaultChannel
	glog.V(logLevel).Infof("The default channel retrieved from packagemanifest is:  %v", pkgDefaultChannel)

	// Create the operatorgroup
	var ogBuilderCreated *olm.OperatorGroupBuilder
	ogBuilder := olm.NewOperatorGroupBuilder(inittools.APIClient, operatorGroup, targetNs)
	if ogBuilder.Exists() {
		glog.V(logLevel).Infof("Operatorgroup %s already exists", operatorGroup)
	} else {
		glog.V(logLevel).Infof("Creating '%s' operatorgroup", operatorGroup)
		ogBuilderCreated, err = ogBuilder.Create()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create operatorgroup '%s': %v", operatorGroup, err)
		}
	}

	// Create the subscription
	subBuilder := olm.NewSubscriptionBuilder(inittools.APIClient, subName, targetNs, catalogSrc, targetNs, pkg)
	glog.V(logLevel).Infof("Setting the subscription channel to: '%s'", pkgDefaultChannel)
	subBuilder.WithChannel(pkgDefaultChannel)
	subBuilder.WithInstallPlanApproval(installPlanApproval)
	glog.V(logLevel).Infof("Creating the subscription, i.e Deploy the operator")
	createdSub, err := subBuilder.Create()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create subscription '%s': %v", subName, err)
	}
	glog.V(logLevel).Infof("Newly created subscription: %s was successfully created", subName)
	//FIXME: can we remove this?
	if createdSub.Exists() {
		glog.V(logLevel).Infof("The newly created subscription '%s' in namespace '%v' has the current CSV: %v",
			subName, targetNs, createdSub.Definition.Status.CurrentCSV)
	}

	return createdSub, ogBuilderCreated, nil
}

func (d deploy) GetBundleConfig() (*BundleConfig, error) {

	log.Print("Getting the GPU bundle configs")

	var bundleConfig BundleConfig

	if err := envconfig.Process("gpu", &bundleConfig); err != nil {
		return nil, fmt.Errorf("failed to instantiate subscription configs: %v", err)
	}

	return &bundleConfig, nil
}

func (d deploy) GetMachineSetConfig() (*MachineSetConfig, error) {

	log.Print("Getting the GPU machineset config")

	var machineSetConfig MachineSetConfig

	if err := envconfig.Process("gpu", &machineSetConfig); err != nil {
		return nil, fmt.Errorf("failed to instantiate machineset config: %v", err)
	}

	return &machineSetConfig, nil
}

func (d deploy) CreateAndLabelNamespaceIfNeeded(logLevel glog.Level, targetNs string,
	labels map[string]string) (*namespace.Builder, error) {

	nsBuilder := namespace.NewBuilder(inittools.APIClient, targetNs)
	if nsBuilder.Exists() {
		glog.V(logLevel).Infof("The namespace '%s' already exists", targetNs)
	} else {
		glog.V(logLevel).Infof("Creating the namespace: %s", targetNs)
		createdNsBuilder, err := nsBuilder.Create()
		if err != nil {
			return nil, fmt.Errorf("failed to create namespace %s: %v", targetNs, err)
		}
		glog.V(logLevel).Infof("Successfully created namespace '%s'", targetNs)

		glog.V(logLevel).Infof("Labeling the newly created namespace '%s'", targetNs)
		nsBuilder, err = createdNsBuilder.WithMultipleLabels(labels).Update()
		if err != nil {
			return nil, fmt.Errorf("failed to label namespace %s with labels %v: %v", targetNs, labels, err)
		}
		glog.V(logLevel).Infof("Successfully labeled the namespace %s", targetNs)
	}

	return nsBuilder, nil
}
