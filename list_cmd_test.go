package tornado

import (
	"reflect"
	"testing"
)

func TestGetListCommandFacotry(t *testing.T) {
	then := getMakeListCommandFacotryTestCase()
	have, err := getListCommandFactory(then.in)()
	if err != nil {
		t.Error("makeListCommandFactory return err\n")
		t.Error(err)
	}
	if !reflect.DeepEqual(have, then.want) {
		t.Errorf("have %#v, but want %#v\n", have, then.want)
		t.Error("set fileName to listCommand in makeListCommandFactory")
	}
}

func getMakeListCommandFacotryTestCase() struct {
	in   string
	want listCommand
} {
	fileName := ".test"
	return struct {
		in   string
		want listCommand
	}{
		in: fileName,
		want: listCommand{
			fileName: fileName,
		},
	}
}

func TestListCommandSynophis(t *testing.T) {
	given := getListCommand()
	want := "list all todos"
	have := given.Synopsis()
	if have != want {
		t.Errorf("have %#v, but want %#v\n", have, want)
		t.Errorf("edit message to be returned in cmdList.Synopsis")
	}
}

func TestListCommandHelp(t *testing.T) {
	given := getListCommand()
	want := "list"
	have := given.Help()
	if have != want {
		t.Errorf("have %#v, but want %#v\n", have, want)
		t.Errorf("edit message to be returned in cmdList.Help")
	}
}

func TestListCommandRun(t *testing.T) {
	given := getListCommand()
	then := getListcommandRunTestCase()
	have := given.Run(then.in)
	if have != then.want {
		t.Errorf("have %#v, but want %#v\n", have, then.want)
		t.Errorf("it should be success when in is %#v", then.in)
	}
}

func getListcommandRunTestCase() struct {
	in   []string
	want int
} {
	return struct {
		in   []string
		want int
	}{
		in:   []string{},
		want: statusOK,
	}
}

func getListCommand() listCommand {
	return listCommand{}
}
