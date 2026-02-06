package unifi

import (
	"context"
	"encoding/json"
)

// AdditionalFieldsCapable is implemented by generated types that can capture additional JSON fields.
type AdditionalFieldsCapable interface {
	KnownJSONFields() map[string]struct{}
	SetAdditionalFields(map[string]json.RawMessage)
}

type rawV1Response struct {
	Meta Meta              `json:"meta"`
	Data []json.RawMessage `json:"data"`
}

func (c *client) getWithAdditionalFieldsV1(ctx context.Context, path string, dst AdditionalFieldsCapable) error {
	var respBody rawV1Response
	if err := c.Get(ctx, path, nil, &respBody); err != nil {
		return err
	}
	if len(respBody.Data) != 1 {
		return ErrNotFound
	}
	raw := respBody.Data[0]
	if err := json.Unmarshal(raw, dst); err != nil {
		return err
	}
	return CaptureAdditionalFields(dst, raw)
}

func (c *client) getWithAdditionalFieldsV2(ctx context.Context, path string, dst AdditionalFieldsCapable) error {
	var raw json.RawMessage
	if err := c.Get(ctx, path, nil, &raw); err != nil {
		return err
	}
	if err := json.Unmarshal(raw, dst); err != nil {
		return err
	}
	return CaptureAdditionalFields(dst, raw)
}

// CaptureAdditionalFields finds additional fields in raw JSON and stores them on the target.
func CaptureAdditionalFields(dst AdditionalFieldsCapable, b []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	known := dst.KnownJSONFields()
	var extra map[string]json.RawMessage
	for k, v := range raw {
		if _, ok := known[k]; !ok {
			if extra == nil {
				extra = make(map[string]json.RawMessage)
			}
			extra[k] = v
		}
	}
	if len(extra) > 0 {
		dst.SetAdditionalFields(extra)
	}
	return nil
}
