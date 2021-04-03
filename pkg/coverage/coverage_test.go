package coverage

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/Liu-Chunhui/line-coverage/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCalculate(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	execPath := filepath.Dir(filename)

	tests := []struct {
		name            string //profileFilename string, module string, basePath string
		profileFilename string
		module          string
		base            string
		expected        []*Result
	}{
		{
			name:            "test",
			profileFilename: filepath.Join(execPath, "../../test/data/testcodefile.out"),
			module:          "github.com/Liu-Chunhui/line-coverage",
			base:            filepath.Join(execPath, "../../"),
			expected: []*Result{
				{
					Target:         "github.com/Liu-Chunhui/line-coverage/test/data/testcodefile",
					CoveredLines:   26,
					UncoveredLines: 5,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Calculate(tt.profileFilename, tt.module, tt.base)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestCalculateOverall(t *testing.T) {
	tests := []struct {
		name     string
		results  []*Result
		expected float64
	}{
		{
			name: "ResultFromTestFile",
			results: []*Result{
				{
					Target:         testFile,
					CoveredLines:   26,
					UncoveredLines: 5,
				},
			},
			expected: float64(26) / float64(5+26),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := CalculateOverall(tt.results)
			fmt.Println(got)
			assert.True(t, test.FloatEquals(tt.expected, got))
		})
	}
}
