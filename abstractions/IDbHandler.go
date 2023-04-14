package abstractions

import (
	"context"
	"database/sql"
)

type IDbHandler interface {
	// HandlerName(string) string
	BeginTransaction(ctx context.Context, opts *sql.TxOptions) error
	Commit() error
	Rollback() error
	Mutation(ctx context.Context, query string, args ...any) (sql.Result, error)
	ReadSingle(ctx context.Context, query string, args ...any) *sql.Row
	Read(ctx context.Context, query string, args ...any) (*sql.Rows, error)
}
