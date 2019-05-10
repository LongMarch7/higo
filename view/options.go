package view

type viewOpt struct {
	dir string
}
type VOption func(o *viewOpt)

func Dir(dir string) VOption {
	return func(o *viewOpt) {
		o.dir = dir
	}
}
