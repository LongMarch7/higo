package app

import (
    "context"
    "github.com/go-kit/kit/log"
    "github.com/go-kit/kit/sd"
    "time"
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
}
type COption func(o *ClientOpt)