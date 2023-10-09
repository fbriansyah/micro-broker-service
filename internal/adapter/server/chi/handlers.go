package chi

import "net/http"

func (adapter *ChiAdapter) Login(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "",
		Data:    "ok",
	}

	adapter.writeJSON(w, http.StatusOK, payload)
}
