package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "irgsh-go-cli"
	app.Usage = "CLI tool for managing IRGSH jobs"
	app.Action = func(c *cli.Context) error {
		fmt.Println("pong")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
