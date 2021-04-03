package coverage

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/Liu-Chunhui/line-coverage/pkg/fileparser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMapProfileToBranch(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	execPath := filepath.Dir(filename)
	testfile := filepath.Join(execPath, "../../test/data/testcodefile")

	lines, _ := fileparser.ReadLines(testfile)
	base := filepath.Join(execPath, "../../")

	tests := []struct {
		name        string
		profileLine string
		expected    *branch
	}{
		{
			name:        "uncoveredLines",
			profileLine: "github.com/Liu-Chunhui/line-coverage/test/data/testcodefile:41.21,47.3 2 0",
			expected: &branch{
				Start:   42,
				Finish:  46,
				Covered: false,
			},
		},
		{
			name:        "finishingLineIs},nil",
			profileLine: "github.com/Liu-Chunhui/line-coverage/test/data/testcodefile:65.2,69.8 1 1",
			expected: &branch{
				Start:   65,
				Finish:  69,
				Covered: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			profile, err := mapLineToCoverageProfile(tt.profileLine, "github.com/Liu-Chunhui/line-coverage", base)
			require.NoError(t, err)

			target, branch := convertProfileToBranch(profile, lines)
			require.NotNil(t, target)
			assert.NotEmpty(t, target)
			assert.Equal(t, tt.expected, branch)
		})
	}

}
