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

type DynamicDNS struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	CustomService string                     `json:"custom_service,omitempty"`                                                                                                                                                                                                                                                                               // ^[^"' ]+$
	HostName      string                     `json:"host_name,omitempty"`                                                                                                                                                                                                                                                                                    // ^[^"' ]+$
	Interface     string                     `json:"interface,omitempty" validate:"omitempty,oneof=wan wan2"`                                                                                                                                                                                                                                                // wan|wan2
	Login         string                     `json:"login,omitempty"`                                                                                                                                                                                                                                                                                        // ^[^"' ]+$
	Options       []string                   `json:"options,omitempty"`                                                                                                                                                                                                                                                                                      // ^[^"' ]+$
	Server        string                     `json:"server"`                                                                                                                                                                                                                                                                                                 // ^[^"' ]+$|^$
	Service       string                     `json:"service,omitempty" validate:"omitempty,oneof=afraid changeip cloudflare cloudxns ddnss dhis dnsexit dnsomatic dnspark dnspod dslreports dtdns duckdns duiadns dyn dyndns dynv6 easydns freemyip googledomains loopia namecheap noip nsupdate ovh sitelutions spdyn strato tunnelbroker zoneedit custom"` // afraid|changeip|cloudflare|cloudxns|ddnss|dhis|dnsexit|dnsomatic|dnspark|dnspod|dslreports|dtdns|duckdns|duiadns|dyn|dyndns|dynv6|easydns|freemyip|googledomains|loopia|namecheap|noip|nsupdate|ovh|sitelutions|spdyn|strato|tunnelbroker|zoneedit|custom
	XPassword     string                     `json:"x_password,omitempty"`                                                                                                                                                                                                                                                                                   // ^[^"' ]+$
	ExtraFields   map[string]json.RawMessage `json:"-"`
}

func (dst *DynamicDNS) UnmarshalJSON(b []byte) error {
	type Alias DynamicDNS
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
			"custom_service": {},
			"host_name":      {},
			"interface":      {},
			"login":          {},
			"options":        {},
			"server":         {},
			"service":        {},
			"x_password":     {},
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

func (src DynamicDNS) MarshalJSON() ([]byte, error) {
	type Alias DynamicDNS
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

func (c *client) listDynamicDNS(ctx context.Context, site string) ([]DynamicDNS, error) {
	var respBody struct {
		Meta Meta         `json:"meta"`
		Data []DynamicDNS `json:"data"`
	}

	err := c.Get(ctx, fmt.Sprintf("s/%s/rest/dynamicdns", site), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody.Data, nil
}

func (c *client) getDynamicDNS(ctx context.Context, site, id string) (*DynamicDNS, error) {
	var respBody struct {
		Meta Meta         `json:"meta"`
		Data []DynamicDNS `json:"data"`
	}

	err := c.Get(ctx, fmt.Sprintf("s/%s/rest/dynamicdns/%s", site, id), nil, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, ErrNotFound
	}

	d := respBody.Data[0]
	return &d, nil
}

func (c *client) deleteDynamicDNS(ctx context.Context, site, id string) error {
	err := c.Delete(ctx, fmt.Sprintf("s/%s/rest/dynamicdns/%s", site, id), struct{}{}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) createDynamicDNS(ctx context.Context, site string, d *DynamicDNS) (*DynamicDNS, error) {
	var respBody struct {
		Meta Meta         `json:"meta"`
		Data []DynamicDNS `json:"data"`
	}

	err := c.Post(ctx, fmt.Sprintf("s/%s/rest/dynamicdns", site), d, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, ErrNotFound
	}

	new := respBody.Data[0]

	return &new, nil
}

func (c *client) updateDynamicDNS(ctx context.Context, site string, d *DynamicDNS) (*DynamicDNS, error) {
	var respBody struct {
		Meta Meta         `json:"meta"`
		Data []DynamicDNS `json:"data"`
	}

	err := c.Put(ctx, fmt.Sprintf("s/%s/rest/dynamicdns/%s", site, d.ID), d, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, ErrNotFound
	}

	new := respBody.Data[0]

	return &new, nil
}
