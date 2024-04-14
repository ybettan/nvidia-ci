package deploy

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetMachineSetConfig", func() {

	var deploy Deploy

	BeforeEach(func() {

		deploy = NewDeploy()
	})

	It("should fail if it fails to process the env variables", func() {

		_, err := deploy.GetMachineSetConfig()
		Expect(err).To(HaveOccurred())
	})

	It("should work as expected", func() {

		const instanceType = "example-gpu-hw"

		os.Setenv("GPU_MACHINESET_INSTANCE_TYPE", instanceType)

		machineSetConfig, err := deploy.GetMachineSetConfig()
		Expect(err).NotTo(HaveOccurred())
		Expect(machineSetConfig.InstanceType).To(Equal(instanceType))
	})
})

var _ = Describe("GetSubscriptionConfig", func() {

	var deploy Deploy

	BeforeEach(func() {

		deploy = NewDeploy()
	})

	It("should fail if it fails to process the env variables", func() {

		_, err := deploy.GetSubscriptionConfig()
		Expect(err).To(HaveOccurred())
	})

	It("should work as expected", func() {

		const (
			catalogSource = "example-catalogsource"
			channel       = "example-channel"
		)

		os.Setenv("GPU_CATALOGSOURCE", catalogSource)
		os.Setenv("GPU_SUBSCRIPTION_CHANNEL", channel)

		subConfig, err := deploy.GetSubscriptionConfig()
		Expect(err).NotTo(HaveOccurred())
		Expect(subConfig.CatalogSource).To(Equal(catalogSource))
		Expect(subConfig.SubscriptionChannel).To(Equal(channel))
	})
})

var _ = Describe("GetBundleConfig", func() {

	var deploy Deploy

	BeforeEach(func() {

		deploy = NewDeploy()
	})

	It("should fail if it fails to process the env variables", func() {

		_, err := deploy.GetBundleConfig()
		Expect(err).To(HaveOccurred())
	})

	It("should work as expected", func() {

		const bundleImage = "registry/org/image:tag"

		os.Setenv("GPU_BUNDLE_IMAGE", bundleImage)

		bundleConfig, err := deploy.GetBundleConfig()
		Expect(err).NotTo(HaveOccurred())
		Expect(bundleConfig.BundleImage).To(Equal(bundleImage))
	})
})
