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

type FirewallGroup struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	GroupMembers []string                   `json:"group_members,omitempty"`
	GroupType    string                     `json:"group_type,omitempty" validate:"omitempty,oneof=address-group port-group ipv6-address-group"` // address-group|port-group|ipv6-address-group
	Name         string                     `json:"name,omitempty" validate:"omitempty,gte=1,lte=64"`                                            // .{1,64}
	ExtraFields  map[string]json.RawMessage `json:"-"`
}

func (dst *FirewallGroup) UnmarshalJSON(b []byte) error {
	type Alias FirewallGroup
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
			"group_members":  {},
			"group_type":     {},
			"name":           {},
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

func (src FirewallGroup) MarshalJSON() ([]byte, error) {
	type Alias FirewallGroup
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

func (c *client) listFirewallGroup(ctx context.Context, site string) ([]FirewallGroup, error) {
	var respBody struct {
		Meta Meta            `json:"meta"`
		Data []FirewallGroup `json:"data"`
	}

	err := c.Get(ctx, fmt.Sprintf("s/%s/rest/firewallgroup", site), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody.Data, nil
}

func (c *client) getFirewallGroup(ctx context.Context, site, id string) (*FirewallGroup, error) {
	var respBody struct {
		Meta Meta            `json:"meta"`
		Data []FirewallGroup `json:"data"`
	}

	err := c.Get(ctx, fmt.Sprintf("s/%s/rest/firewallgroup/%s", site, id), nil, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, ErrNotFound
	}

	d := respBody.Data[0]
	return &d, nil
}

func (c *client) deleteFirewallGroup(ctx context.Context, site, id string) error {
	err := c.Delete(ctx, fmt.Sprintf("s/%s/rest/firewallgroup/%s", site, id), struct{}{}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) createFirewallGroup(ctx context.Context, site string, d *FirewallGroup) (*FirewallGroup, error) {
	var respBody struct {
		Meta Meta            `json:"meta"`
		Data []FirewallGroup `json:"data"`
	}

	err := c.Post(ctx, fmt.Sprintf("s/%s/rest/firewallgroup", site), d, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, ErrNotFound
	}

	new := respBody.Data[0]

	return &new, nil
}

func (c *client) updateFirewallGroup(ctx context.Context, site string, d *FirewallGroup) (*FirewallGroup, error) {
	var respBody struct {
		Meta Meta            `json:"meta"`
		Data []FirewallGroup `json:"data"`
	}

	err := c.Put(ctx, fmt.Sprintf("s/%s/rest/firewallgroup/%s", site, d.ID), d, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, ErrNotFound
	}

	new := respBody.Data[0]

	return &new, nil
}
