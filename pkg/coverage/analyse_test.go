package coverage

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testFile = "example.go"
)

func TestAnalyse(t *testing.T) {
	tests := []struct {
		name           string
		target         Target
		expectedResult *Result
		expectedErr    string
	}{
		{
			name: "OneCoveredLine",
			target: Target{
				File: testFile,
				Branches: []Branch{
					{
						From:    57,
						To:      57,
						Covered: true,
					},
				},
			},
			expectedResult: &Result{
				File:           testFile,
				CoveredLines:   1,
				UncoveredLines: 0,
			},
		},
		{
			name: "OneCoveredLineTwoUncoveredLines",
			target: Target{
				File: testFile,
				Branches: []Branch{
					{
						From:    57,
						To:      57,
						Covered: true,
					},
					{
						From:    58,
						To:      59,
						Covered: false,
					},
				},
			},
			expectedResult: &Result{
				File:           testFile,
				CoveredLines:   1,
				UncoveredLines: 2,
			},
		},
		{
			name: "WhenTargetHasEmptyBranchThenShouldReturnErr",
			target: Target{
				File: testFile,
			},
			expectedErr: "target contains empty branches. Target: example.go",
		},
		{
			name: "WhenBranchEndingLineIsBeforeStartingLineThenShouldReturnErr",
			target: Target{
				File: testFile,
				Branches: []Branch{
					{
						From: 20,
						To:   10,
					},
				},
			},
			expectedErr: "branch starting line is behind ending line. Target file: example.go, Branch: {From:20 To:10 Covered:false}",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Analyse(tt.target)
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
