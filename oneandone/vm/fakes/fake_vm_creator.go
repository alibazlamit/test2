package fakes

import (
	"github.com/bosh-oneandone-cpi/oneandone/resource"
	"github.com/bosh-oneandone-cpi/oneandone/vm"
)

type FakeVMCreator struct {
	Configuration vm.InstanceConfiguration
	Metadata      vm.InstanceMetadata

	CreateInstanceResult *resource.Instance
	CreateInstanceError  error
	CreateInstanceCalled bool
}

func (f *FakeVMCreator) CreateInstance(icfg vm.InstanceConfiguration,
	md vm.InstanceMetadata) (*resource.Instance, error) {

	f.CreateInstanceCalled = true
	f.Configuration = icfg
	f.Metadata = md
	return f.CreateInstanceResult, f.CreateInstanceError
}
