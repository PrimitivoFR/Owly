package main

// This file is a slightly (because we need it to work with our own services)
// modified version of the one given as example in the ghz repo.

import (
	"fmt"
	"os"

	"github.com/bojand/ghz/printer"
	"github.com/bojand/ghz/runner"
)

func main() {
	report, err := runner.Run(
		"user.UserService.LoginUser",
		"localhost:50051",
		runner.WithProtoFile("../../protofiles/user.proto", []string{}),
		runner.WithDataFromFile("../evans-python/json/login.json"),
		runner.WithInsecure(true),
	)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	printer := printer.ReportPrinter{
		Out:    os.Stdout,
		Report: report,
	}

	printer.Print("pretty")
}
