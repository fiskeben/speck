package input

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// FileReader reads user input from a file.
type FileReader struct {
	reader io.Reader
}

func newFileReader(fileName string) (Reader, error) {
	reader, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("could not open file for reading")
	}
	return FileReader{
		reader: reader,
	}, nil
}

func (f FileReader) Read() (*string, error) {
	contents, err := ioutil.ReadAll(f.reader)
	if err != nil {
		return nil, fmt.Errorf("unable to read input file: %s", err.Error())
	}
	res := string(contents)

	return &res, nil
}
