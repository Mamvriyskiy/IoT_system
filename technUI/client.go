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
	//user UserResponse
	TOKEN string
)

func menu() {
	fmt.Println(`
		1)Создать дом
		2)Удалить дом
		3)Обновить имя дома
		4)Добавить устройство
		5)Создать устройство
		6)Добавить участника
		7)Удалить участника
		8)Поменять уровень доступа участника
		9)Просмотреть статистика
		10)Запустить устройство
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

	_, err = requestServer("POST", "/auth/sign-up", data)
	if err != nil {
		return err
	}

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

	TOKEN = result

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
	}

	_, err = requestServer("POST", "/api/home", data)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func deleteHome() error {
	data := map[string]string{
		// "homeId": string(user.ID),
	}

	_, err := requestServer("DELETE", "/api/home", data)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func updateHome() error {
	var newHomeName string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите имя дома: ")
	fmt.Fscan(reader, &newHomeName)

	data := map[string]string{
		"name": newHomeName,
	}
	_, err := requestServer("PUT", "/api/home", data)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func requestServer(typeReq, path string, data map[string]string) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	
	req, err := http.NewRequest(typeReq, "http://localhost:8000" + path, bytes.NewBuffer(jsonData))
	req.Header.Set("Authorization", "Bearer "+ TOKEN)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	var responseBody []byte
	if resp.StatusCode == http.StatusOK {
		responseBody, err = io.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		fmt.Println("Тело ответа:", string(responseBody))
	} else {
		fmt.Println("Сервер вернул статус:", resp.Status)
	}

	var userResponse UserResponse
	if len(responseBody) > 0 {
		if err := json.Unmarshal(responseBody, &userResponse); err != nil {
			return "", err
		}
	}

	return userResponse.Token, nil
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
			err := updateHome()
			if err != nil {
				fmt.Println("Error:", err)
			}
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
