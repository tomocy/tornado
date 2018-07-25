package tornado

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/cli"
)

const (
	fileName = ".tornado"
)

const (
	statusOK = iota
	statusErr
)

var (
	errNoSuchFile = errors.New("no such file")
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
		"list": getListCommandFactory(getFilePath(fileName)),
	}

	return tornado
}

func getFilePath(fileName string) string {
	dir, err := os.Getwd()
	if err != nil {
		homeDir := os.Getenv("HOME")
		return filepath.Join(homeDir, fileName)
	}

	return filepath.Join(dir, fileName)
}

func (t tornado) Run() int {
	code, err := t.cli.Run()
	if err != nil {
		fmt.Println(err)
		return statusErr
	}
	return code
}
