package repository

import (
	"context"
	"fmt"
	"github.com/nekruzrabiev/simple-app/internal/domain"
	"time"
)

type refreshSessionPostgres struct {
	db *store
}

func newRefreshSessionPostgres(db *store) *refreshSessionPostgres {
	return &refreshSessionPostgres{db: db}
}

const createSessionSql = `
	INSERT INTO %s (token, expires_at, user_id)
	VALUES ($1, $2, $3);
`

func (r *refreshSessionPostgres) Create(ctx context.Context, refreshSession domain.RefreshSession) error {
	createSessionQuery := fmt.Sprintf(createSessionSql, refreshSessionTable)

	_, err := r.db.ExecContext(ctx, createSessionQuery, refreshSession.Token, refreshSession.ExpiresAt, refreshSession.UserId)

	return err
}

const getSessionByTokenSql = `
	SELECT
		token,
		expires_at,
		user_id
	FROM %s
	WHERE
		token = $1
		AND expires_at > $2;
`

func (r *refreshSessionPostgres) GetByToken(ctx context.Context, token string, expireAfter time.Time) (domain.RefreshSession, error) {
	var session domain.RefreshSession
	getByTokenQuery := fmt.Sprintf(getSessionByTokenSql, refreshSessionTable)

	err := r.db.GetContext(ctx, &session, getByTokenQuery, token, expireAfter)

	return session, err
}
