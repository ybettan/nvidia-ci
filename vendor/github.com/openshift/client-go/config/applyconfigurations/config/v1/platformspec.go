// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/config/v1"
)

// PlatformSpecApplyConfiguration represents an declarative configuration of the PlatformSpec type for use
// with apply.
type PlatformSpecApplyConfiguration struct {
	Type         *v1.PlatformType                        `json:"type,omitempty"`
	AWS          *AWSPlatformSpecApplyConfiguration      `json:"aws,omitempty"`
	Azure        *v1.AzurePlatformSpec                   `json:"azure,omitempty"`
	GCP          *v1.GCPPlatformSpec                     `json:"gcp,omitempty"`
	BareMetal    *v1.BareMetalPlatformSpec               `json:"baremetal,omitempty"`
	OpenStack    *v1.OpenStackPlatformSpec               `json:"openstack,omitempty"`
	Ovirt        *v1.OvirtPlatformSpec                   `json:"ovirt,omitempty"`
	VSphere      *VSpherePlatformSpecApplyConfiguration  `json:"vsphere,omitempty"`
	IBMCloud     *v1.IBMCloudPlatformSpec                `json:"ibmcloud,omitempty"`
	Kubevirt     *v1.KubevirtPlatformSpec                `json:"kubevirt,omitempty"`
	EquinixMetal *v1.EquinixMetalPlatformSpec            `json:"equinixMetal,omitempty"`
	PowerVS      *PowerVSPlatformSpecApplyConfiguration  `json:"powervs,omitempty"`
	AlibabaCloud *v1.AlibabaCloudPlatformSpec            `json:"alibabaCloud,omitempty"`
	Nutanix      *NutanixPlatformSpecApplyConfiguration  `json:"nutanix,omitempty"`
	External     *ExternalPlatformSpecApplyConfiguration `json:"external,omitempty"`
}

// PlatformSpecApplyConfiguration constructs an declarative configuration of the PlatformSpec type for use with
// apply.
func PlatformSpec() *PlatformSpecApplyConfiguration {
	return &PlatformSpecApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *PlatformSpecApplyConfiguration) WithType(value v1.PlatformType) *PlatformSpecApplyConfiguration {
	b.Type = &value
	return b
}

// WithAWS sets the AWS field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AWS field is set to the value of the last call.
func (b *PlatformSpecApplyConfiguration) WithAWS(value *AWSPlatformSpecApplyConfiguration) *PlatformSpecApplyConfiguration {
	b.AWS = value
	return b
}

// WithAzure sets the Azure field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Azure field is set to the value of the last call.
func (b *PlatformSpecApplyConfiguration) WithAzure(value v1.AzurePlatformSpec) *PlatformSpecApplyConfiguration {
	b.Azure = &value
	return b
}

// WithGCP sets the GCP field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GCP field is set to the value of the last call.
func (b *PlatformSpecApplyConfiguration) WithGCP(value v1.GCPPlatformSpec) *PlatformSpecApplyConfiguration {
	b.GCP = &value
	return b
}

// WithBareMetal sets the BareMetal field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BareMetal field is set to the value of the last call.
func (b *PlatformSpecApplyConfiguration) WithBareMetal(value v1.BareMetalPlatformSpec) *PlatformSpecApplyConfiguration {
	b.BareMetal = &value
	return b
}

// WithOpenStack sets the OpenStack field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OpenStack field is set to the value of the last call.
func (b *PlatformSpecApplyConfiguration) WithOpenStack(value v1.OpenStackPlatformSpec) *PlatformSpecApplyConfiguration {
	b.OpenStack = &value
	return b
}

// WithOvirt sets the Ovirt field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Ovirt field is set to the value of the last call.
func (b *PlatformSpecApplyConfiguration) WithOvirt(value v1.OvirtPlatformSpec) *PlatformSpecApplyConfiguration {
	b.Ovirt = &value
	return b
}

// WithVSphere sets the VSphere field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VSphere field is set to the value of the last call.
func (b *PlatformSpecApplyConfiguration) WithVSphere(value *VSpherePlatformSpecApplyConfiguration) *PlatformSpecApplyConfiguration {
	b.VSphere = value
	return b
}

// WithIBMCloud sets the IBMCloud field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IBMCloud field is set to the value of the last call.
func (b *PlatformSpecApplyConfiguration) WithIBMCloud(value v1.IBMCloudPlatformSpec) *PlatformSpecApplyConfiguration {
	b.IBMCloud = &value
	return b
}

// WithKubevirt sets the Kubevirt field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kubevirt field is set to the value of the last call.
func (b *PlatformSpecApplyConfiguration) WithKubevirt(value v1.KubevirtPlatformSpec) *PlatformSpecApplyConfiguration {
	b.Kubevirt = &value
	return b
}

// WithEquinixMetal sets the EquinixMetal field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EquinixMetal field is set to the value of the last call.
func (b *PlatformSpecApplyConfiguration) WithEquinixMetal(value v1.EquinixMetalPlatformSpec) *PlatformSpecApplyConfiguration {
	b.EquinixMetal = &value
	return b
}

// WithPowerVS sets the PowerVS field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PowerVS field is set to the value of the last call.
func (b *PlatformSpecApplyConfiguration) WithPowerVS(value *PowerVSPlatformSpecApplyConfiguration) *PlatformSpecApplyConfiguration {
	b.PowerVS = value
	return b
}

// WithAlibabaCloud sets the AlibabaCloud field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AlibabaCloud field is set to the value of the last call.
func (b *PlatformSpecApplyConfiguration) WithAlibabaCloud(value v1.AlibabaCloudPlatformSpec) *PlatformSpecApplyConfiguration {
	b.AlibabaCloud = &value
	return b
}

// WithNutanix sets the Nutanix field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Nutanix field is set to the value of the last call.
func (b *PlatformSpecApplyConfiguration) WithNutanix(value *NutanixPlatformSpecApplyConfiguration) *PlatformSpecApplyConfiguration {
	b.Nutanix = value
	return b
}

// WithExternal sets the External field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the External field is set to the value of the last call.
func (b *PlatformSpecApplyConfiguration) WithExternal(value *ExternalPlatformSpecApplyConfiguration) *PlatformSpecApplyConfiguration {
	b.External = value
	return b
}
