package pg

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/north70/course-platform-pkg/db"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
)

type key string

const (
	TxKey key = "tx"
)

type pg struct {
	dbc *pgxpool.Pool
}

func NewDB(dbc *pgxpool.Pool) db.DB {
	return &pg{
		dbc: dbc,
	}
}

func (p *pg) ScanOne(ctx context.Context, dest interface{}, q db.Query) error {
	row, err := p.Query(ctx, q)
	if err != nil {
		return err
	}

	return pgxscan.ScanOne(dest, row)
}

func (p *pg) ScanAll(ctx context.Context, dest interface{}, q db.Query) error {
	rows, err := p.Query(ctx, q)
	if err != nil {
		return err
	}

	return pgxscan.ScanAll(dest, rows)
}

func (p *pg) Exec(ctx context.Context, q db.Query) (pgconn.CommandTag, error) {
	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.Exec(ctx, q.QueryRaw, q.Args...)
	}

	return p.dbc.Exec(ctx, q.QueryRaw, q.Args...)
}

func (p *pg) Query(ctx context.Context, q db.Query) (pgx.Rows, error) {
	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.Query(ctx, q.QueryRaw, q.Args...)
	}

	return p.dbc.Query(ctx, q.QueryRaw, q.Args...)
}

func (p *pg) QueryRow(ctx context.Context, q db.Query) pgx.Row {
	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.QueryRow(ctx, q.QueryRaw, q.Args...)
	}

	return p.dbc.QueryRow(ctx, q.QueryRaw, q.Args...)
}

func (p *pg) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	return p.dbc.BeginTx(ctx, txOptions)
}

func (p *pg) Ping(ctx context.Context) error {
	return p.dbc.Ping(ctx)
}

func (p *pg) Close() {
	p.dbc.Close()
}

func MakeContextTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, TxKey, tx)
}
