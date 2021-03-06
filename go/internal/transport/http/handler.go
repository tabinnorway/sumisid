package http

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router        *mux.Router
	ClubService   ClubService
	PersonService PersonService
	Server        *http.Server
}

func NewHandler(dcService ClubService, pService PersonService) *Handler {
	h := &Handler{
		ClubService:   dcService,
		PersonService: pService,
	}
	h.Router = mux.NewRouter()
	h.Router.Use(JSONMiddleware)
	h.Router.Use(LoggingMiddleware)
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

	h.Router.HandleFunc("/google_login", GoogleLogin).Methods("GET")
	h.Router.HandleFunc("/google_callback", GoogleCallback).Methods("GET")

	h.Router.Handle("/", http.FileServer(http.Dir("./views")))

	h.Router.HandleFunc("/api/v1/people", h.GetAllPerson).Methods("GET")
	h.Router.HandleFunc("/api/v1/people", h.PostPerson).Methods("POST")
	h.Router.HandleFunc("/api/v1/people/{id}", h.PutPerson).Methods("PUT")
	h.Router.HandleFunc("/api/v1/people/{id}", h.GetPerson).Methods("GET")
	h.Router.HandleFunc("/api/v1/people/{id}", h.DeletePerson).Methods("DELETE")

	h.Router.HandleFunc("/api/v1/clubs", h.GetAllClub).Methods("GET")
	h.Router.HandleFunc("/api/v1/clubs/{id}", h.GetClub).Methods("GET")
	h.Router.HandleFunc("/api/v1/clubs", h.PostClub).Methods("POST")
	h.Router.HandleFunc("/api/v1/clubs/{id}", h.PutClub).Methods("PUT")
	h.Router.HandleFunc("/api/v1/clubs/{id}", h.DeleteClub).Methods("DELETE")
}

func GoogleLogin(w http.ResponseWriter, req *http.Request) {
	googleConfig := SetupConfig()
	url := googleConfig.AuthCodeURL("randomstate")
	http.Redirect(w, req, url, http.StatusSeeOther)
}

func GoogleCallback(w http.ResponseWriter, req *http.Request) {
	state := req.URL.Query()["state"][0]
	if state != "randomstate" {
		fmt.Println(w, "states don't match")
		return
	}
	code := req.URL.Query()["code"][0]
	googleConfig := SetupConfig()

	token, err := googleConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Fprintln(w, "Code-Token Exchange failed")
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Fprintln(w, "Failed to fetch user information from Google")
	}

	userData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(w, "Could not parse user data JSON")
	}

	fmt.Fprintln(w, string(userData))
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
	// until an OS Interrupt signal arrives, wait 30 seconds max
	ctx, cancal := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancal()
	h.Server.Shutdown(ctx)

	log.Println("Server shut down gracefully")
	return nil
}
