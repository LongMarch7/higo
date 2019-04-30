package main

import (
    "fmt"
    "github.com/LongMarch7/higo/util/token"
)

func main() {
    fmt.Println(token.NewToken("test"))
    str, err := token.Sign(&token.DefaultClaims{Pwd:"admin"})
    if err == nil{
        ret, retErr := token.Parse(str,&token.DefaultClaims{})
        if retErr == nil{
            fmt.Println(ret.(*token.DefaultClaims).Pwd)
        }
    }
}
