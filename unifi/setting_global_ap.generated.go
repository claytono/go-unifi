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

const SettingGlobalApKey = "global_ap"

type SettingGlobalAp struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	ApExclusions    []string                   `json:"ap_exclusions,omitempty" validate:"omitempty,mac"`                                  // ^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$
	NaChannelSize   int                        `json:"na_channel_size,omitempty" validate:"omitempty,oneof=20 40 80 160"`                 // 20|40|80|160
	NaTxPower       int                        `json:"na_tx_power,omitempty"`                                                             // [0-9]|[1-4][0-9]
	NaTxPowerMode   string                     `json:"na_tx_power_mode,omitempty" validate:"omitempty,oneof=auto medium high low custom"` // auto|medium|high|low|custom
	NgChannelSize   int                        `json:"ng_channel_size,omitempty" validate:"omitempty,oneof=20 40"`                        // 20|40
	NgTxPower       int                        `json:"ng_tx_power,omitempty"`                                                             // [0-9]|[1-4][0-9]
	NgTxPowerMode   string                     `json:"ng_tx_power_mode,omitempty" validate:"omitempty,oneof=auto medium high low custom"` // auto|medium|high|low|custom
	SixEChannelSize int                        `json:"6e_channel_size,omitempty" validate:"omitempty,oneof=20 40 80 160"`                 // 20|40|80|160
	SixETxPower     int                        `json:"6e_tx_power,omitempty"`                                                             // [0-9]|[1-4][0-9]
	SixETxPowerMode string                     `json:"6e_tx_power_mode,omitempty" validate:"omitempty,oneof=auto medium high low custom"` // auto|medium|high|low|custom
	ExtraFields     map[string]json.RawMessage `json:"-"`
}

func (dst *SettingGlobalAp) UnmarshalJSON(b []byte) error {
	type Alias SettingGlobalAp
	aux := &struct {
		NaChannelSize   emptyStringInt `json:"na_channel_size"`
		NaTxPower       emptyStringInt `json:"na_tx_power"`
		NgChannelSize   emptyStringInt `json:"ng_channel_size"`
		NgTxPower       emptyStringInt `json:"ng_tx_power"`
		SixEChannelSize emptyStringInt `json:"6e_channel_size"`
		SixETxPower     emptyStringInt `json:"6e_tx_power"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.NaChannelSize = int(aux.NaChannelSize)
	dst.NaTxPower = int(aux.NaTxPower)
	dst.NgChannelSize = int(aux.NgChannelSize)
	dst.NgTxPower = int(aux.NgTxPower)
	dst.SixEChannelSize = int(aux.SixEChannelSize)
	dst.SixETxPower = int(aux.SixETxPower)

	// Capture extra fields not in the struct
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(b, &raw); err == nil {
		known := map[string]struct{}{
			"_id":              {},
			"site_id":          {},
			"attr_hidden":      {},
			"attr_hidden_id":   {},
			"attr_no_delete":   {},
			"attr_no_edit":     {},
			"key":              {},
			"ap_exclusions":    {},
			"na_channel_size":  {},
			"na_tx_power":      {},
			"na_tx_power_mode": {},
			"ng_channel_size":  {},
			"ng_tx_power":      {},
			"ng_tx_power_mode": {},
			"6e_channel_size":  {},
			"6e_tx_power":      {},
			"6e_tx_power_mode": {},
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

func (src SettingGlobalAp) MarshalJSON() ([]byte, error) {
	type Alias SettingGlobalAp
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

// GetSettingGlobalAp Experimental! This function is not yet stable and may change in the future.
func (c *client) GetSettingGlobalAp(ctx context.Context, site string) (*SettingGlobalAp, error) {
	s, f, err := c.GetSetting(ctx, site, SettingGlobalApKey)
	if err != nil {
		return nil, err
	}
	if s.Key != SettingGlobalApKey {
		return nil, fmt.Errorf("unexpected setting key received. Requested: %q, received: %q", SettingGlobalApKey, s.Key)
	}
	return f.(*SettingGlobalAp), nil
}

// UpdateSettingGlobalAp Experimental! This function is not yet stable and may change in the future.
func (c *client) UpdateSettingGlobalAp(ctx context.Context, site string, s *SettingGlobalAp) (*SettingGlobalAp, error) {
	s.Key = SettingGlobalApKey
	result, err := c.SetSetting(ctx, site, SettingGlobalApKey, s)
	if err != nil {
		return nil, err
	}
	return result.(*SettingGlobalAp), nil
}
