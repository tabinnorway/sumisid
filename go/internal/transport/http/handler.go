package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type DiveClubService interface{}

type Handler struct {
	Router  *mux.Router
	Service DiveClubService
	Server  *http.Server
}

func NewHandler(service DiveClubService) *Handler {
	h := &Handler{
		Service: service,
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
	fmt.Println("***\n*** mapping routes")
	fmt.Println("*** /api/ping")
	h.Router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi there\n")
	})
	fmt.Println("*** /api/ping")
	h.Router.HandleFunc("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong\n")
	})
	fmt.Println("*** finished mapping routes\n***")
}

func (h *Handler) Serve() error {
	if err := h.Server.ListenAndServe(); err != nil {
		fmt.Println("*** Listen and serve failed: ", err)
		return err
	}
	return nil
}
