package adder

import (
	"Nikolay200669/test_grpc/pkg/api"
	"context"
)

// GRPCServer ::
type GRPCServer struct{}

// Add ::
func (s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{Result: req.GetX() * req.GetY()}, nil
}
