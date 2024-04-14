package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type UserResponse struct {
	ID     int    `json: "id"`
	Token  string `json: "token"`
	HomeID int    `json: "homeId"`
}

var (
	user UserResponse
)

func menu() {
	fmt.Println(`
		1)Создать дом
		2)Удалить дом
		3)Добавить устройство
		4)Создать устройство
		5)Добавить участника
		6)Удалить участника
		7)Поменять уровень доступа участника
		8)Просмотреть статистика
		9)Запустить устройство
		0)Завершить работу
	`)
}

func registerUser() error {
	var password, username, email string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите имя пользователя: ")
	_, err := fmt.Fscanf(reader, "%s\n", &username)
	if err != nil {
		return err
	}

	fmt.Print("Введите почту: ")
	_, err = fmt.Fscanf(reader, "%s\n", &email)
	if err != nil {
		return err
	}

	fmt.Print("Введите пароль: ")
	_, err = fmt.Fscanf(reader, "%s\n", &password)
	if err != nil {
		return err
	}

	data := map[string]string{
		"password": password,
		"login":    username,
		"email":    email,
	}

	result, err := requestServer("POST", "/auth/sign-up", data)
	if err != nil {
		return err
	}

	user.ID = result.ID

	return nil
}

func auth() error {
	var password, username string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите имя пользователя: ")
	_, err := fmt.Fscanf(reader, "%s\n", &username)
	if err != nil {
		return err
	}

	fmt.Print("Введите пароль: ")
	_, err = fmt.Fscanf(reader, "%s\n", &password)
	if err != nil {
		return err
	}

	data := map[string]string{
		"password": password,
		"login":    username,
	}

	result, err := requestServer("POST", "/auth/sign-in", data)
	if err != nil {
		return err
	}

	user.Token = result.Token

	return nil
}

func createHome() error {
	var nameHome string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите имя дома: ")
	_, err := fmt.Fscanf(reader, "%s\n", &nameHome)
	if err != nil {
		return err
	}

	data := map[string]string{
		"name": nameHome,
		"ownerId": string(user.ID),
	}

	result, err := requestServer("POST", "/home/create", data)
	if err != nil {
		fmt.Println(err)
		return err
	}

	user.HomeID = result.HomeID
	return nil
}

func deleteHome() error {
	data := map[string]string{
		"homeId": string(user.ID),
	}

	result, err := requestServer("POST", "/home/create", data)
	if err != nil {
		fmt.Println(err)
		return err
	}

	user.HomeID = 0
	return nil
}

func requestServer(typeReq, path string, data map[string]string) (UserResponse, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return UserResponse{}, err
	}

	req, err := http.NewRequest(typeReq, "http://localhost:8000" + path, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err)
		return UserResponse{}, err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return UserResponse{}, err
	}
	defer resp.Body.Close()

	var responseBody []byte
	if resp.StatusCode == http.StatusOK {
		responseBody, err = io.ReadAll(resp.Body)
		if err != nil {
			return UserResponse{}, err
		}
		fmt.Println("Тело ответа:", string(responseBody))
	} else {
		fmt.Println("Сервер вернул статус:", resp.Status)
	}

	var userResponse UserResponse
	if len(responseBody) > 0 {
		if err := json.Unmarshal(responseBody, &userResponse); err != nil {
			return UserResponse{}, err
		}
	}

	return userResponse, nil
}

func clientGO() {
	for {
		var num int
		menu()
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Выберите номер: ")
		_, err := fmt.Fscanf(reader, "%d\n", &num)
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch num {
		case 1:
			err := createHome()
			if err != nil {
				fmt.Println("Error:", err)
			}
		case 2:
			err := deleteHome()
			if err != nil {
				fmt.Println("Error:", err)
			}
		case 3:
		case 4:
		case 5:
		case 6:
		case 7:
		case 8:
		case 9:
		default:
			break
		}
	}
}

func main() {
	var err error
	fmt.Println(`
		Меню:
		1)Регистрация
		2)Войти
	`)
	var num int
	fmt.Scan(&num)
	switch num {
	case 1:
		err = registerUser()
		if err != nil {
			fmt.Println("Error:", err)
		}
	case 2:
		err = auth()
		if err != nil {
			fmt.Println("Error:", err)
		}
	}

	if err == nil {
		clientGO()
	} else {
		fmt.Println("Ошибка!")
	}
}
