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

const SettingRadiusKey = "radius"

type SettingRadius struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	AccountingEnabled     bool                       `json:"accounting_enabled"`
	AcctPort              int                        `json:"acct_port,omitempty"` // [1-9][0-9]{0,3}|[1-5][0-9]{4}|[6][0-4][0-9]{3}|[6][5][0-4][0-9]{2}|[6][5][5][0-2][0-9]|[6][5][5][3][0-5]
	AuthPort              int                        `json:"auth_port,omitempty"` // [1-9][0-9]{0,3}|[1-5][0-9]{4}|[6][0-4][0-9]{3}|[6][5][0-4][0-9]{2}|[6][5][5][0-2][0-9]|[6][5][5][3][0-5]
	ConfigureWholeNetwork bool                       `json:"configure_whole_network"`
	Enabled               bool                       `json:"enabled"`
	InterimUpdateInterval int                        `json:"interim_update_interval,omitempty"` // ^([6-9][0-9]|[1-9][0-9]{2,3}|[1-7][0-9]{4}|8[0-5][0-9]{3}|86[0-3][0-9][0-9]|86400)$
	TunneledReply         bool                       `json:"tunneled_reply"`
	XSecret               string                     `json:"x_secret,omitempty"` // ^[^\\"' ]{1,48}$
	ExtraFields           map[string]json.RawMessage `json:"-"`
}

func (dst *SettingRadius) UnmarshalJSON(b []byte) error {
	type Alias SettingRadius
	aux := &struct {
		AcctPort              emptyStringInt `json:"acct_port"`
		AuthPort              emptyStringInt `json:"auth_port"`
		InterimUpdateInterval emptyStringInt `json:"interim_update_interval"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.AcctPort = int(aux.AcctPort)
	dst.AuthPort = int(aux.AuthPort)
	dst.InterimUpdateInterval = int(aux.InterimUpdateInterval)

	// Capture extra fields not in the struct
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(b, &raw); err == nil {
		known := map[string]struct{}{
			"_id":                     {},
			"site_id":                 {},
			"attr_hidden":             {},
			"attr_hidden_id":          {},
			"attr_no_delete":          {},
			"attr_no_edit":            {},
			"key":                     {},
			"accounting_enabled":      {},
			"acct_port":               {},
			"auth_port":               {},
			"configure_whole_network": {},
			"enabled":                 {},
			"interim_update_interval": {},
			"tunneled_reply":          {},
			"x_secret":                {},
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

func (src SettingRadius) MarshalJSON() ([]byte, error) {
	type Alias SettingRadius
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

// GetSettingRadius Experimental! This function is not yet stable and may change in the future.
func (c *client) GetSettingRadius(ctx context.Context, site string) (*SettingRadius, error) {
	s, f, err := c.GetSetting(ctx, site, SettingRadiusKey)
	if err != nil {
		return nil, err
	}
	if s.Key != SettingRadiusKey {
		return nil, fmt.Errorf("unexpected setting key received. Requested: %q, received: %q", SettingRadiusKey, s.Key)
	}
	return f.(*SettingRadius), nil
}

// UpdateSettingRadius Experimental! This function is not yet stable and may change in the future.
func (c *client) UpdateSettingRadius(ctx context.Context, site string, s *SettingRadius) (*SettingRadius, error) {
	s.Key = SettingRadiusKey
	result, err := c.SetSetting(ctx, site, SettingRadiusKey, s)
	if err != nil {
		return nil, err
	}
	return result.(*SettingRadius), nil
}
