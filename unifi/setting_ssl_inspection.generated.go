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

const SettingSslInspectionKey = "ssl_inspection"

type SettingSslInspection struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	State            string                     `json:"state,omitempty" validate:"omitempty,oneof=off simple advanced"` // off|simple|advanced
	AdditionalFields map[string]json.RawMessage `json:"_additional_properties,omitempty"`
}

func (dst *SettingSslInspection) UnmarshalJSON(b []byte) error {
	type Alias SettingSslInspection
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

func (dst *SettingSslInspection) KnownJSONFields() map[string]struct{} {
	return map[string]struct{}{
		"_id":            {},
		"site_id":        {},
		"attr_hidden":    {},
		"attr_hidden_id": {},
		"attr_no_delete": {},
		"attr_no_edit":   {},
		"key":            {},
		"state":          {},
	}
}

func (dst *SettingSslInspection) SetAdditionalFields(ef map[string]json.RawMessage) {
	dst.AdditionalFields = ef
}

// GetSettingSslInspection Experimental! This function is not yet stable and may change in the future.
func (c *client) GetSettingSslInspection(ctx context.Context, site string) (*SettingSslInspection, error) {
	s, f, err := c.GetSetting(ctx, site, SettingSslInspectionKey)
	if err != nil {
		return nil, err
	}
	if s.Key != SettingSslInspectionKey {
		return nil, fmt.Errorf("unexpected setting key received. Requested: %q, received: %q", SettingSslInspectionKey, s.Key)
	}
	return f.(*SettingSslInspection), nil
}

// UpdateSettingSslInspection Experimental! This function is not yet stable and may change in the future.
func (c *client) UpdateSettingSslInspection(ctx context.Context, site string, s *SettingSslInspection) (*SettingSslInspection, error) {
	s.Key = SettingSslInspectionKey
	result, err := c.SetSetting(ctx, site, SettingSslInspectionKey, struct {
		*SettingSslInspection
		AdditionalFields *struct{} `json:"_additional_properties,omitempty"`
	}{SettingSslInspection: s})
	if err != nil {
		return nil, err
	}
	return result.(*SettingSslInspection), nil
}
