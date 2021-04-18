package percentage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisplay(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		expected string
	}{
		{
			name:     "0.8386",
			value:    0.8386,
			expected: "83.9%",
		},
		{
			name:     "0.8384",
			value:    0.8384,
			expected: "83.8%",
		},
		{
			name:     "0.4615",
			value:    0.4615,
			expected: "46.2%",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := Display(tt.value)
			assert.Equal(t, tt.expected, got)
		})
	}
}
