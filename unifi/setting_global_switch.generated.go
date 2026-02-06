// Code generated from ace.jar fields *.json files
// DO NOT EDIT.

package unifi

import (
	"context"
	"encoding/json"
	"fmt"
)

// just to fix compile issues with the import
var (
	_ context.Context
	_ fmt.Formatter
	_ json.Marshaler
)

const SettingGlobalSwitchKey = "global_switch"

type SettingGlobalSwitch struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	AclDeviceIsolation             []string                            `json:"acl_device_isolation,omitempty"`
	AclL3Isolation                 []SettingGlobalSwitchAclL3Isolation `json:"acl_l3_isolation,omitempty"`
	DHCPSnoop                      bool                                `json:"dhcp_snoop"`
	Dot1XFallbackNetworkID         string                              `json:"dot1x_fallback_networkconf_id"` // [\d\w]+|
	Dot1XPortctrlEnabled           bool                                `json:"dot1x_portctrl_enabled"`
	FloodKnownProtocols            bool                                `json:"flood_known_protocols"`
	FlowctrlEnabled                bool                                `json:"flowctrl_enabled"`
	ForwardUnknownMcastRouterPorts bool                                `json:"forward_unknown_mcast_router_ports"`
	JumboframeEnabled              bool                                `json:"jumboframe_enabled"`
	RADIUSProfileID                string                              `json:"radiusprofile_id"`
	StpVersion                     string                              `json:"stp_version,omitempty" validate:"omitempty,oneof=stp rstp disabled"` // stp|rstp|disabled
	SwitchExclusions               []string                            `json:"switch_exclusions,omitempty" validate:"omitempty,mac"`               // ^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$
	AdditionalFields               map[string]json.RawMessage          `json:"_additional_properties,omitempty"`
}

func (dst *SettingGlobalSwitch) UnmarshalJSON(b []byte) error {
	type Alias SettingGlobalSwitch
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}

	return nil
}

func (dst *SettingGlobalSwitch) KnownJSONFields() map[string]struct{} {
	return map[string]struct{}{
		"_id":                                {},
		"site_id":                            {},
		"attr_hidden":                        {},
		"attr_hidden_id":                     {},
		"attr_no_delete":                     {},
		"attr_no_edit":                       {},
		"key":                                {},
		"acl_device_isolation":               {},
		"acl_l3_isolation":                   {},
		"dhcp_snoop":                         {},
		"dot1x_fallback_networkconf_id":      {},
		"dot1x_portctrl_enabled":             {},
		"flood_known_protocols":              {},
		"flowctrl_enabled":                   {},
		"forward_unknown_mcast_router_ports": {},
		"jumboframe_enabled":                 {},
		"radiusprofile_id":                   {},
		"stp_version":                        {},
		"switch_exclusions":                  {},
	}
}

func (dst *SettingGlobalSwitch) SetAdditionalFields(ef map[string]json.RawMessage) {
	dst.AdditionalFields = ef
}

type SettingGlobalSwitchAclL3Isolation struct {
	DestinationNetworks []string                   `json:"destination_networks,omitempty"`
	SourceNetwork       string                     `json:"source_network,omitempty"`
	AdditionalFields    map[string]json.RawMessage `json:"_additional_properties,omitempty"`
}

func (dst *SettingGlobalSwitchAclL3Isolation) UnmarshalJSON(b []byte) error {
	type Alias SettingGlobalSwitchAclL3Isolation
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}

	return nil
}

func (dst *SettingGlobalSwitchAclL3Isolation) KnownJSONFields() map[string]struct{} {
	return map[string]struct{}{
		"destination_networks": {},
		"source_network":       {},
	}
}

func (dst *SettingGlobalSwitchAclL3Isolation) SetAdditionalFields(ef map[string]json.RawMessage) {
	dst.AdditionalFields = ef
}

// GetSettingGlobalSwitch Experimental! This function is not yet stable and may change in the future.
func (c *client) GetSettingGlobalSwitch(ctx context.Context, site string) (*SettingGlobalSwitch, error) {
	s, f, err := c.GetSetting(ctx, site, SettingGlobalSwitchKey)
	if err != nil {
		return nil, err
	}
	if s.Key != SettingGlobalSwitchKey {
		return nil, fmt.Errorf("unexpected setting key received. Requested: %q, received: %q", SettingGlobalSwitchKey, s.Key)
	}
	return f.(*SettingGlobalSwitch), nil
}

// UpdateSettingGlobalSwitch Experimental! This function is not yet stable and may change in the future.
func (c *client) UpdateSettingGlobalSwitch(ctx context.Context, site string, s *SettingGlobalSwitch) (*SettingGlobalSwitch, error) {
	s.Key = SettingGlobalSwitchKey
	result, err := c.SetSetting(ctx, site, SettingGlobalSwitchKey, struct {
		*SettingGlobalSwitch
		AdditionalFields *struct{} `json:"_additional_properties,omitempty"`
	}{SettingGlobalSwitch: s})
	if err != nil {
		return nil, err
	}
	return result.(*SettingGlobalSwitch), nil
}
