// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.1

package querytest

import (
	"database/sql"
)

type Foo struct {
	A int32
	B int32
	C sql.NullInt32
}
