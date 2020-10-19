package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	authpb "primitivofr/owly/server/auth/authpb"
	"runtime"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var rootCmd = &cobra.Command{
	Use:   "doggo",
	Short: "Doggo allows a simplier use of evans by automatically passing headers. Local use only for now",
	Long:  `by @AppliNH :)`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to doggo 🐶")
		if len(args) < 3 {
			fmt.Println("I need three arguments, in this order:")
			fmt.Println("username, password, port")
		}
		RunEvansWithToken(args[0], args[1], args[2])

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func loginUserEvans(username string, password string) string {
	conn, err := grpc.Dial("localhost:50054", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	c := authpb.NewAuthServiceClient(conn)
	req := &authpb.LoginUserRequest{
		Username: username,
		Password: password,
	}
	res, err := c.LoginUser(context.Background(), req)
	if err != nil {
		panic(err)
	}

	return res.Result.AccessToken

}

func RunEvansWithToken(username string, password string, port string) {
	token := loginUserEvans(username, password)
	log.Println(token)
	var whichOrWhere string
	if runtime.GOOS == "windows" {
		whichOrWhere = "where"

	} else {
		whichOrWhere = "which"
	}
	evansLocationBytes, err := exec.Command(whichOrWhere, "evans").Output()

	if err != nil {
		panic(err)
	}

	var evansLocation string = string(evansLocationBytes)
	evansLocation = strings.ToValidUTF8(evansLocation, "")
	evansLocation = strings.TrimSuffix(evansLocation, "\n")

	err = syscall.Exec(evansLocation, []string{"", "-r", "repl", "--header", `authorization=` + token, "-p", port}, os.Environ())
	if err != nil {
		panic(err)
	}

}
