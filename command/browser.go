package command

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var openTimelineCommand = &cobra.Command{
	Use:   "open [username]",
	Short: "Open a timeline in a browser.",
	Long: `Opens a browser and loads the specified user's timeline.
If the user is not specified, it will open your own timeline.`,
	Run: openTimeline,
}

func openTimeline(cmd *cobra.Command, args []string) {
	var username string

	if len(args) == 1 {
		username = args[0]
	} else {
		username = viper.GetString("username")
	}

	url := fmt.Sprintf("https://micro.blog/%s", username)

	if err := openbrowser(url); err != nil {
		fmt.Printf("Failed to open browser: %s\n", err.Error())
		os.Exit(1)
	}
}

func openbrowser(url string) error {

	switch runtime.GOOS {
	case "linux":
		return exec.Command("xdg-open", url).Start()
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		return exec.Command("open", url).Start()
	default:
		return fmt.Errorf("unsupported platform '%s'", runtime.GOOS)
	}
}
