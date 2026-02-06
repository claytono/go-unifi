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

const SettingTrafficFlowKey = "traffic_flow"

type SettingTrafficFlow struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	EnabledAllowedTraffic        bool                       `json:"enabled_allowed_traffic"`
	GatewayDNSEnabled            bool                       `json:"gateway_dns_enabled"`
	UnifiDeviceManagementEnabled bool                       `json:"unifi_device_management_enabled"`
	UnifiServicesEnabled         bool                       `json:"unifi_services_enabled"`
	AdditionalFields             map[string]json.RawMessage `json:"_additional_properties,omitempty"`
}

func (dst *SettingTrafficFlow) UnmarshalJSON(b []byte) error {
	type Alias SettingTrafficFlow
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

func (dst *SettingTrafficFlow) KnownJSONFields() map[string]struct{} {
	return map[string]struct{}{
		"_id":                             {},
		"site_id":                         {},
		"attr_hidden":                     {},
		"attr_hidden_id":                  {},
		"attr_no_delete":                  {},
		"attr_no_edit":                    {},
		"key":                             {},
		"enabled_allowed_traffic":         {},
		"gateway_dns_enabled":             {},
		"unifi_device_management_enabled": {},
		"unifi_services_enabled":          {},
	}
}

func (dst *SettingTrafficFlow) SetAdditionalFields(ef map[string]json.RawMessage) {
	dst.AdditionalFields = ef
}

// GetSettingTrafficFlow Experimental! This function is not yet stable and may change in the future.
func (c *client) GetSettingTrafficFlow(ctx context.Context, site string) (*SettingTrafficFlow, error) {
	s, f, err := c.GetSetting(ctx, site, SettingTrafficFlowKey)
	if err != nil {
		return nil, err
	}
	if s.Key != SettingTrafficFlowKey {
		return nil, fmt.Errorf("unexpected setting key received. Requested: %q, received: %q", SettingTrafficFlowKey, s.Key)
	}
	return f.(*SettingTrafficFlow), nil
}

// UpdateSettingTrafficFlow Experimental! This function is not yet stable and may change in the future.
func (c *client) UpdateSettingTrafficFlow(ctx context.Context, site string, s *SettingTrafficFlow) (*SettingTrafficFlow, error) {
	s.Key = SettingTrafficFlowKey
	result, err := c.SetSetting(ctx, site, SettingTrafficFlowKey, struct {
		*SettingTrafficFlow
		AdditionalFields *struct{} `json:"_additional_properties,omitempty"`
	}{SettingTrafficFlow: s})
	if err != nil {
		return nil, err
	}
	return result.(*SettingTrafficFlow), nil
}
