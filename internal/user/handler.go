package user

import (
	"encoding/json"
	"net/http"
)

type handler struct {
	service *service
}

func NewHandler(service *service) *handler {
	return &handler{service: service}
}

func (h *handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req UserDataRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.service.SignUp(r.Context(), &req); err != nil {
		http.Error(w, "Create user error", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *handler) SignIn(w http.ResponseWriter, r *http.Request) {
	var req UserDataRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	token, err := h.service.SignIn(r.Context(), &req)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(token))
}
