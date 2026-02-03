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

const SettingAutoSpeedtestKey = "auto_speedtest"

type SettingAutoSpeedtest struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	CronExpr    string                     `json:"cron_expr,omitempty"`
	Enabled     bool                       `json:"enabled"`
	ExtraFields map[string]json.RawMessage `json:"-"`
}

func (dst *SettingAutoSpeedtest) UnmarshalJSON(b []byte) error {
	type Alias SettingAutoSpeedtest
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
			"_id":            {},
			"site_id":        {},
			"attr_hidden":    {},
			"attr_hidden_id": {},
			"attr_no_delete": {},
			"attr_no_edit":   {},
			"key":            {},
			"cron_expr":      {},
			"enabled":        {},
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

func (src SettingAutoSpeedtest) MarshalJSON() ([]byte, error) {
	type Alias SettingAutoSpeedtest
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

// GetSettingAutoSpeedtest Experimental! This function is not yet stable and may change in the future.
func (c *client) GetSettingAutoSpeedtest(ctx context.Context, site string) (*SettingAutoSpeedtest, error) {
	s, f, err := c.GetSetting(ctx, site, SettingAutoSpeedtestKey)
	if err != nil {
		return nil, err
	}
	if s.Key != SettingAutoSpeedtestKey {
		return nil, fmt.Errorf("unexpected setting key received. Requested: %q, received: %q", SettingAutoSpeedtestKey, s.Key)
	}
	return f.(*SettingAutoSpeedtest), nil
}

// UpdateSettingAutoSpeedtest Experimental! This function is not yet stable and may change in the future.
func (c *client) UpdateSettingAutoSpeedtest(ctx context.Context, site string, s *SettingAutoSpeedtest) (*SettingAutoSpeedtest, error) {
	s.Key = SettingAutoSpeedtestKey
	result, err := c.SetSetting(ctx, site, SettingAutoSpeedtestKey, s)
	if err != nil {
		return nil, err
	}
	return result.(*SettingAutoSpeedtest), nil
}
