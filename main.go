package main

// this comment is for a test commit
// another test

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	var app = &cli.App{
		Name:  "cli-calc",
		Usage: "An interactive calculator based entirely in the Command-Line Interface.",
		Action: func(c *cli.Context) error {
			fmt.Println("hello world.")
			return nil
		},
	}
	var err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
