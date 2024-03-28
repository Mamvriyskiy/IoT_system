package repository

import (
	pkg "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"github.com/jmoiron/sqlx"
)

type IUserRepo interface {
	CreateUser(user pkg.User) (int, error)
	GetUserByEmail(email string) (int, error)
	GetPasswordById(id int) (string, error)
}

type IHomeRepo interface {
	CreateHome(idUser int, home pkg.Home) (int, error)
	DeleteHome(idUser int, home pkg.Home) error
	UpdateHome(idUser int, home pkg.Home) error
}

type Repository struct {
	IUserRepo
	IHomeRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		IUserRepo: NewUserPostgres(db),
		IHomeRepo: NewHomePostgres(db),
	}
}
