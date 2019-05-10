package cookie

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
    if flag != c.F {
        c.F = flag
    }

    if token != c.T {
        c.T = token
    }
}
