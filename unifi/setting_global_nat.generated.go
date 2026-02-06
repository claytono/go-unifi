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

const SettingGlobalNatKey = "global_nat"

type SettingGlobalNat struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	ExcludedNetworkIDs []string                   `json:"excluded_network_ids,omitempty"`
	Mode               string                     `json:"mode,omitempty" validate:"omitempty,oneof=auto custom off"` // auto|custom|off
	AdditionalFields   map[string]json.RawMessage `json:"_additional_properties,omitempty"`
}

func (dst *SettingGlobalNat) UnmarshalJSON(b []byte) error {
	type Alias SettingGlobalNat
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

func (dst *SettingGlobalNat) KnownJSONFields() map[string]struct{} {
	return map[string]struct{}{
		"_id":                  {},
		"site_id":              {},
		"attr_hidden":          {},
		"attr_hidden_id":       {},
		"attr_no_delete":       {},
		"attr_no_edit":         {},
		"key":                  {},
		"excluded_network_ids": {},
		"mode":                 {},
	}
}

func (dst *SettingGlobalNat) SetAdditionalFields(ef map[string]json.RawMessage) {
	dst.AdditionalFields = ef
}

// GetSettingGlobalNat Experimental! This function is not yet stable and may change in the future.
func (c *client) GetSettingGlobalNat(ctx context.Context, site string) (*SettingGlobalNat, error) {
	s, f, err := c.GetSetting(ctx, site, SettingGlobalNatKey)
	if err != nil {
		return nil, err
	}
	if s.Key != SettingGlobalNatKey {
		return nil, fmt.Errorf("unexpected setting key received. Requested: %q, received: %q", SettingGlobalNatKey, s.Key)
	}
	return f.(*SettingGlobalNat), nil
}

// UpdateSettingGlobalNat Experimental! This function is not yet stable and may change in the future.
func (c *client) UpdateSettingGlobalNat(ctx context.Context, site string, s *SettingGlobalNat) (*SettingGlobalNat, error) {
	s.Key = SettingGlobalNatKey
	result, err := c.SetSetting(ctx, site, SettingGlobalNatKey, struct {
		*SettingGlobalNat
		AdditionalFields *struct{} `json:"_additional_properties,omitempty"`
	}{SettingGlobalNat: s})
	if err != nil {
		return nil, err
	}
	return result.(*SettingGlobalNat), nil
}
