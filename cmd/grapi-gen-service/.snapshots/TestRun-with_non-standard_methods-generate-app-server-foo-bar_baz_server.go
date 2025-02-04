package foo

import (
	"context"

	"github.com/Ryo-not-rio/grapi/pkg/grapiserver"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	foo_pb "testapp/api/foo"
)

// BarBazServiceServer is a composite interface of foo_pb.BarBazServiceServer and grapiserver.Server.
type BarBazServiceServer interface {
	foo_pb.BarBazServiceServer
	grapiserver.Server
}

// NewBarBazServiceServer creates a new BarBazServiceServer instance.
func NewBarBazServiceServer() BarBazServiceServer {
	return &barBazServiceServerImpl{}
}

type barBazServiceServerImpl struct {
}

func (s *barBazServiceServerImpl) ListBarBazs(ctx context.Context, req *foo_pb.ListBarBazsRequest) (*foo_pb.ListBarBazsResponse, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *barBazServiceServerImpl) CreateBarBaz(ctx context.Context, req *foo_pb.CreateBarBazRequest) (*foo_pb.BarBaz, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *barBazServiceServerImpl) DeleteBarBaz(ctx context.Context, req *foo_pb.DeleteBarBazRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *barBazServiceServerImpl) Rename(ctx context.Context, req *foo_pb.RenameRequest) (*foo_pb.RenameResponse, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *barBazServiceServerImpl) MoveMove(ctx context.Context, req *foo_pb.MoveMoveRequest) (*foo_pb.MoveMoveResponse, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
