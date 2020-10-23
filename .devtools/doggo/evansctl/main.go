package evansctl

import (
	grpcAuth "doggo/grpccli/auth"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
)

// RunEvansWithToken opens an evans connection with an authorization header
// containing the token which belongs to the select user
func RunEvansWithToken(username string, password string, port string) {
	res := grpcAuth.LoginUser(username, password)
	token := res.Result.AccessToken
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

	if runtime.GOOS != "windows" {
		// This won't work in windows
		err = syscall.Exec(evansLocation, []string{"", "-r", "repl", "--header", `authorization=` + token, "-p", port}, os.Environ())
		if err != nil {
			panic(err)
		}
	} else {
		cmd := exec.Command("cmd.exe", "/C", "start", evansLocation, "-r", "repl", "--header", `authorization=`+token, "-p", port)
		if err := cmd.Run(); err != nil {
			log.Println("Error:", err)
		}
	}

}
