package command

import (
	"fmt"
	"os"

	api "github.com/fiskeben/microdotblog"
	"github.com/fiskeben/microdotblog-cli/editor"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Execute performs the root command.
func Execute() error {
	return rootCommand.Execute()
}

var client api.APIClient

var rootCommand = &cobra.Command{
	Use: "mcro [command]",
	Run: func(cmd *cobra.Command, args []string) {
		feed, err := client.GetPosts()
		if err != nil {
			fmt.Println("Error getting feed")
			os.Exit(1)
		}

		for i, p := range feed.Items {
			if i > 0 {
				fmt.Print("\n")
			}
			fmt.Printf("%d. %s %s\n%s\n%s\n", i, p.Author.Name, p.DatePublished, p.ContentHTML, p.URL)
		}
	},
}

var postCommand = &cobra.Command{
	Use: "post [flags]",
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

func init() {
	initConfig()
	client = api.NewAPIClient(viper.GetString("token"))
	args := os.Args[1:]
	rootCommand.Flags().Parse(args)
	rootCommand.SetArgs(args)
	rootCommand.AddCommand(postCommand)
}

func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.AddConfigPath(home)
	viper.AddConfigPath(".")
	viper.SetConfigName(".microdotblog")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}
