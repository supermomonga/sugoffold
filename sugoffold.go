package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "sugoffold"
	app.Version = Version
	app.Usage = ""
	app.Author = "supermomonga"
	app.Email = "sugoikanari@gmail.com"
	app.Commands = Commands

	app.Run(os.Args)
}
