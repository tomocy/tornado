package tornado

import (
	"fmt"

	"github.com/mitchellh/cli"
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
	appName string
	version string
	cli     *cli.CLI
}

func newTornado() *tornado {
	tornado := &tornado{
		appName: "tornado",
		version: "1.0.0",
	}
	tornado.cli = cli.NewCLI(tornado.appName, tornado.version)
	tornado.cli.Commands = map[string]cli.CommandFactory{
		"list": listCommandFacotry,
	}

	return tornado
}

func (t tornado) Run() int {
	code, err := t.cli.Run()
	if err != nil {
		fmt.Println(err)
		return 1
	}
	return code
}
