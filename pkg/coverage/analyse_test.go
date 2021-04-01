package coverage

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/Liu-Chunhui/line-coverage/pkg/fileparser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testFile = "github.com/yesino/example.go"
)

func TestMapProfileToBranch(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	execPath := filepath.Dir(filename)
	testfile := filepath.Join(execPath, "../../test/data/testcodefile")

	lines, _ := fileparser.ReadLines(testfile)
	base := filepath.Join(execPath, "../../")

	profileLine := "github.com/yesino/test/data/testcodefile:41.21,47.3 2 0"
	expected := &Branch{
		Start:   42,
		Finish:  46,
		Covered: false,
	}

	profile, err := ConvertToCoverageProfile(profileLine, "github.com/yesino", base)

	got, err := ConvertProfileToBranch(profile, lines)
	require.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestAnalyse(t *testing.T) {
	tests := []struct {
		name           string
		target         string
		branches       []*Branch
		expectedResult *Result
		expectedErr    string
	}{
		{
			name:   "OneCoveredLine",
			target: testFile,

			branches: []*Branch{
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
			branches: []*Branch{
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
			branches: []*Branch{
				{
					Start:   20,
					Finish:  10,
					Covered: false,
				},
			},
			expectedErr: "branch starting line is behind ending line. Target file: github.com/yesino/example.go, Branch: &{Start:20 Finish:10 Covered:false}",
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
