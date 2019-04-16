package tansport

import (
    "context"
    "encoding/json"
    "net/http"
)

type errorer interface {
    error() error
}

func DefaultHttpDecodeRequest(_ context.Context, req *http.Request) (interface{}, error) {
    return req, nil
}

func JsonEncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
    if e, ok := response.(errorer); ok && e.error() != nil {
        // Not a Go kit transport error, but a business-logic error.
        // Provide those as HTTP errors.
        encodeError(ctx, e.error(), w)
        return nil
    }
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
    if err == nil {
        return
    }
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(http.StatusInternalServerError)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "error": err.Error(),
    })
}
