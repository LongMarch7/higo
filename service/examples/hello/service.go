package hello

import (
	"bytes"
	"context"
	"github.com/LongMarch7/higo/view"
)

// HelloService describes the service.
type HelloService interface {
	// Add your methods here
	HelloWorld(ctx context.Context, s string) (rs string, err string)
}

type basicHelloService struct{}

func (b *basicHelloService) HelloWorld(ctx context.Context, s string) (rs string, err string) {
	// TODO implement the business logic of HelloWorld
	out := &bytes.Buffer{}
	data := make(map[string]interface{})
	data["name"] = s
	view.Template.Render(out,"examples/hello", data)
	return out.String(), ""
}

// NewBasicHelloService returns a naive, stateless implementation of HelloService.
func NewBasicHelloService() HelloService {
	return &basicHelloService{}
}

// New returns a HelloService with all of the expected middleware wired in.
func NewService() HelloService {
	return NewBasicHelloService()
}
