package grpcclient

import (
	"context"
	"io"

	dmbiller "github.com/fbriansyah/micro-broker-service/internal/application/domain/biller"
	dmproduct "github.com/fbriansyah/micro-broker-service/internal/application/domain/product"
	"github.com/fbriansyah/micro-broker-service/internal/port"
	"github.com/fbriansyah/micro-broker-service/util"
	"github.com/fbriansyah/micro-payment-proto/protogen/go/payment"
	"github.com/google/uuid"
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
		BaseAmount: int64(response.DetailBill.BaseAmount),
		FineAmount: int64(response.DetailBill.FineAmount),
		Amount:     int64(response.TotalAmount),
		Name:       response.Name,
	}, nil
}

func (a *PaymentClientAdapter) Payment(ctx context.Context, params dmbiller.PaymentParam) (dmbiller.Transaction, error) {
	response, err := a.client.Payment(ctx, &payment.PaymentRequest{
		UserId: params.UserID,
		InqId:  params.InquiryID,
	})
	if err != nil {
		return dmbiller.Transaction{}, err
	}
	return dmbiller.Transaction{
		BillNumber:          response.BillNumber,
		ProductCode:         response.ProductCode,
		Name:                response.Name,
		TotalAmount:         response.TotalAmount,
		RefferenceNumber:    response.RefferenceNumber,
		TransactionDatetime: util.FromDateTime(response.TransactionDatetime),
	}, nil
}

// GetBalance from userid. If error, its will return -1
func (a *PaymentClientAdapter) GetBalance(ctx context.Context, userID uuid.UUID) (int64, error) {
	resp, err := a.client.GetBalance(ctx, &payment.GetBalanceRequest{UserId: userID.String()})
	if err != nil {
		return -1, err
	}

	return int64(resp.Balance), nil
}

func (a *PaymentClientAdapter) GetListProduct(ctx context.Context) ([]dmproduct.Product, error) {
	products := []dmproduct.Product{}

	stream, err := a.client.ListProduct(ctx, &payment.ListProductRequest{})
	if err != nil {
		return products, err
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return products, err
		}

		products = append(products, dmproduct.Product{
			Code: res.Code,
			Name: res.Name,
		})
	}

	return products, nil
}
