// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// calc GRPC server
//
// Command:
// $ goa gen goa.design/goa/examples/calc/design -o
// $(GOPATH)/src/goa.design/goa/examples/calc

package server

import (
	"context"

	"goa.design/goa"
	calcsvc "goa.design/goa/examples/calc/gen/calc"
	"goa.design/goa/examples/calc/gen/grpc/calc/pb"
	goagrpc "goa.design/goa/grpc"
)

// Server implements the pb.CalcServer interface.
type Server struct {
	AddH goagrpc.UnaryHandler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the expr.
type ErrorNamer interface {
	ErrorName() string
}

// New instantiates the server struct with the calc service endpoints.
func New(e *calcsvc.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		AddH: NewAddHandler(e.Add, uh),
	}
}

// NewAddHandler creates a gRPC handler which serves the "calc" service "add"
// endpoint.
func NewAddHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeAddRequest, EncodeAddResponse)
	}
	return h
}

// Add implements the "Add" method in pb.CalcServer interface.
func (s *Server) Add(ctx context.Context, message *pb.AddRequest) (*pb.AddResponse, error) {
	resp, err := s.AddH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*pb.AddResponse), nil
}
