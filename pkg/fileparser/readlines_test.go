package fileparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Liu-Chunhui/line-coverage/test"
)

func TestReadLines(t *testing.T) {
	testfile, _ := test.CreateTempFile("testcode.go", []byte(testcodefile))

	got, err := ReadLines(testfile.Name())
	require.NoError(t, err)
	assert.Equal(t, 59, len(got))  // total 59 lines
	assert.Equal(t, "\n", got[41]) // line 70 is new line
}

var testcodefile = `package fileparser

import (
	"bufio"
	"io"
	"os"
	"regexp"
)

// ReadLines converts a file to lines.
// Each line should contain '\n' in the end
// len(lines) should match file lines
func ReadLines(filename string, excludingPatterns ...string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, err
		}

		execluded, err := skipLine(line, excludingPatterns...)
		if err != nil {
			return nil, err
		}

		if !execluded {
			lines = append(lines, string(line))
		}
	}

	return lines, nil
}

func skipLine(line string, patterns ...string) (bool, error) {
	for _, p := range patterns {
		match, err := regexp.Match(p, []byte(line))
		if err != nil {
			return true, err
		}

		if match {
			return true, nil
		}
	}

	return false, nil
}
`
