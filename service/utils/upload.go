package utils

import (
    "encoding/json"
    "github.com/LongMarch7/higo/base"
    "github.com/LongMarch7/higo/middleware/hystrix"
    "github.com/LongMarch7/higo/middleware/ratelimit"
    "github.com/LongMarch7/higo/service/web"
    "github.com/LongMarch7/higo/util/define"
    "github.com/LongMarch7/higo/util/global"
    "github.com/go-kit/kit/endpoint"
    "google.golang.org/grpc/grpclog"
    "io"
    "net/http"
    "context"
    "os"
    "path"
    "strconv"
    "time"
)

func MakeUploadHandler(e endpoint.Endpoint, pattern string) func(http.ResponseWriter, *http.Request) {
    ratelimit.NewLimiter().AddRateLimitForMethod(pattern)
    hystrix.NewHystrix().AddHystrixForMethod(pattern)
    clientProxy := web.ApiCallProxy(e)
    return func(res http.ResponseWriter, req *http.Request) {
        // TODO implement the business logic of ApiCall
        ctx := req.Context()
        ctx = context.WithValue(ctx, define.ReqPatternName, pattern)
        _, err := clientProxy(ctx, pattern)
        base.SetCookie(ctx, res)
        if err == nil {
            ret := Upload(req)
            content, err := json.Marshal(ret)
            if err == nil {
                base.JsonRender(res, content)
                return
            }
        }
        base.JsonRender(res, []byte("{\"code\": -1, \"msg\": \"解析错误\",\"data\":\"/error\"}"))
    }
}

func Upload(req *http.Request) map[string]interface{}{
    //保存上传的图片
    _, h, err := req.FormFile("file")
    if err != nil {
        grpclog.Error(err)
    }
    fileName := h.Filename
    fileSuffix := path.Ext(fileName)
    newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix // + "_" + filename
    date_url := time.Now().Format("2006-01-02") + "/"
    err = os.MkdirAll(global.UploadPath + date_url, 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
    if err != nil {
        grpclog.Error(err)
    }
    path1 := global.UploadPath + date_url + newname //h.Filename
    Url := "/upload/" + date_url + newname
    err = SaveToFile( req,"file", path1) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
    if err != nil {
        grpclog.Error(err)
    }
    return map[string]interface{}{"state": "SUCCESS", "link": Url, "title": fileName, "original": fileName}
}

func SaveToFile(req *http.Request,fromfile, tofile string) error {
    file, _, err := req.FormFile(fromfile)
    if err != nil {
        return err
    }
    defer file.Close()
    f, err := os.OpenFile(tofile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
    if err != nil {
        return err
    }
    defer f.Close()
    io.Copy(f, file)
    return nil
}