package controller

import (
   "context"
   "github.com/LongMarch7/higo/controller/base"
    "golang.org/x/tools/go/ssa/interp/testdata/src/errors"
    "strings"
   "reflect"
)

func ControllerCall(ctx context.Context, pattern string)(rs string , err error){
    strs := strings.Split(pattern,":")
    if len(strs) != 2 {
        return "",errors.New("not found")
    }
    c := base.GetController(strs[0])
    params := make([]reflect.Value,1)
    params[0] = reflect.ValueOf(ctx)
    cRef := reflect.ValueOf(c)
    rets :=cRef.MethodByName(strs[1]).Call(params)
    rs = rets[0].Interface().(string)
    retErr := rets[1].Interface()
    if retErr == nil{
        err = nil
    }else{
        err = retErr.(error)
    }
    return
}
