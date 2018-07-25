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

func TestTornadoRunSuccessfully(t *testing.T) {
	given := new(tornado)
	want := getExpectationAsRunningSuccessfully()
	have := given.Run()

	if have != want {
		t.Errorf("have %#v, but want %#v", have, want)
	}
}

func getExpectationAsRunningSuccessfully() int {
	return 0
}
