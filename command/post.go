package command

import (
	"fmt"
	"os"

	"github.com/fiskeben/speck/blog"
	"github.com/fiskeben/speck/input"
	"github.com/spf13/cobra"
)

var postCommand = &cobra.Command{
	Use: "post",
	Run: post,
}

func post(cmd *cobra.Command, args []string) {
	saveFile := cmd.Flag("save").Value.String()

	var fromFile *string
	if len(args) > 0 {
		fromFile = &args[0]
	}

	r := input.NewReader(fromFile)
	post, err := r.Read()

	if err != nil {
		fail(err)
	}

	if post == nil {
		fail(fmt.Errorf("unable to post"))
	}

	if saveFile != "" {
		if err = writePostToFile(saveFile, *post); err != nil {
			fail(err)
		}
	}

	response, err := blog.Post(client, *post, dryRun)
	if err != nil {
		fail(err)
	}
	fmt.Println(*response)
}

func fail(reason error) {
	fmt.Println(reason.Error())
	os.Exit(2)
}

func writePostToFile(path, contents string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Failed to create file '%s' to save post to: %s", path, err.Error())
	}

	_, err = file.WriteString(contents)
	if err != nil {
		return fmt.Errorf("Failed to save post to file: %s", err.Error())
	}

	return nil
}
