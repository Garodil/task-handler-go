package main

import (
	"strconv"
	"sync"
)

// Класс задачи
type todo struct {
	sync.Mutex        // Мьютекс для последовательных изменений в случае параллельных запросов
	id         int    // Идентификатор
	title      string // Название задачи
	completed  bool   // Выполнена ли задача
}

// Назначает, что задача выполнена (completed -> true)
func (td *todo) Complete() {
	td.Lock()
	defer td.Unlock()
	td.completed = true
}

// Назначает, что задача не выполнена (completed -> false)
func (td *todo) Reset() {
	td.Lock()
	defer td.Unlock()
	td.completed = false
}

// Назначает новое название задаче
func (td *todo) Rename(title string) {
	td.Lock()
	defer td.Unlock()
	td.title = title
}

// Возвращает задачу как словарь
func (td *todo) Get() map[string]string {
	var id string = strconv.Itoa(td.id)
	var title string = td.title
	var completed string = strconv.FormatBool(td.completed)

	return map[string]string{
		"id":        id,
		"title":     title,
		"completed": completed,
	}
}
