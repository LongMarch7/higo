package tansport

import (
    "github.com/go-kit/kit/endpoint"
    "context"
    "github.com/gorilla/mux"
    "net/http"
    httptransport "github.com/go-kit/kit/transport/http"
    stdprometheus "github.com/prometheus/client_golang/prometheus/promhttp"
)

type HealthResponse struct {
    Status bool
}
func MakeHealthEndpoint() endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        return HealthResponse{Status: true }, nil
    }
}

func MakeMonitoringHttpHandler(r *mux.Router, endpoint endpoint.Endpoint) http.Handler {
    options := []httptransport.ServerOption{
        httptransport.ServerErrorEncoder(encodeError),
    }

    //GET /health
    r.Methods("GET").Path("/health").Handler(httptransport.NewServer(
        endpoint,
        DefaultHttpDecodeRequest,
        JsonEncodeResponse,
        options...,
    ))

    // GET /metrics
    r.Path("/metrics").Handler(stdprometheus.Handler())
    return r
}