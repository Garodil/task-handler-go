package main

import (
	"log"
	"net/http"
	"strconv"
)

// Отправляет все задачи в клиент
func Get(w http.ResponseWriter, r *http.Request) {
	list := List.Get()

	Respond(w, http.StatusOK, list)
}

// Создаёт новую задачу по запросу от клиента
func Post(w http.ResponseWriter, r *http.Request) {
	var request PostRequest
	err := request.Decode(r.Body, &request)
	if err != nil {
		log.Println("ошибка парса тела запроса " + err.Error())
		RespondError(w, http.StatusBadRequest, "ошибка парса")
		return
	}

	if request.Title == "" {
		log.Println("ошибка создания задачи: нет параметра 'title'")
		RespondError(w, http.StatusBadRequest, "нет 'title'")
		return
	}

	todo := List.Add(request.Title)

	// Возврат id созданной задачи
	Respond(w, http.StatusCreated, map[string]string{"id": todo.Get()["id"]})
}

// Отправляет одну задачу по идентификатору на клиент
func GetById(w http.ResponseWriter, r *http.Request) {
	pathId := r.PathValue("id")

	// Валидация pathId
	todoId, err := strconv.Atoi(pathId)
	if err != nil {
		log.Println("Неверный id в запросе: " + err.Error())
		RespondError(w, http.StatusNotFound, "неверный id")
		return
	}

	if todoId > List.LastId {
		RespondError(w, http.StatusNotFound, "id выходит за пределы списка задач")
		return
	}

	todo := List.Find(todoId)
	if todo == nil {
		RespondError(w, http.StatusNotFound, "задание не найдено")
		return
	}

	Respond(w, http.StatusOK, todo.Get())
}

// Изменяет задание
func Put(w http.ResponseWriter, r *http.Request) {
	pathId := r.PathValue("id")

	var request PutRequest
	err := request.Decode(r.Body, &request)
	if err != nil {
		log.Println("ошибка парса тела запроса " + err.Error())
		RespondError(w, http.StatusBadRequest, "ошибка парса")
		return
	}

	if request.Title == "" && request.Completed == "" {
		log.Println("пустой запрос")
		RespondError(w, http.StatusBadRequest, "пустое тело запроса")
		return
	}

	// Валидация pathId
	todoId, err := strconv.Atoi(pathId)
	if err != nil {
		log.Println("Неверный id в запросе: " + err.Error())
		RespondError(w, http.StatusNotFound, "неверный id")
		return
	}

	if todoId > List.LastId {
		RespondError(w, http.StatusNotFound, "id выходит за пределы списка задач")
		return
	}

	// Валидация Completed

	todo := List.Find(todoId)

	if request.Completed == "true" {
		todo.Complete()
	} else if request.Completed == "false" {
		todo.Reset()
	} else if request.Completed != "" {
		RespondError(w, http.StatusBadRequest, "неверное значение 'completed'")
	}

	if request.Title != "" {
		todo.Rename(request.Title)
	}

	Respond(w, http.StatusOK, map[string]string{"ok": "true"})
}

// Удаляет задание
func Delete(w http.ResponseWriter, r *http.Request) {
	pathId := r.PathValue("id")

	// Валидация pathId
	todoId, err := strconv.Atoi(pathId)
	if err != nil {
		log.Println("Неверный id в запросе: " + err.Error())
		RespondError(w, http.StatusNotFound, "неверный id")
		return
	}

	if todoId > List.LastId {
		RespondError(w, http.StatusNotFound, "id выходит за пределы списка задач")
		return
	}

	err = List.Delete(todoId)
	if err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}

	RespondStatus(w, http.StatusNoContent)
}
