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

const SettingDpiKey = "dpi"

type SettingDpi struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	Enabled               bool                       `json:"enabled"`
	FingerprintingEnabled bool                       `json:"fingerprintingEnabled"`
	AdditionalFields      map[string]json.RawMessage `json:"_additional_properties,omitempty"`
}

func (dst *SettingDpi) UnmarshalJSON(b []byte) error {
	type Alias SettingDpi
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

func (dst *SettingDpi) KnownJSONFields() map[string]struct{} {
	return map[string]struct{}{
		"_id":                   {},
		"site_id":               {},
		"attr_hidden":           {},
		"attr_hidden_id":        {},
		"attr_no_delete":        {},
		"attr_no_edit":          {},
		"key":                   {},
		"enabled":               {},
		"fingerprintingEnabled": {},
	}
}

func (dst *SettingDpi) SetAdditionalFields(ef map[string]json.RawMessage) { dst.AdditionalFields = ef }

// GetSettingDpi Experimental! This function is not yet stable and may change in the future.
func (c *client) GetSettingDpi(ctx context.Context, site string) (*SettingDpi, error) {
	s, f, err := c.GetSetting(ctx, site, SettingDpiKey)
	if err != nil {
		return nil, err
	}
	if s.Key != SettingDpiKey {
		return nil, fmt.Errorf("unexpected setting key received. Requested: %q, received: %q", SettingDpiKey, s.Key)
	}
	return f.(*SettingDpi), nil
}

// UpdateSettingDpi Experimental! This function is not yet stable and may change in the future.
func (c *client) UpdateSettingDpi(ctx context.Context, site string, s *SettingDpi) (*SettingDpi, error) {
	s.Key = SettingDpiKey
	result, err := c.SetSetting(ctx, site, SettingDpiKey, struct {
		*SettingDpi
		AdditionalFields *struct{} `json:"_additional_properties,omitempty"`
	}{SettingDpi: s})
	if err != nil {
		return nil, err
	}
	return result.(*SettingDpi), nil
}
