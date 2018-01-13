package editor

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// Edit creates a temporary file and opens it with the user's editor
// as specified by $EDITOR. If $EDITOR is not set, vi will be used.
func Edit(data *string) (*string, error) {
	fileContents := `
----------
Type your post above this line.
The line and everything below it will be discarded before the post is published.
`
	if data != nil {
		fileContents = *data + fileContents
	}

	tempFile, err := makeTemporaryFile(fileContents)
	if err != nil {
		return nil, err
	}

	err = executeEditor(tempFile)
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadFile(tempFile.Name())
	if err != nil {
		return nil, err
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
		return nil, nil
	}

	return &result, nil
}

func makeTemporaryFile(data string) (*os.File, error) {
	tmpDir := os.TempDir()
	tmpFile, err := ioutil.TempFile(tmpDir, "micro.blog.post.draft.")
	if err != nil {
		fmt.Printf("Error %s while creating tempFile", err.Error())
		return nil, err
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

func executeEditor(file *os.File) error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor, err := exec.LookPath("vi")
		if err != nil {
			fmt.Printf("Error %s while looking up for %s!!", editor, "vi")
			return err
		}
	}

	cmd := exec.Command(editor, file.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		fmt.Printf("Start failed: %s", err.Error())
		return err
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Printf("Could not wait for editor to finish: %s\n", err.Error())
		return err
	}
	return nil
}
