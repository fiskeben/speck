package command

import (
	"fmt"
	"os"

	api "github.com/fiskeben/microdotblog"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var dryRun bool

// Execute performs the root command.
func Execute() error {
	return rootCommand.Execute()
}

var client api.APIClient

var rootCommand = &cobra.Command{
	Use: "speck",
	Run: timeline,
}

func init() {
	initConfig()
	client = api.NewAPIClient(viper.GetString("token"))
	args := os.Args[1:]
	rootCommand.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "verbose output")
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
	viper.SetConfigName(".speck")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}
