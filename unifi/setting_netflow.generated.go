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

const SettingNetflowKey = "netflow"

type SettingNetflow struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	AutoEngineIDEnabled bool                       `json:"auto_engine_id_enabled"`
	Enabled             bool                       `json:"enabled"`
	EngineID            int                        `json:"engine_id,omitempty"` // ^$|[1-9][0-9]*
	ExportFrequency     int                        `json:"export_frequency,omitempty"`
	NetworkIDs          []string                   `json:"network_ids,omitempty"`
	Port                int                        `json:"port,omitempty"` // 102[4-9]|10[3-9][0-9]|1[1-9][0-9]{2}|[2-9][0-9]{3}|[1-5][0-9]{4}|[6][0-4][0-9]{3}|[6][5][0-4][0-9]{2}|[6][5][5][0-2][0-9]|[6][5][5][3][0-5]
	RefreshRate         int                        `json:"refresh_rate,omitempty"`
	SamplingMode        string                     `json:"sampling_mode,omitempty" validate:"omitempty,oneof=off hash random deterministic"` // off|hash|random|deterministic
	SamplingRate        int                        `json:"sampling_rate,omitempty"`                                                          // [2-9]|[1-9][0-9]{1,3}|1[0-5][0-9]{3}|16[0-2][0-9]{2}|163[0-7][0-9]|1638[0-3]|^$
	Server              string                     `json:"server,omitempty"`                                                                 // .{0,252}[^\.]$
	Version             int                        `json:"version,omitempty" validate:"omitempty,oneof=5 9 10"`                              // 5|9|10
	ExtraFields         map[string]json.RawMessage `json:"-"`
}

func (dst *SettingNetflow) UnmarshalJSON(b []byte) error {
	type Alias SettingNetflow
	aux := &struct {
		EngineID        emptyStringInt `json:"engine_id"`
		ExportFrequency emptyStringInt `json:"export_frequency"`
		Port            emptyStringInt `json:"port"`
		RefreshRate     emptyStringInt `json:"refresh_rate"`
		SamplingRate    emptyStringInt `json:"sampling_rate"`
		Version         emptyStringInt `json:"version"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.EngineID = int(aux.EngineID)
	dst.ExportFrequency = int(aux.ExportFrequency)
	dst.Port = int(aux.Port)
	dst.RefreshRate = int(aux.RefreshRate)
	dst.SamplingRate = int(aux.SamplingRate)
	dst.Version = int(aux.Version)

	// Capture extra fields not in the struct
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(b, &raw); err == nil {
		known := map[string]struct{}{
			"_id":                    {},
			"site_id":                {},
			"attr_hidden":            {},
			"attr_hidden_id":         {},
			"attr_no_delete":         {},
			"attr_no_edit":           {},
			"key":                    {},
			"auto_engine_id_enabled": {},
			"enabled":                {},
			"engine_id":              {},
			"export_frequency":       {},
			"network_ids":            {},
			"port":                   {},
			"refresh_rate":           {},
			"sampling_mode":          {},
			"sampling_rate":          {},
			"server":                 {},
			"version":                {},
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

func (src SettingNetflow) MarshalJSON() ([]byte, error) {
	type Alias SettingNetflow
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

// GetSettingNetflow Experimental! This function is not yet stable and may change in the future.
func (c *client) GetSettingNetflow(ctx context.Context, site string) (*SettingNetflow, error) {
	s, f, err := c.GetSetting(ctx, site, SettingNetflowKey)
	if err != nil {
		return nil, err
	}
	if s.Key != SettingNetflowKey {
		return nil, fmt.Errorf("unexpected setting key received. Requested: %q, received: %q", SettingNetflowKey, s.Key)
	}
	return f.(*SettingNetflow), nil
}

// UpdateSettingNetflow Experimental! This function is not yet stable and may change in the future.
func (c *client) UpdateSettingNetflow(ctx context.Context, site string, s *SettingNetflow) (*SettingNetflow, error) {
	s.Key = SettingNetflowKey
	result, err := c.SetSetting(ctx, site, SettingNetflowKey, s)
	if err != nil {
		return nil, err
	}
	return result.(*SettingNetflow), nil
}
