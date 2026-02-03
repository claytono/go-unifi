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

const SettingTeleportKey = "teleport"

type SettingTeleport struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	Enabled     bool                       `json:"enabled"`
	SubnetCidr  string                     `json:"subnet_cidr"` // ^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\/([8-9]|[1-2][0-9]|3[0-2])$|^$
	ExtraFields map[string]json.RawMessage `json:"-"`
}

func (dst *SettingTeleport) UnmarshalJSON(b []byte) error {
	type Alias SettingTeleport
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
			"enabled":        {},
			"subnet_cidr":    {},
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

func (src SettingTeleport) MarshalJSON() ([]byte, error) {
	type Alias SettingTeleport
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

// GetSettingTeleport Experimental! This function is not yet stable and may change in the future.
func (c *client) GetSettingTeleport(ctx context.Context, site string) (*SettingTeleport, error) {
	s, f, err := c.GetSetting(ctx, site, SettingTeleportKey)
	if err != nil {
		return nil, err
	}
	if s.Key != SettingTeleportKey {
		return nil, fmt.Errorf("unexpected setting key received. Requested: %q, received: %q", SettingTeleportKey, s.Key)
	}
	return f.(*SettingTeleport), nil
}

// UpdateSettingTeleport Experimental! This function is not yet stable and may change in the future.
func (c *client) UpdateSettingTeleport(ctx context.Context, site string, s *SettingTeleport) (*SettingTeleport, error) {
	s.Key = SettingTeleportKey
	result, err := c.SetSetting(ctx, site, SettingTeleportKey, s)
	if err != nil {
		return nil, err
	}
	return result.(*SettingTeleport), nil
}
