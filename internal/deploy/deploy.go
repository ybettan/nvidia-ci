package deploy

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
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

type Deploy interface {
	GetMachineSetConfig() (*MachineSetConfig, error)
	GetSubscriptionConfig() (*SubscriptionConfig, error)
	GetBundleConfig() (*BundleConfig, error)
}

type deployGPU struct{}

func NewDeploy() Deploy {
	return deployGPU{}
}

func (dgpu deployGPU) GetMachineSetConfig() (*MachineSetConfig, error) {

	log.Print("Getting the GPU machineset config")

	var machineSetConfig MachineSetConfig

	if err := envconfig.Process("gpu", &machineSetConfig); err != nil {
		return nil, fmt.Errorf("failed to instantiate machineset config: %v", err)
	}

	return &machineSetConfig, nil
}

func (dgpu deployGPU) GetSubscriptionConfig() (*SubscriptionConfig, error) {

	log.Print("Getting the GPU subscription config")

	var subConfig SubscriptionConfig

	if err := envconfig.Process("gpu", &subConfig); err != nil {
		return nil, fmt.Errorf("failed to instantiate subscription config: %v", err)
	}

	return &subConfig, nil
}

func (dgpu deployGPU) GetBundleConfig() (*BundleConfig, error) {

	log.Print("Getting the GPU bundle configs")

	var bundleConfig BundleConfig

	if err := envconfig.Process("gpu", &bundleConfig); err != nil {
		return nil, fmt.Errorf("failed to instantiate subscription configs: %v", err)
	}

	return &bundleConfig, nil
}
