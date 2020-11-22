package cmd

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var modCmdDev = &cobra.Command{
	Use:   "dev",
	Short: "Interact with the architecture in the dev mode.",
	Run: func(cmd *cobra.Command, args []string) {

		loadFileAndRun(cmd, args, "/.devtools/cmd_docker_dev.doggo")

	},
}

var modCmdTest = &cobra.Command{
	Use:   "test",
	Short: "Interact with the architecture in the test mode.",
	Run: func(cmd *cobra.Command, args []string) {

		loadFileAndRun(cmd, args, "/.devtools/cmd_docker_test.doggo")

	},
}

var modCmdLocalProd = &cobra.Command{
	Use:   "localprod",
	Short: "Interact with the architecture in the localprod mode.",
	Run: func(cmd *cobra.Command, args []string) {

		loadFileAndRun(cmd, args, "/.devtools/cmd_docker_localprod.doggo")

	},
}

func loadFileAndRun(cmd *cobra.Command, args []string, filePath string) {

	// Run previous parent (run, stop, restart,...)
	prev := cmd.Parent()
	cmd.Root().Use = "dev"
	prev.Run(cmd, args)

	dat, err := ioutil.ReadFile(owlyPath + filePath)
	if err != nil {
		panic(err)
	}

	cmdStr := string(dat)
	cmdSlice := strings.Split(cmdStr, " ")

	service, _ := cmd.Flags().GetString("service")
	finalArgs := make([]string, 0)

	if len(service) > 0 {
		finalArgs = []string{service}
	}

	doggoDockerComposeCmd := append(doggoDockerComposeCmd, append(finalArgs, args...)...)
	cmdSlice = append(cmdSlice, doggoDockerComposeCmd...)

	cmdEx := exec.Command(cmdSlice[0], cmdSlice[1:]...)
	cmdEx.Stdin = os.Stdin
	cmdEx.Stdout = os.Stdout
	cmdEx.Stderr = os.Stderr
	cmdEx.Dir = owlyPath
	cmdEx.Run()

}
