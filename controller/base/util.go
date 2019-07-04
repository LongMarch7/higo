package base

import (
    "bytes"
    "encoding/json"
    "github.com/LongMarch7/higo/util/define"
    "github.com/LongMarch7/higo/view"
    "google.golang.org/grpc/grpclog"
    "context"
)

type HtmlReturn struct{
    Code      int          `json:"code"`
    Msg       string       `json:"msg"`
    Data      string       `json:"data"`
}

type HtmlLayuiReturn struct{
    Code      int               `json:"code"`
    Msg       string            `json:"msg"`
    Count     int               `json:"count"`
    Data      interface{}       `json:"data"`
}

//type FroalaReturn struct{
//    State    string     `json:"state"`
//    Link     string     `json:"link"`
//    Title    string     `json:"title"`
//    Original string     `json:"original"`
//}

func NewHtmlRet(code int, message string, data string) string{
    ret := HtmlReturn{Code:code,Msg:message,Data:data}
    str, err := json.Marshal(ret)
    if err != nil {
        grpclog.Error(err.Error())
        return "{\"code\": -1, \"msg\": \"解析错误\",\"data\":\"/error\"}"
    }
    return string(str)
}

//func NewFroalaRet(state string, link string, title string, original string) string{
//    ret := FroalaReturn{State:state, Link:link, Title:title, Original:original}
//    str, err := json.Marshal(ret)
//    if err != nil {
//        grpclog.Error(err.Error())
//        return "{\"state\": \"ERROR\", \"link\": \"\", \"title\": \"\", \"original\": \"\"}"
//    }
//    return string(str)
//}

func NewLayuiRet(code int, message string, count int, data interface{}) string{
    ret := HtmlLayuiReturn{Code:code, Msg:message, Count:count, Data:data}
    str, err := json.Marshal(ret)
    if err != nil {
        grpclog.Error(err.Error())
        return "{\"code\":-1,\"msg\":\"解析错误\",\"count\":0,\"data\":[]}"
    }
    return string(str)
}


func GetParamByCtx(ctx context.Context) *Parameter{
    params := ctx.Value(define.ParameterName)
    if params == nil{
        grpclog.Error("Not found " + define.ParameterName)
        return nil
    }
    parameter := params.(*Parameter)
    return parameter
}

func JumpToUrl(title string, url string)(rs string , err error){
    out := &bytes.Buffer{}
    data := make(map[string]interface{})
    data["title"] = title
    data["url"] = url
    view.NewView().Render(out, "jump",data)
    return out.String(), nil
}