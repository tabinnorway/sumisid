package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tabinnorway/sumisid/go/internal/diveclub"
)

type DiveClubService interface {
	CreateDiveClub(ctx context.Context, dc diveclub.DiveClub) (diveclub.DiveClub, error)
	UpdateDiveClub(ctx context.Context, id int, newDc diveclub.DiveClub) (diveclub.DiveClub, error)
	GetDiveClub(ctx context.Context, id int) (diveclub.DiveClub, error)
	DeleteDiveClub(ctx context.Context, id int) error
}

func (h *Handler) PostDiveClub(w http.ResponseWriter, r *http.Request) {
	var newDc diveclub.DiveClub
	if err := json.NewDecoder(r.Body).Decode(&newDc); err != nil {
		log.Print(err)
	}
	dc, err := h.Service.CreateDiveClub(r.Context(), newDc)
	if err != nil {
		log.Print(err)
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dc); err != nil {
		panic(err)
	}
}

func (h *Handler) PutDiveClub(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var newDc diveclub.DiveClub
	if err := json.NewDecoder(r.Body).Decode(&newDc); err != nil {
		log.Print(err)
	}

	dc, err := h.Service.UpdateDiveClub(r.Context(), id, newDc)
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dc); err != nil {
		panic(err)
	}
}

func (h *Handler) GetDiveClub(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dc, err := h.Service.GetDiveClub(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dc); err != nil {
		panic(err)
	}
}
func (h *Handler) DeleteDiveClub(w http.ResponseWriter, r *http.Request) {

}
