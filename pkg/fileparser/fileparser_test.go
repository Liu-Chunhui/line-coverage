package fileparser

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadLines(t *testing.T) {
	exec, _ := os.Executable()
	execPath := filepath.Dir(exec)
	testfile := filepath.Join(execPath, "../../test/data/testcodefile")

	got, err := ReadLines(testfile)
	require.NoError(t, err)
	assert.Equal(t, 71, len(got))  // total 71 lines
	assert.Equal(t, "\n", got[47]) // line 47 is new line
}
