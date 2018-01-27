package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var followersCommand = &cobra.Command{
	Use:   "followers",
	Short: "Get a list of your followers",
	Long:  "Lists your followers as [username] [real name]",
	Run:   followers,
}

var followCommand = &cobra.Command{
	Use:   "follow [username]",
	Short: "Start following another user",
	Long:  "Start following the specified user",
	Run:   follow,
}

var unfollowCommand = &cobra.Command{
	Use:   "unfollow [username]",
	Short: "Stop following a user",
	Long:  "Stop following the specified user",
	Run:   unfollow,
}

func followers(cmd *cobra.Command, args []string) {
	followers, err := client.Followers(viper.GetString("username"))
	if err != nil {
		fmt.Printf("Error getting list of followers: %s\n", err.Error())
		os.Exit(1)
	}

	for _, follower := range followers {
		fmt.Printf("%s %s\n", follower.Username, follower.Name)
	}
}

func follow(cmd *cobra.Command, args []string) {
	username := args[0]
	if err := client.Follow(username); err != nil {
		fmt.Printf("Could not follow %s: %s\n", username, err.Error())
		os.Exit(2)
	}
}

func unfollow(cmd *cobra.Command, args []string) {
	username := args[0]
	if err := client.Unfollow(username); err != nil {
		fmt.Printf("Could not unfollow %s: %s\n", username, err.Error())
		os.Exit(2)
	}
}
