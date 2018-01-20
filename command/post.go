package command

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fiskeben/speck/blog"
	"github.com/fiskeben/speck/editor"
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

	var post *string
	var err error

	if fromFile == nil {
		post, err = getPostFromUser()
	} else {
		post, err = getPostFromFile(*fromFile)
	}

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

func getPostFromUser() (*string, error) {
	var post *string
	var err error

	for {
		post, err = editor.Edit(post)
		if err != nil {
			return nil, err
		}
		if post == nil {
			return nil, fmt.Errorf("post is empty, nothing will be posted to micro.blog")
		}

		postStr := *post
		if len(postStr) > 280 {
			fmt.Println("The post is too long, please make it shorter.")
		} else {
			break
		}
	}
	return post, nil
}

func getPostFromFile(path string) (*string, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Unable to read '%s': %s", path, err.Error())
	}
	res := string(contents)

	return &res, nil
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
