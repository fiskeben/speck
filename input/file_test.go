package input

import (
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_newFileReader(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want Reader
	}{
		{
			"get a file reader",
			args{
				fileName: "./test.md",
			},
			FileReader{
				reader: strings.NewReader("hooray"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := newFileReader(tt.args.fileName); !reflect.DeepEqual(got, tt.want) && err == nil {
				t.Errorf("newFileReader() = %v, want %v", got, tt.want)
			}
		})
	}
}

type failingReader struct{}

func (f failingReader) Read(b []byte) (int, error) {
	return 0, fmt.Errorf("just failing")
}

func TestFileReader_Read(t *testing.T) {
	postdata := "This is a test post"

	type fields struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "reads contents of file",
			fields: fields{
				reader: strings.NewReader(postdata),
			},
			want:    postdata,
			wantErr: false,
		},
		{
			name: "fails to read file",
			fields: fields{
				reader: failingReader{},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := FileReader{
				reader: tt.fields.reader,
			}
			got, err := f.Read()
			if (err != nil) != tt.wantErr {
				t.Errorf("FileReader.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != "" && *got != tt.want {
				t.Errorf("FileReader.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}
