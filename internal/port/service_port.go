package port

import (
	"context"

	dmbiller "github.com/fbriansyah/micro-broker-service/internal/application/domain/biller"
	dmproduct "github.com/fbriansyah/micro-broker-service/internal/application/domain/product"
	dmsession "github.com/fbriansyah/micro-broker-service/internal/application/domain/session"
	dmuser "github.com/fbriansyah/micro-broker-service/internal/application/domain/user"
)

type BrokerServicePort interface {
	// Login send login request to auth microservice
	Login(ctx context.Context, username, password string) (dmuser.User, dmsession.Session, error)
	// Register send register request to auth microservice
	Register(ctx context.Context, user dmuser.User, password string) (dmuser.User, error)
	// Inquiry check token validity to session microservice.
	// If valid, send inquiry request to payment microservice.
	Inquiry(ctx context.Context, billNumber, productCode, token string) (dmbiller.Bill, error)
	// Payment check token validity to session microservice.
	// If valid, send Payment request to payment microservice.
	Payment(ctx context.Context, amount int64, inqID, token string) (dmbiller.Transaction, error)
	// GetPayloadData Validate and get payload from token
	GetPayloadData(ctx context.Context, token string) (dmsession.SessionPayload, error)
	// GetBalance return user balance, if error, it will return -1
	GetBalance(ctx context.Context, token string) (int64, error)
	// GetListProduct return all product
	GetListProduct(ctx context.Context) ([]dmproduct.Product, error)
}
