package dmbiller

type InquiryParam struct {
	UserID      string
	BillNumber  string
	ProductCode string
}

type PaymentParam struct {
	UserID    string
	InquiryID string
	Amount    int64
}
