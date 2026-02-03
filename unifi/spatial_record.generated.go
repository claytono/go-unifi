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

	Devices     []SpatialRecordDevices     `json:"devices,omitempty"`
	Name        string                     `json:"name,omitempty" validate:"omitempty,gte=1,lte=128"` // .{1,128}
	ExtraFields map[string]json.RawMessage `json:"-"`
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
			"devices":        {},
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

func (src SpatialRecord) MarshalJSON() ([]byte, error) {
	type Alias SpatialRecord
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

type SpatialRecordDevices struct {
	MAC         string                     `json:"mac,omitempty" validate:"omitempty,mac"` // ^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$
	Position    SpatialRecordPosition      `json:"position,omitempty"`
	ExtraFields map[string]json.RawMessage `json:"-"`
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

	// Capture extra fields not in the struct
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(b, &raw); err == nil {
		known := map[string]struct{}{
			"mac":      {},
			"position": {},
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

func (src SpatialRecordDevices) MarshalJSON() ([]byte, error) {
	type Alias SpatialRecordDevices
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

type SpatialRecordPosition struct {
	X           float64                    `json:"x,omitempty"` // (^([-]?[\d]+)$)|(^([-]?[\d]+[.]?[\d]+)$)
	Y           float64                    `json:"y,omitempty"` // (^([-]?[\d]+)$)|(^([-]?[\d]+[.]?[\d]+)$)
	Z           float64                    `json:"z,omitempty"` // (^([-]?[\d]+)$)|(^([-]?[\d]+[.]?[\d]+)$)
	ExtraFields map[string]json.RawMessage `json:"-"`
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

	// Capture extra fields not in the struct
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(b, &raw); err == nil {
		known := map[string]struct{}{
			"x": {},
			"y": {},
			"z": {},
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

func (src SpatialRecordPosition) MarshalJSON() ([]byte, error) {
	type Alias SpatialRecordPosition
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
	var respBody struct {
		Meta Meta            `json:"meta"`
		Data []SpatialRecord `json:"data"`
	}

	err := c.Get(ctx, fmt.Sprintf("s/%s/rest/spatialrecord/%s", site, id), nil, &respBody)
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

	err := c.Post(ctx, fmt.Sprintf("s/%s/rest/spatialrecord", site), d, &respBody)
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

	err := c.Put(ctx, fmt.Sprintf("s/%s/rest/spatialrecord/%s", site, d.ID), d, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, ErrNotFound
	}

	new := respBody.Data[0]

	return &new, nil
}
