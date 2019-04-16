package app

import (
    "context"
    "github.com/LongMarch7/higo/middleware/hystrix"
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
    serviceStruct        interface{}
    advertiseAddress     string
    advertisePort        string
    logger               log.Logger
}

type SOption func(o *ServerOpt)

type ClientOpt struct {
    consulAddr      string
    prefix          string
    ctx             context.Context
    factory         sd.Factory
    retryTime       time.Duration
    retryCount      int
    hOptions        []hystrix.HOption
    passingOnly     bool
    logger          log.Logger
}
type COption func(o *ClientOpt)