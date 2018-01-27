package command

import (
	"fmt"
	"os"

	api "github.com/fiskeben/microdotblog"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var version string
var dryRun bool
var client api.APIClient

// Execute performs the root command.
func Execute() error {
	return rootCommand.Execute()
}

var rootCommand = &cobra.Command{
	Use:     "speck",
	Run:     timeline,
	Version: version,
}

func init() {
	initConfig()
	client = api.NewAPIClient(viper.GetString("token"))

	args := os.Args[1:]

	rootCommand.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "verbose output")

	var tmpLimit int
	rootCommand.Flags().IntVarP(&tmpLimit, "limit", "l", 10, "Limit the number of items to list.")
	rootCommand.Flags().Parse(args)
	rootCommand.SetArgs(args)

	rootCommand.AddCommand(postCommand)
	rootCommand.AddCommand(replyCommand)
	rootCommand.AddCommand(followersCommand)
	rootCommand.AddCommand(followCommand)
	rootCommand.AddCommand(unfollowCommand)
}

func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.AddConfigPath(home)
	viper.AddConfigPath(".")
	viper.SetConfigName(".speck")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}
