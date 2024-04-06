package repository

import (
	pkg "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"github.com/jmoiron/sqlx"
)

//go:generate mockgen -source=repository.go -destination=mocks/mocks.go

type IUserRepo interface {
	CreateUser(user pkg.User) (int, error)
	GetUser(login, password string) (pkg.User, error)
}

type IHomeRepo interface {
	CreateHome(idUser int, home pkg.Home) (int, error)
	DeleteHome(idUser int, home pkg.Home) error
	UpdateHome(idUser int, home pkg.Home) error
}

type IAccessHomeRepo interface {
	AddUser(homeID, userID int, access pkg.AccessHome) (int, error)
	DeleteUser(idUser int, access pkg.AccessHome) error
	UpdateLevel(idUser int, access pkg.AccessHome) error
	UpdateStatus(idUser int, access pkg.AccessHome) error
	GetListUserHome(idHome int, access pkg.AccessHome) ([]pkg.User, error)
}

type IDeviceRepo interface {
	CreateDevice(homeID int, device pkg.Devices) (int, error)
	DeleteDevice(idDevice int, device pkg.Devices) error
	UpdateDevice(idDevice int, device pkg.Devices) error
	AddHomeDevice(idHome, idDevice int, input pkg.Devices) error
	DeleteHomeDevice(idHome, idDevice int, input pkg.Devices) error
}

type IHistoryDeviceRepo interface {
	CreateDeviceHistory(deviceID int, history pkg.DevicesHistory) (int, error)
	UpdateDeviceHistory(idDevice int, history pkg.DevicesHistory) error
	GetDeviceHistory(idDevice int) ([]pkg.DevicesHistory, error)
}

type Repository struct {
	IUserRepo
	IHomeRepo
	IAccessHomeRepo
	IDeviceRepo
	IHistoryDeviceRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		IUserRepo:          NewUserPostgres(db),
		IHomeRepo:          NewHomePostgres(db),
		IAccessHomeRepo:    NewAccessHomePostgres(db),
		IDeviceRepo:        NewDevicePostgres(db),
		IHistoryDeviceRepo: NewDeviceHistoryPostgres(db),
	}
}
