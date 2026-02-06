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

const SettingNtpKey = "ntp"

type SettingNtp struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	NtpServer1        string                     `json:"ntp_server_1,omitempty"`
	NtpServer2        string                     `json:"ntp_server_2,omitempty"`
	NtpServer3        string                     `json:"ntp_server_3,omitempty"`
	NtpServer4        string                     `json:"ntp_server_4,omitempty"`
	SettingPreference string                     `json:"setting_preference,omitempty" validate:"omitempty,oneof=auto manual"` // auto|manual
	AdditionalFields  map[string]json.RawMessage `json:"_additional_properties,omitempty"`
}

func (dst *SettingNtp) UnmarshalJSON(b []byte) error {
	type Alias SettingNtp
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

func (dst *SettingNtp) KnownJSONFields() map[string]struct{} {
	return map[string]struct{}{
		"_id":                {},
		"site_id":            {},
		"attr_hidden":        {},
		"attr_hidden_id":     {},
		"attr_no_delete":     {},
		"attr_no_edit":       {},
		"key":                {},
		"ntp_server_1":       {},
		"ntp_server_2":       {},
		"ntp_server_3":       {},
		"ntp_server_4":       {},
		"setting_preference": {},
	}
}

func (dst *SettingNtp) SetAdditionalFields(ef map[string]json.RawMessage) { dst.AdditionalFields = ef }

// GetSettingNtp Experimental! This function is not yet stable and may change in the future.
func (c *client) GetSettingNtp(ctx context.Context, site string) (*SettingNtp, error) {
	s, f, err := c.GetSetting(ctx, site, SettingNtpKey)
	if err != nil {
		return nil, err
	}
	if s.Key != SettingNtpKey {
		return nil, fmt.Errorf("unexpected setting key received. Requested: %q, received: %q", SettingNtpKey, s.Key)
	}
	return f.(*SettingNtp), nil
}

// UpdateSettingNtp Experimental! This function is not yet stable and may change in the future.
func (c *client) UpdateSettingNtp(ctx context.Context, site string, s *SettingNtp) (*SettingNtp, error) {
	s.Key = SettingNtpKey
	result, err := c.SetSetting(ctx, site, SettingNtpKey, struct {
		*SettingNtp
		AdditionalFields *struct{} `json:"_additional_properties,omitempty"`
	}{SettingNtp: s})
	if err != nil {
		return nil, err
	}
	return result.(*SettingNtp), nil
}
