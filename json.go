package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Десериализация из io.ReadCloser
func (request *MainRequest) Decode(body io.ReadCloser, result any) error {
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&result)
	if err != nil {
		return err
	}

	return nil
}

// Десериализация из json в result
func FormatJson(input []byte, result any) {
	json.Unmarshal(input, &result)
}

// Сериализация в json
func ParseJson(input any) []byte {
	bytes, err := json.Marshal(&input)
	if err != nil {
		return []byte{}
	}

	return bytes
}

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
