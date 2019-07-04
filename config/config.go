package config

import (
    "encoding/json"
    "fmt"
    "github.com/LongMarch7/higo/util/define"
    "github.com/LongMarch7/higo/util/global"
    "github.com/LongMarch7/higo/util/log"
    "google.golang.org/grpc/grpclog"
    "io/ioutil"
    "strings"
)

type Middleware struct{
    //ZipkinName          string    `json:"zipkin_name"`
    ZipkinUrl           string    `json:"zipkin_url"`
    ZipkinhostPort      string    `json:"zipkinhost_port"`
    ZipkinDebug         bool      `json:"zipkin_debug"`
    ZipkinMaxLogs       int       `json:"zipkin_max_logs"`

    RatelimitInterval   int64     `json:"ratelimit_interval"`
    RatelimitBurst      int       `json:"ratelimit_burst"`

    //PrometheusSubsystem string    `json:"prometheus_subsystem"`
    //PrometheusName      string    `json:"prometheus_name"`

    //HystrixName         string    `json:"hystrix_name"`
    HystrixTimeout      int       `json:"hystrix_timeout"`
    HystrixMCR          int       `json:"hystrix_mcr"`    //maxConcurrentRequests
    HystrixRVT          int       `json:"hystrix_rvt"`    //requestVolumeThreshold
    HystrixSW           int       `json:"hystrix_sw"`     //sleepWindow
    HystrixEPT          int       `json:"hystrix_ept"`    //errorPercentThreshold

    //LogName             string    `json:"log_name"`
    //LogMethodName       string    `json:"log_method_name"`
}
type ServiceList struct {
    Name         string  `json:"name"`
    Addr         string  `json:"addr"`
    //Port         string  `json:"port"`
    Count        string  `json:"count"`   //connect max num
    AdAddr       string  `json:"ad_addr"` //consul advertise address
    //AdPort       string  `json:"ad_port"` //consul advertise port
    TemplatePath string   `json:"template_path"`
}
type Sql struct{
    Driver          string `json:"driver"`
    User            string  `json:"user"`
    Pwd             string  `json:"pwd"`
    Net             string  `json:"net"`
    Addr            string  `json:"addr"`
    Port            string  `json:"port"`
    Db              string  `json:"db"`
    MaxOpenConn     int     `json:"max_open_conn"`
    MaxIdleConn     int     `json:"max_idle_conn"`
    Show            bool    `json:"show"`
    Level           string  `json:"level"`
    File            string  `json:"file"`
}
type Memcache struct{
    MaxIdleConn     int      `json:"max_idle_conn"`
    Server          []string `json:"server"`
}
type Config struct{
    Domain       string        `json:"domain"`
    SslKey       string        `json:"ssl_key"`
    SslCrt       string        `json:"ssl_crt"`
    Port         string        `json:"port"`
    ConsulServer string        `json:"consul_server"`
    ServiceList  []ServiceList `json:"service_list"`
    Sql          Sql           `json:"sql"`
    Memcache     Memcache      `json:"memcache"`
    RetryCount   int           `json:"retry_count"`
    RetryTime    int64         `json:"retry_time"`
    CliMw        Middleware    `json:"cli_middleware"`
    SvrMw        Middleware    `json:"svr_middleware"`
    Logger       zap.Config    `json:"logger"`
    RootPath     string        `json:"root_path"`
    StaticPath   string        `json:"static_path"`
    UploadPath   string        `json:"upload_path"`
    TemplatePath   string      `json:"template_path"`
}

type Configer struct {
    Name       string
    Port       string
    AdPort     string
    Config     Config
}

var ConfigFilePath   string

func ConfigInit(mode string, name string, path string, port string, adPort string) *Configer{
    config :=new(Configer)
    if strings.Contains(mode,define.InitModeStr) {
        global.AppMode = define.InitMode
    }else if strings.Contains(mode,define.SvrModeStr){
        global.AppMode = define.SvrMode
        if len(port) == 0 || len(adPort) ==0 {
            fmt.Println("Not set port or ad port")
            return nil
        }
        config.Port = port
        config.AdPort = adPort
    }else if strings.Contains(mode,define.CliModeStr){
        global.AppMode = define.CliMode
    }else{
        fmt.Println("Must set mode")
        return nil
    }
    if len(name)==0 && global.AppMode != define.InitMode {
        fmt.Println("Not set server name")
        return nil
    }
    config.Name = name
    global.SeverName = name
    if len(path) == 0 {
        path = "config.json"
    }
    ConfigFilePath = path
    config.Config = Config{}
    err := Load(path, &config.Config)
    if err != nil {
        fmt.Println("Read config failed")
        return nil
    }
    return config
}

func Load(filename string, v interface{}) error{
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        grpclog.Error(err.Error())
        return err
    }
    err = json.Unmarshal(data, v)
    if err != nil {
        grpclog.Error(err.Error())
        return err
    }
    return nil
}
