package prepared

import (
	"context"
	"database/sql"
)

type dbtx interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db dbtx) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db dbtx) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createCityStmt, err = db.PrepareContext(ctx, createCity); err != nil {
		return nil, err
	}
	if q.createVenueStmt, err = db.PrepareContext(ctx, createVenue); err != nil {
		return nil, err
	}
	if q.deleteVenueStmt, err = db.PrepareContext(ctx, deleteVenue); err != nil {
		return nil, err
	}
	if q.getCityStmt, err = db.PrepareContext(ctx, getCity); err != nil {
		return nil, err
	}
	if q.getVenueStmt, err = db.PrepareContext(ctx, getVenue); err != nil {
		return nil, err
	}
	if q.listCitiesStmt, err = db.PrepareContext(ctx, listCities); err != nil {
		return nil, err
	}
	if q.listVenuesStmt, err = db.PrepareContext(ctx, listVenues); err != nil {
		return nil, err
	}
	if q.updateCityNameStmt, err = db.PrepareContext(ctx, updateCityName); err != nil {
		return nil, err
	}
	if q.updateVenueNameStmt, err = db.PrepareContext(ctx, updateVenueName); err != nil {
		return nil, err
	}
	if q.venueCountByCityStmt, err = db.PrepareContext(ctx, venueCountByCity); err != nil {
		return nil, err
	}
	return &q, nil
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                   dbtx
	tx                   *sql.Tx
	createCityStmt       *sql.Stmt
	createVenueStmt      *sql.Stmt
	deleteVenueStmt      *sql.Stmt
	getCityStmt          *sql.Stmt
	getVenueStmt         *sql.Stmt
	listCitiesStmt       *sql.Stmt
	listVenuesStmt       *sql.Stmt
	updateCityNameStmt   *sql.Stmt
	updateVenueNameStmt  *sql.Stmt
	venueCountByCityStmt *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                   tx,
		tx:                   tx,
		createCityStmt:       q.createCityStmt,
		createVenueStmt:      q.createVenueStmt,
		deleteVenueStmt:      q.deleteVenueStmt,
		getCityStmt:          q.getCityStmt,
		getVenueStmt:         q.getVenueStmt,
		listCitiesStmt:       q.listCitiesStmt,
		listVenuesStmt:       q.listVenuesStmt,
		updateCityNameStmt:   q.updateCityNameStmt,
		updateVenueNameStmt:  q.updateVenueNameStmt,
		venueCountByCityStmt: q.venueCountByCityStmt,
	}
}
