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

const SettingRoamingAssistantKey = "roaming_assistant"

type SettingRoamingAssistant struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	Enabled          bool                       `json:"enabled"`
	Rssi             int                        `json:"rssi,omitempty"` // ^-([6-7][0-9]|80)$
	AdditionalFields map[string]json.RawMessage `json:"_additional_properties,omitempty"`
}

func (dst *SettingRoamingAssistant) UnmarshalJSON(b []byte) error {
	type Alias SettingRoamingAssistant
	aux := &struct {
		Rssi emptyStringInt `json:"rssi"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.Rssi = int(aux.Rssi)

	return nil
}

func (dst *SettingRoamingAssistant) KnownJSONFields() map[string]struct{} {
	return map[string]struct{}{
		"_id":            {},
		"site_id":        {},
		"attr_hidden":    {},
		"attr_hidden_id": {},
		"attr_no_delete": {},
		"attr_no_edit":   {},
		"key":            {},
		"enabled":        {},
		"rssi":           {},
	}
}

func (dst *SettingRoamingAssistant) SetAdditionalFields(ef map[string]json.RawMessage) {
	dst.AdditionalFields = ef
}

// GetSettingRoamingAssistant Experimental! This function is not yet stable and may change in the future.
func (c *client) GetSettingRoamingAssistant(ctx context.Context, site string) (*SettingRoamingAssistant, error) {
	s, f, err := c.GetSetting(ctx, site, SettingRoamingAssistantKey)
	if err != nil {
		return nil, err
	}
	if s.Key != SettingRoamingAssistantKey {
		return nil, fmt.Errorf("unexpected setting key received. Requested: %q, received: %q", SettingRoamingAssistantKey, s.Key)
	}
	return f.(*SettingRoamingAssistant), nil
}

// UpdateSettingRoamingAssistant Experimental! This function is not yet stable and may change in the future.
func (c *client) UpdateSettingRoamingAssistant(ctx context.Context, site string, s *SettingRoamingAssistant) (*SettingRoamingAssistant, error) {
	s.Key = SettingRoamingAssistantKey
	result, err := c.SetSetting(ctx, site, SettingRoamingAssistantKey, struct {
		*SettingRoamingAssistant
		AdditionalFields *struct{} `json:"_additional_properties,omitempty"`
	}{SettingRoamingAssistant: s})
	if err != nil {
		return nil, err
	}
	return result.(*SettingRoamingAssistant), nil
}
