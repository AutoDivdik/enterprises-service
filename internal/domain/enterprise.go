package domain

import "github.com/google/uuid"

type Enterprise struct {
	ID              uuid.UUID
	Name            string
	Country         string
	MaintenanceYear int
	Phone           string
	Fax             string
	TypeOfOwnership *TypeOfOwnership
}

func NewEnterprise(name, country string, maintenanceYear int, phone, fax string) *Enterprise {
	return &Enterprise{
		ID:              uuid.New(),
		Name:            name,
		Country:         country,
		MaintenanceYear: maintenanceYear,
		Phone:           phone,
		Fax:             fax,
	}
}
