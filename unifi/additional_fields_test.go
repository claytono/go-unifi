package unifi

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// mockAdditionalFieldsCapable is a test double for AdditionalFieldsCapable
type mockAdditionalFieldsCapable struct {
	knownFields      map[string]struct{}
	additionalFields map[string]json.RawMessage
}

func (m *mockAdditionalFieldsCapable) KnownJSONFields() map[string]struct{} {
	return m.knownFields
}

func (m *mockAdditionalFieldsCapable) SetAdditionalFields(ef map[string]json.RawMessage) {
	m.additionalFields = ef
}

func TestCaptureAdditionalFields(t *testing.T) {
	tests := []struct {
		name        string
		jsonData    string
		knownFields map[string]struct{}
		wantExtra   map[string]string // simplified: just check keys and string values
		wantNil     bool
	}{
		{
			name:     "captures unknown fields",
			jsonData: `{"name":"test","_id":"abc123","unknown_field":"value","another":42}`,
			knownFields: map[string]struct{}{
				"name": {},
				"_id":  {},
			},
			wantExtra: map[string]string{
				"unknown_field": `"value"`,
				"another":       `42`,
			},
		},
		{
			name:     "no additional fields when all known",
			jsonData: `{"name":"test","_id":"abc123"}`,
			knownFields: map[string]struct{}{
				"name": {},
				"_id":  {},
			},
			wantNil: true,
		},
		{
			name:        "all fields unknown",
			jsonData:    `{"foo":"bar","baz":123}`,
			knownFields: map[string]struct{}{},
			wantExtra: map[string]string{
				"foo": `"bar"`,
				"baz": `123`,
			},
		},
		{
			name:     "handles nested objects as raw JSON",
			jsonData: `{"name":"test","nested":{"a":1,"b":2}}`,
			knownFields: map[string]struct{}{
				"name": {},
			},
			wantExtra: map[string]string{
				"nested": `{"a":1,"b":2}`,
			},
		},
		{
			name:     "handles arrays as raw JSON",
			jsonData: `{"name":"test","items":[1,2,3]}`,
			knownFields: map[string]struct{}{
				"name": {},
			},
			wantExtra: map[string]string{
				"items": `[1,2,3]`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mockAdditionalFieldsCapable{
				knownFields: tt.knownFields,
			}

			err := CaptureAdditionalFields(m, []byte(tt.jsonData))
			require.NoError(t, err)

			if tt.wantNil {
				assert.Nil(t, m.additionalFields)
				return
			}

			require.NotNil(t, m.additionalFields)
			assert.Len(t, m.additionalFields, len(tt.wantExtra))

			for k, wantVal := range tt.wantExtra {
				gotVal, ok := m.additionalFields[k]
				assert.True(t, ok, "expected key %q in additional fields", k)
				assert.JSONEq(t, wantVal, string(gotVal))
			}
		})
	}
}

func TestCaptureAdditionalFields_InvalidJSON(t *testing.T) {
	m := &mockAdditionalFieldsCapable{
		knownFields: map[string]struct{}{},
	}

	err := CaptureAdditionalFields(m, []byte(`not valid json`))
	assert.Error(t, err)
}

// Test that generated types implement AdditionalFieldsCapable
func TestGeneratedTypesImplementAdditionalFieldsCapable(t *testing.T) {
	// Verify a few representative generated types implement the interface
	var _ AdditionalFieldsCapable = &Network{}
	var _ AdditionalFieldsCapable = &Device{}
	var _ AdditionalFieldsCapable = &WLAN{}
	var _ AdditionalFieldsCapable = &User{}
	var _ AdditionalFieldsCapable = &SettingMgmt{}
}

// Test that KnownJSONFields returns expected fields for a generated type
func TestNetwork_KnownJSONFields(t *testing.T) {
	n := &Network{}
	known := n.KnownJSONFields()

	// Check a few expected fields
	assert.Contains(t, known, "_id")
	assert.Contains(t, known, "site_id")
	assert.Contains(t, known, "name")
	assert.Contains(t, known, "purpose")

	// Should not contain made-up fields
	assert.NotContains(t, known, "not_a_real_field")
}

// Test SetAdditionalFields on a generated type
func TestNetwork_SetAdditionalFields(t *testing.T) {
	n := &Network{}
	extra := map[string]json.RawMessage{
		"custom_field": json.RawMessage(`"custom_value"`),
	}

	n.SetAdditionalFields(extra)

	assert.Equal(t, extra, n.AdditionalFields)
}

// Test MarshalJSON includes additional fields
func TestNetwork_MarshalJSON_WithAdditionalFields(t *testing.T) {
	n := &Network{
		Name: "test-network",
	}
	n.AdditionalFields = map[string]json.RawMessage{
		"custom": json.RawMessage(`"value"`),
	}

	b, err := json.Marshal(n)
	require.NoError(t, err)

	var m map[string]interface{}
	err = json.Unmarshal(b, &m)
	require.NoError(t, err)

	assert.Equal(t, "test-network", m["name"])
	assert.Contains(t, m, "_additional_properties")
}

// Test MarshalJSON without additional fields doesn't add _additional_properties
func TestNetwork_MarshalJSON_WithoutAdditionalFields(t *testing.T) {
	n := &Network{
		Name: "test-network",
	}

	b, err := json.Marshal(n)
	require.NoError(t, err)

	var m map[string]interface{}
	err = json.Unmarshal(b, &m)
	require.NoError(t, err)

	assert.Equal(t, "test-network", m["name"])
	assert.NotContains(t, m, "_additional_properties")
}
