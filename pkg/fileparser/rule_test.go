package fileparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSkipLine(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		rules    []string
		expected bool
	}{
		{
			name:     "include",
			line:     "github.com/Liu-Chunhui/line-coverage/pkg/fileparser/readlines.go:12.51,14.16 2 1\n",
			rules:    CoverageProfileExcludingRules,
			expected: false,
		},
		{
			name:     "exclude",
			line:     "mode: set\n",
			rules:    CoverageProfileExcludingRules,
			expected: true,
		},
		{
			name:     "new line",
			line:     "\n",
			rules:    CoverageProfileExcludingRules,
			expected: true,
		},
		{
			name:     "empty line",
			line:     "    \n",
			rules:    CoverageProfileExcludingRules,
			expected: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := MatchPattern(tt.line, tt.rules...)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, got)
		})
	}
}
