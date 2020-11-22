package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var dockerCleanCmd = &cobra.Command{
	Use:   "docker_clean",
	Short: "Clean volume, container, system cache out of Docker.",
	Long: `
	Clean volume, container, system cache out of Docker.`,
	Run: func(cmd *cobra.Command, args []string) {

		cmdEx := exec.Command("docker", "system", "prune", "-f")
		cmdEx.Stdin = os.Stdin
		cmdEx.Stdout = os.Stdout
		cmdEx.Stderr = os.Stderr
		cmdEx.Run()

		containers := []string{"docker", "container", "prune", "-f"}
		volumes := []string{"docker", "volume", "prune", "-f"}
		cmds := [][]string{containers, volumes}

		for _, i := range cmds {
			cmdEx.Args = i
			cmdEx.Run()
		}

	},
}

func init() {
	rootCmd.AddCommand(dockerCleanCmd)
}
