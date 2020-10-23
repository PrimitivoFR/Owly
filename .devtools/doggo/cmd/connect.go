package cmd

import (
	"doggo/evansctl"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Opens an Evans session at desired port, using a user's token.",
	Long: `
	Opens an Evans session at desired port, using a user's token from
	provided username and password`,
	Run: func(cmd *cobra.Command, args []string) {

		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		port, _ := cmd.Flags().GetString("port")

		if username == "" || password == "" || port == "" {
			fmt.Println("Error! Was expecting username, password and port  :(")
			fmt.Println("Have a look at doggo connect --help")
			os.Exit(1)
		} else {
			evansctl.RunEvansWithToken(username, password, port)
		}

	},
}

func init() {
	rootCmd.AddCommand(connectCmd)

	connectCmd.Flags().StringP("username", "U", "", "Target user's username")
	connectCmd.Flags().StringP("password", "P", "", "Target user's password")
	connectCmd.Flags().StringP("port", "p", "", "gRPC targeted port")
}
