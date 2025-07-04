package model

// User is a simple example. If needed, I can create two models, one for DB and one for client use (like DTO)
type User struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Lastname string `json:"lastname" db:"lastname"`
}
