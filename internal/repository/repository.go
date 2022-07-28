package repository

import (
	"context"
	"github.com/nekruzrabiev/simple-app/internal/domain"
	"time"
)

type Repositories struct {
	Transactor     Transactor
	User           User
	RefreshSession RefreshSession
}

func NewRepositories(db *store) *Repositories {
	return &Repositories{
		Transactor:     newTransactorPostgres(db),
		User:           newUserPostgres(db),
		RefreshSession: newRefreshSessionPostgres(db),
	}
}

// Transactor runs logic inside a single database transaction
type Transactor interface {
	// WithinTransaction runs a function within a database transaction.
	//
	// Transaction is propagated in the context,
	// so it is important to propagate it to underlying repositories.
	// Function commits if error is nil, and rollbacks if not.
	// It returns the same error.
	WithinTransaction(context.Context, func(ctx context.Context) error) error
}

//User
type UpdateUserInput struct {
	Id   int
	Name string
}

//Refresh Session
type RefreshSession interface {
	Create(ctx context.Context, refreshSession domain.RefreshSession) error
	GetByToken(ctx context.Context, token string, expireAfter time.Time) (domain.RefreshSession, error)
}

type User interface {
	Create(ctx context.Context, phone string) (int, error)
	Update(ctx context.Context, input UpdateUserInput) error
	Get(ctx context.Context, userId int) (domain.User, error)
	Delete(ctx context.Context, userId int) error
}
