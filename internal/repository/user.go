package repository

import (
	"context"
	"fmt"
	"github.com/nekruzrabiev/simple-app/internal/domain"
)

type userPostgres struct {
	db *store
}

func newUserPostgres(db *store) *userPostgres {
	return &userPostgres{db: db}
}

const createUserSql = `
	INSERT INTO %s (full_name, phone)
	VALUES ('', $1)
	RETURNING ID;
`

func (r *userPostgres) Create(ctx context.Context, phone string) (int, error) {
	var id int
	createUserQuery := fmt.Sprintf(createUserSql, userTable)
	err := r.db.GetContext(ctx, &id, createUserQuery, phone)

	return id, err
}

func (r *userPostgres) Update(ctx context.Context, input UpdateUserInput) error {
	return nil
}
func (r *userPostgres) Get(ctx context.Context, userId int) (domain.User, error) {
	return domain.User{}, nil
}
func (r *userPostgres) Delete(ctx context.Context, userId int) error {
	return nil
}
