package domain

import "github.com/google/uuid"

type TypeOfOwnership struct {
	ID   uuid.UUID
	Name string
}

func NewTypeOfOwnership(name string) *TypeOfOwnership {
	return &TypeOfOwnership{
		ID:   uuid.New(),
		Name: name,
	}
}
