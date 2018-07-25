package tornado

import clilib "github.com/mitchellh/cli"

const (
	appName = "tornado"
	version = "1.0.0"
)

type Tornado interface {
	Run() int
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
	return 0
}
