package chi

import (
	"context"
	"errors"
	"net/http"
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

	session, err := adapter.brokerService.Login(context.Background(), req.Username, req.Password)

	if err != nil {
		adapter.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "",
		Data:    session,
	}

	adapter.writeJSON(w, http.StatusOK, payload)
}
