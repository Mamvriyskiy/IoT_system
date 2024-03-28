package service

import (
	pkg "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/repository"
)

type Authorization interface {
	CreateUser(user pkg.User) (int, error)
}

type Services struct {
	Authorization
}

func NewServices(repo *repository.Repository) *Services {
	return &Services{
		Authorization : NewAuthService(repo.Authorization),
	}
}
