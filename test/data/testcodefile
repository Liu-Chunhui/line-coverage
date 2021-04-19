package fileparser

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"regexp"
)

// ReadLines converts a file to lines.
// Each line should contain '\n' in the end
// len(lines) should match file lines
func ReadLines(filename string, excludingPatterns ...string) ([]string, error) {
	file, err := os.Open(filepath.Clean(filename))
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
			lines = append(lines, line)
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
