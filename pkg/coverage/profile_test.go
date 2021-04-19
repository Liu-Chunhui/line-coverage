package coverage

import (
	"testing"

	"github.com/Liu-Chunhui/line-coverage/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadProfiles(t *testing.T) {
	module := "github.com/Liu-Chunhui/line-coverage"
	testfile, _ := test.CreateTempFile("coverage.out", []byte(coverageout)) // 44 profiles
	got, err := loadProfiles(testfile.Name(), module, "base")
	require.NoError(t, err)
	assert.Equal(t, 19, len(got))
}

func TestMapLineToCoverageProfile(t *testing.T) {
	line := "github.com/Liu-Chunhui/line-coverage/pkg/percentage/percentage.go:8.36,13.2 3 3"
	module := "github.com/Liu-Chunhui/line-coverage"

	expected := &coverageProfile{
		Target:         "github.com/Liu-Chunhui/line-coverage/pkg/percentage/percentage.go",
		TargetFilename: "/line-coverage/pkg/percentage/percentage.go",
		StartLine:      8,
		StartPosition:  36,
		FinishLine:     13,
		FinishPosition: 2,
		Statements:     3,
		Executions:     3,
	}

	got, err := mapLineToCoverageProfile(line, module, "/line-coverage")
	require.NoError(t, err)
	assert.Equal(t, expected, got)
}

var coverageout = `mode: atomic
github.com/Liu-Chunhui/line-coverage/pkg/percentage/percentage.go:8.36,13.2 3 3
github.com/Liu-Chunhui/line-coverage/pkg/fileparser/readlines.go:14.80,16.16 2 1
github.com/Liu-Chunhui/line-coverage/pkg/fileparser/readlines.go:19.2,24.6 4 1
github.com/Liu-Chunhui/line-coverage/pkg/fileparser/readlines.go:44.2,44.19 1 1
github.com/Liu-Chunhui/line-coverage/pkg/fileparser/readlines.go:16.16,18.3 1 0
github.com/Liu-Chunhui/line-coverage/pkg/fileparser/readlines.go:24.6,26.17 2 60
github.com/Liu-Chunhui/line-coverage/pkg/fileparser/readlines.go:34.3,35.17 2 59
github.com/Liu-Chunhui/line-coverage/pkg/fileparser/readlines.go:39.3,39.17 1 59
github.com/Liu-Chunhui/line-coverage/pkg/fileparser/readlines.go:26.17,27.21 1 1
github.com/Liu-Chunhui/line-coverage/pkg/fileparser/readlines.go:31.4,31.19 1 0
github.com/Liu-Chunhui/line-coverage/pkg/fileparser/readlines.go:27.21,28.10 1 1
github.com/Liu-Chunhui/line-coverage/pkg/fileparser/readlines.go:35.17,37.4 1 0
github.com/Liu-Chunhui/line-coverage/pkg/fileparser/readlines.go:39.17,41.4 1 59
github.com/Liu-Chunhui/line-coverage/pkg/fileparser/readlines.go:47.62,48.29 1 63
github.com/Liu-Chunhui/line-coverage/pkg/fileparser/readlines.go:59.2,59.19 1 60
github.com/Liu-Chunhui/line-coverage/pkg/fileparser/readlines.go:48.29,50.17 2 6
github.com/Liu-Chunhui/line-coverage/pkg/fileparser/readlines.go:54.3,54.12 1 6
github.com/Liu-Chunhui/line-coverage/pkg/fileparser/readlines.go:50.17,52.4 1 0
github.com/Liu-Chunhui/line-coverage/pkg/fileparser/readlines.go:54.12,56.4 1 3
`
