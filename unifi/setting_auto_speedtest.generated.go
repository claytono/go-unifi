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

	CronExpr         string                     `json:"cron_expr,omitempty"`
	Enabled          bool                       `json:"enabled"`
	AdditionalFields map[string]json.RawMessage `json:"_additional_properties,omitempty"`
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

	return nil
}

func (dst *SettingAutoSpeedtest) KnownJSONFields() map[string]struct{} {
	return map[string]struct{}{
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
}

func (dst *SettingAutoSpeedtest) SetAdditionalFields(ef map[string]json.RawMessage) {
	dst.AdditionalFields = ef
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
	result, err := c.SetSetting(ctx, site, SettingAutoSpeedtestKey, struct {
		*SettingAutoSpeedtest
		AdditionalFields *struct{} `json:"_additional_properties,omitempty"`
	}{SettingAutoSpeedtest: s})
	if err != nil {
		return nil, err
	}
	return result.(*SettingAutoSpeedtest), nil
}
