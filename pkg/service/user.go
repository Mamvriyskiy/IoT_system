package service

import (
	"crypto/sha1"
	"fmt"
	pkg "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/repository"
)

const salt = "hfdjmaxckdk20"

type UserService struct {
	repo repository.IUserRepo
}

func NewUserService(repo repository.IUserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user pkg.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *UserService) CheckUser(user pkg.User) (bool, int, error) {
	var err error
	if user.Email == "" {
		// *TODO: log
		return false, -1, err
	}

	id, err := s.repo.GetUserByEmail(user.Email)
	if err != nil {
		// *TODO: log
		return false, id, err
	}

	pswd, err := s.repo.GetPasswordById(id)
	if err != nil {
		// *TODO: log
		return false, id, err
	}

	return s.comparePassword(user.Password, pswd), id, err
}

func (s *UserService) comparePassword(pswd, hash string) bool {
	newHash := s.generatePasswordHash(pswd)
	return newHash == hash
}

func (s *UserService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
