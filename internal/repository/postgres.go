package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	userTable           = "users"
	refreshSessionTable = "refresh_sessions"
)

type store struct {
	*sqlx.DB
}

type ConfigPostgres struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg *ConfigPostgres) (*store, error) {
	dataSource := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)

	db, err := sqlx.Connect("pgx", dataSource)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &store{db}, nil
}

func (s *store) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	tx := extractTx(ctx)
	if tx != nil {
		return tx.SelectContext(ctx, dest, query, args...)
	}
	return s.DB.SelectContext(ctx, dest, query, args...)
}

func (s *store) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	tx := extractTx(ctx)
	if tx != nil {
		return tx.GetContext(ctx, dest, query, args...)
	}
	return s.DB.GetContext(ctx, dest, query, args...)
}

func (s *store) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	tx := extractTx(ctx)
	if tx != nil {
		return tx.ExecContext(ctx, query, args...)
	}
	return s.DB.ExecContext(ctx, query, args...)
}

func (s *store) QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	tx := extractTx(ctx)
	if tx != nil {
		return tx.QueryxContext(ctx, query, args...)
	}
	return s.DB.QueryxContext(ctx, query, args...)
}

func (s *store) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	tx := extractTx(ctx)
	if tx != nil {
		return tx.QueryRowxContext(ctx, query, args...)
	}
	return s.DB.QueryRowxContext(ctx, query, args...)
}

type txKey struct{}

// injectTx injects transaction to context
func injectTx(ctx context.Context, tx *sqlx.Tx) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

// extractTx extracts transaction from context
func extractTx(ctx context.Context) *sqlx.Tx {
	if tx, ok := ctx.Value(txKey{}).(*sqlx.Tx); ok {
		return tx
	}
	return nil
}
