package service

import (
	pkg "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/repository"
)

type IUser interface {
	CreateUser(user pkg.User) (int, error)
	CheckUser(user pkg.User) (bool, int, error)
}

type IHome interface {
	CreateHome(idUser int, home pkg.Home) (int, error)
	DeleteHome(idUser int, home pkg.Home) error
	UpdateHome(idUser int, home pkg.Home) error
}

type Services struct {
	IUser
	IHome
}

func NewServices(repo *repository.Repository) *Services {
	return &Services{
		IUser: NewUserService(repo.IUserRepo),
		IHome: NewHomeService(repo.IHomeRepo),
	}
}
