package grpcclient

import (
	"context"

	dmbiller "github.com/fbriansyah/micro-broker-service/internal/application/domain/biller"
	"github.com/fbriansyah/micro-broker-service/internal/port"
	"github.com/fbriansyah/micro-payment-proto/protogen/go/payment"
	"google.golang.org/grpc"
)

type PaymentClientAdapter struct {
	client port.PaymentPort
}

func NewPaymentClientAdapter(conn *grpc.ClientConn) *PaymentClientAdapter {
	return &PaymentClientAdapter{
		client: payment.NewPaymentServiceClient(conn),
	}
}

func (a *PaymentClientAdapter) Inquiry(ctx context.Context, params dmbiller.InquiryParam) (dmbiller.Bill, error) {
	response, err := a.client.Inquiry(ctx, &payment.InquiryRequest{
		UserId:      params.UserID,
		BillNumber:  params.BillNumber,
		ProductCode: params.ProductCode,
	})

	if err != nil {
		return dmbiller.Bill{}, err
	}

	return dmbiller.Bill{
		InquiryID:  response.InqId,
		BillNumber: response.BillNumber,
		Amount:     int64(response.TotalAmount),
		Name:       response.Name,
	}, nil
}

func (a *PaymentClientAdapter) Payment(ctx context.Context, params dmbiller.PaymentParam) error {
	a.client.Payment(ctx, &payment.PaymentRequest{
		UserId: params.UserID,
		InqId:  params.InquiryID,
	})
	return nil
}
