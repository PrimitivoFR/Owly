package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "doggo",
	Short: "Doggo is your best friend for the Owly project",
	Long: `
	Doggo is your best friend for the Owly project`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("No command passed :(")
		fmt.Println("Use doggo --help to see more")

	},
}

func init() {
	fmt.Println("Welcome to Doggo 🐶")
}

// Execute executes the CLI
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
