package controller

import (
   "context"
   "github.com/LongMarch7/higo/controller/base"
    "strings"
   "reflect"
)

func ControllerCall(ctx context.Context, pattern string)(rs string , err string){
    strs := strings.Split(pattern,":")
    if len(strs) != 2 {
        return "","not found"
    }
    c := base.GetController(strs[0])
    params := make([]reflect.Value,1)
    params[0] = reflect.ValueOf(ctx)
    cRef := reflect.ValueOf(c)
    rets :=cRef.MethodByName(strs[1]).Call(params)
    return rets[0].Interface().(string),rets[1].Interface().(string)
}
