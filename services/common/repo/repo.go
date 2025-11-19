package repo

import (
	"context"

	"github.com/ak-repo/stream-hub/pkg/db"
	"github.com/jackc/pgx/v5"
)

type CommonRepository struct {
	DB *db.PostgresDB
}

func NewCommonRepository(pgDB *db.PostgresDB) *CommonRepository {
	return &CommonRepository{DB: pgDB}
}

// QueryRow executes a query that returns a single row
func (r *CommonRepository) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return r.DB.Pool.QueryRow(ctx, query, args...)
}
