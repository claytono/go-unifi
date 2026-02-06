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

type BroadcastGroup struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	MemberTable      []string                   `json:"member_table,omitempty"`
	Name             string                     `json:"name,omitempty"`
	AdditionalFields map[string]json.RawMessage `json:"_additional_properties,omitempty"`
}

func (dst *BroadcastGroup) UnmarshalJSON(b []byte) error {
	type Alias BroadcastGroup
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

func (dst *BroadcastGroup) KnownJSONFields() map[string]struct{} {
	return map[string]struct{}{
		"_id":            {},
		"site_id":        {},
		"attr_hidden":    {},
		"attr_hidden_id": {},
		"attr_no_delete": {},
		"attr_no_edit":   {},
		"member_table":   {},
		"name":           {},
	}
}

func (dst *BroadcastGroup) SetAdditionalFields(ef map[string]json.RawMessage) {
	dst.AdditionalFields = ef
}

func (c *client) listBroadcastGroup(ctx context.Context, site string) ([]BroadcastGroup, error) {
	var respBody struct {
		Meta Meta             `json:"meta"`
		Data []BroadcastGroup `json:"data"`
	}

	err := c.Get(ctx, fmt.Sprintf("s/%s/rest/broadcastgroup", site), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody.Data, nil
}

func (c *client) getBroadcastGroup(ctx context.Context, site, id string) (*BroadcastGroup, error) {
	path := fmt.Sprintf("s/%s/rest/broadcastgroup/%s", site, id)

	if c.includeAdditionalFields {
		var item BroadcastGroup
		if err := c.getWithAdditionalFieldsV1(ctx, path, &item); err != nil {
			return nil, err
		}
		return &item, nil
	}

	var respBody struct {
		Meta Meta             `json:"meta"`
		Data []BroadcastGroup `json:"data"`
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

func (c *client) deleteBroadcastGroup(ctx context.Context, site, id string) error {
	err := c.Delete(ctx, fmt.Sprintf("s/%s/rest/broadcastgroup/%s", site, id), struct{}{}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) createBroadcastGroup(ctx context.Context, site string, d *BroadcastGroup) (*BroadcastGroup, error) {
	var respBody struct {
		Meta Meta             `json:"meta"`
		Data []BroadcastGroup `json:"data"`
	}

	err := c.Post(ctx, fmt.Sprintf("s/%s/rest/broadcastgroup", site), struct {
		*BroadcastGroup
		AdditionalFields *struct{} `json:"_additional_properties,omitempty"`
	}{BroadcastGroup: d}, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, ErrNotFound
	}

	new := respBody.Data[0]

	return &new, nil
}

func (c *client) updateBroadcastGroup(ctx context.Context, site string, d *BroadcastGroup) (*BroadcastGroup, error) {
	var respBody struct {
		Meta Meta             `json:"meta"`
		Data []BroadcastGroup `json:"data"`
	}

	err := c.Put(ctx, fmt.Sprintf("s/%s/rest/broadcastgroup/%s", site, d.ID), struct {
		*BroadcastGroup
		AdditionalFields *struct{} `json:"_additional_properties,omitempty"`
	}{BroadcastGroup: d}, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, ErrNotFound
	}

	new := respBody.Data[0]

	return &new, nil
}
