package handlers

import (
	"encoding/json"
	"net/http"
)

func DoStatus(w http.ResponseWriter, data interface{}, httpStatus int) {
	w.WriteHeader(httpStatus)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func Created(w http.ResponseWriter, data interface{}) {
	DoStatus(w, data, http.StatusCreated)
}

func OK(w http.ResponseWriter, data interface{}) {
	DoStatus(w, data, http.StatusOK)
}

func NoContent(w http.ResponseWriter) {
	DoStatus(w, nil, http.StatusNoContent)
}

func NotFound(w http.ResponseWriter) {
	DoStatus(w, nil, http.StatusNotFound)
}
