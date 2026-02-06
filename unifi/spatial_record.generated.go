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

type SpatialRecord struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Devices          []SpatialRecordDevices     `json:"devices,omitempty"`
	Name             string                     `json:"name,omitempty" validate:"omitempty,gte=1,lte=128"` // .{1,128}
	AdditionalFields map[string]json.RawMessage `json:"_additional_properties,omitempty"`
}

func (dst *SpatialRecord) UnmarshalJSON(b []byte) error {
	type Alias SpatialRecord
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

func (dst *SpatialRecord) KnownJSONFields() map[string]struct{} {
	return map[string]struct{}{
		"_id":            {},
		"site_id":        {},
		"attr_hidden":    {},
		"attr_hidden_id": {},
		"attr_no_delete": {},
		"attr_no_edit":   {},
		"devices":        {},
		"name":           {},
	}
}

func (dst *SpatialRecord) SetAdditionalFields(ef map[string]json.RawMessage) {
	dst.AdditionalFields = ef
}

type SpatialRecordDevices struct {
	MAC              string                     `json:"mac,omitempty" validate:"omitempty,mac"` // ^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$
	Position         SpatialRecordPosition      `json:"position,omitempty"`
	AdditionalFields map[string]json.RawMessage `json:"_additional_properties,omitempty"`
}

func (dst *SpatialRecordDevices) UnmarshalJSON(b []byte) error {
	type Alias SpatialRecordDevices
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

func (dst *SpatialRecordDevices) KnownJSONFields() map[string]struct{} {
	return map[string]struct{}{
		"mac":      {},
		"position": {},
	}
}

func (dst *SpatialRecordDevices) SetAdditionalFields(ef map[string]json.RawMessage) {
	dst.AdditionalFields = ef
}

type SpatialRecordPosition struct {
	X                float64                    `json:"x,omitempty"` // (^([-]?[\d]+)$)|(^([-]?[\d]+[.]?[\d]+)$)
	Y                float64                    `json:"y,omitempty"` // (^([-]?[\d]+)$)|(^([-]?[\d]+[.]?[\d]+)$)
	Z                float64                    `json:"z,omitempty"` // (^([-]?[\d]+)$)|(^([-]?[\d]+[.]?[\d]+)$)
	AdditionalFields map[string]json.RawMessage `json:"_additional_properties,omitempty"`
}

func (dst *SpatialRecordPosition) UnmarshalJSON(b []byte) error {
	type Alias SpatialRecordPosition
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

func (dst *SpatialRecordPosition) KnownJSONFields() map[string]struct{} {
	return map[string]struct{}{
		"x": {},
		"y": {},
		"z": {},
	}
}

func (dst *SpatialRecordPosition) SetAdditionalFields(ef map[string]json.RawMessage) {
	dst.AdditionalFields = ef
}

func (c *client) listSpatialRecord(ctx context.Context, site string) ([]SpatialRecord, error) {
	var respBody struct {
		Meta Meta            `json:"meta"`
		Data []SpatialRecord `json:"data"`
	}

	err := c.Get(ctx, fmt.Sprintf("s/%s/rest/spatialrecord", site), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody.Data, nil
}

func (c *client) getSpatialRecord(ctx context.Context, site, id string) (*SpatialRecord, error) {
	path := fmt.Sprintf("s/%s/rest/spatialrecord/%s", site, id)

	if c.includeAdditionalFields {
		var item SpatialRecord
		if err := c.getWithAdditionalFieldsV1(ctx, path, &item); err != nil {
			return nil, err
		}
		return &item, nil
	}

	var respBody struct {
		Meta Meta            `json:"meta"`
		Data []SpatialRecord `json:"data"`
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

func (c *client) deleteSpatialRecord(ctx context.Context, site, id string) error {
	err := c.Delete(ctx, fmt.Sprintf("s/%s/rest/spatialrecord/%s", site, id), struct{}{}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) createSpatialRecord(ctx context.Context, site string, d *SpatialRecord) (*SpatialRecord, error) {
	var respBody struct {
		Meta Meta            `json:"meta"`
		Data []SpatialRecord `json:"data"`
	}

	err := c.Post(ctx, fmt.Sprintf("s/%s/rest/spatialrecord", site), struct {
		*SpatialRecord
		AdditionalFields *struct{} `json:"_additional_properties,omitempty"`
	}{SpatialRecord: d}, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, ErrNotFound
	}

	new := respBody.Data[0]

	return &new, nil
}

func (c *client) updateSpatialRecord(ctx context.Context, site string, d *SpatialRecord) (*SpatialRecord, error) {
	var respBody struct {
		Meta Meta            `json:"meta"`
		Data []SpatialRecord `json:"data"`
	}

	err := c.Put(ctx, fmt.Sprintf("s/%s/rest/spatialrecord/%s", site, d.ID), struct {
		*SpatialRecord
		AdditionalFields *struct{} `json:"_additional_properties,omitempty"`
	}{SpatialRecord: d}, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, ErrNotFound
	}

	new := respBody.Data[0]

	return &new, nil
}
