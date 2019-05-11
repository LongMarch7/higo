package web

import (
	"context"
	"github.com/LongMarch7/higo/controller"
)

// WebService describes the service.
type WebService interface {
	// Add your methods here
	HtmlCall(ctx context.Context, pattern string) (rs string, err string)
	ApiCall(ctx context.Context, pattern string) (rs string, err string)
}
type basicWebService struct{}

func (b *basicWebService) HtmlCall(ctx context.Context, pattern string) (rs string, err string) {
	// TODO implement the business logic of HtmlCall
	return controller.ControllerCall(ctx,pattern)
}
func (b *basicWebService) ApiCall(ctx context.Context, pattern string) (rs string, err string) {
	// TODO implement the business logic of ApiCall
	return controller.ControllerCall(ctx,pattern)
}

// NewBasicWebService returns a naive, stateless implementation of WebService.
func NewBasicWebService() WebService {
	return &basicWebService{}
}

// New returns a WebService with all of the expected middleware wired in.
func NewService() WebService {
	return NewBasicWebService()
}
