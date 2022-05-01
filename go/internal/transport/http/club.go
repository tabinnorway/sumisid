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
	services "github.com/tabinnorway/sumisid/go/internal/services"
)

type ClubService interface {
	CreateClub(ctx context.Context, dc services.Club) (services.Club, error)
	UpdateClub(ctx context.Context, id int, newDc services.Club) (services.Club, error)
	GetAllClub(ctx context.Context) ([]services.Club, error)
	GetClub(ctx context.Context, id int) (services.Club, error)
	DeleteClub(ctx context.Context, id int) error
}

func (h *Handler) PostClub(w http.ResponseWriter, r *http.Request) {
	var newDc services.Club
	if err := json.NewDecoder(r.Body).Decode(&newDc); err != nil {
		log.Print(err)
	}
	dc, err := h.ClubService.CreateClub(r.Context(), newDc)
	if err != nil {
		log.Print(err)
	}

	if err := json.NewEncoder(w).Encode(dc); err != nil {
		panic(err)
	}
}

func (h *Handler) PutClub(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	oldDc, err := h.ClubService.GetClub(r.Context(), id)
	if err != nil || oldDc.Id != id {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var newDc services.Club
	if err := json.NewDecoder(r.Body).Decode(&newDc); err != nil {
		log.Print(err)
	}

	dc, err := h.ClubService.UpdateClub(r.Context(), id, newDc)
	if err != nil {
		panic(err)
	}
	if err := json.NewEncoder(w).Encode(dc); err != nil {
		panic(err)
	}
	println(dc.Name)
}

func (h *Handler) GetAllClub(w http.ResponseWriter, r *http.Request) {
	dcs, err := h.ClubService.GetAllClub(r.Context())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			NotFoundError(w, "Club")
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}

	if err := json.NewEncoder(w).Encode(dcs); err != nil {
		panic(err)
	}
}

func (h *Handler) GetClub(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dc, err := h.ClubService.GetClub(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			NotFoundError(w, "Club")
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}

	if err := json.NewEncoder(w).Encode(dc); err != nil {
		panic(err)
	}
}

func (h *Handler) DeleteClub(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = h.ClubService.GetClub(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			NotFoundError(w, "Club")
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.ClubService.DeleteClub(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	EntityWasDeleted(w, "Club")
}
