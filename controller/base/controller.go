package base

import (
    "context"
    "encoding/json"
    "github.com/LongMarch7/higo/util/cookie"
    "github.com/LongMarch7/higo/util/define"
    "google.golang.org/grpc/grpclog"
    "google.golang.org/grpc/metadata"
    "net/url"
)

type Controller interface {
}

type Parameter struct {
    Method           string
    MuxParams        map[string]string
    GetParams        url.Values
    PostFormParams   url.Values
    PostJsonParams   map[string] interface{}
    Cookie           cookie.Cookie
    ContentType      string
    Pattern          string
}

var controller = make(map[string] Controller, 32)

func AddController(name string, obj Controller){
    controller[name] = obj
}

func GetController(name string) (Controller, bool){
    value, ok := controller[name]
    return value, ok
}

func NewParameter(ctx context.Context) context.Context{
    parameter :=&Parameter{
        MuxParams: make(map[string]string),
        GetParams: make(url.Values),
        PostFormParams: make(url.Values),
        PostJsonParams: make(map[string]interface{}),
        Cookie: cookie.Cookie{},
    }
    md, ok := metadata.FromIncomingContext(ctx)
    if ok {
        if value,ok1 := md[define.MuxParamsName]; ok1{
            if ret := json.Unmarshal([]byte(value[0]), &parameter.MuxParams); ret != nil{
                grpclog.Error("Parsing mux error :" + value[0])
            }
        }
        if value,ok1 := md[define.GetParamsName]; ok1{
            if ret := json.Unmarshal([]byte(value[0]), &parameter.GetParams); ret != nil{
                grpclog.Error("Parsing get error :" + value[0])
            }
        }
        if value,ok1 := md[define.PostParamsFormName]; ok1{
            if ret := json.Unmarshal([]byte(value[0]), &parameter.PostFormParams); ret != nil{
                grpclog.Error("Parsing form post error:" + value[0])
            }
        }
        if value,ok1 := md[define.PostParamsJsonName]; ok1{
            if ret := json.Unmarshal([]byte(value[0]), &parameter.PostJsonParams); ret != nil{
                grpclog.Error("Parsing json post error:" + value[0])
            }
        }
        if value,ok1 := md[define.ReqCookieName]; ok1{
            //if ret := json.Unmarshal([]byte(value[0]), &parameter.Cookie); ret != nil{
            //    grpclog.Error("Parsing cookie error:" + value[0])
            //}
            cookie.UnMarshal(value[0], &parameter.Cookie)
        }
        if value,ok1 := md[define.ReqMethodName]; ok1{
            parameter.Method  = value[0]
        }
        if value,ok1 := md[define.ContentType]; ok1{
            parameter.ContentType  = value[0]
        }
        if value,ok1 := md[define.ReqPatternName]; ok1{
            ctx = context.WithValue(ctx,define.ReqPatternName, value[0])
            parameter.Pattern = value[0]
        }
    }
    grpclog.Info("request_data ---||---",parameter)
    return context.WithValue(ctx,define.ParameterName, parameter)
}