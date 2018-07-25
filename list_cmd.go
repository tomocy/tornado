package tornado

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

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
	todos, err := lc.getTodos()
	if err != nil {
		if err == errNoSuchFile {
			fmt.Println("there is nothing to do!")
			return statusOK
		}

		fmt.Println(err)
		return statusErr
	}

	if len(todos) <= 0 {
		fmt.Println("there is nothing to do!")
		return statusOK
	}
	for _, todo := range todos {
		fmt.Println(todo)
	}
	return statusOK
}

func (lc listCommand) getTodos() ([]string, error) {
	_, err := os.Stat(fileName)
	if err != nil {
		return nil, errNoSuchFile
	}
	f, err := os.Open(lc.fileName)
	if err != nil {
		return nil, fmt.Errorf("cannot open file %s\n%s", lc.fileName, err)
	}
	defer f.Close()

	todos := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		todos = append(todos, scanner.Text())
	}

	return todos, nil
}
