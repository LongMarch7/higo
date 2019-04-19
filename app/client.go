package app

import (
    zaplog "github.com/LongMarch7/go-web/plugin/zap-log"
    "time"
)

type Client struct {
    opts                  ClientOpt
}

func defaultClientConfig() ClientOpt{
    return ClientOpt{
        consulAddr: "http://localhost:8500",
        prefix: "bookServer",
        retryTime: time.Second * 3,
        retryCount: 3,
        passingOnly: true,
        logger: zaplog.NewDefaultLogger(),
    }
}

func NewClient(opts ...COption) *Client{
    opt := defaultClientConfig()
    for _, o := range opts {
        o(&opt)
    }
    return &Client{
        opts: opt,
    }
}