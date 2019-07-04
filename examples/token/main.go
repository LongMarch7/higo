package main

import (
    "fmt"
    "github.com/LongMarch7/higo/util/token"
)

func main() {
    fmt.Println(token.NewTokenWithSalt("123456"))
    str, err := token.Sign(&token.DefaultClaims{Pwd:"admin"})
    fmt.Println(str)
    if err == nil{
        ret, retErr := token.Parse(str,&token.DefaultClaims{})
        if retErr == nil{
            fmt.Println(ret.(*token.DefaultClaims).Pwd)
        }
    }
}
