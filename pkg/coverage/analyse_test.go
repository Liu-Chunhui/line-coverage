package coverage

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Liu-Chunhui/line-coverage/pkg/fileparser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testFile = "github.com/yesino/example.go"
)

func TestMapProfileToBranch(t *testing.T) {
	exec, _ := os.Executable()
	execPath := filepath.Dir(exec)
	testfile := filepath.Join(execPath, "../../test/data/testcodefile")
	lines, _ := fileparser.ReadLines(testfile)
	profileLine := "github.com/yesino/line-coverage/test/data/testcodefile:41.21,47.3 2 0"
	expected := &Branch{
		Target:  "github.com/yesino/line-coverage/test/data/testcodefile",
		Start:   42,
		Finish:  46,
		Covered: false,
	}

	got, err := ConvertProfileToBranch(profileLine, lines)
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
					Start:  20,
					Finish: 10,
				},
			},
			expectedErr: "branch starting line is behind ending line. Target file: github.com/yesino/example.go, Branch: &{Target: Start:20 Finish:10 Covered:false}",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Analyse(tt.target, tt.branches)
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
