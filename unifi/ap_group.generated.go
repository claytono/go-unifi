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

type APGroup struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	DeviceMACs       []string                   `json:"device_macs,omitempty"`
	Name             string                     `json:"name,omitempty"`
	AdditionalFields map[string]json.RawMessage `json:"_additional_properties,omitempty"`
}

func (dst *APGroup) UnmarshalJSON(b []byte) error {
	type Alias APGroup
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

func (dst *APGroup) KnownJSONFields() map[string]struct{} {
	return map[string]struct{}{
		"_id":            {},
		"site_id":        {},
		"attr_hidden":    {},
		"attr_hidden_id": {},
		"attr_no_delete": {},
		"attr_no_edit":   {},
		"device_macs":    {},
		"name":           {},
	}
}

func (dst *APGroup) SetAdditionalFields(ef map[string]json.RawMessage) { dst.AdditionalFields = ef }

func (c *client) listAPGroup(ctx context.Context, site string) ([]APGroup, error) {
	var respBody []APGroup

	err := c.Get(ctx, fmt.Sprintf("%s/site/%s/apgroups", c.apiPaths.ApiV2Path, site), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func (c *client) getAPGroup(ctx context.Context, site, id string) (*APGroup, error) {
	path := fmt.Sprintf("%s/site/%s/apgroups/%s", c.apiPaths.ApiV2Path, site, id)

	if c.includeAdditionalFields {
		var item APGroup
		if err := c.getWithAdditionalFieldsV2(ctx, path, &item); err != nil {
			return nil, err
		}
		if item.ID == "" {
			return nil, ErrNotFound
		}
		return &item, nil
	}

	var respBody APGroup

	err := c.Get(ctx, path, nil, &respBody)

	if err != nil {
		return nil, err
	}
	if respBody.ID == "" {
		return nil, ErrNotFound
	}
	return &respBody, nil
}

func (c *client) deleteAPGroup(ctx context.Context, site, id string) error {
	err := c.Delete(ctx, fmt.Sprintf("%s/site/%s/apgroups/%s", c.apiPaths.ApiV2Path, site, id), struct{}{}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) createAPGroup(ctx context.Context, site string, d *APGroup) (*APGroup, error) {
	var respBody APGroup

	err := c.Post(ctx, fmt.Sprintf("%s/site/%s/apgroups", c.apiPaths.ApiV2Path, site), struct {
		*APGroup
		AdditionalFields *struct{} `json:"_additional_properties,omitempty"`
	}{APGroup: d}, &respBody)
	if err != nil {
		return nil, err
	}

	return &respBody, nil
}

func (c *client) updateAPGroup(ctx context.Context, site string, d *APGroup) (*APGroup, error) {
	var respBody APGroup

	err := c.Put(ctx, fmt.Sprintf("%s/site/%s/apgroups/%s", c.apiPaths.ApiV2Path, site, d.ID), struct {
		*APGroup
		AdditionalFields *struct{} `json:"_additional_properties,omitempty"`
	}{APGroup: d}, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody, nil
}
