// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// divider GRPC server
//
// Command:
// $ goa gen goa.design/goa/examples/error/design -o
// $(GOPATH)/src/goa.design/goa/examples/error

package server

import (
	"context"

	"goa.design/goa"
	dividersvc "goa.design/goa/examples/error/gen/divider"
	"goa.design/goa/examples/error/gen/grpc/divider/pb"
	goagrpc "goa.design/goa/grpc"
	"google.golang.org/grpc/codes"
)

// Server implements the pb.DividerServer interface.
type Server struct {
	IntegerDivideH goagrpc.UnaryHandler
	DivideH        goagrpc.UnaryHandler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the expr.
type ErrorNamer interface {
	ErrorName() string
}

// New instantiates the server struct with the divider service endpoints.
func New(e *dividersvc.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		IntegerDivideH: NewIntegerDivideHandler(e.IntegerDivide, uh),
		DivideH:        NewDivideHandler(e.Divide, uh),
	}
}

// NewIntegerDivideHandler creates a gRPC handler which serves the "divider"
// service "integer_divide" endpoint.
func NewIntegerDivideHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeIntegerDivideRequest, EncodeIntegerDivideResponse)
	}
	return h
}

// IntegerDivide implements the "IntegerDivide" method in pb.DividerServer
// interface.
func (s *Server) IntegerDivide(ctx context.Context, message *pb.IntegerDivideRequest) (*pb.IntegerDivideResponse, error) {
	resp, err := s.IntegerDivideH.Handle(ctx, message)
	if err != nil {
		if en, ok := err.(ErrorNamer); ok {
			switch en.ErrorName() {
			case "has_remainder":
				return nil, goagrpc.NewStatusError(codes.Unknown, err)
			case "div_by_zero":
				return nil, goagrpc.NewStatusError(codes.InvalidArgument, err)
			case "timeout":
				return nil, goagrpc.NewStatusError(codes.DeadlineExceeded, err)
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*pb.IntegerDivideResponse), nil
}

// NewDivideHandler creates a gRPC handler which serves the "divider" service
// "divide" endpoint.
func NewDivideHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeDivideRequest, EncodeDivideResponse)
	}
	return h
}

// Divide implements the "Divide" method in pb.DividerServer interface.
func (s *Server) Divide(ctx context.Context, message *pb.DivideRequest) (*pb.DivideResponse, error) {
	resp, err := s.DivideH.Handle(ctx, message)
	if err != nil {
		if en, ok := err.(ErrorNamer); ok {
			switch en.ErrorName() {
			case "div_by_zero":
				return nil, goagrpc.NewStatusError(codes.InvalidArgument, err)
			case "timeout":
				return nil, goagrpc.NewStatusError(codes.DeadlineExceeded, err)
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*pb.DivideResponse), nil
}
