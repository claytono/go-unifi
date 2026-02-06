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

const SettingDashboardKey = "dashboard"

type SettingDashboard struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	LayoutPreference string                     `json:"layout_preference,omitempty" validate:"omitempty,oneof=auto manual"` // auto|manual
	Widgets          []SettingDashboardWidgets  `json:"widgets,omitempty"`
	AdditionalFields map[string]json.RawMessage `json:"_additional_properties,omitempty"`
}

func (dst *SettingDashboard) UnmarshalJSON(b []byte) error {
	type Alias SettingDashboard
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

func (dst *SettingDashboard) KnownJSONFields() map[string]struct{} {
	return map[string]struct{}{
		"_id":               {},
		"site_id":           {},
		"attr_hidden":       {},
		"attr_hidden_id":    {},
		"attr_no_delete":    {},
		"attr_no_edit":      {},
		"key":               {},
		"layout_preference": {},
		"widgets":           {},
	}
}

func (dst *SettingDashboard) SetAdditionalFields(ef map[string]json.RawMessage) {
	dst.AdditionalFields = ef
}

type SettingDashboardWidgets struct {
	Enabled          bool                       `json:"enabled"`
	Name             string                     `json:"name,omitempty" validate:"omitempty,oneof=critical_traffic_prioritization cybersecure traffic_identification wifi_technology wifi_channels wifi_client_experience wifi_tx_retries most_active_apps_aps_clients most_active_apps_clients most_active_aps_clients most_active_apps_aps most_active_apps v2_most_active_aps v2_most_active_clients wifi_connectivity ap_radio_density wifi_channel_preset_configuration"` // critical_traffic_prioritization|cybersecure|traffic_identification|wifi_technology|wifi_channels|wifi_client_experience|wifi_tx_retries|most_active_apps_aps_clients|most_active_apps_clients|most_active_aps_clients|most_active_apps_aps|most_active_apps|v2_most_active_aps|v2_most_active_clients|wifi_connectivity|ap_radio_density|wifi_channel_preset_configuration
	AdditionalFields map[string]json.RawMessage `json:"_additional_properties,omitempty"`
}

func (dst *SettingDashboardWidgets) UnmarshalJSON(b []byte) error {
	type Alias SettingDashboardWidgets
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

func (dst *SettingDashboardWidgets) KnownJSONFields() map[string]struct{} {
	return map[string]struct{}{
		"enabled": {},
		"name":    {},
	}
}

func (dst *SettingDashboardWidgets) SetAdditionalFields(ef map[string]json.RawMessage) {
	dst.AdditionalFields = ef
}

// GetSettingDashboard Experimental! This function is not yet stable and may change in the future.
func (c *client) GetSettingDashboard(ctx context.Context, site string) (*SettingDashboard, error) {
	s, f, err := c.GetSetting(ctx, site, SettingDashboardKey)
	if err != nil {
		return nil, err
	}
	if s.Key != SettingDashboardKey {
		return nil, fmt.Errorf("unexpected setting key received. Requested: %q, received: %q", SettingDashboardKey, s.Key)
	}
	return f.(*SettingDashboard), nil
}

// UpdateSettingDashboard Experimental! This function is not yet stable and may change in the future.
func (c *client) UpdateSettingDashboard(ctx context.Context, site string, s *SettingDashboard) (*SettingDashboard, error) {
	s.Key = SettingDashboardKey
	result, err := c.SetSetting(ctx, site, SettingDashboardKey, struct {
		*SettingDashboard
		AdditionalFields *struct{} `json:"_additional_properties,omitempty"`
	}{SettingDashboard: s})
	if err != nil {
		return nil, err
	}
	return result.(*SettingDashboard), nil
}
