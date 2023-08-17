package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
	"github.com/vinybergamo/cli-skeleton/command"
	"github.com/vinybergamo/docker-image-labeler/commands"
)

var AppName = "docker-image-labeler"

var Version string

func main() {
	os.Exit(Run(os.Args[1:]))
}

func Run(args []string) int {
	commandMeta := command.SetupRun(AppName, Version, args)
	commandMeta.Ui = command.HumanZerologUiWithFields(commandMeta.Ui, make(map[string]interface{}, 0))
	c := cli.NewCLI(AppName, Version)
	c.Args = os.Args[1:]
	c.Commands = command.Commands(commandMeta, Commands)
	exitCode, err := c.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
		return 1
	}

	return exitCode
}

func Commands(meta command.Meta) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"relabel": func() (cli.Command, error) {
			return &commands.RelabelCommand{Meta: meta}, nil
		},
		"version": func() (cli.Command, error) {
			return &command.VersionCommand{Meta: meta}, nil
		},
	}
}
