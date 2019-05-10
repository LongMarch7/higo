package admin

import (
    "bytes"
    "encoding/json"
    "fmt"
    "github.com/LongMarch7/higo/controller/base"
    "github.com/LongMarch7/higo/view"
    "golang.org/x/net/context"
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

func (a* adminController)Index(ctx context.Context) (rs string , err string){
    out := &bytes.Buffer{}
    view.NewView().Render(out, name + "/index",nil)
    return out.String(), ""
}

func (a* adminController)Login(ctx context.Context) (rs string , err string){
    //out := &bytes.Buffer{}
    //view.NewView().Render(out, name + "/index",nil)
    params := ctx.Value("Parameter")
    fmt.Println(params)
    if params == nil{
        return "", "parameter get error"
    }
    parameter := params.(*base.Parameter)
    muxParame := parameter.MuxParams
    var mapValues = make(map[string]string)
    muxErr := json.Unmarshal([]byte(muxParame),&mapValues)
    if muxErr != nil{
        return "", "mux Unmarshal error"
    }
    fmt.Println(mapValues["name"])
    header := metadata.Pairs("res_cookie", mapValues["name"])
    grpc.SendHeader(ctx, header)
    return "{\"code\": 0, \"result\": \"login success\"}", ""
}

