package app

import (
    "context"
    "github.com/LongMarch7/higo/middleware"
    "github.com/go-kit/kit/log"
    "github.com/go-kit/kit/sd"
    "time"
    grpc_transport "github.com/go-kit/kit/transport/grpc"
)

type ServerOpt struct {
    consulAddr           string
    prefix               string
    serverAddr           string
    serverPort           int
    ctx                  context.Context
    netType              string
    maxThreadCount       string
    advertiseAddress     string
    advertisePort        string
    logger               log.Logger
}

type SOption func(o *ServerOpt)

func SConsulAddr(consulAddr string) SOption {
    return func(o *ServerOpt) {
        o.consulAddr = consulAddr
    }
}

func SPrefix(prefix string) SOption {
    return func(o *ServerOpt) {
        o.prefix = prefix
    }
}

func SServerAddr(server string) SOption {
    return func(o *ServerOpt) {
        o.serverAddr = server
    }
}

func SServerPort(serverPort int) SOption {
    return func(o *ServerOpt) {
        o.serverPort = serverPort
    }
}

func SCtx(ctx context.Context) SOption {
    return func(o *ServerOpt) {
        o.ctx = ctx
    }
}

func SMaxThreadCount(maxThreadCount  string) SOption {
    return func(o *ServerOpt) {
        o.maxThreadCount = maxThreadCount
    }
}

func SLogger(logger  log.Logger) SOption {
    return func(o *ServerOpt) {
        o.logger = logger
    }
}

func SAdvertiseAddress(advertiseAddress  string) SOption {
    return func(o *ServerOpt) {
        o.advertiseAddress = advertiseAddress
    }
}

func SAdvertisePort(advertisePort  string) SOption {
    return func(o *ServerOpt) {
        o.advertisePort = advertisePort
    }
}

func SNetType(netType  string) SOption {
    return func(o *ServerOpt) {
        o.netType = netType
    }
}

type ClientOpt struct {
    consulAddr      string
    prefix          string
    factory         sd.Factory
    retryTime       time.Duration
    retryCount      int
    passingOnly     bool
    logger          log.Logger
    middleware      *middleware.Middleware
    encodeFunc      grpc_transport.EncodeRequestFunc
    decodeFunc      grpc_transport.DecodeResponseFunc
    zipkinName      string
}
type COption func(o *ClientOpt)

func CConsulAddr(consulAddr  string) COption {
    return func(o *ClientOpt) {
        o.consulAddr = consulAddr
    }
}

func CPrefix(prefix  string) COption {
    return func(o *ClientOpt) {
        o.prefix = prefix
    }
}

func CFactory(factory  sd.Factory) COption {
    return func(o *ClientOpt) {
        o.factory = factory
    }
}

func CRetryCount(retryCount int) COption {
    return func(o *ClientOpt) {
        o.retryCount = retryCount
    }
}

func CRetryTime(retryTime time.Duration) COption {
    return func(o *ClientOpt) {
        o.retryTime = retryTime
    }
}

func CPassingOnly(passingOnly bool) COption {
    return func(o *ClientOpt) {
        o.passingOnly = passingOnly
    }
}

func CLogger(logger log.Logger) COption {
    return func(o *ClientOpt) {
        o.logger = logger
    }
}

func CMiddleware(middleware *middleware.Middleware) COption {
    return func(o *ClientOpt) {
        o.middleware = middleware
    }
}

func CEncodeFunc(encodeFunc grpc_transport.EncodeRequestFunc) COption {
    return func(o *ClientOpt) {
        o.encodeFunc = encodeFunc
    }
}

func CDecodeFunc(decodeFunc grpc_transport.DecodeResponseFunc) COption {
    return func(o *ClientOpt) {
        o.decodeFunc = decodeFunc
    }
}

func CZipkinName(zipkinName string) COption {
    return func(o *ClientOpt) {
        o.zipkinName = zipkinName
    }
}