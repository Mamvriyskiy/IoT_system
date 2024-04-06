package service

import (
	pkg "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/repository"
	"fmt"
)

type DeviceService struct {
	repo repository.IDeviceRepo
}

func NewDeviceService(repo repository.IDeviceRepo) *DeviceService {
	return &DeviceService{repo: repo}
}

func (s *DeviceService) CreateDevice(homeID int, device pkg.Devices) (int, error) {
	fmt.Println("+")
	return s.repo.CreateDevice(homeID, device)
}

func (s *DeviceService) DeleteDevice(idDevice int, device pkg.Devices) error {
	return s.repo.DeleteDevice(idDevice, device)
}
