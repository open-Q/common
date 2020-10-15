package golang

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewGenericFlag(t *testing.T) {
	v := interface{}(123)
	f := NewGenericFlag(&v)
	require.NotNil(t, f)
}

func TestGenericFlag_Value(t *testing.T) {
	tt := []struct {
		name  string
		value interface{}
	}{
		{
			name: "with nil value",
		},
		{
			name:  "with bool value",
			value: true,
		},
		{
			name:  "with int value",
			value: 123,
		},
		{
			name:  "with float value",
			value: 123.56,
		},
		{
			name:  "with struct value",
			value: struct{}{},
		},
		{
			name:  "with array value",
			value: []string{"hello", "world"},
		},
	}
	for i := range tt {
		tc := &tt[i]
		t.Run(tc.name, func(t *testing.T) {
			flag := NewGenericFlag(tc.value)
			require.Equal(t, tc.value, flag.Value())
		})
	}
}
