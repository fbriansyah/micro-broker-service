package application

import (
	"context"
	"errors"

	dmbiller "github.com/fbriansyah/micro-broker-service/internal/application/domain/biller"
	dmproduct "github.com/fbriansyah/micro-broker-service/internal/application/domain/product"
	dmsession "github.com/fbriansyah/micro-broker-service/internal/application/domain/session"
	dmuser "github.com/fbriansyah/micro-broker-service/internal/application/domain/user"
	"github.com/fbriansyah/micro-broker-service/internal/port"
)

var ErrorInvalidToken = errors.New("invalid token")

type BrokerService struct {
	authClient    port.AuthAdapterPort
	sessionClient port.SessionAdapterPort
	paymentClient port.PaymentAdapterPort
}

type BrokerClientConfig struct {
	AuthClient    port.AuthAdapterPort
	SessionClient port.SessionAdapterPort
	PaymentClient port.PaymentAdapterPort
}

func NewBrokerSerice(clientConfig BrokerClientConfig) *BrokerService {
	return &BrokerService{
		authClient:    clientConfig.AuthClient,
		sessionClient: clientConfig.SessionClient,
		paymentClient: clientConfig.PaymentClient,
	}
}

func (s *BrokerService) Login(ctx context.Context, username, password string) (dmuser.User, dmsession.Session, error) {
	usr, err := s.authClient.Login(ctx, username, password)
	if err != nil {
		return dmuser.User{}, dmsession.Session{}, nil
	}

	session, err := s.sessionClient.CreateSession(ctx, usr.ID)
	if err != nil {
		return dmuser.User{}, dmsession.Session{}, err
	}

	return usr, session, nil
}

func (s *BrokerService) Register(ctx context.Context, user dmuser.User, password string) (dmuser.User, error) {
	return s.authClient.Register(ctx, user, password)
}

func (s *BrokerService) Inquiry(ctx context.Context, billNumber, productCode, token string) (dmbiller.Bill, error) {
	payload, err := s.sessionClient.GetPayloadData(ctx, token)
	if err != nil {
		return dmbiller.Bill{}, ErrorInvalidToken
	}

	arg := dmbiller.InquiryParam{
		UserID:      payload.UserID.String(),
		BillNumber:  billNumber,
		ProductCode: productCode,
	}

	bill, err := s.paymentClient.Inquiry(ctx, arg)
	if err != nil {
		return dmbiller.Bill{}, err
	}

	return bill, nil
}

func (s *BrokerService) Payment(ctx context.Context, amount int64, inqID, token string) (dmbiller.Transaction, error) {
	payload, err := s.sessionClient.GetPayloadData(ctx, token)
	if err != nil {
		return dmbiller.Transaction{}, ErrorInvalidToken
	}

	trx, err := s.paymentClient.Payment(ctx, dmbiller.PaymentParam{
		UserID:    payload.UserID.String(),
		InquiryID: inqID,
		Amount:    amount,
	})
	if err != nil {
		return dmbiller.Transaction{}, ErrorInvalidToken
	}

	return trx, nil
}

func (s *BrokerService) GetPayloadData(ctx context.Context, token string) (dmsession.SessionPayload, error) {
	return s.sessionClient.GetPayloadData(ctx, token)
}

func (s *BrokerService) GetBalance(ctx context.Context, token string) (int64, error) {
	payload, err := s.GetPayloadData(ctx, token)
	if err != nil {
		return -1, err
	}

	return s.paymentClient.GetBalance(ctx, payload.UserID)
}

func (s *BrokerService) GetListProduct(ctx context.Context) ([]dmproduct.Product, error) {
	return s.paymentClient.GetListProduct(ctx)
}
