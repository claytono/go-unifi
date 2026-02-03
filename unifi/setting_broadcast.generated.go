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

const SettingBroadcastKey = "broadcast"

type SettingBroadcast struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	SoundAfterEnabled   bool                       `json:"sound_after_enabled"`
	SoundAfterResource  string                     `json:"sound_after_resource,omitempty"`
	SoundAfterType      string                     `json:"sound_after_type,omitempty" validate:"omitempty,oneof=sample media"` // sample|media
	SoundBeforeEnabled  bool                       `json:"sound_before_enabled"`
	SoundBeforeResource string                     `json:"sound_before_resource,omitempty"`
	SoundBeforeType     string                     `json:"sound_before_type,omitempty" validate:"omitempty,oneof=sample media"` // sample|media
	ExtraFields         map[string]json.RawMessage `json:"-"`
}

func (dst *SettingBroadcast) UnmarshalJSON(b []byte) error {
	type Alias SettingBroadcast
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
			"_id":                   {},
			"site_id":               {},
			"attr_hidden":           {},
			"attr_hidden_id":        {},
			"attr_no_delete":        {},
			"attr_no_edit":          {},
			"key":                   {},
			"sound_after_enabled":   {},
			"sound_after_resource":  {},
			"sound_after_type":      {},
			"sound_before_enabled":  {},
			"sound_before_resource": {},
			"sound_before_type":     {},
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

func (src SettingBroadcast) MarshalJSON() ([]byte, error) {
	type Alias SettingBroadcast
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

// GetSettingBroadcast Experimental! This function is not yet stable and may change in the future.
func (c *client) GetSettingBroadcast(ctx context.Context, site string) (*SettingBroadcast, error) {
	s, f, err := c.GetSetting(ctx, site, SettingBroadcastKey)
	if err != nil {
		return nil, err
	}
	if s.Key != SettingBroadcastKey {
		return nil, fmt.Errorf("unexpected setting key received. Requested: %q, received: %q", SettingBroadcastKey, s.Key)
	}
	return f.(*SettingBroadcast), nil
}

// UpdateSettingBroadcast Experimental! This function is not yet stable and may change in the future.
func (c *client) UpdateSettingBroadcast(ctx context.Context, site string, s *SettingBroadcast) (*SettingBroadcast, error) {
	s.Key = SettingBroadcastKey
	result, err := c.SetSetting(ctx, site, SettingBroadcastKey, s)
	if err != nil {
		return nil, err
	}
	return result.(*SettingBroadcast), nil
}
