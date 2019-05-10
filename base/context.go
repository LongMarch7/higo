// Package base is the basic building cblock of utron. The main structure here is
// Context, but for some reasons to avoid confusion since there is a lot of
// context packages I decided to name this package base instead.
package base

import (
    "google.golang.org/grpc/metadata"
    "net/http"
    gorilla_context "github.com/gorilla/context"
    "context"
)

// Content holds http response content type strings
var Content = struct {
    Type        string
    TextPlain   string
    TextHTML    string
    Application struct {
        Form, JSON, MultipartForm string
    }
}{
    "Content-Type", "text/plain", "text/html",
    struct {
        Form, JSON, MultipartForm string
    }{
        "application/x-www-form-urlencoded",
        "application/json",
        "multipart/form-data",
    },
}

type BaseContext struct {
    Params      map[string]string
    Method      string
    GrpcHeader  metadata.MD
    GrpcTrailer metadata.MD
    request     *http.Request
    response    http.ResponseWriter
}

const StrucName = "baseContext"

// NewContext creates new context for the given w and r
func NewContext(w http.ResponseWriter, r *http.Request) *BaseContext {
    ctx := &BaseContext{
        Params:   make(map[string]string),
        Method: r.Method,
        request:  r,
        response: w,
    }
    //ctx.Init()
    return ctx
}
//
//// Init initializes the context
//func (c *Context) Init() {
//    c.Params = mux.Vars(c.request)
//}

//// TextPlain renders text/plain response
//func (c *Context) TextPlain() {
//    c.SetHeader(Content.Type, Content.TextPlain)
//}
//
//// JSON renders JSON response
//func (c *Context) JSON() {
//    c.SetHeader(Content.Type, Content.Application.JSON)
//}
//
//// HTML renders text/html response
//func (c *Context) HTML() {
//    c.SetHeader(Content.Type, Content.TextHTML)
//}

// Request returns the *http.Request object used by the context
func (c *BaseContext) Request() *http.Request {
    return c.request
}

// Response returns the http.ResponseWriter object used by the context
func (c *BaseContext) Response() http.ResponseWriter {
    return c.response
}

// GetData retrievess any data stored in the request using
// gorilla.Context package
func (c *BaseContext) GetData(key interface{}) interface{} {
    return gorilla_context.Get(c.Request(), key)
}

//SetData stores key value into the request object attached with the context.
//this is a helper method, wraping gorilla/context
func (c *BaseContext) SetData(key, value interface{}) {
    gorilla_context.Set(c.Request(), key, value)
}

// Set sets value in the context object. You can use this to change the following
//
//	 * Request by passing *http.Request
//	 * ResponseWriter by passing http.ResponseVritter
//	 * view by passing View
//	 * response status code by passing an int
func (c *BaseContext) Set(value interface{}) {
    switch value := value.(type) {
    case *http.Request:
        c.request = value
    case http.ResponseWriter:
        c.response = value
    case int:
        c.response.WriteHeader(value)
    }
}

// SetHeader sets response header
func (c *BaseContext) SetHeader(key, value string) {
    c.response.Header().Set(key, value)
}

func (c *BaseContext) HtmlRender(data []byte) {
    HtmlRender(c.response,data)
}

func HtmlRender(res http.ResponseWriter, data []byte) {
    res.Header().Set(Content.Type, Content.TextHTML)
    res.Write(data)
}

func (c *BaseContext) JsonRender(data []byte) {
    JsonRender(c.response,data)
}

func JsonRender(res http.ResponseWriter, data []byte) {
    res.Header().Set(Content.Type, Content.Application.JSON)
    res.Write(data)
}

func (c *BaseContext) TextRender(data []byte) {
    TextRender(c.response,data)
}

func TextRender(res http.ResponseWriter, data []byte) {
    res.Header().Set(Content.Type, Content.TextPlain)
    res.Write(data)
}

func SetCookie(ctx context.Context, res http.ResponseWriter){
    baseCtx := ctx.Value(StrucName)
    if baseCtx != nil {
        cookie := baseCtx.(*BaseContext).GrpcHeader
        if value,ok := cookie["res_cookie"]; ok{
            if len(cookie) > 0{
                http.SetCookie(res, &http.Cookie{
                    Name:  "info",
                    Value: value[0],
                    HttpOnly: true,
                })
            }
        }
    }
}

func (c *BaseContext) Redirect(url string, code int) {
    http.Redirect(c.Response(), c.Request(), url, code)
}
