package dmbiller

type Bill struct {
	InquiryID  string `json:"inquiry_id"`
	Name       string `json:"name"`
	BillNumber string `json:"bill_number"`
	Amount     int64  `json:"amount"`
}
