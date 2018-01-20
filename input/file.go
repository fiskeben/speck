package input

import (
	"fmt"
	"io/ioutil"
)

// FileReader reads user input from a file.
type FileReader struct {
	fileName string
}

func newFileReader(fileName string) Reader {
	return FileReader{
		fileName: fileName,
	}
}

func (f FileReader) Read() (*string, error) {
	path := f.fileName

	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Unable to read '%s': %s", path, err.Error())
	}
	res := string(contents)

	return &res, nil
}
