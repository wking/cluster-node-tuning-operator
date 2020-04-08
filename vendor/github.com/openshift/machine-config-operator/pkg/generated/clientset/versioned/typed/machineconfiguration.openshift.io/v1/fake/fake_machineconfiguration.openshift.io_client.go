// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/openshift/machine-config-operator/pkg/generated/clientset/versioned/typed/machineconfiguration.openshift.io/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeMachineconfigurationV1 struct {
	*testing.Fake
}

func (c *FakeMachineconfigurationV1) ContainerRuntimeConfigs() v1.ContainerRuntimeConfigInterface {
	return &FakeContainerRuntimeConfigs{c}
}

func (c *FakeMachineconfigurationV1) ControllerConfigs() v1.ControllerConfigInterface {
	return &FakeControllerConfigs{c}
}

func (c *FakeMachineconfigurationV1) KubeletConfigs() v1.KubeletConfigInterface {
	return &FakeKubeletConfigs{c}
}

func (c *FakeMachineconfigurationV1) MachineConfigs() v1.MachineConfigInterface {
	return &FakeMachineConfigs{c}
}

func (c *FakeMachineconfigurationV1) MachineConfigPools() v1.MachineConfigPoolInterface {
	return &FakeMachineConfigPools{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeMachineconfigurationV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
