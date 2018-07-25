package tornado

import "github.com/mitchellh/cli"

type listCommand struct {
	fileName string
}

func getListCommandFactory(fileName string) func() (cli.Command, error) {
	listCmd := listCommand{
		fileName: fileName,
	}

	return func() (cli.Command, error) {
		return listCmd, nil
	}
}

func (lc listCommand) Synopsis() string {
	return "list all todos"
}

func (lc listCommand) Help() string {
	return "list"
}

func (lc listCommand) Run(args []string) int {
	return 0
}
