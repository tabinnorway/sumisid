package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router          *mux.Router
	DiveClubService DiveClubService
	PersonService   PersonService
	Server          *http.Server
}

func NewHandler(dcService DiveClubService, pService PersonService) *Handler {
	h := &Handler{
		DiveClubService: dcService,
		PersonService:   pService,
	}
	h.Router = mux.NewRouter()
	h.mapRoutes()

	h.Server = &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: h.Router,
	}
	return h
}

func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong\n")
	})

	h.Router.HandleFunc("/api/v1/people", h.GetAllPerson).Methods("GET")
	h.Router.HandleFunc("/api/v1/people", h.PostPerson).Methods("POST")
	h.Router.HandleFunc("/api/v1/people/{id}", h.PutPerson).Methods("PUT")
	h.Router.HandleFunc("/api/v1/people/{id}", h.GetPerson).Methods("GET")
	h.Router.HandleFunc("/api/v1/people/{id}", h.DeletePerson).Methods("DELETE")

	h.Router.HandleFunc("/api/v1/diveclubs", h.GetAllDiveClub).Methods("GET")
	h.Router.HandleFunc("/api/v1/diveclubs", h.PostDiveClub).Methods("POST")
	h.Router.HandleFunc("/api/v1/diveclubs/{id}", h.PutDiveClub).Methods("PUT")
	h.Router.HandleFunc("/api/v1/diveclubs/{id}", h.GetDiveClub).Methods("GET")
	h.Router.HandleFunc("/api/v1/diveclubs/{id}", h.DeleteDiveClub).Methods("DELETE")
}

func (h *Handler) Serve() error {
	fmt.Println("Listening to ", h.Server.Addr)

	// Run ListenAndServe in a separate go thread
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println("*** Listen and serve failed: ", err)
		}
	}()

	// Wait for an os.Interrupt signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Since ListenAndServe runs in a go thread, we need to wait here
	// until an OS Interrupt signal arrives, wait 15 seconds
	ctx, cancal := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancal()
	h.Server.Shutdown(ctx)

	log.Println("Server shut down gracefully")
	return nil
}
