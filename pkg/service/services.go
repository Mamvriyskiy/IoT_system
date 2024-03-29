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

type IAccessHome interface {
	AddUser(access pkg.AccessHome) (int, error)
	DeleteUser(idUser int, access pkg.AccessHome) error
	UpdateLevel(idUser int, access pkg.AccessHome) error
	UpdateStatus(idUser int, access pkg.AccessHome) error
	GetListUserHome(homeId int, access pkg.AccessHome) ([]pkg.User, error)
}

type IDevice interface {
	CreateDevice(device pkg.Devices) (int, error)
	DeleteDevice(idDevice int, device pkg.Devices) error
	UpdateDevice(idDevice int, device pkg.Devices) error
	AddHomeDevice(idHome int, idDevice int, input pkg.Devices) error
	DeleteHomeDevice(idHome int, idDevice int, input pkg.Devices) error
}

type IHistoryDevice interface {

}

type Services struct {
	IUser
	IHome
	IAccessHome
	IDevice
	IHistoryDevice
}

func NewServices(repo *repository.Repository) *Services {
	return &Services{
		IUser: NewUserService(repo.IUserRepo),
		IHome: NewHomeService(repo.IHomeRepo),
		IAccessHome: NewAccessHomeService(repo.IAccessHomeRepo),
		IHistoryDevice: NewHistoryDeviceService(repot.IHistoryDevice)
	}
}
