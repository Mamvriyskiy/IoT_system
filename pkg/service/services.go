package service

import (
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/repository"
)

type IUser interface {
	CreateUser(user pkg.User) (int, error)
	CheckUser(user pkg.User) (int, error)
	GenerateToken(login, password string) (string, error)
}

type IHome interface {
	CreateHome(idUser int, home pkg.Home) (int, error)
	DeleteHome(homeID int) error
	UpdateHome(home pkg.Home) error
	GetHomeByID(homeID int) (pkg.Home, error)
}

type IAccessHome interface {
	AddUser(homeID, userID int, access pkg.AccessHome) (int, error)
	DeleteUser(idUser int) error
	UpdateLevel(idUser int, access pkg.AccessHome) error
	UpdateStatus(idUser int, access pkg.AccessHome) error
	GetListUserHome(homeID int) ([]pkg.ClientHome, error)
}

type IDevice interface {
	CreateDevice(homeID int, device pkg.Devices) (int, error)
	DeleteDevice(idDevice int) error
	GetDeviceByID(deviceID int) (pkg.Devices, error)
}

type IHistoryDevice interface {
	CreateDeviceHistory(deviceID int, history pkg.DevicesHistory) (int, error)
	GetDeviceHistory(idDevice int) ([]pkg.DevicesHistory, error)
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
		IUser:          NewUserService(repo.IUserRepo),
		IHome:          NewHomeService(repo.IHomeRepo),
		IAccessHome:    NewAccessHomeService(repo.IAccessHomeRepo),
		IDevice:        NewDeviceService(repo.IDeviceRepo),
		IHistoryDevice: NewHistoryDeviceService(repo.IHistoryDeviceRepo),
	}
}
