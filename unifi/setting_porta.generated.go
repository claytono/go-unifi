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

const SettingPortaKey = "porta"

type SettingPorta struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	Ugw3WAN2Enabled  bool                       `json:"ugw3_wan2_enabled"`
	AdditionalFields map[string]json.RawMessage `json:"_additional_properties,omitempty"`
}

func (dst *SettingPorta) UnmarshalJSON(b []byte) error {
	type Alias SettingPorta
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

func (dst *SettingPorta) KnownJSONFields() map[string]struct{} {
	return map[string]struct{}{
		"_id":               {},
		"site_id":           {},
		"attr_hidden":       {},
		"attr_hidden_id":    {},
		"attr_no_delete":    {},
		"attr_no_edit":      {},
		"key":               {},
		"ugw3_wan2_enabled": {},
	}
}

func (dst *SettingPorta) SetAdditionalFields(ef map[string]json.RawMessage) {
	dst.AdditionalFields = ef
}

// GetSettingPorta Experimental! This function is not yet stable and may change in the future.
func (c *client) GetSettingPorta(ctx context.Context, site string) (*SettingPorta, error) {
	s, f, err := c.GetSetting(ctx, site, SettingPortaKey)
	if err != nil {
		return nil, err
	}
	if s.Key != SettingPortaKey {
		return nil, fmt.Errorf("unexpected setting key received. Requested: %q, received: %q", SettingPortaKey, s.Key)
	}
	return f.(*SettingPorta), nil
}

// UpdateSettingPorta Experimental! This function is not yet stable and may change in the future.
func (c *client) UpdateSettingPorta(ctx context.Context, site string, s *SettingPorta) (*SettingPorta, error) {
	s.Key = SettingPortaKey
	result, err := c.SetSetting(ctx, site, SettingPortaKey, struct {
		*SettingPorta
		AdditionalFields *struct{} `json:"_additional_properties,omitempty"`
	}{SettingPorta: s})
	if err != nil {
		return nil, err
	}
	return result.(*SettingPorta), nil
}
