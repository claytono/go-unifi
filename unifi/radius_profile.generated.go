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

type RADIUSProfile struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	AccountingEnabled         bool                       `json:"accounting_enabled"`
	AcctServers               []RADIUSProfileAcctServers `json:"acct_servers,omitempty"`
	AuthServers               []RADIUSProfileAuthServers `json:"auth_servers,omitempty"`
	InterimUpdateEnabled      bool                       `json:"interim_update_enabled"`
	InterimUpdateInterval     int                        `json:"interim_update_interval,omitempty"`                 // ^([6-9][0-9]|[1-9][0-9]{2,3}|[1-7][0-9]{4}|8[0-5][0-9]{3}|86[0-3][0-9][0-9]|86400)$
	Name                      string                     `json:"name,omitempty" validate:"omitempty,gte=1,lte=128"` // .{1,128}
	TlsEnabled                bool                       `json:"tls_enabled"`
	UseUsgAcctServer          bool                       `json:"use_usg_acct_server"`
	UseUsgAuthServer          bool                       `json:"use_usg_auth_server"`
	VLANEnabled               bool                       `json:"vlan_enabled"`
	VLANWLANMode              string                     `json:"vlan_wlan_mode,omitempty" validate:"omitempty,oneof=disabled optional required"` // disabled|optional|required
	XCaCrts                   []RADIUSProfileXCaCrts     `json:"x_ca_crts,omitempty"`
	XClientCrt                string                     `json:"x_client_crt,omitempty"`
	XClientCrtFilename        string                     `json:"x_client_crt_filename,omitempty"`
	XClientPrivateKey         string                     `json:"x_client_private_key,omitempty"`
	XClientPrivateKeyFilename string                     `json:"x_client_private_key_filename,omitempty"`
	XClientPrivateKeyPassword string                     `json:"x_client_private_key_password,omitempty"`
	AdditionalFields          map[string]json.RawMessage `json:"_additional_properties,omitempty"`
}

func (dst *RADIUSProfile) UnmarshalJSON(b []byte) error {
	type Alias RADIUSProfile
	aux := &struct {
		InterimUpdateInterval emptyStringInt `json:"interim_update_interval"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.InterimUpdateInterval = int(aux.InterimUpdateInterval)

	return nil
}

func (dst *RADIUSProfile) KnownJSONFields() map[string]struct{} {
	return map[string]struct{}{
		"_id":                           {},
		"site_id":                       {},
		"attr_hidden":                   {},
		"attr_hidden_id":                {},
		"attr_no_delete":                {},
		"attr_no_edit":                  {},
		"accounting_enabled":            {},
		"acct_servers":                  {},
		"auth_servers":                  {},
		"interim_update_enabled":        {},
		"interim_update_interval":       {},
		"name":                          {},
		"tls_enabled":                   {},
		"use_usg_acct_server":           {},
		"use_usg_auth_server":           {},
		"vlan_enabled":                  {},
		"vlan_wlan_mode":                {},
		"x_ca_crts":                     {},
		"x_client_crt":                  {},
		"x_client_crt_filename":         {},
		"x_client_private_key":          {},
		"x_client_private_key_filename": {},
		"x_client_private_key_password": {},
	}
}

func (dst *RADIUSProfile) SetAdditionalFields(ef map[string]json.RawMessage) {
	dst.AdditionalFields = ef
}

type RADIUSProfileAcctServers struct {
	IP               string                     `json:"ip,omitempty" validate:"omitempty,ipv4"` // ^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$
	Port             int                        `json:"port,omitempty"`                         // ^([1-9][0-9]{0,3}|[1-5][0-9]{4}|[6][0-4][0-9]{3}|[6][5][0-4][0-9]{2}|[6][5][5][0-2][0-9]|[6][5][5][3][0-5])$|^$
	XSecret          string                     `json:"x_secret,omitempty"`
	AdditionalFields map[string]json.RawMessage `json:"_additional_properties,omitempty"`
}

func (dst *RADIUSProfileAcctServers) UnmarshalJSON(b []byte) error {
	type Alias RADIUSProfileAcctServers
	aux := &struct {
		Port emptyStringInt `json:"port"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.Port = int(aux.Port)

	return nil
}

func (dst *RADIUSProfileAcctServers) KnownJSONFields() map[string]struct{} {
	return map[string]struct{}{
		"ip":       {},
		"port":     {},
		"x_secret": {},
	}
}

func (dst *RADIUSProfileAcctServers) SetAdditionalFields(ef map[string]json.RawMessage) {
	dst.AdditionalFields = ef
}

type RADIUSProfileAuthServers struct {
	IP               string                     `json:"ip,omitempty" validate:"omitempty,ipv4"` // ^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$
	Port             int                        `json:"port,omitempty"`                         // ^([1-9][0-9]{0,3}|[1-5][0-9]{4}|[6][0-4][0-9]{3}|[6][5][0-4][0-9]{2}|[6][5][5][0-2][0-9]|[6][5][5][3][0-5])$|^$
	XSecret          string                     `json:"x_secret,omitempty"`
	AdditionalFields map[string]json.RawMessage `json:"_additional_properties,omitempty"`
}

func (dst *RADIUSProfileAuthServers) UnmarshalJSON(b []byte) error {
	type Alias RADIUSProfileAuthServers
	aux := &struct {
		Port emptyStringInt `json:"port"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.Port = int(aux.Port)

	return nil
}

func (dst *RADIUSProfileAuthServers) KnownJSONFields() map[string]struct{} {
	return map[string]struct{}{
		"ip":       {},
		"port":     {},
		"x_secret": {},
	}
}

func (dst *RADIUSProfileAuthServers) SetAdditionalFields(ef map[string]json.RawMessage) {
	dst.AdditionalFields = ef
}

type RADIUSProfileXCaCrts struct {
	Filename         string                     `json:"filename,omitempty"`
	XCaCrt           string                     `json:"x_ca_crt,omitempty"`
	AdditionalFields map[string]json.RawMessage `json:"_additional_properties,omitempty"`
}

func (dst *RADIUSProfileXCaCrts) UnmarshalJSON(b []byte) error {
	type Alias RADIUSProfileXCaCrts
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

func (dst *RADIUSProfileXCaCrts) KnownJSONFields() map[string]struct{} {
	return map[string]struct{}{
		"filename": {},
		"x_ca_crt": {},
	}
}

func (dst *RADIUSProfileXCaCrts) SetAdditionalFields(ef map[string]json.RawMessage) {
	dst.AdditionalFields = ef
}

func (c *client) listRADIUSProfile(ctx context.Context, site string) ([]RADIUSProfile, error) {
	var respBody struct {
		Meta Meta            `json:"meta"`
		Data []RADIUSProfile `json:"data"`
	}

	err := c.Get(ctx, fmt.Sprintf("s/%s/rest/radiusprofile", site), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody.Data, nil
}

func (c *client) getRADIUSProfile(ctx context.Context, site, id string) (*RADIUSProfile, error) {
	path := fmt.Sprintf("s/%s/rest/radiusprofile/%s", site, id)

	if c.includeAdditionalFields {
		var item RADIUSProfile
		if err := c.getWithAdditionalFieldsV1(ctx, path, &item); err != nil {
			return nil, err
		}
		return &item, nil
	}

	var respBody struct {
		Meta Meta            `json:"meta"`
		Data []RADIUSProfile `json:"data"`
	}

	err := c.Get(ctx, path, nil, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, ErrNotFound
	}

	d := respBody.Data[0]
	return &d, nil
}

func (c *client) deleteRADIUSProfile(ctx context.Context, site, id string) error {
	err := c.Delete(ctx, fmt.Sprintf("s/%s/rest/radiusprofile/%s", site, id), struct{}{}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) createRADIUSProfile(ctx context.Context, site string, d *RADIUSProfile) (*RADIUSProfile, error) {
	var respBody struct {
		Meta Meta            `json:"meta"`
		Data []RADIUSProfile `json:"data"`
	}

	err := c.Post(ctx, fmt.Sprintf("s/%s/rest/radiusprofile", site), struct {
		*RADIUSProfile
		AdditionalFields *struct{} `json:"_additional_properties,omitempty"`
	}{RADIUSProfile: d}, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, ErrNotFound
	}

	new := respBody.Data[0]

	return &new, nil
}

func (c *client) updateRADIUSProfile(ctx context.Context, site string, d *RADIUSProfile) (*RADIUSProfile, error) {
	var respBody struct {
		Meta Meta            `json:"meta"`
		Data []RADIUSProfile `json:"data"`
	}

	err := c.Put(ctx, fmt.Sprintf("s/%s/rest/radiusprofile/%s", site, d.ID), struct {
		*RADIUSProfile
		AdditionalFields *struct{} `json:"_additional_properties,omitempty"`
	}{RADIUSProfile: d}, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, ErrNotFound
	}

	new := respBody.Data[0]

	return &new, nil
}
