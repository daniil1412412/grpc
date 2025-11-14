package auth

import (
	"context"
	ssov1 "grpc-serv/gen/go/sso"

	"google.golang.org/grpc"
)

type serverApi struct {
	ssov1.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	ssov1.RegisterAuthServer(gRPC, &serverApi{})
}

func (s *serverApi) Login(ctx context.Context, req *ssov1.AuthRequest) (*ssov1.AuthResponce, error) {
	panic("implement me")
}

func (s *serverApi) Register(ctx context.Context, req *ssov1.RegisterRequest) (*ssov1.RegisterResponce, error) {
	panic("implement me")
}

func (s *serverApi) IsAdmin(ctx context.Context, req *ssov1.IsAdminRequest) (*ssov1.IsAdminResponce, error) {
	panic("implement me")
}
