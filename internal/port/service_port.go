package port

import (
	"context"

	dmbiller "github.com/fbriansyah/micro-broker-service/internal/application/domain/biller"
	dmsession "github.com/fbriansyah/micro-broker-service/internal/application/domain/session"
	dmuser "github.com/fbriansyah/micro-broker-service/internal/application/domain/user"
)

type BrokerServicePort interface {
	Login(ctx context.Context, username, password string) (dmsession.Session, error)
	Register(ctx context.Context, user dmuser.User, password string) (dmuser.User, error)
	Inquiry(ctx context.Context, billNumber, productCode, token string) (dmbiller.Bill, error)
	Payment(ctx context.Context, amount int64, inqID, token string) (dmbiller.Transaction, error)
}
