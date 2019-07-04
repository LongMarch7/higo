package view

type viewOpt struct {
	dir          string
	delimsLeft   string
	delimsRight  string
}
type VOption func(o *viewOpt)

func Dir(dir string) VOption {
	return func(o *viewOpt) {
		o.dir = dir
	}
}

func DelimsLeft(delimsLeft string) VOption {
	return func(o *viewOpt) {
		o.delimsLeft = delimsLeft
	}
}

func DelimsRight(delimsRight string) VOption {
	return func(o *viewOpt) {
		o.delimsRight = delimsRight
	}
}
