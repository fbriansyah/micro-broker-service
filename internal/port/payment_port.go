package port

import (
	"context"

	"github.com/fbriansyah/micro-payment-proto/protogen/go/payment"

	"google.golang.org/grpc"
)

type PaymentPort interface {
	Inquiry(ctx context.Context, in *payment.InquiryRequest, opts ...grpc.CallOption) (*payment.InquiryResponse, error)
	Payment(ctx context.Context, in *payment.PaymentRequest, opts ...grpc.CallOption) (*payment.PaymentResponse, error)
}
