package define

import (
    "google.golang.org/grpc"
)
type GrpcRegister func(s *grpc.Server)
