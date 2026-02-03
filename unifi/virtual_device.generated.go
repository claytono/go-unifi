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

type VirtualDevice struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	HeightInMeters float64                    `json:"heightInMeters,omitempty"`
	Locked         bool                       `json:"locked"`
	MapID          string                     `json:"map_id"`
	Type           string                     `json:"type,omitempty" validate:"omitempty,oneof=uap usg usw"` // uap|usg|usw
	X              string                     `json:"x,omitempty"`
	Y              string                     `json:"y,omitempty"`
	ExtraFields    map[string]json.RawMessage `json:"-"`
}

func (dst *VirtualDevice) UnmarshalJSON(b []byte) error {
	type Alias VirtualDevice
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
			"heightInMeters": {},
			"locked":         {},
			"map_id":         {},
			"type":           {},
			"x":              {},
			"y":              {},
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

func (src VirtualDevice) MarshalJSON() ([]byte, error) {
	type Alias VirtualDevice
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

func (c *client) listVirtualDevice(ctx context.Context, site string) ([]VirtualDevice, error) {
	var respBody struct {
		Meta Meta            `json:"meta"`
		Data []VirtualDevice `json:"data"`
	}

	err := c.Get(ctx, fmt.Sprintf("s/%s/rest/virtualdevice", site), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody.Data, nil
}

func (c *client) getVirtualDevice(ctx context.Context, site, id string) (*VirtualDevice, error) {
	var respBody struct {
		Meta Meta            `json:"meta"`
		Data []VirtualDevice `json:"data"`
	}

	err := c.Get(ctx, fmt.Sprintf("s/%s/rest/virtualdevice/%s", site, id), nil, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, ErrNotFound
	}

	d := respBody.Data[0]
	return &d, nil
}

func (c *client) deleteVirtualDevice(ctx context.Context, site, id string) error {
	err := c.Delete(ctx, fmt.Sprintf("s/%s/rest/virtualdevice/%s", site, id), struct{}{}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) createVirtualDevice(ctx context.Context, site string, d *VirtualDevice) (*VirtualDevice, error) {
	var respBody struct {
		Meta Meta            `json:"meta"`
		Data []VirtualDevice `json:"data"`
	}

	err := c.Post(ctx, fmt.Sprintf("s/%s/rest/virtualdevice", site), d, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, ErrNotFound
	}

	new := respBody.Data[0]

	return &new, nil
}

func (c *client) updateVirtualDevice(ctx context.Context, site string, d *VirtualDevice) (*VirtualDevice, error) {
	var respBody struct {
		Meta Meta            `json:"meta"`
		Data []VirtualDevice `json:"data"`
	}

	err := c.Put(ctx, fmt.Sprintf("s/%s/rest/virtualdevice/%s", site, d.ID), d, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, ErrNotFound
	}

	new := respBody.Data[0]

	return &new, nil
}
