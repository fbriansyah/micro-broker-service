package port

import (
	"context"

	dmbiller "github.com/fbriansyah/micro-broker-service/internal/application/domain/biller"
	dmproduct "github.com/fbriansyah/micro-broker-service/internal/application/domain/product"
	dmsession "github.com/fbriansyah/micro-broker-service/internal/application/domain/session"
	dmuser "github.com/fbriansyah/micro-broker-service/internal/application/domain/user"
	"github.com/google/uuid"
)

// AuthAdapterPort is interface for client adapter
type AuthAdapterPort interface {
	Login(ctx context.Context, username, password string) (dmuser.User, error)
	Register(ctx context.Context, user dmuser.User, password string) (dmuser.User, error)
}

// SessionAdapterPort is interface for client adapter
type SessionAdapterPort interface {
	GetPayloadData(ctx context.Context, token string) (dmsession.SessionPayload, error)
	CreateSession(ctx context.Context, userID string) (dmsession.Session, error)
}

// PaymentAdapterPort is interface for client adapter
type PaymentAdapterPort interface {
	Inquiry(ctx context.Context, params dmbiller.InquiryParam) (dmbiller.Bill, error)
	Payment(ctx context.Context, params dmbiller.PaymentParam) (dmbiller.Transaction, error)
	GetBalance(ctx context.Context, userID uuid.UUID) (int64, error)
	GetListProduct(ctx context.Context) ([]dmproduct.Product, error)
}
