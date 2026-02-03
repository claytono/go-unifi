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

type DHCPOption struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Code        string                     `json:"code,omitempty"` // ^(?!(?:15|42|43|44|51|66|67|252)$)([7-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-4])$
	Name        string                     `json:"name,omitempty"` // ^[A-Za-z0-9-_]{1,25}$
	Signed      bool                       `json:"signed"`
	Type        string                     `json:"type,omitempty" validate:"omitempty,oneof=boolean hexarray integer ipaddress macaddress text"` // ^(boolean|hexarray|integer|ipaddress|macaddress|text)$
	Width       int                        `json:"width,omitempty" validate:"omitempty,oneof=8 16 32"`                                           // ^(8|16|32)$
	ExtraFields map[string]json.RawMessage `json:"-"`
}

func (dst *DHCPOption) UnmarshalJSON(b []byte) error {
	type Alias DHCPOption
	aux := &struct {
		Width emptyStringInt `json:"width"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.Width = int(aux.Width)

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
			"code":           {},
			"name":           {},
			"signed":         {},
			"type":           {},
			"width":          {},
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

func (src DHCPOption) MarshalJSON() ([]byte, error) {
	type Alias DHCPOption
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

func (c *client) listDHCPOption(ctx context.Context, site string) ([]DHCPOption, error) {
	var respBody struct {
		Meta Meta         `json:"meta"`
		Data []DHCPOption `json:"data"`
	}

	err := c.Get(ctx, fmt.Sprintf("s/%s/rest/dhcpoption", site), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody.Data, nil
}

func (c *client) getDHCPOption(ctx context.Context, site, id string) (*DHCPOption, error) {
	var respBody struct {
		Meta Meta         `json:"meta"`
		Data []DHCPOption `json:"data"`
	}

	err := c.Get(ctx, fmt.Sprintf("s/%s/rest/dhcpoption/%s", site, id), nil, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, ErrNotFound
	}

	d := respBody.Data[0]
	return &d, nil
}

func (c *client) deleteDHCPOption(ctx context.Context, site, id string) error {
	err := c.Delete(ctx, fmt.Sprintf("s/%s/rest/dhcpoption/%s", site, id), struct{}{}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) createDHCPOption(ctx context.Context, site string, d *DHCPOption) (*DHCPOption, error) {
	var respBody struct {
		Meta Meta         `json:"meta"`
		Data []DHCPOption `json:"data"`
	}

	err := c.Post(ctx, fmt.Sprintf("s/%s/rest/dhcpoption", site), d, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, ErrNotFound
	}

	new := respBody.Data[0]

	return &new, nil
}

func (c *client) updateDHCPOption(ctx context.Context, site string, d *DHCPOption) (*DHCPOption, error) {
	var respBody struct {
		Meta Meta         `json:"meta"`
		Data []DHCPOption `json:"data"`
	}

	err := c.Put(ctx, fmt.Sprintf("s/%s/rest/dhcpoption/%s", site, d.ID), d, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, ErrNotFound
	}

	new := respBody.Data[0]

	return &new, nil
}
