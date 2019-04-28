package setting

import (
	"context"
	"net/http"

	endpoint "github.com/go-kit/kit/endpoint"
)

func MakeSayHelloHandler(e endpoint.Endpoint, f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	clientProxy := SayHelloProxy(e)
	return func(res http.ResponseWriter, req *http.Request) {
		// TODO implement the business logic of SayHello
		var ctx context.Context
		var s *TestAlias
		// RESULT0:rs string
		// RESULT1:err string
		clientProxy(ctx, s)
		f(res, req)
	}
}
func MakeDeleteuserHandler(e endpoint.Endpoint, f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	clientProxy := DeleteuserProxy(e)
	return func(res http.ResponseWriter, req *http.Request) {
		// TODO implement the business logic of Deleteuser
		var ctx context.Context
		var s string
		// RESULT0:rs string
		// RESULT1:err string
		clientProxy(ctx, s)
		f(res, req)
	}
}
