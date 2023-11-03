package db

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// Client клиент для работы с БД
type Client interface {
	DB() DB
	Close() error
}

// TxManager менеджер транзакций, который выполняет указанный пользователем обработчик в транзакции
type TxManager interface {
	ReadCommitted(ctx context.Context, f func(ctx context.Context) error) error
}

// Query обертка над запросом, хранящая имя запроса и сам запрос
// Имя запроса используется для логирования и потенциально может использоваться еще где-то, например, для трейсинга
type Query struct {
	Name     string
	QueryRaw string
	Args     []interface{}
}

// Transactor интерфейс для работы с транзакциями
type Transactor interface {
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

// SQLExecutor комбинирует NamedExecutor и QueryExecutor
type SQLExecutor interface {
	NamedExecutor
	QueryExecutor
}

// NamedExecutor интерфейс для работы с именованными запросами с помощью тегов в структурах
type NamedExecutor interface {
	ScanOne(ctx context.Context, dest interface{}, q Query) error
	ScanAll(ctx context.Context, dest interface{}, q Query) error
}

// QueryExecutor интерфейс для работы с обычными запросами
type QueryExecutor interface {
	Exec(ctx context.Context, q Query) (pgconn.CommandTag, error)
	Query(ctx context.Context, q Query) (pgx.Rows, error)
	QueryRow(ctx context.Context, q Query) pgx.Row
}

// Pinger интерфейс для проверки соединения с БД
type Pinger interface {
	Ping(ctx context.Context) error
}

// DB интерфейс для работы с БД
type DB interface {
	SQLExecutor
	Transactor
	Pinger
	Close()
}
