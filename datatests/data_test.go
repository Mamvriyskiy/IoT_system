package tests_test

import (
	"os"
	"reflect"
	"testing"

	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/repository"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func createDB() (*sqlx.DB, error) {
	err := initConfig()
	if err != nil {
		return nil, err
	}

	err = godotenv.Load()
	if err != nil {
		return nil, err
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	return db, err
}

func TestHomeData(t *testing.T) {
	db, err := createDB()
	assert.NoError(t, err)

	homePostgres := repository.NewHomePostgres(db)
	assert.NoError(t, err)

	testCases := []struct {
		OwnerID int
		Home    pkg.Home
	}{
		{OwnerID: 1, Home: pkg.Home{Name: "home1"}},
		{OwnerID: 2, Home: pkg.Home{Name: "home2"}},
		{OwnerID: 3, Home: pkg.Home{Name: "home3"}},
	}

	for _, tc := range testCases {
		homeID, err := homePostgres.CreateHome(tc.OwnerID, tc.Home)
		assert.NoError(t, err)

		res, err := homePostgres.GetHomeByID(homeID)
		assert.NoError(t, err)
		assert.Equal(t, tc.Home.Name, res.Name)
		assert.Equal(t, tc.OwnerID, res.OwnerID)
	}
}

func TestDeviceData(t *testing.T) {
	db, err := createDB()
	assert.NoError(t, err)

	devicePostgres := repository.NewDevicePostgres(db)

	testCases := []struct {
		HomeID  int
		Devices pkg.Devices
	}{
		{
			HomeID: 10,
			Devices: pkg.Devices{
				Name: "name1", TypeDevice: "type1", Status: "active", Brand: "apple",
				PowerConsumption: 10, MinParameter: 1, MaxParameter: 10,
			},
		},
		{
			HomeID: 11,
			Devices: pkg.Devices{
				Name: "name2", TypeDevice: "type2", Status: "active", Brand: "apple",
				PowerConsumption: 20, MinParameter: 20, MaxParameter: 30,
			},
		},
		{
			HomeID: 12,
			Devices: pkg.Devices{
				Name: "name3", TypeDevice: "type3", Status: "no active", Brand: "samsung",
				PowerConsumption: 30, MinParameter: 11, MaxParameter: 12,
			},
		},
	}

	for _, tc := range testCases {
		deviceID, err := devicePostgres.CreateDevice(tc.HomeID, tc.Devices)
		assert.NoError(t, err)

		res, err := devicePostgres.GetDeviceByID(deviceID)
		assert.NoError(t, err)
		assert.Equal(t, tc.Devices.Name, res.Name)
		assert.Equal(t, tc.Devices.TypeDevice, res.TypeDevice)
		assert.Equal(t, tc.Devices.Status, res.Status)
		assert.Equal(t, tc.Devices.Brand, res.Brand)
		assert.Equal(t, tc.Devices.PowerConsumption, res.PowerConsumption)
		assert.Equal(t, tc.Devices.MinParameter, res.MinParameter)
		assert.Equal(t, tc.Devices.MaxParameter, res.MaxParameter)
	}
}

func TestAccessData(t *testing.T) {
	db, err := createDB()
	assert.NoError(t, err)

	homePostgres := repository.NewHomePostgres(db)
	accessPostgres := repository.NewAccessHomePostgres(db)
	userPostgres := repository.NewUserPostgres(db)

	testCases := []struct {
		OwnerID  int
		Home     pkg.Home
		Users    []pkg.User
		Access   []pkg.AccessHome
		Expected []pkg.ClientHome
	}{
		{
			OwnerID: 1,
			Home:    pkg.Home{Name: "Misfio32"},
			Users: []pkg.User{
				{Password: "user1pass", Username: "user1", Email: "user1@example.com"},
				{Password: "user2pass", Username: "user2", Email: "user2@example.com"},
			},
			Access: []pkg.AccessHome{
				{AccessStatus: "granted", AccessLevel: 1},
				{AccessStatus: "granted", AccessLevel: 2},
			},
			Expected: []pkg.ClientHome{
				{Username: "user1", AccessLevel: 1, AccessStatus: "granted"},
				{Username: "user2", AccessLevel: 2, AccessStatus: "granted"},
			},
		},
	}

	for _, tc := range testCases {
		homeID, err := homePostgres.CreateHome(tc.OwnerID, tc.Home)
		assert.NoError(t, err)

		for i := range tc.Users {
			tc.Users[i].ID, err = userPostgres.CreateUser(tc.Users[i])
			assert.NoError(t, err)
		}

		for j := range tc.Access {
			tc.Access[j].ID, err = accessPostgres.AddUser(homeID, tc.Users[j].ID, tc.Access[j])
			assert.NoError(t, err)
		}

		res, err := accessPostgres.GetListUserHome(homeID)
		if !reflect.DeepEqual(tc.Expected, res) {
			t.Errorf("Ожидаемый: %v, Фактический: %v", tc.Expected, res)
		}
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
