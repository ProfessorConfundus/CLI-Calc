package main

/*
This is a calculator based entirely in the Command Line Interface. Made by ProfessorConfundus @ https://github.com/ProfessorConfundus/
*/

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
				arithmicsInit(c.Args())
			case 2:
				fmt.Println("Comparisons mode selected.")
				comparisonsInit(c.Args())
			case 3:
				fmt.Println("Fractions, decimals & percentages mode selected.")
				fdpInit(c.Args())
			case 4:
				fmt.Println("Trigonometry mode selected.")
				trigonometryInit(c.Args())
			default:
				fmt.Println("That's not a mode that currently exists.")
			}
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "arithmics",
				Aliases: []string{"a", "1"},
				Usage:   "Quickly compute arithmic problems without using the interactive mode.",
				Action: func(c *cli.Context) error {
					arithmicsInit(c.Args())
					return nil
				},
			},
			{
				Name:    "comparisons",
				Aliases: []string{"c", "2"},
				Usage:   "Quickly compare values without using the interactive mode.",
				Action: func(c *cli.Context) error {
					comparisonsInit(c.Args())
					return nil
				},
			},
			{
				Name:    "fdp",
				Aliases: []string{"2"},
				Usage:   "Quickly compute fractions, decimals or percentages without using the interactive mode.",
				Action: func(c *cli.Context) error {
					fdpInit(c.Args())
					return nil
				},
			},
			{
				Name:    "trigonometry",
				Aliases: []string{"t", "4"},
				Usage:   "Quickly compute trigonometry-related problems without using the interactive mode.",
				Action: func(c *cli.Context) error {
					trigonometryInit(c.Args())
					return nil
				},
			},
		},
	}
	var err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
