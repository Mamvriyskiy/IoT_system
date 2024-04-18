package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	pkg "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
)

type UserResponse struct {
	ID       int    `json: "id"`
	Token    string `json: "token"`
	HomeID   int    `json: "homeId"`
	HomeName string `json: "name"`
}

var (
	TOKEN string
)

func menu() {
	fmt.Println(`
		1)Создать дом
		2)Удалить дом
		3)Обновить имя дома
		4)Добавить устройство
		5)Удалить устройство
		6)Добавить участника
		7)Удалить участника
		8)Просмотреть статистику
		9)Запустить устройство
		0)Завершить работу
	`)
}

func checkStat() error {
	var device string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите имя устройства: ")
	_, err := fmt.Fscanf(reader, "%s\n", &device)
	if err != nil {
		return err
	}

	data := pkg.AddHistory{
		Name:       device,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = requestServer("GET", "/api/history", jsonData)

	if err != nil {
		return err
	}

	return nil
}

func runDevice() error {
	var device string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите имя устройства: ")
	_, err := fmt.Fscanf(reader, "%s\n", &device)
	if err != nil {
		return err
	}

	data := pkg.AddHistory{
		Name:       device,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = requestServer("POST", "/api/history", jsonData)

	if err != nil {
		return err
	}

	return nil
}

func deleteUser() error {
	var email string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите почту пользователя: ")
	_, err := fmt.Fscanf(reader, "%s\n", &email)
	if err != nil {
		return err
	}

	data := pkg.AddUserHome{
		Email:       email,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = requestServer("POST", "/api/access", jsonData)

	if err != nil {
		return err
	}

	return nil
}

func addUser() error {
	var email string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите почту пользователя: ")
	_, err := fmt.Fscanf(reader, "%s\n", &email)
	if err != nil {
		return err
	}

	var level int
	fmt.Print("Введите уровень доступа: ")
	_, err = fmt.Fscanf(reader, "%d\n", &level)
	if err != nil {
		return err
	}

	data := pkg.AddUserHome{
		Email:       email,
		AccessLevel: level,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = requestServer("POST", "/api/access", jsonData)

	if err != nil {
		return err
	}

	return nil
}

func deleteDevice() error {
	var deviceName string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите имя устройства: ")
	_, err := fmt.Fscanf(reader, "%s\n", &deviceName)
	if err != nil {
		return err
	}

	data := pkg.Devices{
		Name: deviceName,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = requestServer("DELETE", "/api/device", jsonData)
	if err != nil {
		return err
	}

	return nil
}

func addDevice() error {
	var deviceName string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите имя устройства: ")
	_, err := fmt.Fscanf(reader, "%s\n", &deviceName)
	if err != nil {
		return err
	}

	data := pkg.Devices{
		Name:             deviceName,
		TypeDevice:       "vacuum cleaner",
		Brand:            "apple",
		Status:           "wait",
		PowerConsumption: 100,
		MinParameter:     10,
		MaxParameter:     30,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = requestServer("POST", "/api/device", jsonData)
	if err != nil {
		return err
	}

	return nil
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

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = requestServer("POST", "/auth/sign-up", jsonData)
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

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	result, err := requestServer("POST", "/auth/sign-in", jsonData)
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

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = requestServer("POST", "/api/home", jsonData)
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

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = requestServer("DELETE", "/api/home", jsonData)
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

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = requestServer("PUT", "/api/home", jsonData)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func getListHomeByUser() error {
	data := map[string]string{}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = requestServer("GET", "/api/home", jsonData)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func requestServer(typeReq, path string, jsonData []byte) (string, error) {
	req, err := http.NewRequest(typeReq, "http://localhost:8000"+path, bytes.NewBuffer(jsonData))
	req.Header.Set("Authorization", "Bearer "+TOKEN)
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
			err := addDevice()
			if err != nil {
				fmt.Println("Error:", err)
			}
		case 5:
			err := deleteDevice()
			if err != nil {
				fmt.Println("Error:", err)
			}
		case 6:
			err := addUser()
			if err != nil {
				fmt.Println("Error:", err)
			}
		case 7:
			err := deleteUser()
			if err != nil {
				fmt.Println("Error:", err)
			}
		case 8:
			err := checkStat()
			if err != nil {
				fmt.Println("Error:", err)
			}
		case 9:
			err := runDevice()
			if err != nil {
				fmt.Println("Error:", err)
			}
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
