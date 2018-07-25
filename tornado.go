package tornado

import (
	"fmt"

	clilib "github.com/mitchellh/cli"
)

const (
	appName = "tornado"
	version = "1.0.0"
)

type Tornado interface {
	Run() int
}

func New() Tornado {
	return newTornado()
}

type tornado struct {
	cli *clilib.CLI
}

func newTornado() *tornado {
	cli := clilib.NewCLI(appName, version)
	cli.Commands = map[string]clilib.CommandFactory{
		"list": listCommandFacotry,
	}
	return &tornado{
		cli: cli,
	}
}

func (t tornado) Run() int {
	code, err := t.cli.Run()
	if err != nil {
		fmt.Println(err)
		return 1
	}
	return code
}
