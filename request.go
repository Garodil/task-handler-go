package main

// Основной класс запроса
type MainRequest struct{}

// Класс POST запроса
type PostRequest struct {
	*MainRequest
	Title string `json:"title"` // Название задачи
}

// Класс PUT запроса
type PutRequest struct {
	*MainRequest
	Title     string `json:"title"`     // Новое название для задачи
	Completed string `json:"completed"` // Должно быть true или false
}
