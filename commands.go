package main

import (
	"github.com/codegangsta/cli"
	"log"
	"os"
)

var Commands = []cli.Command{
	commandGenerate,
	commandList,
}

var commandGenerate = cli.Command{
	Name:  "generate",
	Usage: "",
	Description: `
`,
	Action: doGenerate,
}

var commandList = cli.Command{
	Name:  "list",
	Usage: "",
	Description: `
`,
	Action: doList,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doGenerate(c *cli.Context) {
	println("generate")
}
func doList(c *cli.Context) {
	println("list")
}
