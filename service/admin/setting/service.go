package setting

import (
	"context"
)

// SettingService describes the service.
type SettingService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	SayHello(ctx context.Context, s string) (rs string, err error)
	Deleteuser(ctx context.Context, s string) (rs string, err error)
}

type basicSettingService struct{}

func (b *basicSettingService) SayHello(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of SayHello
	rs = "hello " + s
	return rs, nil
}

// NewBasicSettingService returns a naive, stateless implementation of SettingService.
func NewBasicSettingService() SettingService {
	return &basicSettingService{}
}

// New returns a SettingService with all of the expected middleware wired in.
func NewService() SettingService {
	return NewBasicSettingService()
}

func (b *basicSettingService) Deleteuser(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of Deleteuser
	return rs, err
}
