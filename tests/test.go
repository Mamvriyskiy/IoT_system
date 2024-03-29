package tests 

import (
    //"errors"
    "testing"
	mocks_service "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/service/mocks"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/service"
    "github.com/golang/mock/gomock"
)

func TestCreateUser(t *testing.T) {
    // Создаем новый экземпляр контроллера gomock
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    // Создаем заглушку IUserRepo
    mockRepo := mocks_service.NewMockIUserRepo(ctrl)

    // Устанавливаем ожидания для вызова CreateUser метода в IUserRepo
    user := pkg.User{ /* Заполните данными пользователя */}
    mockRepo.EXPECT().CreateUser(user).Return(123, nil)

    // Создаем экземпляр UserService с использованием заглушки IUserRepo
    userService := service.NewUserService(mockRepo)

    // Вызываем метод CreateUser на userService
    userID, err := userService.CreateUser(user)

    // Проверяем, что результаты соответствуют ожиданиям
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    if userID != 123 {
        t.Errorf("Expected userID 123, got %d", userID)
    }
}

// func TestCreateUser(t *testing.T) {
// 	type mockBehavior func(s *mock_service.MockCreateUser, user pkg.User)

// 	testTable := []struct {
// 		inputUser pkg.User
// 		mockBehavior mockBehavior
// 		expectedId int
// 	} {
// 		{
// 			inputUser : pkg.User{
// 				username: : "username",
// 				email: "qwerty@mail.ru",
// 				password: "qwerty",
// 			},
// 			mockBehavior: func(s *mock_service.MockCreateUser, user pkg.User) {
// 				s.EXPECT().CreateUser(user).Retrun(1, nil)
// 			}
// 			expectedId: 1
// 		}
// 	}

// 	for _, testCase := range testTable {
// 		t.Run(testCase.inputUser, func(t *testing T) {
// 			c := gomock.NewController(t)
// 			defer c.Finish()


// 		})
// 	}
// }

 