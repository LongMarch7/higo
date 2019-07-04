package admin

import (
    "bytes"
    "encoding/json"
    "fmt"
    "github.com/LongMarch7/higo/controller/base"
    "github.com/LongMarch7/higo/util/define"
    "github.com/LongMarch7/higo/view"
    "golang.org/x/net/context"
    "golang.org/x/tools/go/ssa/interp/testdata/src/errors"
    "google.golang.org/grpc"
    "google.golang.org/grpc/metadata"
)

type adminController struct {
}
var controller = &adminController{}
var name = "admin"
func init(){
    base.AddController(name, controller)
}

func (a* adminController)Index(ctx context.Context) (rs string , err error){
    out := &bytes.Buffer{}
    view.NewView().Render(out, name + "/index",nil)
    return out.String(), nil
}

func (a* adminController)Login(ctx context.Context) (rs string , err error){
    //out := &bytes.Buffer{}
    //view.NewView().Render(out, name + "/index",nil)
    params := ctx.Value(define.ParameterName)
    fmt.Println(params)
    if params == nil{
        return "", errors.New("parameter get error")
    }
    parameter := params.(*base.Parameter)
    muxParame := parameter.MuxParams
    fmt.Println(muxParame["name"])
    header := metadata.Pairs(define.ResCookieName, muxParame["name"])
    grpc.SendHeader(ctx, header)
    return "{\"code\": 0, \"result\": \"login success\"}", nil
}

