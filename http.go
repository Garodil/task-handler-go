package main

import (
	"log"
	"net/http"
)

// Сервер
var Host host = host{http.NewServeMux(), "localhost:9000"}

// Класс сервера
type host struct {
	*http.ServeMux        // Защита маршрутов
	ip             string // Адрес сервера
}

// Запускает сервер
func (h *host) Host() {
	log.Println("Hosted")
	if err := http.ListenAndServe(h.ip, h); err != nil {
		log.Fatalln(err)
	}
}

// Создаёт маршруты
func (h *host) Route() {
	h.HandleFunc("GET /todos", Get)            // возвращение всех задач из списка задач
	h.HandleFunc("POST /todos", Post)          // создание новой задачи в список задач
	h.HandleFunc("GET /todos/{id}", GetById)   // возвращение одной задачи по идентификатору
	h.HandleFunc("PUT /todos/{id}", Put)       // обновление задачи
	h.HandleFunc("DELETE /todos/{id}", Delete) // удаление задачи
}
