package input

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// Editor reads user input by opening an editor (as defined by $EDITOR with
// fallback to vi).
type Editor struct {
	result  *string
	warning *string
}

func newEditor(existingContent *string) (Reader, error) {
	return Editor{
		result: existingContent,
	}, nil
}

func (e Editor) Read() (*string, error) {
	for {
		err := e.getPostFromUser()
		if err != nil {
			return nil, err
		}
		if e.result == nil {
			return nil, fmt.Errorf("post is empty, nothing will be posted to micro.blog")
		}

		postStr := *e.result
		if len(postStr) > 280 {
			errorMessage := "The post is too long, please make it shorter."
			e.warning = &errorMessage
		} else {
			return &postStr, nil
		}
	}
}

func (e *Editor) getPostFromUser() error {
	fileContents := `
----------
Type your post above this line.
The line and everything below it will be discarded before the post is published.
`
	if e.warning != nil {
		fileContents = fmt.Sprintf("%s\n\nWarning: %s", fileContents, *e.warning)
	}

	if e.result != nil {
		fileContents = *e.result + fileContents
	}

	tempFile, err := makeTemporaryFile(fileContents)
	if err != nil {
		return err
	}

	err = executeEditor(tempFile)
	if err != nil {
		return err
	}

	bytes, err := ioutil.ReadFile(tempFile.Name())
	if err != nil {
		return fmt.Errorf("failed to read temporary file: %s", err.Error())
	}

	if err = clearTemporaryFile(tempFile); err != nil {
		fmt.Printf("Failed to clean up temporary file: %s\n", err.Error())
	}

	result := string(bytes)
	separatorPos := strings.LastIndex(result, "----------")
	if separatorPos > -1 {
		result = result[:separatorPos-1]
	}

	if len(result) == 0 {
		return nil
	}

	e.result = &result

	return nil
}

func makeTemporaryFile(data string) (*os.File, error) {
	tmpDir := os.TempDir()
	tmpFile, err := ioutil.TempFile(tmpDir, "micro.blog.post.draft.")
	if err != nil {
		return nil, fmt.Errorf("Error %s while creating tempFile", err.Error())
	}

	_, err = tmpFile.WriteString(data)
	if err != nil {
		fmt.Printf("Failed to write to temp file: %s\n", err.Error())
	}

	return tmpFile, nil
}

func clearTemporaryFile(file *os.File) error {
	return os.RemoveAll(file.Name())
}

var execCommand = exec.Command

func executeEditor(file *os.File) error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor, err := exec.LookPath("vi")
		if err != nil {
			return fmt.Errorf("Error %s while looking up for %s", editor, "vi")
		}
	}

	cmd := execCommand(editor, file.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("start failed: %s", err.Error())
	}
	err = cmd.Wait()
	if err != nil {
		return fmt.Errorf("could not wait for editor to finish: %s", err.Error())
	}
	return nil
}
