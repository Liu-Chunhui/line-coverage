package fileparser

import (
	"bufio"
	"io"
	"os"
)

// ReadLines converts a file to lines.
// Each line should contain '\n' in the end
// len(lines) should match file lines
func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	reader := bufio.NewReader(file)

	for {
		bytes, err := reader.ReadString('\n')
		if err == io.EOF {
			lines = append(lines, "\n")
			break
		}
		lines = append(lines, string(bytes))
	}

	return lines, nil
}
