package repository

import (
	"context"
	"github.com/nekruzrabiev/simple-app/internal/domain"
	"github.com/nekruzrabiev/simple-app/internal/service"
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

//Refresh Session
type RefreshSession interface {
	Create(ctx context.Context, refreshSession domain.RefreshSession) error
	GetByToken(ctx context.Context, token string, expireAfter time.Time) (domain.RefreshSession, error)
}

type User interface {
	Create(ctx context.Context, user domain.User) (int, error)
	UpdateName(ctx context.Context, input service.UserUpdateInput) error
	Get(ctx context.Context, userId int) (*domain.User, error)
	Contains(ctx context.Context, email string) (bool, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
}
