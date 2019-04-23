package setting

import (
	"context"
)

type Test struct {
	Test1 string
	Test2 int
	Test3 bool
	Test4 []string
}

// SettingService describes the service.
type SettingService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	SayHello(ctx context.Context, s []*TestAlias) (rs string, err string)
	Deleteuser(ctx context.Context, s string) (rs string, err string)
}
type basicSettingService struct{}

func (b *basicSettingService) SayHello(ctx context.Context, s []*TestAlias) (rs string, err string) {
	// TODO implement the business logic of SayHello
	return rs, err
}
func (b *basicSettingService) Deleteuser(ctx context.Context, s string) (rs string, err string) {
	// TODO implement the business logic of Deleteuser
	return rs, err
}

// NewBasicSettingService returns a naive, stateless implementation of SettingService.
func NewBasicSettingService() SettingService {
	return &basicSettingService{}
}

// New returns a SettingService with all of the expected middleware wired in.
func NewService() SettingService {
	return NewBasicSettingService()
}
