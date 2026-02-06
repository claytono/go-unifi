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

const SettingLocaleKey = "locale"

type SettingLocale struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	Timezone         string                     `json:"timezone,omitempty"`
	AdditionalFields map[string]json.RawMessage `json:"_additional_properties,omitempty"`
}

func (dst *SettingLocale) UnmarshalJSON(b []byte) error {
	type Alias SettingLocale
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

func (dst *SettingLocale) KnownJSONFields() map[string]struct{} {
	return map[string]struct{}{
		"_id":            {},
		"site_id":        {},
		"attr_hidden":    {},
		"attr_hidden_id": {},
		"attr_no_delete": {},
		"attr_no_edit":   {},
		"key":            {},
		"timezone":       {},
	}
}

func (dst *SettingLocale) SetAdditionalFields(ef map[string]json.RawMessage) {
	dst.AdditionalFields = ef
}

// GetSettingLocale Experimental! This function is not yet stable and may change in the future.
func (c *client) GetSettingLocale(ctx context.Context, site string) (*SettingLocale, error) {
	s, f, err := c.GetSetting(ctx, site, SettingLocaleKey)
	if err != nil {
		return nil, err
	}
	if s.Key != SettingLocaleKey {
		return nil, fmt.Errorf("unexpected setting key received. Requested: %q, received: %q", SettingLocaleKey, s.Key)
	}
	return f.(*SettingLocale), nil
}

// UpdateSettingLocale Experimental! This function is not yet stable and may change in the future.
func (c *client) UpdateSettingLocale(ctx context.Context, site string, s *SettingLocale) (*SettingLocale, error) {
	s.Key = SettingLocaleKey
	result, err := c.SetSetting(ctx, site, SettingLocaleKey, struct {
		*SettingLocale
		AdditionalFields *struct{} `json:"_additional_properties,omitempty"`
	}{SettingLocale: s})
	if err != nil {
		return nil, err
	}
	return result.(*SettingLocale), nil
}
