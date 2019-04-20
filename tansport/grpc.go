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

func DefaultGrpcEncodeRequestFunc(_ context.Context,rsp interface{}) (request interface{}, err error){
    return rsp, nil
}
func DefaultGrpcDecodeResponseFunc(_ context.Context,rsp interface{}) (response interface{}, err error){
    return rsp, nil
}
