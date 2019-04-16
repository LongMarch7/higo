package tansport

import (
    "context"
)

func DefaultGrpcDecodeRequest(_ context.Context, req interface{}) (interface{}, error) {
    return req, nil
}

func DefaultGrpcEncodeResponse(_ context.Context, rsp interface{}) (interface{}, error) {
    return rsp, nil
}
