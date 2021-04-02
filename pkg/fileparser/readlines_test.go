package fileparser

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadLines(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	execPath := filepath.Dir(filename)
	testfile := filepath.Join(execPath, "../../test/data/testcodefile")

	got, err := ReadLines(testfile)
	require.NoError(t, err)
	assert.Equal(t, 71, len(got))  // total 71 lines
	assert.Equal(t, "\n", got[47]) // line 47 is new line
}

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

			got, err := skipLine(tt.line, tt.rules...)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, got)
		})
	}
}
