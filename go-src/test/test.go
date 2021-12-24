package test

import (
	"context"

	server "github.com/hideyoshidan/grpc/server"
)

// Test
type Test struct {
	server.UnimplementedSampleServer
	testMsg string
}

// NewTest generate instance of Test struct
func NewTest() server.SampleServer {
	return &Test{}
}

// Sample return string
func (s *Test) Sample(con context.Context, in *server.SampleRequest) (*server.SampleResponse, error) {
	return &server.SampleResponse{
		ResponseMsg: s.testMsg,
	}, nil
}
