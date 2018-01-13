package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func timeline(cmd *cobra.Command, args []string) {
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
}
