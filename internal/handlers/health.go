package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type HealthHandler struct {
	db *sql.DB
}

type HealthResponse struct {
	Status string `json:"status"`
}

func NewHealthHandler(db *sql.DB) http.Handler {
	return &HealthHandler{
		db: db,
	}
}

type Health struct {
	ID   int
	Name string
}

func (h *HealthHandler) HealthGet(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	rows, err := h.db.Query("select 1")
	if err != nil {
		log.Printf("Error querying database: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	healthResponse := HealthResponse{
		Status: "OK",
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(healthResponse); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.HealthGet(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
