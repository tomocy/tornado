package tornado

type Tornado interface {
	Run() int
}

type tornado struct {
}

func (t tornado) Run() int {
	return 0
}
