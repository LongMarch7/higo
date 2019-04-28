// Package base is the basic building cblock of utron. The main structure here is
// Context, but for some reasons to avoid confusion since there is a lot of
// context packages I decided to name this package base instead.
package base

import (
    "github.com/gorilla/sessions"
    "net/http"

    "github.com/gorilla/context"
    "github.com/gorilla/mux"
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

type Context struct {

    Params map[string]string

    // Data keeps values that are going to be passed to the view as context
    Data map[string]interface{}

    SessionStore sessions.Store

    request    *http.Request
    response   http.ResponseWriter
    //out        io.ReadWriter
}

// NewContext creates new context for the given w and r
func NewContext(w http.ResponseWriter, r *http.Request) *Context {
    ctx := &Context{
        Params:   make(map[string]string),
        Data:     make(map[string]interface{}),
        request:  r,
        response: w,
        //out:      &bytes.Buffer{},
    }
    ctx.Init()
    return ctx
}

// Init initializes the context
func (c *Context) Init() {
    c.Params = mux.Vars(c.request)
}

// Write writes the data to the context, data is written to the http.ResponseWriter
// upon calling Commit().
//
// data will only be used when Template is not specified and there is no View set. You can use
// this for creating APIs (which does not depend on views like JSON APIs)
//func (c *Context) Write(data []byte) (int, error) {
//    return c.out.Write(data)
//}

// TextPlain renders text/plain response
func (c *Context) TextPlain() {
    c.SetHeader(Content.Type, Content.TextPlain)
}

// JSON renders JSON response
func (c *Context) JSON() {
    c.SetHeader(Content.Type, Content.Application.JSON)
}

// HTML renders text/html response
func (c *Context) HTML() {
    c.SetHeader(Content.Type, Content.TextHTML)
}

// Request returns the *http.Request object used by the context
func (c *Context) Request() *http.Request {
    return c.request
}

// Response returns the http.ResponseWriter object used by the context
func (c *Context) Response() http.ResponseWriter {
    return c.response
}

// GetData retrievess any data stored in the request using
// gorilla.Context package
func (c *Context) GetData(key interface{}) interface{} {
    return context.Get(c.Request(), key)
}

//SetData stores key value into the request object attached with the context.
//this is a helper method, wraping gorilla/context
func (c *Context) SetData(key, value interface{}) {
    context.Set(c.Request(), key, value)
}

// Set sets value in the context object. You can use this to change the following
//
//	 * Request by passing *http.Request
//	 * ResponseWriter by passing http.ResponseVritter
//	 * view by passing View
//	 * response status code by passing an int
func (c *Context) Set(value interface{}) {
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
func (c *Context) SetHeader(key, value string) {
    c.response.Header().Set(key, value)
}

func (c *Context) HtmlRender(data []byte) {
    c.HTML()
    c.response.Write(data)
}

func (c *Context) JsonRender(data []byte) {
    c.JSON()
    c.response.Write(data)
}

func (c *Context) TextRender(data []byte) {
    c.TextPlain()
    c.response.Write(data)
}

func (c *Context) Redirect(url string, code int) {
    http.Redirect(c.Response(), c.Request(), url, code)
}
