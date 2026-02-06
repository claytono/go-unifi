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

type Dashboard struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	ControllerVersion string                     `json:"controller_version,omitempty"`
	Desc              string                     `json:"desc,omitempty"`
	IsPublic          bool                       `json:"is_public"`
	Modules           []DashboardModules         `json:"modules,omitempty"`
	Name              string                     `json:"name,omitempty"`
	AdditionalFields  map[string]json.RawMessage `json:"_additional_properties,omitempty"`
}

func (dst *Dashboard) UnmarshalJSON(b []byte) error {
	type Alias Dashboard
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

func (dst *Dashboard) KnownJSONFields() map[string]struct{} {
	return map[string]struct{}{
		"_id":                {},
		"site_id":            {},
		"attr_hidden":        {},
		"attr_hidden_id":     {},
		"attr_no_delete":     {},
		"attr_no_edit":       {},
		"controller_version": {},
		"desc":               {},
		"is_public":          {},
		"modules":            {},
		"name":               {},
	}
}

func (dst *Dashboard) SetAdditionalFields(ef map[string]json.RawMessage) { dst.AdditionalFields = ef }

type DashboardModules struct {
	Config           string                     `json:"config,omitempty"`
	ID               string                     `json:"id"`
	ModuleID         string                     `json:"module_id"`
	Restrictions     string                     `json:"restrictions,omitempty"`
	AdditionalFields map[string]json.RawMessage `json:"_additional_properties,omitempty"`
}

func (dst *DashboardModules) UnmarshalJSON(b []byte) error {
	type Alias DashboardModules
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

func (dst *DashboardModules) KnownJSONFields() map[string]struct{} {
	return map[string]struct{}{
		"config":       {},
		"id":           {},
		"module_id":    {},
		"restrictions": {},
	}
}

func (dst *DashboardModules) SetAdditionalFields(ef map[string]json.RawMessage) {
	dst.AdditionalFields = ef
}

func (c *client) listDashboard(ctx context.Context, site string) ([]Dashboard, error) {
	var respBody struct {
		Meta Meta        `json:"meta"`
		Data []Dashboard `json:"data"`
	}

	err := c.Get(ctx, fmt.Sprintf("s/%s/rest/dashboard", site), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody.Data, nil
}

func (c *client) getDashboard(ctx context.Context, site, id string) (*Dashboard, error) {
	path := fmt.Sprintf("s/%s/rest/dashboard/%s", site, id)

	if c.includeAdditionalFields {
		var item Dashboard
		if err := c.getWithAdditionalFieldsV1(ctx, path, &item); err != nil {
			return nil, err
		}
		return &item, nil
	}

	var respBody struct {
		Meta Meta        `json:"meta"`
		Data []Dashboard `json:"data"`
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

func (c *client) deleteDashboard(ctx context.Context, site, id string) error {
	err := c.Delete(ctx, fmt.Sprintf("s/%s/rest/dashboard/%s", site, id), struct{}{}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) createDashboard(ctx context.Context, site string, d *Dashboard) (*Dashboard, error) {
	var respBody struct {
		Meta Meta        `json:"meta"`
		Data []Dashboard `json:"data"`
	}

	err := c.Post(ctx, fmt.Sprintf("s/%s/rest/dashboard", site), struct {
		*Dashboard
		AdditionalFields *struct{} `json:"_additional_properties,omitempty"`
	}{Dashboard: d}, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, ErrNotFound
	}

	new := respBody.Data[0]

	return &new, nil
}

func (c *client) updateDashboard(ctx context.Context, site string, d *Dashboard) (*Dashboard, error) {
	var respBody struct {
		Meta Meta        `json:"meta"`
		Data []Dashboard `json:"data"`
	}

	err := c.Put(ctx, fmt.Sprintf("s/%s/rest/dashboard/%s", site, d.ID), struct {
		*Dashboard
		AdditionalFields *struct{} `json:"_additional_properties,omitempty"`
	}{Dashboard: d}, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, ErrNotFound
	}

	new := respBody.Data[0]

	return &new, nil
}
