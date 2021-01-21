package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var owlyPath string

var doggoDockerComposeCmd []string

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the architecture with the specified mode.",
	Long: `
	Run the architecture with the specified mode`,
	Run: func(cmd *cobra.Command, args []string) {

		if cmd.Root().Use == "doggo" {
			// If command is "doggo [sub-command]" then show help
			cmd.Help()
			os.Exit(0)
		} else {
			// source : https://www.twitchquotes.com/copypastas/1027
			asciiArtRun :=
				`
┈┈┈┈╱▏┈┈┈┈┈╱▔▔▔▔╲┈┈┈┈┈
┈┈┈┈▏▏┈┈┈┈┈▏╲▕▋▕▋▏┈┈┈┈
┈┈┈┈╲╲┈┈┈┈┈▏┈▏┈▔▔▔▆┈┈┈
┈┈┈┈┈╲▔▔▔▔▔╲╱┈╰┳┳┳╯┈┈┈
┈┈╱╲╱╲▏┈┈┈┈┈┈▕▔╰━╯┈┈┈┈
┈┈▔╲╲╱╱▔╱▔▔╲╲╲╲┈┈┈┈┈┈┈
┈┈┈┈╲╱╲╱┈┈┈┈╲╲▂╲▂┈┈┈┈┈
┈┈┈┈┈┈┈┈┈┈┈┈┈╲╱╲╱┈┈┈┈┈
`
			fmt.Println(asciiArtRun)
			doggoDockerComposeCmd = append(doggoDockerComposeCmd, "up")
		}
	},
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops the architecture with the specified mode.",
	Long: `
	Stops the architecture with the specified mode`,
	Run: func(cmd *cobra.Command, args []string) {

		if cmd.Root().Use == "doggo" {
			// If command is "doggo [sub-command]" then show help
			cmd.Help()
			os.Exit(0)
		} else {

			// source : https://www.asciiart.eu/animals/dogs
			asciiArtStop :=
				`
			__      _
			o'')}____//
			 -_/      )
			 (_(_/-(_/
`
			fmt.Println(asciiArtStop)
			doggoDockerComposeCmd = append(doggoDockerComposeCmd, "stop")

		}
	},
}

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart a service of the architecture, or the whole architecture, with the specified mode.",
	Long: `
	Restart a service of the architecture, or the whole architecture, with the specified mode`,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Root().Use == "doggo" {
			// If command is "doggo [sub-command]" then show help
			cmd.Help()
			os.Exit(0)
		} else {
			doggoDockerComposeCmd = append(doggoDockerComposeCmd, "restart")
		}

	},
}

func init() {
	owlyPath = os.Getenv("OWLY_PATH")

	if owlyPath == "" {
		panic("You need to set your OWLY_PATH env variable that points to the OWLY project directory. \n Don't forget to remove the '/' at the end of the path.")
	}

	// Needed, because the last one you add with AddCommand becomes the parent command.
	// ==> https://stackoverflow.com/questions/60420572/add-subcommand-in-go-cobra-with-multiple-parents

	modCmdDev.Flags().StringP("service", "s", "", "Targeted service")
	modCmdTest.Flags().StringP("service", "s", "", "Targeted service")
	modCmdLocalProd.Flags().StringP("service", "s", "", "Targeted service")

	modCmdDev.Flags().StringP("args", "a", "", "Additional arguments to provide")
	modCmdTest.Flags().StringP("args", "a", "", "Additional arguments to provide")
	modCmdLocalProd.Flags().StringP("args", "a", "", "Additional arguments to provide")

	var runCmdDev = *modCmdDev
	var stopCmdDev = *modCmdDev
	var restartCmdDev = *modCmdDev

	var runCmdTest = *modCmdTest
	var stopCmdTest = *modCmdTest
	var restartCmdTest = *modCmdTest

	var runCmdLocalProd = *modCmdLocalProd
	var stopCmdLocalProd = *modCmdLocalProd
	var restartCmdLocalProd = *modCmdLocalProd

	runCmd.AddCommand(&runCmdDev)
	stopCmd.AddCommand(&stopCmdDev)
	restartCmd.AddCommand(&restartCmdDev)

	runCmd.AddCommand(&runCmdTest)
	stopCmd.AddCommand(&stopCmdTest)
	restartCmd.AddCommand(&restartCmdTest)

	runCmd.AddCommand(&runCmdLocalProd)
	stopCmd.AddCommand(&stopCmdLocalProd)
	restartCmd.AddCommand(&restartCmdLocalProd)

	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(stopCmd)
	rootCmd.AddCommand(restartCmd)

}
