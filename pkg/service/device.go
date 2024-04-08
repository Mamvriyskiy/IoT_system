package service

import (
	pkg "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/repository"
)

type DeviceService struct {
	repo repository.IDeviceRepo
}

func NewDeviceService(repo repository.IDeviceRepo) *DeviceService {
	return &DeviceService{repo: repo}
}

func (s *DeviceService) CreateDevice(homeID int, device pkg.Devices) (int, error) {
	return s.repo.CreateDevice(homeID, device)
}

func (s *DeviceService) DeleteDevice(idDevice int) error {
	return s.repo.DeleteDevice(idDevice)
}

func (s *DeviceService) GetDeviceByID(deviceID int) (pkg.Devices, error) {
	return s.repo.GetDeviceByID(deviceID)
}
