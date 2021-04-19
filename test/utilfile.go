package test

import (
	"io/ioutil"
	"os"
)

func CreateTempFile(fileName string, content []byte) (f *os.File, err error) {
	file, _ := ioutil.TempFile("", fileName)
	defer func() {
		_ = file.Close()
	}()

	_, err = file.Write(content)
	if err != nil {
		return nil, err
	}
	return file, nil
}
