package coverage

import (
	"fmt"
	"testing"

	"github.com/Liu-Chunhui/line-coverage/test"
	"github.com/stretchr/testify/assert"
)

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
					File:           testFile,
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
