package grpcclient

import (
	"context"

	dmsession "github.com/fbriansyah/micro-broker-service/internal/application/domain/session"
	"github.com/fbriansyah/micro-broker-service/internal/port"
	"github.com/fbriansyah/micro-broker-service/util"
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

// GetPayloadData validate token and extract payload from token
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

func (a *SessionClientAdapter) CreateSession(ctx context.Context, userID string) (dmsession.Session, error) {
	session, err := a.client.CreateSession(ctx, &session.UserID{
		UserId: userID,
	})

	if err != nil {
		return dmsession.Session{}, err
	}

	return dmsession.Session{
		Id:                    session.Id,
		UserId:                userID,
		AccessToken:           session.AccessToken,
		RefreshToken:          session.RefreshToken,
		AccessTokenExpiresAt:  util.FromDateTime(session.AccessTokenExpiresAt),
		RefreshTokenExpiresAt: util.FromDateTime(session.RefreshTokenExpiresAt),
	}, nil
}
