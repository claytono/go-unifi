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

const SettingSuperCloudaccessKey = "super_cloudaccess"

type SettingSuperCloudaccess struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	DeviceAuth       string                     `json:"device_auth,omitempty"`
	DeviceID         string                     `json:"device_id"`
	Enabled          bool                       `json:"enabled"`
	UbicUuid         string                     `json:"ubic_uuid,omitempty"`
	XCertificateArn  string                     `json:"x_certificate_arn,omitempty"`
	XCertificatePem  string                     `json:"x_certificate_pem,omitempty"`
	XPrivateKey      string                     `json:"x_private_key,omitempty"`
	AdditionalFields map[string]json.RawMessage `json:"_additional_properties,omitempty"`
}

func (dst *SettingSuperCloudaccess) UnmarshalJSON(b []byte) error {
	type Alias SettingSuperCloudaccess
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

func (dst *SettingSuperCloudaccess) KnownJSONFields() map[string]struct{} {
	return map[string]struct{}{
		"_id":               {},
		"site_id":           {},
		"attr_hidden":       {},
		"attr_hidden_id":    {},
		"attr_no_delete":    {},
		"attr_no_edit":      {},
		"key":               {},
		"device_auth":       {},
		"device_id":         {},
		"enabled":           {},
		"ubic_uuid":         {},
		"x_certificate_arn": {},
		"x_certificate_pem": {},
		"x_private_key":     {},
	}
}

func (dst *SettingSuperCloudaccess) SetAdditionalFields(ef map[string]json.RawMessage) {
	dst.AdditionalFields = ef
}

// GetSettingSuperCloudaccess Experimental! This function is not yet stable and may change in the future.
func (c *client) GetSettingSuperCloudaccess(ctx context.Context, site string) (*SettingSuperCloudaccess, error) {
	s, f, err := c.GetSetting(ctx, site, SettingSuperCloudaccessKey)
	if err != nil {
		return nil, err
	}
	if s.Key != SettingSuperCloudaccessKey {
		return nil, fmt.Errorf("unexpected setting key received. Requested: %q, received: %q", SettingSuperCloudaccessKey, s.Key)
	}
	return f.(*SettingSuperCloudaccess), nil
}

// UpdateSettingSuperCloudaccess Experimental! This function is not yet stable and may change in the future.
func (c *client) UpdateSettingSuperCloudaccess(ctx context.Context, site string, s *SettingSuperCloudaccess) (*SettingSuperCloudaccess, error) {
	s.Key = SettingSuperCloudaccessKey
	result, err := c.SetSetting(ctx, site, SettingSuperCloudaccessKey, struct {
		*SettingSuperCloudaccess
		AdditionalFields *struct{} `json:"_additional_properties,omitempty"`
	}{SettingSuperCloudaccess: s})
	if err != nil {
		return nil, err
	}
	return result.(*SettingSuperCloudaccess), nil
}
