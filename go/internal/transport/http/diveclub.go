package http

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	diveclub "github.com/tabinnorway/sumisid/go/internal/services"
)

type DiveClubService interface {
	CreateDiveClub(ctx context.Context, dc diveclub.DiveClub) (diveclub.DiveClub, error)
	UpdateDiveClub(ctx context.Context, id int, newDc diveclub.DiveClub) (diveclub.DiveClub, error)
	GetAllDiveClub(ctx context.Context) ([]diveclub.DiveClub, error)
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

func (h *Handler) GetAllDiveClub(w http.ResponseWriter, r *http.Request) {
	dcs, err := h.Service.GetAllDiveClub(r.Context())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			NotFoundError(w, "Club")
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dcs); err != nil {
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
		if errors.Is(err, sql.ErrNoRows) {
			NotFoundError(w, "Club")
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dc); err != nil {
		panic(err)
	}
}

func (h *Handler) DeleteDiveClub(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = h.Service.GetDiveClub(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			NotFoundError(w, "Club")
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.Service.DeleteDiveClub(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	EntityWasDeleted(w, "Club")
	return
}
