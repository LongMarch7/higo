package app

import (
    consulsd "github.com/go-kit/kit/sd/consul"
    "github.com/go-kit/kit/log"
    "os"
    "github.com/hashicorp/consul/api"
    "github.com/go-kit/kit/sd"
    "strconv"
)

type RegisterConfig struct{
    consulAddress string
    prefix string
    service string
    port int
    advertiseAddress string
    advertisePort string
    logger log.Logger
    maxThreadCount string
}

func Register(config RegisterConfig) (registar sd.Registrar) {
    var client consulsd.Client
    {
        consulConfig := api.DefaultConfig()
        consulConfig.Address = config.consulAddress
        consulClient, err := api.NewClient(consulConfig)
        if err != nil {
            config.logger.Log("err", err)
            os.Exit(1)
        }
        client = consulsd.NewClient(consulClient)
    }

    check := api.AgentServiceChecks{
        &api.AgentServiceCheck{
            HTTP:     "http://" + config.advertiseAddress + ":" + config.advertisePort + "/health",
            Interval: "10s",
            Timeout:  "1s",
            Notes:    "Basic health checks",
        },
    }

    asr := api.AgentServiceRegistration{
        ID:      "PORT" + strconv.Itoa(config.port), //unique service ID
        Name:    config.prefix,
        Address: config.service,
        Port:    config.port,
        Tags:    []string{config.prefix,
                            "server="+ config.advertiseAddress + ":" + config.advertisePort,
                            "maxThreadCount=" + config.maxThreadCount},
        Checks:   check,
    }
    registar = consulsd.NewRegistrar(client, &asr,  config.logger)
    return
}

