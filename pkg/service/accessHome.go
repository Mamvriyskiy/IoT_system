package service

import (
	pkg "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/repository"
)

type AccessHomeService struct {
	repo repository.IAccessHomeRepo
}

func NewAccessHomeService(repo repository.IAccessHomeRepo) *AccessHomeService {
	return &AccessHomeService{repo: repo}
}

func (s *AccessHomeService) AddUser(homeID, userID int, access pkg.AccessHome) (int, error) {
	return s.repo.AddUser(homeID, userID, access)
}

func (s *AccessHomeService) DeleteUser(idUser int) error {
	return s.repo.DeleteUser(idUser)
}

func (s *AccessHomeService) UpdateLevel(idUser int, access pkg.AccessHome) error {
	return s.repo.UpdateLevel(idUser, access)
}

func (s *AccessHomeService) UpdateStatus(idUser int, access pkg.AccessHome) error {
	return s.repo.UpdateStatus(idUser, access)
}

func (s *AccessHomeService) GetListUserHome(idHome int) ([]pkg.ClientHome, error) {
	return s.repo.GetListUserHome(idHome)
}
