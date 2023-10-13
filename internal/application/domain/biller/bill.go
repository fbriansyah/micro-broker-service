package dmbiller

type Bill struct {
	InquiryID  string `json:"inquiry_id"`
	Name       string `json:"name"`
	BillNumber string `json:"bill_number"`
	BaseAmount int64  `json:"base_amount"`
	FineAmount int64  `json:"fine_amount"`
	Amount     int64  `json:"amount"`
}
