package web

import (
	"context"
	"github.com/LongMarch7/higo/middleware/hystrix"
	"github.com/LongMarch7/higo/middleware/ratelimit"
	"net/http"
	"github.com/LongMarch7/higo/base"
	"github.com/LongMarch7/higo/util/define"
	"github.com/LongMarch7/higo/util/error/html"
	"github.com/LongMarch7/higo/util/error/json"
	endpoint "github.com/go-kit/kit/endpoint"
	"strings"
)

func MakeHtmlCallHandler(e endpoint.Endpoint, pattern string) func(http.ResponseWriter, *http.Request) {
	ratelimit.NewLimiter().AddRateLimitForMethod(pattern)
	hystrix.NewHystrix().AddHystrixForMethod(pattern)
	clientProxy := HtmlCallProxy(e)
	return func(res http.ResponseWriter, req *http.Request) {
		// TODO implement the business logic of HtmlCall
		ctx := req.Context()
		ctx = context.WithValue(ctx, define.ReqPatternName, pattern)
		rs, err := clientProxy(ctx, pattern)
		base.SetCookie(ctx, res)
		if err == nil {
			base.HtmlRender(res, []byte(rs))
		} else {
			errorHtml :=strings.Replace(html.NotFoundError,"{[.content]}","网络繁忙", -1)
			base.HtmlRender(res, []byte(errorHtml))
		}
	}
}
func MakeApiCallHandler(e endpoint.Endpoint, pattern string) func(http.ResponseWriter, *http.Request) {
	ratelimit.NewLimiter().AddRateLimitForMethod(pattern)
	hystrix.NewHystrix().AddHystrixForMethod(pattern)
	clientProxy := ApiCallProxy(e)
	return func(res http.ResponseWriter, req *http.Request) {
		// TODO implement the business logic of ApiCall
		ctx := req.Context()
		ctx = context.WithValue(ctx, define.ReqPatternName, pattern)
		rs, err := clientProxy(ctx, pattern)
		base.SetCookie(ctx, res)
		if err == nil {
			base.JsonRender(res, []byte(rs))
		} else {
			base.JsonRender(res, []byte(json.NotFoundError))
		}
	}
}
