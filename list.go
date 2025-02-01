package main

import (
	"errors"
	"strconv"
	"sync"
)

var List = CreateTodoList()

// Класс списка задач
type list struct {
	sync.Mutex        // Мьютекс для последовательных добавлений в список в случае параллельных запросов
	LastId     int    // Последнее созданное id
	todos      []todo // Массив задач
}

// Создаёт список задач
func CreateTodoList() *list {
	return &list{}
}

// Добавляет задачу в список задач и возвращает указатель на эту задачу
func (l *list) Add(title string) *todo {
	l.Lock()
	defer l.Unlock()
	l.LastId++
	l.todos = append(l.todos, todo{id: l.LastId, title: title, completed: false})
	return &l.todos[len(l.todos)-1]
}

// Возвращает список задач в виде массива с строковыми словарями, может вернуть пустой массив
func (l *list) Get() []map[string]string {
	var todos []map[string]string = []map[string]string{}

	for i := range l.todos {
		var index string = strconv.Itoa(i)
		var id string = strconv.Itoa(l.todos[i].id)
		var title string = l.todos[i].title
		var completed string = strconv.FormatBool(l.todos[i].completed)

		todo := map[string]string{
			"index":     index,
			"id":        id,
			"title":     title,
			"completed": completed,
		}
		todos = append(todos, todo)
	}
	return todos
}

// Осуществляет линейный поиск на задачу по идентификатору и передаёт указатель на неё, в случае, если задача не найдена - возвращает пустое значение
func (l *list) Find(id int) *todo {
	for index := range l.todos {
		if id == l.todos[index].id {
			return &l.todos[index]
		}
	}
	return nil
}

// Осуществляет линейный поиск на задачу по идентификатору и удаляет её
func (l *list) Delete(id int) error {
	for index := range l.todos {
		if id == l.todos[index].id {
			l.todos = append(l.todos[:index], l.todos[index+1:]...)
			return nil
		}
	}
	return errors.New("задание не найдено")
}
