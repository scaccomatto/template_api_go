package data

import "github.com/google/uuid"

type Data struct {
	Id    uuid.UUID
	Name  string
	Value int
}
