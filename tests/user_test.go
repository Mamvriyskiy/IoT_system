package tests_test

import (
	"testing"

	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	mocks_service "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/repository/mocks"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/service"
	"github.com/golang/mock/gomock"
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

	mockRepo.EXPECT().
		GetPasswordByID(123).
		Return(
			"6866646a6d6178636b646b323065e84be33532fb784c48129675f9eff3a682b27168c0ea744b2cf58ee02337c5",
			nil)

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
