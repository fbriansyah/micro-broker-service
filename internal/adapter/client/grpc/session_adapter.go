package grpcclient

import (
	"context"

	dmsession "github.com/fbriansyah/micro-broker-service/internal/application/domain/session"
	"github.com/fbriansyah/micro-broker-service/internal/port"
	"github.com/fbriansyah/micro-payment-proto/protogen/go/session"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type SessionClientAdapter struct {
	client port.SessionPort
}

func NewSessionClientAdapter(conn *grpc.ClientConn) *SessionClientAdapter {
	client := session.NewSessionServiceClient(conn)
	return &SessionClientAdapter{
		client: client,
	}
}

func (a *SessionClientAdapter) GetPayloadData(ctx context.Context, token string) (dmsession.SessionPayload, error) {
	payload, err := a.client.GetPayloadFromToken(ctx, &session.Token{
		AccessToken: token,
	})

	if err != nil {
		return dmsession.SessionPayload{}, err
	}

	return dmsession.SessionPayload{
		UserID: uuid.MustParse(payload.UserId),
	}, nil
}
