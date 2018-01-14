package command

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fiskeben/speck/editor"
	"github.com/spf13/cobra"
)

var postCommand = &cobra.Command{
	Use: "post",
	Run: post,
}

func post(cmd *cobra.Command, args []string) {
	saveFile := cmd.Flag("save").Value.String()
	var post *string

	if len(args) == 1 {
		post = getPostFromFile(args[0], saveFile)
	} else {
		post = getPostFromUser(saveFile)
	}
	if post == nil {
		os.Exit(1)
	}

	if dryRun {
		fmt.Println("Dry run mode - your post will not be published.")
		fmt.Println("Here's what your post would be like:")
		fmt.Println("")
		fmt.Println(*post)
		fmt.Println("Run again without the --dry-run flag to publish your words.")
		return
	}

	posted, err := client.Post(*post)
	if err != nil {
		fmt.Printf("Error creating post: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("Your post was created: %s\n", posted.URL)
}

func getPostFromUser(saveFile string) *string {
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
			if saveFile != "" {
				writePostToFile(saveFile, *post)
			}
			return post
		}
	}
}

func getPostFromFile(path, saveFile string) *string {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Unable to read %s: %s\n", path, err.Error())
		return nil
	}
	res := string(contents)

	if saveFile != "" {
		writePostToFile(saveFile, res)
	}

	return &res
}

func writePostToFile(path, contents string) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Printf("Failed to create file '' to save post to: %s\n", err.Error())
	}

	_, err = file.WriteString(contents)
	if err != nil {
		fmt.Printf("Failed to save post to file: %s\n", err.Error())
	}
}
