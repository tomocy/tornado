package tornado

import "testing"

func TestRunSuccessfully(t *testing.T) {
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
