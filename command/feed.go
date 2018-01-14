package command

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func timeline(cmd *cobra.Command, args []string) {
	feed, err := client.GetPosts()
	if err != nil {
		fmt.Println("Error getting feed")
		os.Exit(1)
	}

	limitStr := cmd.Flag("limit").Value.String()
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		fmt.Printf("Failed to parse limit '%s', falling back to limit to ten items (error: %s)\n", limitStr, err.Error())
		limit = 10
	}

	items := feed.Items[:limit]
	width := terminalWidth()
	line := strings.Repeat("-", width)

	for i, p := range items {
		if i > 0 {
			fmt.Printf("\n%s\n", line)
		}
		fmt.Printf("%s (%s) wrote:\n%s\nPosted %s - %s\n", p.Author.Name, p.Author.MicroblogProperties.Username, p.ContentHTML, p.DatePublished, p.URL)
	}
}

func terminalWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 100
	}

	data := strings.TrimSpace(string(out))
	parts := strings.Split(data, " ")

	width, err := strconv.Atoi(parts[1])
	if err != nil {
		return 100
	}

	return width
}
