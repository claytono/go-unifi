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

type DNSRecord struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Enabled     bool                       `json:"enabled"`
	Key         string                     `json:"key,omitempty" validate:"omitempty,gte=1,lte=256"`                                    // .{1,256}
	Port        int                        `json:"port,omitempty"`                                                                      // ^[0-9][0-9]?$|^
	Priority    int                        `json:"priority,omitempty"`                                                                  // ^[0-9][0-9]?$|^
	RecordType  string                     `json:"record_type,omitempty" validate:"omitempty,oneof=A AAAA CNAME MX NS PTR SOA SRV TXT"` // A|AAAA|CNAME|MX|NS|PTR|SOA|SRV|TXT
	Ttl         int                        `json:"ttl,omitempty"`                                                                       // ^[0-9][0-9]?$|^
	Value       string                     `json:"value,omitempty" validate:"omitempty,gte=1,lte=256"`                                  // .{1,256}
	Weight      int                        `json:"weight,omitempty"`                                                                    // ^[0-9][0-9]?$|^
	ExtraFields map[string]json.RawMessage `json:"-"`
}

func (dst *DNSRecord) UnmarshalJSON(b []byte) error {
	type Alias DNSRecord
	aux := &struct {
		Port     emptyStringInt `json:"port"`
		Priority emptyStringInt `json:"priority"`
		Ttl      emptyStringInt `json:"ttl"`
		Weight   emptyStringInt `json:"weight"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.Port = int(aux.Port)
	dst.Priority = int(aux.Priority)
	dst.Ttl = int(aux.Ttl)
	dst.Weight = int(aux.Weight)

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
			"enabled":        {},
			"key":            {},
			"port":           {},
			"priority":       {},
			"record_type":    {},
			"ttl":            {},
			"value":          {},
			"weight":         {},
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

func (src DNSRecord) MarshalJSON() ([]byte, error) {
	type Alias DNSRecord
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

func (c *client) listDNSRecord(ctx context.Context, site string) ([]DNSRecord, error) {
	var respBody []DNSRecord

	err := c.Get(ctx, fmt.Sprintf("%s/site/%s/static-dns", c.apiPaths.ApiV2Path, site), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func (c *client) getDNSRecord(ctx context.Context, site, id string) (*DNSRecord, error) {
	var respBody DNSRecord

	err := c.Get(ctx, fmt.Sprintf("%s/site/%s/static-dns/%s", c.apiPaths.ApiV2Path, site, id), nil, &respBody)

	if err != nil {
		return nil, err
	}
	if respBody.ID == "" {
		return nil, ErrNotFound
	}
	return &respBody, nil
}

func (c *client) deleteDNSRecord(ctx context.Context, site, id string) error {
	err := c.Delete(ctx, fmt.Sprintf("%s/site/%s/static-dns/%s", c.apiPaths.ApiV2Path, site, id), struct{}{}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) createDNSRecord(ctx context.Context, site string, d *DNSRecord) (*DNSRecord, error) {
	var respBody DNSRecord

	err := c.Post(ctx, fmt.Sprintf("%s/site/%s/static-dns", c.apiPaths.ApiV2Path, site), d, &respBody)
	if err != nil {
		return nil, err
	}

	return &respBody, nil
}

func (c *client) updateDNSRecord(ctx context.Context, site string, d *DNSRecord) (*DNSRecord, error) {
	var respBody DNSRecord

	err := c.Put(ctx, fmt.Sprintf("%s/site/%s/static-dns/%s", c.apiPaths.ApiV2Path, site, d.ID), d, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody, nil
}
