package test

import "context"

type TestStruc struct {
	Test1 string
	Test2 int
	Test3 bool
	Test4 []string
}

// SettingService describes the service.
type TestService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	SayHello(ctx context.Context, s *TestStrucAlias) (rs string, err string)
	Deleteuser(ctx context.Context, s string) (rs string, err string)
	TestArray(ctx context.Context, s []*TestStrucAlias) (rs string, err string)
}

type basicTestService struct{}

func (b *basicTestService) SayHello(ctx context.Context, s *TestStrucAlias) (rs string, err string) {
	// TODO implement the business logic of SayHello
	rs = "hello " + s.Test1
	return rs, err
}
func (b *basicTestService) Deleteuser(ctx context.Context, s string) (rs string, err string) {
	// TODO implement the business logic of Deleteuser
	return rs, err
}

// NewBasicTestService returns a naive, stateless implementation of TestService.
func NewBasicTestService() TestService {
	return &basicTestService{}
}

// New returns a TestService with all of the expected middleware wired in.
func NewService() TestService {
	return NewBasicTestService()
}

func (b *basicTestService) TestArray(ctx context.Context, s []*TestStrucAlias) (rs string, err string) {
	// TODO implement the business logic of TestArray
	return rs, err
}
