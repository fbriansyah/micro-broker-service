package dmbiller

import "time"

type Transaction struct {
	BillNumber          string    `json:"bill_number"`
	ProductCode         string    `json:"product_code"`
	Name                string    `json:"name"`
	TotalAmount         float64   `json:"total_amount"`
	RefferenceNumber    string    `json:"refference_number"`
	TransactionDatetime time.Time `json:"transaction_date"`
}
