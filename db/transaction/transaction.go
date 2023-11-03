package transaction

import (
	"context"
	"fmt"
	"github.com/north70/course-platform-pkg/db"
	"github.com/north70/course-platform-pkg/db/pg"

	"github.com/jackc/pgx/v4"
)

type manager struct {
	db db.Transactor
}

// NewTransactionManager создает новый менеджер транзакций, который удовлетворяет интерфейсу db.TxManager
func NewTransactionManager(db db.Transactor) db.TxManager {
	return &manager{
		db: db,
	}
}

// Transaction основная функция, которая выполняет указанный пользователем обработчик в транзакции
func (m *manager) transaction(ctx context.Context, opts pgx.TxOptions, fn func(ctx context.Context) error) (err error) {
	// Если это вложенная транзакция, пропускаем инициацию новой транзакции и выполняем обработчик.
	tx, ok := ctx.Value(pg.TxKey).(pgx.Tx)
	if ok {
		return fn(ctx)
	}

	// Стартуем новую транзакцию.
	tx, err = m.db.BeginTx(ctx, opts)
	if err != nil {
		return fmt.Errorf("can't begin transaction: %w", err)
	}

	// Кладем транзакцию в контекст.
	ctx = pg.MakeContextTx(ctx, tx)

	// Настраиваем функцию отсрочки для отката или коммита транзакции.
	defer func() {
		// восстанавливаемся после паники
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered: %v", r)
		}

		// откатываем транзакцию, если произошла ошибка
		if err != nil {
			if errRollback := tx.Rollback(ctx); errRollback != nil {
				err = fmt.Errorf("errRollback - %v: %w", errRollback, err)
			}

			return
		}

		// если ошибок не было, коммитим транзакцию
		err = tx.Commit(ctx)
		if err != nil {
			err = fmt.Errorf("tx commit failed: %w", err)
		}
	}()

	// Выполните код внутри транзакции.
	// Если функция терпит неудачу, возвращаем ошибку, и функция отсрочки выполняет откат
	// или в противном случае транзакция коммитится.
	if err = fn(ctx); err != nil {
		err = fmt.Errorf("failed executing code inside transaction: %w", err)
	}

	return err
}

func (m *manager) ReadCommitted(ctx context.Context, f func(ctx context.Context) error) error {
	txOpts := pgx.TxOptions{IsoLevel: pgx.ReadCommitted}
	return m.transaction(ctx, txOpts, f)
}
