// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.12.0

package querytest

import (
	"database/sql"

	"github.com/google/uuid"
)

type Organisation struct {
	PartyID   uuid.UUID
	Name      string
	LegalName sql.NullString
}

type Party struct {
	PartyID uuid.UUID
	Name    string
}

type Person struct {
	PartyID   uuid.UUID
	Name      string
	FirstName string
	LastName  string
}
