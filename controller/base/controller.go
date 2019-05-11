package base

import (
    "context"
    "github.com/LongMarch7/higo/util/define"
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

func NewParameter(ctx context.Context) context.Context{
    parameter :=new(Parameter)
    md, ok := metadata.FromIncomingContext(ctx)
    if ok {
        if value,ok1 := md[define.MuxParamsName]; ok1{
            parameter.MuxParams  = value[0]
        }
        if value,ok1 := md[define.GetParamsName]; ok1{
            parameter.GetParams  = value[0]
        }
        if value,ok1 := md[define.PostParamsName]; ok1{
            parameter.PostParams  = value[0]
        }
        if value,ok1 := md[define.ReqCookieName]; ok1{
            parameter.Cookie  = value[0]
        }
        if value,ok1 := md[define.ReqMethodName]; ok1{
            parameter.Method  = value[0]
        }
        if value,ok1 := md[define.ReqPatternName]; ok1{
            ctx = context.WithValue(ctx,define.PatternName, value[0])
        }
    }
    return context.WithValue(ctx,define.ParameterName, parameter)
}