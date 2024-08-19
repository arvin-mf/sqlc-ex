package repository

import (
	"context"
	"database/sql"
	"sj/internal/db/sqlc"
)

type UserRepository interface {
	AddUser(arg sqlc.AddUserParams) (sql.Result, error)
	EmailExists(email string) (int64, error)
	FindByEmail(email string) (sqlc.User, error)
}

type userRepository struct {
	db *sqlc.Queries
}

func NewUserRepository(db *sqlc.Queries) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) AddUser(arg sqlc.AddUserParams) (sql.Result, error) {
	return r.db.AddUser(context.Background(), arg)
}

func (r *userRepository) EmailExists(email string) (int64, error) {
	return r.db.EmailExists(context.Background(), email)
}

func (r *userRepository) FindByEmail(email string) (sqlc.User, error) {
	return r.db.FindByEmail(context.Background(), email)
}
