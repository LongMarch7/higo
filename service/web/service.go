package web

import (
	"context"

	"github.com/LongMarch7/higo/controller"
	"github.com/LongMarch7/higo/util/define"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// WebService describes the service.
type WebService interface {
	// Add your methods here
	HtmlCall(ctx context.Context, pattern string) (rs string, err error)
	ApiCall(ctx context.Context, pattern string) (rs string, err error)
}
type basicWebService struct{}

func (b *basicWebService) HtmlCall(ctx context.Context, pattern string) (rs string, err error) {
	// TODO implement the business logic of HtmlCall
	header := metadata.Pairs(define.ResTypeName, "html")
	grpc.SetHeader(ctx, header)
	return controller.ControllerCall(ctx, pattern)
}
func (b *basicWebService) ApiCall(ctx context.Context, pattern string) (rs string, err error) {
	// TODO implement the business logic of ApiCall
	header := metadata.Pairs(define.ResTypeName, "json")
	grpc.SetHeader(ctx, header)
	return controller.ControllerCall(ctx, pattern)
}

// NewBasicWebService returns a naive, stateless implementation of WebService.
func NewBasicWebService() WebService {
	return &basicWebService{}
}

// New returns a WebService with all of the expected middleware wired in.
func NewService() WebService {
	return NewBasicWebService()
}
