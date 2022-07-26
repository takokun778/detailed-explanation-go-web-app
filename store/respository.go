package store

import (
	"context"
	"database/sql"
	"todo/clock"

	"github.com/jmoiron/sqlx"
)

type Beginner interface {
	BeginTx(context.Context, *sql.TxOptions) (*sql.Tx, error)
}

type Preparer interface {
	PreparexContext(context.Context, string) (*sqlx.Stmt, error)
}

type Execer interface {
	ExecContext(context.Context, string, ...any) (sql.Result, error)
	NamedExecContext(context.Context, string, interface{}) (sql.Result, error)
}

type Queryer interface {
	Preparer
	QueryxContext(context.Context, string, ...any) (*sqlx.Rows, error)
	QueryRowxContext(context.Context, string, ...any) *sqlx.Row
	GetContext(context.Context, interface{}, string, ...any) error
	SelectContext(context.Context, interface{}, string, ...any) error
}

var (
	_ Beginner = (*sqlx.DB)(nil)
	_ Preparer = (*sqlx.DB)(nil)
	_ Queryer  = (*sqlx.DB)(nil)
	_ Execer   = (*sqlx.DB)(nil)
	_ Execer   = (*sqlx.Tx)(nil)
)

type Repository struct {
	Clocker clock.Clocker
}
