// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.12.0

package querytest

import (
	"database/sql"
)

type Bar struct {
	ID    int32
	Title sql.NullString
}

type Foo struct {
	ID int32
}
