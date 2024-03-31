package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	responseWithJson(w, 200, struct{}{})
}

func handlerError(w http.ResponseWriter, r *http.Request) {
	responseWithError(w, 400, "Something went wrong")
}

func responseWithError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Println("Responding with a 5XX error", message)
	}

	type errorResponse struct {
		Error string `json:"error"` //Vai criar um json com um campo error
	}

	responseWithJson(w, code, errorResponse{
		Error: message,
	})
}

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Fatal("Error to marshal response %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
