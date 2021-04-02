package coverage

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadProfiles(t *testing.T) {
	module := "github.com/yesino/line-coverage"

	_, filename, _, _ := runtime.Caller(0)
	execPath := filepath.Dir(filename)
	base := filepath.Join(execPath, "../../")

	testfile := filepath.Join(execPath, "../../test/data/test.out") // 44 profiles

	got, err := loadProfiles(testfile, module, base)
	require.NoError(t, err)
	assert.Equal(t, 44, len(got))
}

func TestMapLineToCoverageProfile(t *testing.T) {
	line := "github.com/yesino/line-coverage/test/testdata.go:30.50,32.9 2 4"
	module := "github.com/yesino/line-coverage"

	_, filename, _, _ := runtime.Caller(0)
	execPath := filepath.Dir(filename)
	base := filepath.Join(execPath, "../../")
	expected := &coverageProfile{
		Target:         "github.com/yesino/line-coverage/test/testdata.go",
		TargetFilename: "/Users/yesino/Documents/GitHub/line-coverage/test/testdata.go",
		StartLine:      30,
		StartPosition:  50,
		FinishLine:     32,
		FinishPosition: 9,
		Statements:     2,
		Executions:     4,
	}

	got, err := mapLineToCoverageProfile(line, module, base)
	require.NoError(t, err)
	assert.Equal(t, expected, got)
}
