package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/daviddengcn/go-colortext"

	"github.com/urfave/cli"
)

var magStack = "mag_stack"

var (
	errNoArg          = errors.New("No argument provided for this command")
	errReadTemplate   = errors.New("Could not read template. Please contact author")
	errCreatingConfig = errors.New("Could not create config.json file")
	errNgNew          = errors.New("Could not create angular project")
	errNgRename       = errors.New("Could not remove angular project to angular")
	// errCreatingFile   = errors.New("Could not create file")
)

func errCreatingFile(filename string) error {
	return errors.New("Could not create file: " + filename)
}
func errWritingFile(filename string) error {
	return errors.New("Could not write file: " + filename)
}

func main() {
	app := cli.NewApp()
	app.Name = "mag-stack"
	app.Description = "a cli for creating a MongoDB - Angular 2+ - Go app"
	// app.CustomAppHelpTemplate
	app.Commands = commands
	app.EnableBashCompletion = true
	app.ExitErrHandler = errorHandler
	app.Run(os.Args)
}

var commands = []cli.Command{
	{
		Name:    "new",
		Aliases: []string{"n"},
		Usage:   "create new project",
		Action:  newP,
	},
	{
		Name:    "build",
		Aliases: []string{"b"},
		Usage:   "builds the angular and go files",
		Action:  buildAction,
	},
}

var errorHandler cli.ExitErrHandlerFunc = func(c *cli.Context, e error) {
	ct.Foreground(ct.Red, true)
	if e == nil {
		fmt.Println("no command specified")
	} else {
		fmt.Println(e.Error())
	}
	ct.ResetColor()
}
