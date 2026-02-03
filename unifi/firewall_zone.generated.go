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

type FirewallZone struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Name        string                     `json:"name,omitempty"`
	NetworkIDs  []string                   `json:"network_ids"`
	ExtraFields map[string]json.RawMessage `json:"-"`
}

func (dst *FirewallZone) UnmarshalJSON(b []byte) error {
	type Alias FirewallZone
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
			"name":           {},
			"network_ids":    {},
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

func (src FirewallZone) MarshalJSON() ([]byte, error) {
	type Alias FirewallZone
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

func (c *client) listFirewallZone(ctx context.Context, site string) ([]FirewallZone, error) {
	var respBody []FirewallZone

	err := c.Get(ctx, fmt.Sprintf("%s/site/%s/firewall/zone", c.apiPaths.ApiV2Path, site), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func (c *client) getFirewallZone(ctx context.Context, site, id string) (*FirewallZone, error) {
	var respBody FirewallZone

	err := c.Get(ctx, fmt.Sprintf("%s/site/%s/firewall/zone/%s", c.apiPaths.ApiV2Path, site, id), nil, &respBody)

	if err != nil {
		return nil, err
	}
	if respBody.ID == "" {
		return nil, ErrNotFound
	}
	return &respBody, nil
}

func (c *client) deleteFirewallZone(ctx context.Context, site, id string) error {
	err := c.Delete(ctx, fmt.Sprintf("%s/site/%s/firewall/zone/%s", c.apiPaths.ApiV2Path, site, id), struct{}{}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) createFirewallZone(ctx context.Context, site string, d *FirewallZone) (*FirewallZone, error) {
	var respBody FirewallZone

	err := c.Post(ctx, fmt.Sprintf("%s/site/%s/firewall/zone", c.apiPaths.ApiV2Path, site), d, &respBody)
	if err != nil {
		return nil, err
	}

	return &respBody, nil
}

func (c *client) updateFirewallZone(ctx context.Context, site string, d *FirewallZone) (*FirewallZone, error) {
	var respBody FirewallZone

	err := c.Put(ctx, fmt.Sprintf("%s/site/%s/firewall/zone/%s", c.apiPaths.ApiV2Path, site, d.ID), d, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody, nil
}
