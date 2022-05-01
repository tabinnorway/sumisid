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

type PersonService interface {
	CreatePerson(ctx context.Context, dc services.Person) (services.Person, error)
	UpdatePerson(ctx context.Context, id int, newDc services.Person) (services.Person, error)
	GetAllPerson(ctx context.Context) ([]services.Person, error)
	GetPerson(ctx context.Context, id int) (services.Person, error)
	DeletePerson(ctx context.Context, id int) error
}

func (h *Handler) PostPerson(w http.ResponseWriter, r *http.Request) {
	var newDc services.Person
	if err := json.NewDecoder(r.Body).Decode(&newDc); err != nil {
		log.Print(err)
	}
	dc, err := h.PersonService.CreatePerson(r.Context(), newDc)
	if err != nil {
		log.Print(err)
	}

	if err := json.NewEncoder(w).Encode(dc); err != nil {
		panic(err)
	}
}

func (h *Handler) PutPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	oldPerson, err := h.PersonService.GetPerson(r.Context(), id)
	if err != nil || oldPerson.Id != id {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var newPerson services.Person
	if err := json.NewDecoder(r.Body).Decode(&newPerson); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
		return
	}

	dc, err := h.PersonService.UpdatePerson(r.Context(), id, newPerson)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(dc); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetAllPerson(w http.ResponseWriter, r *http.Request) {
	dcs, err := h.PersonService.GetAllPerson(r.Context())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			NotFoundError(w, "Person")
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

func (h *Handler) GetPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p, err := h.PersonService.GetPerson(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			NotFoundError(w, "Person")
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}

	if p.MainClubId > 0 {
		p.MainClub, err = h.ClubService.GetClub(r.Context(), p.MainClubId)
		if err != nil {
			p.MainClub = services.Club{}
		}
	}

	if err := json.NewEncoder(w).Encode(p); err != nil {
		panic(err)
	}
}

func (h *Handler) DeletePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = h.PersonService.GetPerson(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			NotFoundError(w, "Person")
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.PersonService.DeletePerson(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	EntityWasDeleted(w, "Person")
}
