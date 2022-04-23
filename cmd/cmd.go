package main

import (
	"fmt"
	"log"
	"os"

	"github.com/g-lib/passlib/pwd"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "passlib-cli",
		Usage: "CLI tool for passlib",
		Description: `Passlib, a password hashing library which provides cross-platform implementations 
of over 30 password hashing algorithms, 
as well as a framework for managing existing password hashes.!`,
		Authors: []*cli.Author{
			{
				Name:  "Tacey Wong",
				Email: "xinyong.w@gmail.com",
			},
			{
				Name: "All Contributors",
			},
		},
		Copyright:       "(c) 2022 - Forever Tacey Wong under MIT License",
		HideHelpCommand: true,
	}
	app.Action = func(c *cli.Context) error {
		fmt.Printf("%s - %s\n", app.Name, app.Usage)
		fmt.Printf("Please type `%s -h/--help` for the help of usage\n", app.Name)
		return nil
	}
	app.Commands = []*cli.Command{
		{
			Name: "pwd",
			// Aliases: []string{"c"},
			Usage:  "Generate one or more random passwords",
			Action: pwdCMD,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func pwdCMD(c *cli.Context) error {
	fmt.Println(pwd.GenWordQuickly(6))
	return nil

}
