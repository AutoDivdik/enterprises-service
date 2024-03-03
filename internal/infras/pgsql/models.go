// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package postgresql

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type EnterprisesEnterprise struct {
	ID                uuid.UUID    `json:"id"`
	Name              string       `json:"name"`
	Country           string       `json:"country"`
	Maintenanceyear   int32        `json:"maintenanceyear"`
	Phone             string       `json:"phone"`
	Fax               string       `json:"fax"`
	TypeOfOwnershipID uuid.UUID    `json:"type_of_ownership_id"`
	Created           time.Time    `json:"created"`
	Updated           sql.NullTime `json:"updated"`
}

type EnterprisesTypesOfOwnership struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
