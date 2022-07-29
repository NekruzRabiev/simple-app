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
	INSERT INTO %s (full_name, email, password)
	VALUES ($1, $2, $2)
	RETURNING id;
`

func (r *userPostgres) Create(ctx context.Context, user domain.User) (int, error) {
	var id int
	createUserQuery := fmt.Sprintf(createUserSql, userTable)
	err := r.db.GetContext(ctx, &id, createUserQuery, user.FullName, user.Email, user.Password)

	return id, err
}

const updateUserFullNameSql = `
	UPDATE %s
	SET
		full_name = $1
	WHERE
		id = $2;
`

func (r *userPostgres) UpdateName(ctx context.Context, id int, name string) error {
	query := fmt.Sprintf(updateUserFullNameSql, userTable)
	_, err := r.db.ExecContext(ctx, query, name, id)
	return err
}

const getUserByIdSql = `
	SELECT
		u.id,
		u.full_name,
		u.email
	FROM
		%s AS u
	WHERE
		u.id = $1
	LIMIT
		1;
`

func (r *userPostgres) Get(ctx context.Context, userId int) (*domain.User, error) {
	query := fmt.Sprintf(getUserByEmailSql, userTable)

	user := new(domain.User)

	err := r.db.GetContext(ctx, user, query, userId)
	return user, err
}

const containsUserSql = `
	SELECT
		EXISTS (
			SELECT
				u.id
			FROM
				%s AS u
			WHERE
				u.email = $1
			LIMIT
				1
		);
`

func (r *userPostgres) Contains(ctx context.Context, email string) (bool, error) {
	query := fmt.Sprintf(containsUserSql, userTable)

	var contains bool

	err := r.db.GetContext(ctx, &contains, query, email)
	return contains, err
}

const getUserByEmailSql = `
	SELECT
		u.id,
		u.password
	FROM
		%s AS u
	WHERE
		u.email = $1
	LIMIT
		1;
`

func (r *userPostgres) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := fmt.Sprintf(getUserByEmailSql, userTable)

	user := new(domain.User)
	err := r.db.GetContext(ctx, user, query, email)
	return user, err
}
