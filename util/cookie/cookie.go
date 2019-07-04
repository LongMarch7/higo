package cookie

import (
    "encoding/base64"
    "encoding/json"
    "google.golang.org/grpc/grpclog"
)

type Cookie struct{
    T      string `json:"t"`    // token
    U      string `json:"u"`    // user id
    F      int    `json:"f"`    // flag
    S      string `json:"s"`    // pwd
}

func (c* Cookie)UpdateCookie(token string, user string, pwd string, flag int){
    if user != c.U {
        c.U = user
    }

    if pwd != c.S {
        c.S = pwd
    }
    if flag != c.F && flag >0 {
        c.F = flag
    }

    if token != c.T {
        c.T = token
    }
}

func (c* Cookie)Marshal()  string{
    ret, err :=json.Marshal(c)
    if err != nil{
        return ""
    }
    return base64.StdEncoding.EncodeToString([]byte(ret))
}
func UnMarshal(str string, cookie* Cookie)  (bool){
    signedByte, err := base64.StdEncoding.DecodeString(str)
    if err != nil{
        return false
    }
    if ret := json.Unmarshal(signedByte, cookie); ret != nil{
        grpclog.Error("Parsing cookie error:" + str)
        return false
    }
    return true
}
