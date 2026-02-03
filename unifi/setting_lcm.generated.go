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

const SettingLcmKey = "lcm"

type SettingLcm struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	Brightness  int                        `json:"brightness,omitempty"` // [1-9]|[1-9][0-9]|100
	Enabled     bool                       `json:"enabled"`
	IDleTimeout int                        `json:"idle_timeout,omitempty"` // [1-9][0-9]|[1-9][0-9][0-9]|[1-2][0-9][0-9][0-9]|3[0-5][0-9][0-9]|3600
	Sync        bool                       `json:"sync"`
	TouchEvent  bool                       `json:"touch_event"`
	ExtraFields map[string]json.RawMessage `json:"-"`
}

func (dst *SettingLcm) UnmarshalJSON(b []byte) error {
	type Alias SettingLcm
	aux := &struct {
		Brightness  emptyStringInt `json:"brightness"`
		IDleTimeout emptyStringInt `json:"idle_timeout"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.Brightness = int(aux.Brightness)
	dst.IDleTimeout = int(aux.IDleTimeout)

	// Capture extra fields not in the struct
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(b, &raw); err == nil {
		known := map[string]struct{}{
			"_id":            {},
			"site_id":        {},
			"attr_hidden":    {},
			"attr_hidden_id": {},
			"attr_no_delete": {},
			"attr_no_edit":   {},
			"key":            {},
			"brightness":     {},
			"enabled":        {},
			"idle_timeout":   {},
			"sync":           {},
			"touch_event":    {},
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

func (src SettingLcm) MarshalJSON() ([]byte, error) {
	type Alias SettingLcm
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

// GetSettingLcm Experimental! This function is not yet stable and may change in the future.
func (c *client) GetSettingLcm(ctx context.Context, site string) (*SettingLcm, error) {
	s, f, err := c.GetSetting(ctx, site, SettingLcmKey)
	if err != nil {
		return nil, err
	}
	if s.Key != SettingLcmKey {
		return nil, fmt.Errorf("unexpected setting key received. Requested: %q, received: %q", SettingLcmKey, s.Key)
	}
	return f.(*SettingLcm), nil
}

// UpdateSettingLcm Experimental! This function is not yet stable and may change in the future.
func (c *client) UpdateSettingLcm(ctx context.Context, site string, s *SettingLcm) (*SettingLcm, error) {
	s.Key = SettingLcmKey
	result, err := c.SetSetting(ctx, site, SettingLcmKey, s)
	if err != nil {
		return nil, err
	}
	return result.(*SettingLcm), nil
}
