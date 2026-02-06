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

type PortForward struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	DestinationIP      string                      `json:"destination_ip,omitempty"` // ^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$|^any$
	DestinationIPs     []PortForwardDestinationIPs `json:"destination_ips,omitempty"`
	DstPort            string                      `json:"dst_port,omitempty"` // (([1-9][0-9]{0,3}|[1-5][0-9]{4}|[6][0-4][0-9]{3}|[6][5][0-4][0-9]{2}|[6][5][5][0-2][0-9]|[6][5][5][3][0-5])|([1-9][0-9]{0,3}|[1-5][0-9]{4}|[6][0-4][0-9]{3}|[6][5][0-4][0-9]{2}|[6][5][5][0-2][0-9]|[6][5][5][3][0-5])-([1-9][0-9]{0,3}|[1-5][0-9]{4}|[6][0-4][0-9]{3}|[6][5][0-4][0-9]{2}|[6][5][5][0-2][0-9]|[6][5][5][3][0-5]))+(,([1-9][0-9]{0,3}|[1-5][0-9]{4}|[6][0-4][0-9]{3}|[6][5][0-4][0-9]{2}|[6][5][5][0-2][0-9]|[6][5][5][3][0-5])|,([1-9][0-9]{0,3}|[1-5][0-9]{4}|[6][0-4][0-9]{3}|[6][5][0-4][0-9]{2}|[6][5][5][0-2][0-9]|[6][5][5][3][0-5])-([1-9][0-9]{0,3}|[1-5][0-9]{4}|[6][0-4][0-9]{3}|[6][5][0-4][0-9]{2}|[6][5][5][0-2][0-9]|[6][5][5][3][0-5])){0,14}
	Enabled            bool                        `json:"enabled"`
	Fwd                string                      `json:"fwd,omitempty" validate:"omitempty,ipv4"` // ^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$
	FwdPort            string                      `json:"fwd_port,omitempty"`                      // (([1-9][0-9]{0,3}|[1-5][0-9]{4}|[6][0-4][0-9]{3}|[6][5][0-4][0-9]{2}|[6][5][5][0-2][0-9]|[6][5][5][3][0-5])|([1-9][0-9]{0,3}|[1-5][0-9]{4}|[6][0-4][0-9]{3}|[6][5][0-4][0-9]{2}|[6][5][5][0-2][0-9]|[6][5][5][3][0-5])-([1-9][0-9]{0,3}|[1-5][0-9]{4}|[6][0-4][0-9]{3}|[6][5][0-4][0-9]{2}|[6][5][5][0-2][0-9]|[6][5][5][3][0-5]))+(,([1-9][0-9]{0,3}|[1-5][0-9]{4}|[6][0-4][0-9]{3}|[6][5][0-4][0-9]{2}|[6][5][5][0-2][0-9]|[6][5][5][3][0-5])|,([1-9][0-9]{0,3}|[1-5][0-9]{4}|[6][0-4][0-9]{3}|[6][5][0-4][0-9]{2}|[6][5][5][0-2][0-9]|[6][5][5][3][0-5])-([1-9][0-9]{0,3}|[1-5][0-9]{4}|[6][0-4][0-9]{3}|[6][5][0-4][0-9]{2}|[6][5][5][0-2][0-9]|[6][5][5][3][0-5])){0,14}
	Log                bool                        `json:"log"`
	Name               string                      `json:"name,omitempty" validate:"omitempty,gte=1,lte=128"`          // .{1,128}
	PfwdInterface      string                      `json:"pfwd_interface,omitempty"`                                   // wan[2-8]?|both|all
	Proto              string                      `json:"proto,omitempty" validate:"omitempty,oneof=tcp_udp tcp udp"` // tcp_udp|tcp|udp
	Src                string                      `json:"src,omitempty"`                                              // ^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$|^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])-(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$|^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])/([0-9]|[1-2][0-9]|3[0-2])$|^!(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$|^!(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])-(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$|^!(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])/([0-9]|[1-2][0-9]|3[0-2])$|^any$
	SrcFirewallGroupID string                      `json:"src_firewall_group_id"`
	SrcLimitingEnabled bool                        `json:"src_limiting_enabled"`
	SrcLimitingType    string                      `json:"src_limiting_type,omitempty" validate:"omitempty,oneof=ip firewall_group"` // ip|firewall_group
	AdditionalFields   map[string]json.RawMessage  `json:"_additional_properties,omitempty"`
}

func (dst *PortForward) UnmarshalJSON(b []byte) error {
	type Alias PortForward
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

func (dst *PortForward) KnownJSONFields() map[string]struct{} {
	return map[string]struct{}{
		"_id":                   {},
		"site_id":               {},
		"attr_hidden":           {},
		"attr_hidden_id":        {},
		"attr_no_delete":        {},
		"attr_no_edit":          {},
		"destination_ip":        {},
		"destination_ips":       {},
		"dst_port":              {},
		"enabled":               {},
		"fwd":                   {},
		"fwd_port":              {},
		"log":                   {},
		"name":                  {},
		"pfwd_interface":        {},
		"proto":                 {},
		"src":                   {},
		"src_firewall_group_id": {},
		"src_limiting_enabled":  {},
		"src_limiting_type":     {},
	}
}

func (dst *PortForward) SetAdditionalFields(ef map[string]json.RawMessage) { dst.AdditionalFields = ef }

type PortForwardDestinationIPs struct {
	DestinationIP    string                     `json:"destination_ip,omitempty"` // ^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$|^any$
	Interface        string                     `json:"interface,omitempty"`      // wan[2-8]?
	AdditionalFields map[string]json.RawMessage `json:"_additional_properties,omitempty"`
}

func (dst *PortForwardDestinationIPs) UnmarshalJSON(b []byte) error {
	type Alias PortForwardDestinationIPs
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

func (dst *PortForwardDestinationIPs) KnownJSONFields() map[string]struct{} {
	return map[string]struct{}{
		"destination_ip": {},
		"interface":      {},
	}
}

func (dst *PortForwardDestinationIPs) SetAdditionalFields(ef map[string]json.RawMessage) {
	dst.AdditionalFields = ef
}

func (c *client) listPortForward(ctx context.Context, site string) ([]PortForward, error) {
	var respBody struct {
		Meta Meta          `json:"meta"`
		Data []PortForward `json:"data"`
	}

	err := c.Get(ctx, fmt.Sprintf("s/%s/rest/portforward", site), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody.Data, nil
}

func (c *client) getPortForward(ctx context.Context, site, id string) (*PortForward, error) {
	path := fmt.Sprintf("s/%s/rest/portforward/%s", site, id)

	if c.includeAdditionalFields {
		var item PortForward
		if err := c.getWithAdditionalFieldsV1(ctx, path, &item); err != nil {
			return nil, err
		}
		return &item, nil
	}

	var respBody struct {
		Meta Meta          `json:"meta"`
		Data []PortForward `json:"data"`
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

func (c *client) deletePortForward(ctx context.Context, site, id string) error {
	err := c.Delete(ctx, fmt.Sprintf("s/%s/rest/portforward/%s", site, id), struct{}{}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) createPortForward(ctx context.Context, site string, d *PortForward) (*PortForward, error) {
	var respBody struct {
		Meta Meta          `json:"meta"`
		Data []PortForward `json:"data"`
	}

	err := c.Post(ctx, fmt.Sprintf("s/%s/rest/portforward", site), struct {
		*PortForward
		AdditionalFields *struct{} `json:"_additional_properties,omitempty"`
	}{PortForward: d}, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, ErrNotFound
	}

	new := respBody.Data[0]

	return &new, nil
}

func (c *client) updatePortForward(ctx context.Context, site string, d *PortForward) (*PortForward, error) {
	var respBody struct {
		Meta Meta          `json:"meta"`
		Data []PortForward `json:"data"`
	}

	err := c.Put(ctx, fmt.Sprintf("s/%s/rest/portforward/%s", site, d.ID), struct {
		*PortForward
		AdditionalFields *struct{} `json:"_additional_properties,omitempty"`
	}{PortForward: d}, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, ErrNotFound
	}

	new := respBody.Data[0]

	return &new, nil
}
