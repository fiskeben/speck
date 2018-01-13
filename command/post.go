package command

import (
	"fmt"
	"os"

	"github.com/fiskeben/microdotblog-cli/editor"
	"github.com/spf13/cobra"
)

var postCommand = &cobra.Command{
	Use: "post",
	Run: func(cmd *cobra.Command, args []string) {
		post := getPostFromUser()
		if post == nil {
			return
		}
		posted, err := client.Post(*post)
		if err != nil {
			fmt.Printf("Error creating post: %s\n", err.Error())
			os.Exit(1)
		}
		fmt.Printf("Your post was created: %s\n", posted.URL)
	},
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
