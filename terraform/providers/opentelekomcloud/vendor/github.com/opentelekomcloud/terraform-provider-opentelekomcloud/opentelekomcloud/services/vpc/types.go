package vpc

import (
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/eips"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/extensions/layer3/floatingips"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/extensions/layer3/routers"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/networks"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/ports"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/subnets"

	"github.com/opentelekomcloud/terraform-provider-opentelekomcloud/opentelekomcloud/common"
)

// FloatingIPCreateOpts represents the attributes used when creating a new floating ip.
type FloatingIPCreateOpts struct {
	floatingips.CreateOpts
	ValueSpecs map[string]string `json:"value_specs,omitempty"`
}

// ToFloatingIPCreateMap casts a CreateOpts struct to a map.
// It overrides floatingips.ToFloatingIPCreateMap to add the ValueSpecs field.
func (opts FloatingIPCreateOpts) ToFloatingIPCreateMap() (map[string]interface{}, error) {
	return common.BuildRequest(opts, "floatingip")
}

// NetworkCreateOpts represents the attributes used when creating a new network.
type NetworkCreateOpts struct {
	networks.CreateOpts
	ValueSpecs map[string]string `json:"value_specs,omitempty"`
}

// ToNetworkCreateMap casts a CreateOpts struct to a map.
// It overrides networks.ToNetworkCreateMap to add the ValueSpecs field.
func (opts NetworkCreateOpts) ToNetworkCreateMap() (map[string]interface{}, error) {
	return common.BuildRequest(opts, "network")
}

// PortCreateOpts represents the attributes used when creating a new port.
type PortCreateOpts struct {
	ports.CreateOpts
	ValueSpecs map[string]string `json:"value_specs,omitempty"`
}

// ToPortCreateMap casts a CreateOpts struct to a map.
// It overrides ports.ToPortCreateMap to add the ValueSpecs field.
func (opts PortCreateOpts) ToPortCreateMap() (map[string]interface{}, error) {
	return common.BuildRequest(opts, "port")
}

// RouterCreateOpts represents the attributes used when creating a new router.
type RouterCreateOpts struct {
	routers.CreateOpts
	ValueSpecs map[string]string `json:"value_specs,omitempty"`
}

// ToRouterCreateMap casts a CreateOpts struct to a map.
// It overrides routers.ToRouterCreateMap to add the ValueSpecs field.
func (opts RouterCreateOpts) ToRouterCreateMap() (map[string]interface{}, error) {
	return common.BuildRequest(opts, "router")
}

// SubnetCreateOpts represents the attributes used when creating a new subnet.
type SubnetCreateOpts struct {
	subnets.CreateOpts
	ValueSpecs map[string]string `json:"value_specs,omitempty"`
}

// ToSubnetCreateMap casts a CreateOpts struct to a map.
// It overrides subnets.ToSubnetCreateMap to add the ValueSpecs field.
func (opts SubnetCreateOpts) ToSubnetCreateMap() (map[string]interface{}, error) {
	b, err := common.BuildRequest(opts, "subnet")
	if err != nil {
		return nil, err
	}

	if m := b["subnet"].(map[string]interface{}); m["gateway_ip"] == "" {
		m["gateway_ip"] = nil
	}

	return b, nil
}

// EIPCreateOpts represents the attributes used when creating a new eip.
type EIPCreateOpts struct {
	eips.ApplyOpts
	ValueSpecs map[string]string `json:"value_specs,omitempty"`
}
