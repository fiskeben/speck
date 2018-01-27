package command

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fiskeben/speck/blog"
	"github.com/fiskeben/speck/input"
	"github.com/spf13/cobra"
)

var postCommand = &cobra.Command{
	Use:   "post",
	Short: "Create a new post",
	Long:  "Write a new post or supply one as a file and post it to Micro.blog",
	Args:  cobra.NoArgs,
	Run:   post,
}

var replyCommand = &cobra.Command{
	Use:   "reply [post ID]",
	Short: "Reply to a post",
	Long:  "Post a reply to a given post by suppying the ID of the post to reply to.",
	Args:  cobra.ExactArgs(1),
	Run:   reply,
}

var saveFile string

var fromFile string

func init() {
	postCommand.Flags().StringVarP(&saveFile, "save-as", "s", "", "Path to file to save the post to")
	postCommand.Flags().StringVarP(&fromFile, "from-file", "f", "", "Path to file to read from instead of opening an editor")

	replyCommand.Flags().StringVarP(&saveFile, "save-as", "s", "", "Path to file to save the post to")
	replyCommand.Flags().StringVarP(&fromFile, "from-file", "f", "", "Path to file to read from instead of opening an editor")
}

func post(cmd *cobra.Command, args []string) {
	post, err := readPost(fromFile)
	if err != nil {
		fail(err)
	}

	if saveFile != "" {
		if err := writePostToFile(saveFile, post); err != nil {
			fail(err)
		}
	}

	response, err := blog.Post(client, post, dryRun)
	if err != nil {
		fail(err)
	}
	fmt.Println(*response)
}

func reply(cmd *cobra.Command, args []string) {
	replyTo, err := strconv.ParseInt(args[0], 10, 0)
	if err != nil {
		fail(fmt.Errorf("the ID '%v' does not seem to be a number", args[0]))
	}

	post, err := readPost(fromFile)
	if err != nil {
		fail(err)
	}

	if saveFile != "" {
		if err := writePostToFile(saveFile, post); err != nil {
			fail(err)
		}
	}

	response, err := blog.Reply(client, replyTo, post, dryRun)
	if err != nil {
		fail(err)
	}
	fmt.Println(*response)
}

func readPost(fromFile string) (string, error) {
	r, err := input.NewReader(fromFile)
	if err != nil {
		fail(err)
	}

	post, err := r.Read()

	if err != nil {
		return "", err
	}

	if post == nil {
		return "", fmt.Errorf("unable to post from file")
	}

	return *post, nil
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
