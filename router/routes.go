package router

import (
    "github.com/LongMarch7/higo/app"
    "github.com/gorilla/mux"
    "google.golang.org/grpc/grpclog"
    "net/http"
    "strings"
    "sync"
)

type Routs struct{
    Mehods    string
    Pattern   string
    Handler   func(http.ResponseWriter, *http.Request)
}

type Router struct {
    *mux.Router
    routes  []Routs
    Cli     *app.Client
}

var supported = "GET POST PUT PATCH TRACE PATCH DELETE HEAD OPTIONS"

var onceAction sync.Once
var router *Router
func NewRouter() *Router {
    onceAction.Do(func() {
        router = &Router{
            Router: mux.NewRouter(),
        }
    })
    return router
}

func (r* Router)init(routes []Routs) {
    for _, value := range routes {
        route := r.HandleFunc(value.Pattern, value.Handler)
        upMethod := strings.ToUpper(value.Mehods)
        methodSlice := strings.Split(upMethod, "|")
        var methods []string
        for _, v := range methodSlice {
            if !strings.Contains(supported, v) {
                grpclog.Error("not support method " + v)
                continue
            }
            methods = append(methods, v)
        }
        if len(methods) > 0 {
            route.Methods(methods...)
        }
    }
}

func (r* Router)Add(routes []Routs){
    r.init(routes)
    r.routes = append(r.routes, routes...)
}