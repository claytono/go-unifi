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

const SettingSnmpKey = "snmp"

type SettingSnmp struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	Community   string                     `json:"community,omitempty" validate:"omitempty,gte=1,lte=256"` // .{1,256}
	Enabled     bool                       `json:"enabled"`
	EnabledV3   bool                       `json:"enabledV3"`
	Username    string                     `json:"username,omitempty"`   // [a-zA-Z0-9_-]{1,30}
	XPassword   string                     `json:"x_password,omitempty"` // [^'"]{8,32}
	ExtraFields map[string]json.RawMessage `json:"-"`
}

func (dst *SettingSnmp) UnmarshalJSON(b []byte) error {
	type Alias SettingSnmp
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
			"community":      {},
			"enabled":        {},
			"enabledV3":      {},
			"username":       {},
			"x_password":     {},
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

func (src SettingSnmp) MarshalJSON() ([]byte, error) {
	type Alias SettingSnmp
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

// GetSettingSnmp Experimental! This function is not yet stable and may change in the future.
func (c *client) GetSettingSnmp(ctx context.Context, site string) (*SettingSnmp, error) {
	s, f, err := c.GetSetting(ctx, site, SettingSnmpKey)
	if err != nil {
		return nil, err
	}
	if s.Key != SettingSnmpKey {
		return nil, fmt.Errorf("unexpected setting key received. Requested: %q, received: %q", SettingSnmpKey, s.Key)
	}
	return f.(*SettingSnmp), nil
}

// UpdateSettingSnmp Experimental! This function is not yet stable and may change in the future.
func (c *client) UpdateSettingSnmp(ctx context.Context, site string, s *SettingSnmp) (*SettingSnmp, error) {
	s.Key = SettingSnmpKey
	result, err := c.SetSetting(ctx, site, SettingSnmpKey, s)
	if err != nil {
		return nil, err
	}
	return result.(*SettingSnmp), nil
}
