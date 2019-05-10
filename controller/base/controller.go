package base

import (
    "context"
    "google.golang.org/grpc/metadata"
)

type Controller interface {
}

type Parameter struct {
    Method       string
    MuxParams    string
    GetParams    string
    PostParams   string
    Cookie       string
}

var controller = make(map[string] Controller, 32)

func AddController(name string, obj Controller){
    controller[name] = obj
}

func GetController(name string) Controller{
    return controller[name]
}

func NewParameter(ctx context.Context, method string) context.Context{
    parameter :=new(Parameter)
    parameter.Method = method
    md, ok := metadata.FromIncomingContext(ctx)
    if ok {
        if value,ok1 := md["mux_params"]; ok1{
            parameter.MuxParams  = value[0]
        }
        if value,ok1 := md["get_params"]; ok1{
            parameter.GetParams  = value[0]
        }
        if value,ok1 := md["post_params"]; ok1{
            parameter.PostParams  = value[0]
        }
        if value,ok1 := md["req_cookie"]; ok1{
            parameter.Cookie  = value[0]
        }
    }
    return context.WithValue(ctx,"Parameter", parameter)
}