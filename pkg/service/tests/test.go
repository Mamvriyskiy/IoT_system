package tests 

import (
    //"errors"
    "testing"

    "github.com/golang/mock/gomock"
)

func TestCreateUser(t *testing.T) {
    // Создаем новый контроллер моков
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    // Создаем мок-объект для IUser
    mockRepo := NewMockIUser(ctrl)

    // Создаем экземпляр UserService с использованием мока
    userService := &UserService{repo: mockRepo}

    // Подготавливаем данные пользователя для теста
    testUser := pkg.User{
        Username: "testuser",
        Password: "password123",
    }

    // Ожидаем, что метод CreateUser в моке будет вызван с тем же пользователем
    mockRepo.EXPECT().CreateUser(testUser).Return(1, nil)

    // Вызываем метод CreateUser из UserService
    userID, err := userService.CreateUser(testUser)

    // Проверяем ожидаемый результат
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
    if userID != 1 {
        t.Errorf("Expected user ID 1, got %d", userID)
    }
}
