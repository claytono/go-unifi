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

type WLANGroup struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Name             string                     `json:"name,omitempty" validate:"omitempty,gte=1,lte=128"` // .{1,128}
	AdditionalFields map[string]json.RawMessage `json:"_additional_properties,omitempty"`
}

func (dst *WLANGroup) UnmarshalJSON(b []byte) error {
	type Alias WLANGroup
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

func (dst *WLANGroup) KnownJSONFields() map[string]struct{} {
	return map[string]struct{}{
		"_id":            {},
		"site_id":        {},
		"attr_hidden":    {},
		"attr_hidden_id": {},
		"attr_no_delete": {},
		"attr_no_edit":   {},
		"name":           {},
	}
}

func (dst *WLANGroup) SetAdditionalFields(ef map[string]json.RawMessage) { dst.AdditionalFields = ef }

func (c *client) listWLANGroup(ctx context.Context, site string) ([]WLANGroup, error) {
	var respBody struct {
		Meta Meta        `json:"meta"`
		Data []WLANGroup `json:"data"`
	}

	err := c.Get(ctx, fmt.Sprintf("s/%s/rest/wlangroup", site), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody.Data, nil
}

func (c *client) getWLANGroup(ctx context.Context, site, id string) (*WLANGroup, error) {
	path := fmt.Sprintf("s/%s/rest/wlangroup/%s", site, id)

	if c.includeAdditionalFields {
		var item WLANGroup
		if err := c.getWithAdditionalFieldsV1(ctx, path, &item); err != nil {
			return nil, err
		}
		return &item, nil
	}

	var respBody struct {
		Meta Meta        `json:"meta"`
		Data []WLANGroup `json:"data"`
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

func (c *client) deleteWLANGroup(ctx context.Context, site, id string) error {
	err := c.Delete(ctx, fmt.Sprintf("s/%s/rest/wlangroup/%s", site, id), struct{}{}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) createWLANGroup(ctx context.Context, site string, d *WLANGroup) (*WLANGroup, error) {
	var respBody struct {
		Meta Meta        `json:"meta"`
		Data []WLANGroup `json:"data"`
	}

	err := c.Post(ctx, fmt.Sprintf("s/%s/rest/wlangroup", site), struct {
		*WLANGroup
		AdditionalFields *struct{} `json:"_additional_properties,omitempty"`
	}{WLANGroup: d}, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, ErrNotFound
	}

	new := respBody.Data[0]

	return &new, nil
}

func (c *client) updateWLANGroup(ctx context.Context, site string, d *WLANGroup) (*WLANGroup, error) {
	var respBody struct {
		Meta Meta        `json:"meta"`
		Data []WLANGroup `json:"data"`
	}

	err := c.Put(ctx, fmt.Sprintf("s/%s/rest/wlangroup/%s", site, d.ID), struct {
		*WLANGroup
		AdditionalFields *struct{} `json:"_additional_properties,omitempty"`
	}{WLANGroup: d}, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, ErrNotFound
	}

	new := respBody.Data[0]

	return &new, nil
}
