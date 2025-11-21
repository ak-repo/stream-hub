package repo

import (
	"context"

	"github.com/ak-repo/stream-hub/pkg/db"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type CommonRepository struct {
	DB *db.PostgresDB
}

func NewCommonRepository(pgDB *db.PostgresDB) *CommonRepository {
	return &CommonRepository{DB: pgDB}
}

//
// BASIC QUERY HELPERS
//

// QueryRow executes a query that returns a single row
// QueryRow → SELECT … WHERE id=?
func (r *CommonRepository) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return r.DB.Pool.QueryRow(ctx, query, args...)
}

// Query → SELECT * FROM table
func (r *CommonRepository) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	return r.DB.Pool.Query(ctx, query, args...)
}

// Exec → INSERT, UPDATE, DELETE (no returning row)
func (r *CommonRepository) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return r.DB.Pool.Exec(ctx, query, args...)
}

//
// GENERIC CRUD HELPERS (OPTIONAL)
//

// Insert returns inserted ID
func (r *CommonRepository) InsertAndReturnID(ctx context.Context, query string, args ...interface{}) (int64, error) {
	var id int64
	err := r.DB.Pool.QueryRow(ctx, query, args...).Scan(&id)
	return id, err
}

// Exists → check if row exists
func (r *CommonRepository) Exists(ctx context.Context, query string, args ...interface{}) (bool, error) {
	var exists bool
	err := r.DB.Pool.QueryRow(ctx, query, args...).Scan(&exists)
	return exists, err
}

// Count → return integer count
func (r *CommonRepository) Count(ctx context.Context, query string, args ...interface{}) (int64, error) {
	var count int64
	err := r.DB.Pool.QueryRow(ctx, query, args...).Scan(&count)
	return count, err
}

// Delete → return number of affected rows
func (r *CommonRepository) Delete(ctx context.Context, query string, args ...interface{}) (int64, error) {
	tag, err := r.DB.Pool.Exec(ctx, query, args...)
	return tag.RowsAffected(), err
}
