package tests

import (
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	mocks_service "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/repository/mocks"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/service"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks_service.NewMockIUserRepo(ctrl)

	user := pkg.User{
		Username: "username",
		Email:    "qwerty@mail.ru",
		Password: "qwerty",
	}

	mockRepo.EXPECT().CreateUser(gomock.Any()).Return(123, nil)

	userService := service.NewUserService(mockRepo)

	userID, err := userService.CreateUser(user)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if userID != 123 {
		t.Errorf("Expected userID 123, got %d", userID)
	}
}

func TestCheckUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks_service.NewMockIUserRepo(ctrl)

	user := pkg.User{
		Username: "username",
		Email:    "qwerty@mail.ru",
		Password: "qwerty",
	}

	mockRepo.EXPECT().GetPasswordById(123).Return("6866646a6d6178636b646b3230b1b3773a05c0ed0176787a4f1574ff0075f7521e", nil)
	mockRepo.EXPECT().GetUserByEmail("qwerty@mail.ru").Return(123, nil)

	userService := service.NewUserService(mockRepo)

	cmp, userID, err := userService.CheckUser(user)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if userID != 123 {
		t.Errorf("Expected userID 123, got %d", userID)
	}
	if !cmp {
		t.Error("Expected search true, got", cmp)
	}
}
