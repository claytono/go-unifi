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

type HeatMap struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Description      string                     `json:"description,omitempty"`
	MapID            string                     `json:"map_id"`
	Name             string                     `json:"name,omitempty"`                                            // .*[^\s]+.*
	Type             string                     `json:"type,omitempty" validate:"omitempty,oneof=download upload"` // download|upload
	AdditionalFields map[string]json.RawMessage `json:"_additional_properties,omitempty"`
}

func (dst *HeatMap) UnmarshalJSON(b []byte) error {
	type Alias HeatMap
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

func (dst *HeatMap) KnownJSONFields() map[string]struct{} {
	return map[string]struct{}{
		"_id":            {},
		"site_id":        {},
		"attr_hidden":    {},
		"attr_hidden_id": {},
		"attr_no_delete": {},
		"attr_no_edit":   {},
		"description":    {},
		"map_id":         {},
		"name":           {},
		"type":           {},
	}
}

func (dst *HeatMap) SetAdditionalFields(ef map[string]json.RawMessage) { dst.AdditionalFields = ef }

func (c *client) listHeatMap(ctx context.Context, site string) ([]HeatMap, error) {
	var respBody struct {
		Meta Meta      `json:"meta"`
		Data []HeatMap `json:"data"`
	}

	err := c.Get(ctx, fmt.Sprintf("s/%s/rest/heatmap", site), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody.Data, nil
}

func (c *client) getHeatMap(ctx context.Context, site, id string) (*HeatMap, error) {
	path := fmt.Sprintf("s/%s/rest/heatmap/%s", site, id)

	if c.includeAdditionalFields {
		var item HeatMap
		if err := c.getWithAdditionalFieldsV1(ctx, path, &item); err != nil {
			return nil, err
		}
		return &item, nil
	}

	var respBody struct {
		Meta Meta      `json:"meta"`
		Data []HeatMap `json:"data"`
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

func (c *client) deleteHeatMap(ctx context.Context, site, id string) error {
	err := c.Delete(ctx, fmt.Sprintf("s/%s/rest/heatmap/%s", site, id), struct{}{}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) createHeatMap(ctx context.Context, site string, d *HeatMap) (*HeatMap, error) {
	var respBody struct {
		Meta Meta      `json:"meta"`
		Data []HeatMap `json:"data"`
	}

	err := c.Post(ctx, fmt.Sprintf("s/%s/rest/heatmap", site), struct {
		*HeatMap
		AdditionalFields *struct{} `json:"_additional_properties,omitempty"`
	}{HeatMap: d}, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, ErrNotFound
	}

	new := respBody.Data[0]

	return &new, nil
}

func (c *client) updateHeatMap(ctx context.Context, site string, d *HeatMap) (*HeatMap, error) {
	var respBody struct {
		Meta Meta      `json:"meta"`
		Data []HeatMap `json:"data"`
	}

	err := c.Put(ctx, fmt.Sprintf("s/%s/rest/heatmap/%s", site, d.ID), struct {
		*HeatMap
		AdditionalFields *struct{} `json:"_additional_properties,omitempty"`
	}{HeatMap: d}, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, ErrNotFound
	}

	new := respBody.Data[0]

	return &new, nil
}
