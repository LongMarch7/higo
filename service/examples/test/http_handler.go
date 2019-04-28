package test

import (
	"encoding/json"
	base "github.com/LongMarch7/higo/base"
	endpoint "github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
	"net/http"
)

func MakeSayHelloHandler(e endpoint.Endpoint) func(http.ResponseWriter, *http.Request) {
	clientProxy := SayHelloProxy(e)
	return func(res http.ResponseWriter, req *http.Request) {
		// TODO implement the business logic of SayHello
		baseContext := base.NewContext(res, req)
		ctx := baseContext.Request().Context()
		s := new(TestStrucAlias)
		// RESULT0:rs string
		// RESULT1:err string
		vars := mux.Vars(req)
		s.Test1 = vars["serviceName"]
		rs, err := clientProxy(ctx, s)
		var ret = map[string]string{"serviceName":rs,"error":err}
		response, _ := json.Marshal(ret)
		baseContext.JsonRender(response)
	}
}
func MakeDeleteuserHandler(e endpoint.Endpoint) func(http.ResponseWriter, *http.Request) {
	clientProxy := DeleteuserProxy(e)
	return func(res http.ResponseWriter, req *http.Request) {
		// TODO implement the business logic of Deleteuser
		baseContext := base.NewContext(res, req)
		ctx := baseContext.Request().Context()
		var s string
		// RESULT0:rs string
		// RESULT1:err string
		clientProxy(ctx, s)
	}
}
func MakeTestArrayHandler(e endpoint.Endpoint) func(http.ResponseWriter, *http.Request) {
	clientProxy := TestArrayProxy(e)
	return func(res http.ResponseWriter, req *http.Request) {
		// TODO implement the business logic of TestArray
		baseContext := base.NewContext(res, req)
		ctx := baseContext.Request().Context()
		s := make([]*TestStrucAlias, 2)
		// RESULT0:rs string
		// RESULT1:err string
		clientProxy(ctx, s)
	}
}
