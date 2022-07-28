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

func (u *userPostgres) Create(ctx context.Context, phone string) (int, error) {
	var id int
	createUserQuery := fmt.Sprintf(createUserSql, userTable)
	err := u.db.GetContext(ctx, &id, createUserQuery, phone)

	return id, err
}

func (u *userPostgres) Update(ctx context.Context, input UpdateUserInput) error {
	return nil
}
func (u *userPostgres) Get(ctx context.Context, userId int) (domain.User, error) {
	return domain.User{}, nil
}
func (u *userPostgres) Delete(ctx context.Context, userId int) error {
	return nil
}
