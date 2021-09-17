package main

// this comment is for a test commit
// another test

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var modes = map[int8]string{
	1: "1: basic arithmics",
	2: "2: comparisons",
	3: "3: fraction, decimal & percentage conversion",
	4: "4: trigonometry",
}

func main() {
	var app = &cli.App{
		Name:  "cli-calc",
		Usage: "An interactive calculator based entirely in the Command-Line Interface.",
		Action: func(c *cli.Context) error {
			fmt.Println("Welcome to PC's Command Line Calculator! \nWhat would you like to do?")
			for _, mode := range modes {
				fmt.Println(mode)
			}
			var chosenMode int8
			var _, err = fmt.Scan(&chosenMode)
			if err != nil {
				fmt.Println("Invalid input type.")
				return nil
			}
			switch chosenMode {
			case 1:
				fmt.Println("Basic arithmics mode selected.")
				arithmicsInit()
			case 2:
				fmt.Println("Comparisons mode selected.")
				comparisonsInit()
			case 3:
				fmt.Println("Fractions, decimals & percentages mode selected.")
				fdpInit()
			case 4:
				fmt.Println("Trigonometry mode selected.")
				trigonometryInit()
			default:
				fmt.Println("That's not a mode that currently exists.")
			}
			return nil
		},
	}
	var err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
