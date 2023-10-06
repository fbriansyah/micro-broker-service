package port

import (
	"context"

	"github.com/fbriansyah/micro-payment-proto/protogen/go/session"
	"google.golang.org/grpc"
)

type SessionPort interface {
	CreateSession(ctx context.Context, in *session.UserID, opts ...grpc.CallOption) (*session.Session, error)
	RefreshToken(ctx context.Context, in *session.SessionID, opts ...grpc.CallOption) (*session.Session, error)
	DeleteSession(ctx context.Context, in *session.SessionID, opts ...grpc.CallOption) (*session.SessionID, error)
	GetPayloadFromToken(ctx context.Context, in *session.Token, opts ...grpc.CallOption) (*session.Payload, error)
}
