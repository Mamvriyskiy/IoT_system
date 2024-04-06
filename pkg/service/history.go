package service

import (
	pkg "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/repository"
)

type HistoryDeviceService struct {
	repo repository.IHistoryDeviceRepo
}

func NewHistoryDeviceService(repo repository.IHistoryDeviceRepo) *HistoryDeviceService {
	return &HistoryDeviceService{repo: repo}
}

func (s *HistoryDeviceService) CreateDeviceHistory(deviceID int, history pkg.DevicesHistory) (int, error) {
	return s.repo.CreateDeviceHistory(deviceID, history)
}

func (s *HistoryDeviceService) UpdateDeviceHistory(idDevice int, history pkg.DevicesHistory) error {
	return s.repo.UpdateDeviceHistory(idDevice, history)
}

func (s *HistoryDeviceService) GetDeviceHistory(idDevice int) ([]pkg.DevicesHistory, error) {
	return s.repo.GetDeviceHistory(idDevice)
}
