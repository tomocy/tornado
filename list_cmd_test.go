package tornado

import (
	"reflect"
	"testing"
)

func TestListCommandFacotry(t *testing.T) {
	given := listCommandFacotry
	want := getListCommand()
	have, err := given()
	if err != nil {
		t.Error("listCommandFactory returns err\n")
		t.Error(err, "\n")
	}
	if !reflect.DeepEqual(have, want) {
		t.Errorf("have %#v, but want %#v\n", have, want)
		t.Error("fix something wrong in listCommandFactory")
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

func TestListCommandRunSuccessfully(t *testing.T) {
	given := getListCommand()
	then := struct {
		in   []string
		want int
	}{
		in:   []string{"a"},
		want: 0,
	}
	have := given.Run(then.in)
	if have != then.want {
		t.Errorf("have %#v, but want %#v\n", have, then.want)
		t.Errorf("it should be success when in is %#v", then.in)
	}
}

func getListCommand() listCommand {
	return listCommand{}
}
