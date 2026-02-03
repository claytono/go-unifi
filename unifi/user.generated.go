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

type User struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	DevIdOverride int    `json:"dev_id_override,omitempty"`            // non-generated field
	IP            string `json:"ip,omitempty" validate:"omitempty,ip"` // non-generated field

	Blocked                       bool                       `json:"blocked,omitempty"`
	FixedApEnabled                bool                       `json:"fixed_ap_enabled"`
	FixedApMAC                    string                     `json:"fixed_ap_mac,omitempty" validate:"omitempty,mac"` // ^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$
	FixedIP                       string                     `json:"fixed_ip,omitempty"`
	Hostname                      string                     `json:"hostname,omitempty"`
	LastSeen                      int                        `json:"last_seen,omitempty"`
	LocalDNSRecord                string                     `json:"local_dns_record,omitempty"`
	LocalDNSRecordEnabled         bool                       `json:"local_dns_record_enabled"`
	MAC                           string                     `json:"mac,omitempty" validate:"omitempty,mac"` // ^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$
	Name                          string                     `json:"name,omitempty"`
	NetworkID                     string                     `json:"network_id"`
	Note                          string                     `json:"note,omitempty"`
	UseFixedIP                    bool                       `json:"use_fixedip"`
	UserGroupID                   string                     `json:"usergroup_id"`
	VirtualNetworkOverrideEnabled bool                       `json:"virtual_network_override_enabled"`
	VirtualNetworkOverrideID      string                     `json:"virtual_network_override_id"`
	ExtraFields                   map[string]json.RawMessage `json:"-"`
}

func (dst *User) UnmarshalJSON(b []byte) error {
	type Alias User
	aux := &struct {
		LastSeen emptyStringInt `json:"last_seen"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.LastSeen = int(aux.LastSeen)

	// Capture extra fields not in the struct
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(b, &raw); err == nil {
		known := map[string]struct{}{
			"_id":                              {},
			"site_id":                          {},
			"attr_hidden":                      {},
			"attr_hidden_id":                   {},
			"attr_no_delete":                   {},
			"attr_no_edit":                     {},
			"dev_id_override":                  {},
			"ip":                               {},
			"blocked":                          {},
			"fixed_ap_enabled":                 {},
			"fixed_ap_mac":                     {},
			"fixed_ip":                         {},
			"hostname":                         {},
			"last_seen":                        {},
			"local_dns_record":                 {},
			"local_dns_record_enabled":         {},
			"mac":                              {},
			"name":                             {},
			"network_id":                       {},
			"note":                             {},
			"use_fixedip":                      {},
			"usergroup_id":                     {},
			"virtual_network_override_enabled": {},
			"virtual_network_override_id":      {},
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

func (src User) MarshalJSON() ([]byte, error) {
	type Alias User
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

func (c *client) listUser(ctx context.Context, site string) ([]User, error) {
	var respBody struct {
		Meta Meta   `json:"meta"`
		Data []User `json:"data"`
	}

	err := c.Get(ctx, fmt.Sprintf("s/%s/rest/user", site), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody.Data, nil
}

func (c *client) getUser(ctx context.Context, site, id string) (*User, error) {
	var respBody struct {
		Meta Meta   `json:"meta"`
		Data []User `json:"data"`
	}

	err := c.Get(ctx, fmt.Sprintf("s/%s/rest/user/%s", site, id), nil, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, ErrNotFound
	}

	d := respBody.Data[0]
	return &d, nil
}

func (c *client) deleteUser(ctx context.Context, site, id string) error {
	err := c.Delete(ctx, fmt.Sprintf("s/%s/rest/user/%s", site, id), struct{}{}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) createUser(ctx context.Context, site string, d *User) (*User, error) {
	var respBody struct {
		Meta Meta   `json:"meta"`
		Data []User `json:"data"`
	}

	err := c.Post(ctx, fmt.Sprintf("s/%s/rest/user", site), d, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, ErrNotFound
	}

	new := respBody.Data[0]

	return &new, nil
}

func (c *client) updateUser(ctx context.Context, site string, d *User) (*User, error) {
	var respBody struct {
		Meta Meta   `json:"meta"`
		Data []User `json:"data"`
	}

	err := c.Put(ctx, fmt.Sprintf("s/%s/rest/user/%s", site, d.ID), d, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, ErrNotFound
	}

	new := respBody.Data[0]

	return &new, nil
}
