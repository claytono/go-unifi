package unifi

import (
	"context"
	"encoding/json"
	"fmt"
)

func (dst *Network) MarshalJSON() ([]byte, error) {
	type Alias Network
	aux := &struct {
		*Alias

		WANEgressQOS *emptyStringInt `json:"wan_egress_qos,omitempty"`
	}{
		Alias: (*Alias)(dst),
	}

	if dst.Purpose == "wan" {
		// only send QOS when this is a WAN network
		v := emptyStringInt(dst.WANEgressQOS)
		aux.WANEgressQOS = &v
	}

	b, err := json.Marshal(aux)
	if err != nil {
		return nil, err
	}
	if len(dst.ExtraFields) == 0 {
		return b, nil
	}
	var m map[string]json.RawMessage
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	extra, err := json.Marshal(dst.ExtraFields)
	if err != nil {
		return nil, err
	}
	m["_additional_properties"] = extra
	return json.Marshal(m)
}

//func (c *client) DeleteNetwork(ctx context.Context, site, id, name string) error {
//	err := c.Delete(ctx, fmt.Sprintf("s/%s/rest/networkconf/%s", site, id), struct {
//		Name string `json:"name"`
//	}{
//		Name: name,
//	}, nil)
//	if err != nil {
//		return err
//	}
//	return nil
//}

func (c *client) DeleteNetwork(ctx context.Context, site, id string) error {
	err := c.Delete(ctx, fmt.Sprintf("s/%s/rest/networkconf/%s", site, id), nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) ListNetwork(ctx context.Context, site string) ([]Network, error) {
	return c.listNetwork(ctx, site)
}

func (c *client) GetNetwork(ctx context.Context, site, id string) (*Network, error) {
	return c.getNetwork(ctx, site, id)
}

func (c *client) CreateNetwork(ctx context.Context, site string, d *Network) (*Network, error) {
	return c.createNetwork(ctx, site, d)
}

func (c *client) UpdateNetwork(ctx context.Context, site string, d *Network) (*Network, error) {
	return c.updateNetwork(ctx, site, d)
}
