package chi

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

func (adapter *ChiAdapter) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := adapter.readJSON(w, r, &req)
	if err != nil {
		adapter.errorJSON(w, errors.New("invalid param"), http.StatusBadRequest)
		return
	}

	usr, session, err := adapter.brokerService.Login(context.Background(), req.Username, req.Password)
	if err != nil {
		adapter.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "",
		Data: map[string]any{
			"user":    usr,
			"session": session,
		},
	}

	adapter.writeJSON(w, http.StatusOK, payload)
}

func (adapter *ChiAdapter) Inquiry(w http.ResponseWriter, r *http.Request) {
	authHeader := strings.Split(r.Header.Get("authorization"), " ")
	if len(authHeader) != 2 {
		adapter.errorJSON(w, errors.New("authorization header not found"), http.StatusUnauthorized)
		return
	}
	token := authHeader[1]

	var req struct {
		BillNumber  string `json:"bill_number"`
		ProductCode string `json:"product_code"`
	}

	err := adapter.readJSON(w, r, &req)
	if err != nil {
		adapter.errorJSON(w, err)
		return
	}

	bill, err := adapter.brokerService.Inquiry(context.Background(), req.BillNumber, req.ProductCode, token)
	if err != nil {
		adapter.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "",
		Data:    bill,
	}

	adapter.writeJSON(w, http.StatusOK, payload)
}

func (adapter *ChiAdapter) Payment(w http.ResponseWriter, r *http.Request) {
	authHeader := strings.Split(r.Header.Get("authorization"), " ")
	if len(authHeader) != 2 {
		adapter.errorJSON(w, errors.New("authorization header not found"), http.StatusUnauthorized)
		return
	}
	token := authHeader[1]

	var req struct {
		InquiryID string `json:"inquiry_id"`
		Amount    int64  `json:"amount"`
	}

	err := adapter.readJSON(w, r, &req)
	if err != nil {
		adapter.errorJSON(w, err)
		return
	}

	trx, err := adapter.brokerService.Payment(context.Background(), req.Amount, req.InquiryID, token)
	if err != nil {
		adapter.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "",
		Data:    trx,
	}

	adapter.writeJSON(w, 200, payload)
}

func (adapter *ChiAdapter) GetBalance(w http.ResponseWriter, r *http.Request) {
	authHeader := strings.Split(r.Header.Get("authorization"), " ")
	if len(authHeader) != 2 {
		adapter.errorJSON(w, errors.New("authorization header not found"), http.StatusUnauthorized)
		return
	}
	token := authHeader[1]

	balance, err := adapter.brokerService.GetBalance(context.Background(), token)

	payload := jsonResponse{
		Error:   false,
		Message: "",
		Data:    -1,
	}

	if err != nil {
		payload.Error = true
		payload.Message = err.Error()

		adapter.writeJSON(w, 500, payload)
		return
	}

	payload.Data = balance

	adapter.writeJSON(w, 200, payload)
}
