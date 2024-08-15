package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type HomeHandler struct {
	db *sql.DB
}

type HomeResponse struct {
	Status string `json:"status"`
}

func NewHomeHandler(db *sql.DB) http.Handler {
	return &HomeHandler{
		db: db,
	}
}

func (h *HomeHandler) HomeGet(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{Status: "ok"}

	w.Header().Set("Content-Type", "application/json")

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(jsonResponse)

	w.WriteHeader(http.StatusOK)
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.HomeGet(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
