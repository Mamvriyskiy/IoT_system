package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"
	//"fmt"
	"github.com/stretchr/testify/assert"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/handler"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/service"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/repository"
	"github.com/spf13/viper"
	"github.com/joho/godotenv"
	"os"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func TestPingRoute(t *testing.T) {
	if err := initConfig(); err != nil {
		t.Errorf("No initConfig")
		return
	}

	if err := godotenv.Load(); err != nil {
		// *TODO: log
		return
	}

	db, err := repository.NewPostgresDB(&repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	
	if err != nil {
		t.Errorf("No connection DB %s", err)
		return 
	}

	// repos := repository.NewRepository(db)
	// services := service.NewServices(repos)
	// handlers := handler.NewHandler(services)

	router := handlers.InitRouters()

	w := httptest.NewRecorder()
	jsonStr := []byte(`{"password":"qwerty","login":"misfio32","email":"m_ivan_s@mail.ru"}`)

	// Создание HTTP-запроса с помощью строкового ридера
	req, err := http.NewRequest("POST", "/auth/sign-in", bytes.NewBuffer(jsonStr))
	if err != nil {
		//t.Errorf("Error: %s", err)
	}
	router.ServeHTTP(w, req)

	//assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
