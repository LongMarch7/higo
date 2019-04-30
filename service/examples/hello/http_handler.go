package hello

import (
	base "github.com/LongMarch7/higo/base"
	endpoint "github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
	"net/http"
)

func MakeHelloWorldHandler(e endpoint.Endpoint) func(http.ResponseWriter, *http.Request) {
	clientProxy := HelloWorldProxy(e)
	return func(res http.ResponseWriter, req *http.Request) {
		// TODO implement the business logic of HelloWorld
		baseContext := base.NewContext(res, req)
		ctx := baseContext.Request().Context()
		var s string
		vars := mux.Vars(req)
		s = vars["serviceName"]
		// RESULT0:rs string
		// RESULT1:err string
		rs, err :=clientProxy(ctx, s)
		if len(err) == 0 {
			baseContext.HtmlRender([]byte(rs))
		}else{
			baseContext.JsonRender([]byte("not found"))
		}
	}
}
