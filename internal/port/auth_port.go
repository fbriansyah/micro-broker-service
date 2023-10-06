package port

import (
	"context"

	"github.com/fbriansyah/micro-payment-proto/protogen/go/auth"

	"google.golang.org/grpc"
)

type AuthPort interface {
	Login(ctx context.Context, in *auth.LoginRequest, opts ...grpc.CallOption) (*auth.LoginResponse, error)
	CreateUser(ctx context.Context, in *auth.CreateUserRequest, opts ...grpc.CallOption) (*auth.CreateUserResponse, error)
}
