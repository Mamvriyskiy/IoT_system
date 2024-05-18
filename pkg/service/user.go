package service

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/logger"
	pkg "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/repository"
	jwt "github.com/dgrijalva/jwt-go"
)

const (
	salt       = "hfdjmaxckdk20"
	signingKey = "jaskljfkdfndnznmckmdkaf3124kfdlsf"
)

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

func (s *UserService) CheckUser(user pkg.User) (id int, err error) {
	if user.Email == "" {
		logger.Log("Error", "CheckUser", "Empty email:", nil, "")
		return -1, err
	}

	user, err = s.repo.GetUser(user.Username, user.Password)
	if err != nil {
		logger.Log("Error", "GetUser", "Error get user:", err, user.Username, user.Password)
		return id, err
	}

	return user.ID, err
}

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"userid"`
}

func (s *UserService) GenerateToken(login, password string) (string, error) {
	user, err := s.repo.GetUser(login, s.generatePasswordHash(password))
	if err != nil {
		logger.Log("Error", "GetUser", "Error get user:", err, login, "s.generatePasswordHash(password)")
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *UserService) generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return hex.EncodeToString(hash.Sum([]byte(salt)))
}

func (s *UserService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				logger.Log("Error", "token.Method.(*jwt.SigningMethodHMAC)", "Error jwt token:", nil, "")
				return 0, nil
			}

			return []byte(signingKey), nil
		})
	if err != nil {
		logger.Log("Error", "jwt.ParseWithClaims", "Error parse token:", err, accessToken)
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		logger.Log("Error", "token.Claims.(*tokenClaims)", "Error token:", nil)
		return 0, nil
	}

	return claims.UserID, nil
}
