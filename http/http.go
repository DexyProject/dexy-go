package http

import (
	"fmt"
	"net/http"
)

type Handler func(http.ResponseWriter, *http.Request) error

func (fn Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := fn(w, r)
	if err == nil {
		return
	}

	switch e := err.(type) {
	case Error:
		returnError(w, e.Message, e.Code)
	default:
		returnError(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

type Error struct {
	Message       string
	Code          int
}

func NewError(message string, code int) Error {
	return Error{Message: message, Code: code}
}

func (e Error) Error() string {
	return e.Message
}

func NotFound(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusNotFound)
	rw.Write([]byte(fmt.Sprintf("{\"error\": \"%s\", \"request\": \"%s\"}", "not found", r.Method + " " + r.URL.RequestURI())))
}

func returnError(w http.ResponseWriter, err string, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf("{\"error\": \"%v\"}", err)))
}
