// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type Todo struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID string `json:"userId"`
}

type UpdateTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
	TodoID string `json:"todoId"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
