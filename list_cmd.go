package tornado

import "github.com/mitchellh/cli"

type listCommand struct {
}

var (
	listCommandFacotry = func() (cli.Command, error) {
		return listCommand{}, nil
	}
)

func (lc listCommand) Synopsis() string {
	return "list all todos"
}

func (lc listCommand) Help() string {
	return "list"
}

func (lc listCommand) Run(args []string) int {
	return 0
}
