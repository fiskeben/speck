package command

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fiskeben/microdotblog-cli/editor"
	"github.com/spf13/cobra"
)

var postCommand = &cobra.Command{
	Use: "post",
	Run: post,
}

func post(cmd *cobra.Command, args []string) {
	var post *string

	if len(args) == 1 {
		post = getPostFromFile(args[0])
	} else {
		post = getPostFromUser()
	}
	if post == nil {
		os.Exit(1)
	}

	posted, err := client.Post(*post)
	if err != nil {
		fmt.Printf("Error creating post: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("Your post was created: %s\n", posted.URL)
}

func getPostFromUser() *string {
	var post *string
	var err error
	for {
		post, err = editor.Edit(post)
		if err != nil {
			os.Exit(1)
		}
		if post == nil {
			fmt.Println("Post is empty, nothing will be posted to micro.blog.")
			return nil
		}

		postStr := *post
		if len(postStr) > 280 {
			fmt.Println("The post is too long, please make it shorter.")
		} else {
			return post
		}
	}
}

func getPostFromFile(path string) *string {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Unable to read %s: %s\n", path, err.Error())
		return nil
	}
	res := string(contents)
	return &res
}
