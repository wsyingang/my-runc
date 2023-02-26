package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "mydocker"
	app.Usage = "demo for docker"
	app.Commands = []cli.Command{
		initCommand, runCommand,
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatalf("run fail %v", err)
	}
}
