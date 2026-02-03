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

const SettingEtherLightingKey = "ether_lighting"

type SettingEtherLighting struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	NetworkOverrides []SettingEtherLightingNetworkOverrides `json:"network_overrides,omitempty"`
	SpeedOverrides   []SettingEtherLightingSpeedOverrides   `json:"speed_overrides,omitempty"`
	ExtraFields      map[string]json.RawMessage             `json:"-"`
}

func (dst *SettingEtherLighting) UnmarshalJSON(b []byte) error {
	type Alias SettingEtherLighting
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}

	// Capture extra fields not in the struct
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(b, &raw); err == nil {
		known := map[string]struct{}{
			"_id":               {},
			"site_id":           {},
			"attr_hidden":       {},
			"attr_hidden_id":    {},
			"attr_no_delete":    {},
			"attr_no_edit":      {},
			"key":               {},
			"network_overrides": {},
			"speed_overrides":   {},
		}
		for k, v := range raw {
			if _, ok := known[k]; !ok {
				if dst.ExtraFields == nil {
					dst.ExtraFields = make(map[string]json.RawMessage)
				}
				dst.ExtraFields[k] = v
			}
		}
	}

	return nil
}

func (src SettingEtherLighting) MarshalJSON() ([]byte, error) {
	type Alias SettingEtherLighting
	b, err := json.Marshal(Alias(src))
	if err != nil {
		return nil, err
	}
	if len(src.ExtraFields) == 0 {
		return b, nil
	}
	var m map[string]json.RawMessage
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	extra, err := json.Marshal(src.ExtraFields)
	if err != nil {
		return nil, err
	}
	m["_additional_properties"] = extra
	return json.Marshal(m)
}

type SettingEtherLightingNetworkOverrides struct {
	Key         string                     `json:"key,omitempty"`
	RawColorHex string                     `json:"raw_color_hex,omitempty"` // [0-9A-Fa-f]{6}
	ExtraFields map[string]json.RawMessage `json:"-"`
}

func (dst *SettingEtherLightingNetworkOverrides) UnmarshalJSON(b []byte) error {
	type Alias SettingEtherLightingNetworkOverrides
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}

	// Capture extra fields not in the struct
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(b, &raw); err == nil {
		known := map[string]struct{}{
			"key":           {},
			"raw_color_hex": {},
		}
		for k, v := range raw {
			if _, ok := known[k]; !ok {
				if dst.ExtraFields == nil {
					dst.ExtraFields = make(map[string]json.RawMessage)
				}
				dst.ExtraFields[k] = v
			}
		}
	}

	return nil
}

func (src SettingEtherLightingNetworkOverrides) MarshalJSON() ([]byte, error) {
	type Alias SettingEtherLightingNetworkOverrides
	b, err := json.Marshal(Alias(src))
	if err != nil {
		return nil, err
	}
	if len(src.ExtraFields) == 0 {
		return b, nil
	}
	var m map[string]json.RawMessage
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	extra, err := json.Marshal(src.ExtraFields)
	if err != nil {
		return nil, err
	}
	m["_additional_properties"] = extra
	return json.Marshal(m)
}

type SettingEtherLightingSpeedOverrides struct {
	Key         string                     `json:"key,omitempty" validate:"omitempty,oneof=FE GbE 2.5GbE 5GbE 10GbE 25GbE 40GbE 100GbE"` // FE|GbE|2.5GbE|5GbE|10GbE|25GbE|40GbE|100GbE
	RawColorHex string                     `json:"raw_color_hex,omitempty"`                                                              // [0-9A-Fa-f]{6}
	ExtraFields map[string]json.RawMessage `json:"-"`
}

func (dst *SettingEtherLightingSpeedOverrides) UnmarshalJSON(b []byte) error {
	type Alias SettingEtherLightingSpeedOverrides
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}

	// Capture extra fields not in the struct
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(b, &raw); err == nil {
		known := map[string]struct{}{
			"key":           {},
			"raw_color_hex": {},
		}
		for k, v := range raw {
			if _, ok := known[k]; !ok {
				if dst.ExtraFields == nil {
					dst.ExtraFields = make(map[string]json.RawMessage)
				}
				dst.ExtraFields[k] = v
			}
		}
	}

	return nil
}

func (src SettingEtherLightingSpeedOverrides) MarshalJSON() ([]byte, error) {
	type Alias SettingEtherLightingSpeedOverrides
	b, err := json.Marshal(Alias(src))
	if err != nil {
		return nil, err
	}
	if len(src.ExtraFields) == 0 {
		return b, nil
	}
	var m map[string]json.RawMessage
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	extra, err := json.Marshal(src.ExtraFields)
	if err != nil {
		return nil, err
	}
	m["_additional_properties"] = extra
	return json.Marshal(m)
}

// GetSettingEtherLighting Experimental! This function is not yet stable and may change in the future.
func (c *client) GetSettingEtherLighting(ctx context.Context, site string) (*SettingEtherLighting, error) {
	s, f, err := c.GetSetting(ctx, site, SettingEtherLightingKey)
	if err != nil {
		return nil, err
	}
	if s.Key != SettingEtherLightingKey {
		return nil, fmt.Errorf("unexpected setting key received. Requested: %q, received: %q", SettingEtherLightingKey, s.Key)
	}
	return f.(*SettingEtherLighting), nil
}

// UpdateSettingEtherLighting Experimental! This function is not yet stable and may change in the future.
func (c *client) UpdateSettingEtherLighting(ctx context.Context, site string, s *SettingEtherLighting) (*SettingEtherLighting, error) {
	s.Key = SettingEtherLightingKey
	result, err := c.SetSetting(ctx, site, SettingEtherLightingKey, s)
	if err != nil {
		return nil, err
	}
	return result.(*SettingEtherLighting), nil
}
