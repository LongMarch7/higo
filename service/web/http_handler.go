package web

import (
	"github.com/LongMarch7/higo/base"
	"github.com/LongMarch7/higo/util/define"
	"github.com/LongMarch7/higo/util/error/html"
	"github.com/LongMarch7/higo/util/error/json"
	endpoint "github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
	"net/http"
)

func MakeHtmlCallHandler(e endpoint.Endpoint, pattern string) func(http.ResponseWriter, *http.Request) {
	clientProxy := HtmlCallProxy(e)
	return func(res http.ResponseWriter, req *http.Request) {
		// TODO implement the business logic of HtmlCall
		ctx := req.Context()
		ctx = context.WithValue(ctx,define.ReqPatternName,pattern)
		rs, err := clientProxy(ctx, pattern)
		base.SetCookie(ctx,res)
		if len(err) == 0 {
			base.HtmlRender(res, []byte(rs))
		}else{
			base.HtmlRender(res, []byte(html.NotFoundError))
		}
	}
}
func MakeApiCallHandler(e endpoint.Endpoint, pattern string) func(http.ResponseWriter, *http.Request) {
	clientProxy := ApiCallProxy(e)
	return func(res http.ResponseWriter, req *http.Request) {
		// TODO implement the business logic of ApiCall
		ctx := req.Context()
		ctx = context.WithValue(ctx,define.ReqPatternName,pattern)
		rs, err := clientProxy(ctx, pattern)
		base.SetCookie(ctx,res)
		if len(err) == 0 {
			base.JsonRender(res, []byte(rs))
		}else{
			base.JsonRender(res,[]byte(json.NotFoundError))
		}
	}
}
