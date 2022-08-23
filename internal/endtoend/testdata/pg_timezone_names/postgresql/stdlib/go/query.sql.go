// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const getTimezones = `-- name: GetTimezones :many
SELECT name, abbrev, utc_offset, is_dst from pg_catalog.pg_timezone_names
`

type GetTimezonesRow struct {
	Name      sql.NullString
	Abbrev    sql.NullString
	UtcOffset sql.NullInt64
	IsDst     sql.NullBool
}

func (q *Queries) GetTimezones(ctx context.Context) ([]GetTimezonesRow, error) {
	rows, err := q.db.QueryContext(ctx, getTimezones)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTimezonesRow
	for rows.Next() {
		var i GetTimezonesRow
		if err := rows.Scan(
			&i.Name,
			&i.Abbrev,
			&i.UtcOffset,
			&i.IsDst,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
