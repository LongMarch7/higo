package base

import (
	"github.com/LongMarch7/higo/util/define"
	"github.com/gorilla/mux"
	"encoding/json"
	"github.com/LongMarch7/higo/base"
	"golang.org/x/net/context"
	"io/ioutil"
	"net/http"
	"strings"
)

func MakeReqDataMiddleware(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		// TODO implement the business logic of HtmlCall
		ctx := req.Context()
		baseCtx := ctx.Value(define.StrucName)
		if baseCtx == nil {
			baseCtx = base.NewContext(res, req)
		}
		baseContext := baseCtx.(*base.BaseContext)

		muxVals := mux.Vars(req)
		if len(muxVals) > 0 {
			muxStrings, muxErr := json.Marshal(muxVals)
			if muxErr == nil {
				baseContext.Params[define.MuxParamsName]  = string(muxStrings)
			}
		}

		getVals := req.URL.Query()
		if len(getVals) > 0 {
			getStrings, getErr := json.Marshal(getVals)
			if getErr == nil {
				baseContext.Params[define.GetParamsName]  = string(getStrings)
			}
		}
		contType := req.Header.Get("Content-Type")
		if strings.Contains(contType, define.ContentTypeJson){
			posgVals, postErr := ioutil.ReadAll(req.Body)
			req.Body.Close()
			if postErr == nil {
				vals := string(posgVals)
				if len(vals) > 0 {
					baseContext.Params[define.PostParamsJsonName] = vals
				}
			}
			baseContext.Params[define.ContentType]  = define.ContentTypeJson
		} else if strings.Contains(contType, define.ContentTypeForm){
			req.ParseForm()
			posgVals, postErr := json.Marshal(req.PostForm)
			if postErr == nil {
				baseContext.Params[define.PostParamsFormName]  = string(posgVals)
			}
			baseContext.Params[define.ContentType]  = define.ContentTypeForm
		}

		c, err := req.Cookie(define.CookieName)
		if err == nil  && len(c.Value) > 0 {
			baseContext.Params[define.ReqCookieName] = c.Value
		}
		baseContext.Params[define.ReqMethodName] = req.Method
		ctx = context.WithValue(ctx, define.StrucName,baseContext)
		req = req.WithContext(ctx)
		next(res,req)
	}
}

type GrpcClientParameter struct{
	Srv           string
	Method        string
	NewRlyFunc    func() interface{}
}
