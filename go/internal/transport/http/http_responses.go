package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponseMessage struct {
	Message      string
	ErrorMessage string
	StatusCode   int
}

func NotFoundError(w http.ResponseWriter, entityName string) {
	resp := ResponseMessage{
		Message:      "",
		ErrorMessage: fmt.Sprintf("The %s was not found or user does not have access", entityName),
		StatusCode:   http.StatusNotFound,
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		panic(err)
	}
}

func EntityWasDeleted(w http.ResponseWriter, entityName string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
