// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.12.0
// source: query.sql

package querytest

import (
	"context"
)

const listUsersByRole = `-- name: ListUsersByRole :many
SELECT role FROM foo.users WHERE role = $1
`

func (q *Queries) ListUsersByRole(ctx context.Context, role FooTypeUserRole) ([]FooTypeUserRole, error) {
	rows, err := q.db.Query(ctx, listUsersByRole, role)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FooTypeUserRole
	for rows.Next() {
		var role FooTypeUserRole
		if err := rows.Scan(&role); err != nil {
			return nil, err
		}
		items = append(items, role)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
