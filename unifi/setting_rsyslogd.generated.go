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

const SettingRsyslogdKey = "rsyslogd"

type SettingRsyslogd struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	Contents                    []string                   `json:"contents" validate:"omitempty,oneof=device client firewall_default_policy triggers updates admin_activity critical security_detections vpn"` // device|client|firewall_default_policy|triggers|updates|admin_activity|critical|security_detections|vpn
	Debug                       bool                       `json:"debug"`
	Enabled                     bool                       `json:"enabled"`
	IP                          string                     `json:"ip"`
	LogAllContents              bool                       `json:"log_all_contents"`
	NetconsoleEnabled           bool                       `json:"netconsole_enabled"`
	NetconsoleHost              string                     `json:"netconsole_host"`
	NetconsolePort              int                        `json:"netconsole_port,omitempty"` // [1-9][0-9]{0,3}|[1-5][0-9]{4}|[6][0-4][0-9]{3}|[6][5][0-4][0-9]{2}|[6][5][5][0-2][0-9]|[6][5][5][3][0-5]
	Port                        int                        `json:"port,omitempty"`            // [1-9][0-9]{0,3}|[1-5][0-9]{4}|[6][0-4][0-9]{3}|[6][5][0-4][0-9]{2}|[6][5][5][0-2][0-9]|[6][5][5][3][0-5]
	ThisController              bool                       `json:"this_controller"`
	ThisControllerEncryptedOnly bool                       `json:"this_controller_encrypted_only"`
	ExtraFields                 map[string]json.RawMessage `json:"-"`
}

func (dst *SettingRsyslogd) UnmarshalJSON(b []byte) error {
	type Alias SettingRsyslogd
	aux := &struct {
		NetconsolePort emptyStringInt `json:"netconsole_port"`
		Port           emptyStringInt `json:"port"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.NetconsolePort = int(aux.NetconsolePort)
	dst.Port = int(aux.Port)

	// Capture extra fields not in the struct
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(b, &raw); err == nil {
		known := map[string]struct{}{
			"_id":                            {},
			"site_id":                        {},
			"attr_hidden":                    {},
			"attr_hidden_id":                 {},
			"attr_no_delete":                 {},
			"attr_no_edit":                   {},
			"key":                            {},
			"contents":                       {},
			"debug":                          {},
			"enabled":                        {},
			"ip":                             {},
			"log_all_contents":               {},
			"netconsole_enabled":             {},
			"netconsole_host":                {},
			"netconsole_port":                {},
			"port":                           {},
			"this_controller":                {},
			"this_controller_encrypted_only": {},
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

func (src SettingRsyslogd) MarshalJSON() ([]byte, error) {
	type Alias SettingRsyslogd
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

// GetSettingRsyslogd Experimental! This function is not yet stable and may change in the future.
func (c *client) GetSettingRsyslogd(ctx context.Context, site string) (*SettingRsyslogd, error) {
	s, f, err := c.GetSetting(ctx, site, SettingRsyslogdKey)
	if err != nil {
		return nil, err
	}
	if s.Key != SettingRsyslogdKey {
		return nil, fmt.Errorf("unexpected setting key received. Requested: %q, received: %q", SettingRsyslogdKey, s.Key)
	}
	return f.(*SettingRsyslogd), nil
}

// UpdateSettingRsyslogd Experimental! This function is not yet stable and may change in the future.
func (c *client) UpdateSettingRsyslogd(ctx context.Context, site string, s *SettingRsyslogd) (*SettingRsyslogd, error) {
	s.Key = SettingRsyslogdKey
	result, err := c.SetSetting(ctx, site, SettingRsyslogdKey, s)
	if err != nil {
		return nil, err
	}
	return result.(*SettingRsyslogd), nil
}
