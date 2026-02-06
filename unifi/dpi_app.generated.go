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

type DpiApp struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Apps             []int                      `json:"apps,omitempty"`
	Blocked          bool                       `json:"blocked"`
	Cats             []int                      `json:"cats,omitempty"`
	Enabled          bool                       `json:"enabled"`
	Log              bool                       `json:"log"`
	Name             string                     `json:"name,omitempty" validate:"omitempty,gte=1,lte=128"` // .{1,128}
	QOSRateMaxDown   int                        `json:"qos_rate_max_down,omitempty"`                       // -1|[2-9]|[1-9][0-9]{1,4}|100000|10[0-1][0-9]{3}|102[0-3][0-9]{2}|102400
	QOSRateMaxUp     int                        `json:"qos_rate_max_up,omitempty"`                         // -1|[2-9]|[1-9][0-9]{1,4}|100000|10[0-1][0-9]{3}|102[0-3][0-9]{2}|102400
	AdditionalFields map[string]json.RawMessage `json:"_additional_properties,omitempty"`
}

func (dst *DpiApp) UnmarshalJSON(b []byte) error {
	type Alias DpiApp
	aux := &struct {
		Apps           []emptyStringInt `json:"apps"`
		Cats           []emptyStringInt `json:"cats"`
		QOSRateMaxDown emptyStringInt   `json:"qos_rate_max_down"`
		QOSRateMaxUp   emptyStringInt   `json:"qos_rate_max_up"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.Apps = make([]int, len(aux.Apps))
	for i, v := range aux.Apps {
		dst.Apps[i] = int(v)
	}
	dst.Cats = make([]int, len(aux.Cats))
	for i, v := range aux.Cats {
		dst.Cats[i] = int(v)
	}
	dst.QOSRateMaxDown = int(aux.QOSRateMaxDown)
	dst.QOSRateMaxUp = int(aux.QOSRateMaxUp)

	return nil
}

func (dst *DpiApp) KnownJSONFields() map[string]struct{} {
	return map[string]struct{}{
		"_id":               {},
		"site_id":           {},
		"attr_hidden":       {},
		"attr_hidden_id":    {},
		"attr_no_delete":    {},
		"attr_no_edit":      {},
		"apps":              {},
		"blocked":           {},
		"cats":              {},
		"enabled":           {},
		"log":               {},
		"name":              {},
		"qos_rate_max_down": {},
		"qos_rate_max_up":   {},
	}
}

func (dst *DpiApp) SetAdditionalFields(ef map[string]json.RawMessage) { dst.AdditionalFields = ef }

func (c *client) listDpiApp(ctx context.Context, site string) ([]DpiApp, error) {
	var respBody struct {
		Meta Meta     `json:"meta"`
		Data []DpiApp `json:"data"`
	}

	err := c.Get(ctx, fmt.Sprintf("s/%s/rest/dpiapp", site), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody.Data, nil
}

func (c *client) getDpiApp(ctx context.Context, site, id string) (*DpiApp, error) {
	path := fmt.Sprintf("s/%s/rest/dpiapp/%s", site, id)

	if c.includeAdditionalFields {
		var item DpiApp
		if err := c.getWithAdditionalFieldsV1(ctx, path, &item); err != nil {
			return nil, err
		}
		return &item, nil
	}

	var respBody struct {
		Meta Meta     `json:"meta"`
		Data []DpiApp `json:"data"`
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

func (c *client) deleteDpiApp(ctx context.Context, site, id string) error {
	err := c.Delete(ctx, fmt.Sprintf("s/%s/rest/dpiapp/%s", site, id), struct{}{}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) createDpiApp(ctx context.Context, site string, d *DpiApp) (*DpiApp, error) {
	var respBody struct {
		Meta Meta     `json:"meta"`
		Data []DpiApp `json:"data"`
	}

	err := c.Post(ctx, fmt.Sprintf("s/%s/rest/dpiapp", site), struct {
		*DpiApp
		AdditionalFields *struct{} `json:"_additional_properties,omitempty"`
	}{DpiApp: d}, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, ErrNotFound
	}

	new := respBody.Data[0]

	return &new, nil
}

func (c *client) updateDpiApp(ctx context.Context, site string, d *DpiApp) (*DpiApp, error) {
	var respBody struct {
		Meta Meta     `json:"meta"`
		Data []DpiApp `json:"data"`
	}

	err := c.Put(ctx, fmt.Sprintf("s/%s/rest/dpiapp/%s", site, d.ID), struct {
		*DpiApp
		AdditionalFields *struct{} `json:"_additional_properties,omitempty"`
	}{DpiApp: d}, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, ErrNotFound
	}

	new := respBody.Data[0]

	return &new, nil
}
