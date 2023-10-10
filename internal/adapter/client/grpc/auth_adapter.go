package grpcclient

import (
	"context"

	dmuser "github.com/fbriansyah/micro-broker-service/internal/application/domain/user"
	"github.com/fbriansyah/micro-broker-service/internal/port"
	"github.com/fbriansyah/micro-payment-proto/protogen/go/auth"
	"google.golang.org/grpc"
)

type AuthClientAdapter struct {
	client port.AuthPort
}

func NewAuthClientAdapter(conn *grpc.ClientConn) *AuthClientAdapter {
	client := auth.NewAuthServiceClient(conn)
	return &AuthClientAdapter{
		client: client,
	}
}

func (a *AuthClientAdapter) Login(ctx context.Context, username, password string) (dmuser.User, error) {

	resp, err := a.client.Login(ctx, &auth.LoginRequest{
		Username: username,
		Password: password,
	})

	if err != nil {
		return dmuser.User{}, err
	}

	return dmuser.User{
		ID:       resp.Userid,
		Username: username,
		Name:     resp.Name,
	}, nil
}

func (a *AuthClientAdapter) Register(ctx context.Context, user dmuser.User, password string) (dmuser.User, error) {

	usr, err := a.client.CreateUser(ctx, &auth.CreateUserRequest{
		Username: user.Username,
		Password: password,
		Name:     user.Name,
	})

	if err != nil {
		return dmuser.User{}, err
	}

	return dmuser.User{
		ID:       usr.Userid,
		Username: usr.Username,
		Name:     usr.Name,
	}, nil
}
