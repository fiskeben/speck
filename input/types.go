package input

// Reader defines how to read input from the user.
type Reader interface {
	Read() (*string, error)
}

// NewReader returns the correct Reader implementation based on whether an
// input file was supplied or not.
func NewReader(inputFile *string) Reader {
	if inputFile == nil {
		return newEditor(nil)
	}
	return newFileReader(*inputFile)
}
