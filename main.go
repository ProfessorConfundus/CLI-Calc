package main

import (
	"fmt"

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
}
