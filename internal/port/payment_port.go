package port

import (
	"context"

	"github.com/fbriansyah/micro-payment-proto/protogen/go/payment"

	"google.golang.org/grpc"
)

type PaymentPort interface {
	Inquiry(ctx context.Context, in *payment.InquiryRequest, opts ...grpc.CallOption) (*payment.InquiryResponse, error)
	Payment(ctx context.Context, in *payment.PaymentRequest, opts ...grpc.CallOption) (*payment.PaymentResponse, error)
	ListProduct(ctx context.Context, in *payment.ListProductRequest, opts ...grpc.CallOption) (payment.PaymentService_ListProductClient, error)
	GetBalance(ctx context.Context, in *payment.GetBalanceRequest, opts ...grpc.CallOption) (*payment.GetBalanceResponse, error)
}
