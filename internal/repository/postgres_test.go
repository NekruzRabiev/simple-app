package repository

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPostgresDBSuccess(t *testing.T) {
	// Init
	config := ConfigPostgres{
		Host:     "127.0.0.1",
		Port:     "5433",
		DBName:   "simple",
		Username: "postgres",
		Password: "qwerty",
		SSLMode:  "disable",
	}

	// Execution
	db, err := NewPostgresDB(&config)

	// Validation
	assert.Nil(t, db.Ping())
	assert.Nil(t, err)
	assert.NotNil(t, db)
}

func TestNewPostgresDBError(t *testing.T) {
	// Init
	configError := ConfigPostgres{
		Host:     "db",
		Port:     "5432",
		DBName:   "simp",
		Username: "postgres",
		Password: "qwerty",
		SSLMode:  "disable",
	}
	// Execution
	db, err := NewPostgresDB(&configError)

	//validation
	assert.NotNil(t, err, "cannot connect")
	assert.Nil(t, db)

	assert.Panics(t, func() {
		db.Ping()
	})
	assert.Equal(t, (*store)(nil), db)
}
