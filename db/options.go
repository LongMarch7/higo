package db

type dbOpt struct {
	name          string
	dialect       string
	args          string
	show          bool
	level         string
	maxIdleConns  int
	maxOpenConns  int
}
type DOption func(o *dbOpt)

func Name(name string) DOption {
	return func(o *dbOpt) {
		o.name = name
	}
}

func Dialect(dialect string) DOption {
	return func(o *dbOpt) {
		o.dialect = dialect
	}
}

func Show(show bool) DOption {
	return func(o *dbOpt) {
		o.show = show
	}
}

func Level(level string) DOption {
	return func(o *dbOpt) {
		o.level = level
	}
}


func Args(args string) DOption {
	return func(o *dbOpt) {
		o.args = args
	}
}


func MaxIdleConns(mxIdleConns int) DOption {
	return func(o *dbOpt) {
		o.maxIdleConns = mxIdleConns
	}
}

func MaxOpenConns(maxOpenConns int) DOption {
	return func(o *dbOpt) {
		o.maxOpenConns = maxOpenConns
	}
}


