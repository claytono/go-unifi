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

const SettingUswKey = "usw"

type SettingUsw struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	DHCPSnoop        bool                       `json:"dhcp_snoop"`
	AdditionalFields map[string]json.RawMessage `json:"_additional_properties,omitempty"`
}

func (dst *SettingUsw) UnmarshalJSON(b []byte) error {
	type Alias SettingUsw
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

func (dst *SettingUsw) KnownJSONFields() map[string]struct{} {
	return map[string]struct{}{
		"_id":            {},
		"site_id":        {},
		"attr_hidden":    {},
		"attr_hidden_id": {},
		"attr_no_delete": {},
		"attr_no_edit":   {},
		"key":            {},
		"dhcp_snoop":     {},
	}
}

func (dst *SettingUsw) SetAdditionalFields(ef map[string]json.RawMessage) { dst.AdditionalFields = ef }

// GetSettingUsw Experimental! This function is not yet stable and may change in the future.
func (c *client) GetSettingUsw(ctx context.Context, site string) (*SettingUsw, error) {
	s, f, err := c.GetSetting(ctx, site, SettingUswKey)
	if err != nil {
		return nil, err
	}
	if s.Key != SettingUswKey {
		return nil, fmt.Errorf("unexpected setting key received. Requested: %q, received: %q", SettingUswKey, s.Key)
	}
	return f.(*SettingUsw), nil
}

// UpdateSettingUsw Experimental! This function is not yet stable and may change in the future.
func (c *client) UpdateSettingUsw(ctx context.Context, site string, s *SettingUsw) (*SettingUsw, error) {
	s.Key = SettingUswKey
	result, err := c.SetSetting(ctx, site, SettingUswKey, struct {
		*SettingUsw
		AdditionalFields *struct{} `json:"_additional_properties,omitempty"`
	}{SettingUsw: s})
	if err != nil {
		return nil, err
	}
	return result.(*SettingUsw), nil
}
