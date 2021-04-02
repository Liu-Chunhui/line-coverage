package coverage

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testFile = "github.com/yesino/example.go"
)

func TestAnalyse(t *testing.T) {
	tests := []struct {
		name           string
		target         string
		branches       []*branch
		expectedResult *Result
		expectedErr    string
	}{
		{
			name:   "OneCoveredLine",
			target: testFile,

			branches: []*branch{
				{
					Start:   57,
					Finish:  57,
					Covered: true,
				},
			},

			expectedResult: &Result{
				Target:         testFile,
				CoveredLines:   1,
				UncoveredLines: 0,
			},
		},
		{
			name:   "OneCoveredLineTwoUncoveredLines",
			target: testFile,
			branches: []*branch{
				{
					Start:   57,
					Finish:  57,
					Covered: true,
				},
				{
					Start:   58,
					Finish:  59,
					Covered: false,
				},
			},

			expectedResult: &Result{
				Target:         testFile,
				CoveredLines:   1,
				UncoveredLines: 2,
			},
		},
		{
			name:   "WhenTargetHasEmptyBranchThenShouldReturnErr",
			target: testFile,

			expectedErr: "target contains empty branches. Target: github.com/yesino/example.go",
		},
		{
			name:   "WhenBranchEndingLineIsBeforeStartingLineThenShouldReturnErr",
			target: testFile,
			branches: []*branch{
				{
					Start:   20,
					Finish:  10,
					Covered: false,
				},
			},
			expectedErr: "branch starting line is behind ending line. Target file: github.com/yesino/example.go, branch: &{Start:20 Finish:10 Covered:false}",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := AnalyseTargetFile(tt.target, tt.branches)
			if len(tt.expectedErr) > 0 {
				require.Nil(t, got)
				assert.EqualError(t, err, tt.expectedErr)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedResult, got)
			}
		})
	}
}
