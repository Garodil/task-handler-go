package main

import (
	"log"
	"net/http"
)

// Посылает ответ клиенту
func Respond(w http.ResponseWriter, status int, message any) {
	response := ParseJson(message)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err := w.Write(response)
	if err != nil {
		log.Println(err)
	}
}

// Посылает только статус в ответе
func RespondStatus(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
}

// Посылает ответ с ошибкой
func RespondError(w http.ResponseWriter, status int, message string) {
	response := ParseJson(map[string]string{"error": message})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err := w.Write(response)
	if err != nil {
		log.Println(err)
	}
}
