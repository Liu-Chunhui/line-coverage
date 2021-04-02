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

	profileLine := "github.com/yesino/test/data/testcodefile:41.21,47.3 2 0"
	expected := &branch{
		Start:   42,
		Finish:  46,
		Covered: false,
	}

	profile, err := mapLineToCoverageProfile(profileLine, "github.com/yesino", base)

	got, err := convertProfileToBranch(profile, lines)
	require.NoError(t, err)
	assert.Equal(t, expected, got)
}
