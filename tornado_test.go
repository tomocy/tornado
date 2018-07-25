package tornado

import (
	"testing"
)

func TestNewTornado(t *testing.T) {
	given := newTornado()
	thens := getNewTornadoTestCases()
	for _, then := range thens {
		_, have := given.cli.Commands[then.in]
		if have != then.want {
			t.Errorf("have %#v, but want %#v\n", have, then.want)
			if then.want {
				t.Errorf("add %#v command\n", then.in)
			} else {
				t.Errorf("remove %#v command\n", then.in)
			}
		}
	}
}

func getNewTornadoTestCases() []struct {
	in   string
	want bool
} {
	return []struct {
		in   string
		want bool
	}{
		{
			in:   "list",
			want: true,
		},
		{
			in:   "unknown",
			want: false,
		},
	}
}

func TestGetFilePath(t *testing.T) {
	then := getGetFilePathTestCase()
	have := getFilePath(then.in)
	if have != then.want {
		t.Errorf("have %#v, but want %#v\n", have, then.want)
		t.Error("fix something wrong in getFullFileName")
	}
}

func getGetFilePathTestCase() struct {
	in   string
	want string
} {
	fileName := ".test"
	return struct {
		in   string
		want string
	}{
		in:   fileName,
		want: "/Users/socialis/vagrant/go/code/src/github.com/tomocy/tornado/" + fileName,
	}
}
func TestTornadoRun(t *testing.T) {
	given := newTornado()
	thens := getTornadoRunTestCases()
	for _, then := range thens {
		given.cli.Args = then.in
		have := given.Run()
		if have != then.want {
			t.Errorf("have %#v, but want %#v", have, then.want)
		}

		// to reset args
		given = newTornado()
	}
}

func getTornadoRunTestCases() []struct {
	in   []string
	want int
} {
	return []struct {
		in   []string
		want int
	}{
		{
			in: []string{
				"list",
			},
			want: 0,
		},
		{
			in:   []string{},
			want: 127,
		},
	}
}
