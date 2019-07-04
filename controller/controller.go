package controller

import (
   "context"
   "github.com/LongMarch7/higo/controller/base"
    "golang.org/x/tools/go/ssa/interp/testdata/src/errors"
    "google.golang.org/grpc/grpclog"
    "strings"
   "reflect"
)

func ControllerCall(ctx context.Context, pattern string)(rs string , err error){
    strs := strings.Split(pattern,":")
    err = errors.New("Not found")

    if len(strs) != 2 {
        grpclog.Error("Pattern invalid")
        return
    }
    value,ok := base.GetController(strs[0])
    if !ok{
        grpclog.Error("Not found Controller")
        return
    }
    params := make([]reflect.Value,1)
    params[0] = reflect.ValueOf(ctx)
    cRef := reflect.ValueOf(value)
    funcValue :=cRef.MethodByName(strs[1])
    if !funcValue.IsValid() || funcValue.IsNil(){
        grpclog.Error("Not found function")
        return
    }else if funcValue.IsValid(){
        rets := funcValue.Call(params)
        rs = rets[0].Interface().(string)
        retErr := rets[1].Interface()
        if retErr == nil{
            err = nil
        }else{
            err = retErr.(error)
        }
    }
    return
}
