package tests

import (
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	mocks_service "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/repository/mocks"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/service"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestCreateDevice(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks_service.NewMockIDeviceRepo(ctrl)

	device := pkg.Devices{
		Name:             "tea",
		TypeDevice:         "kettle",
		Status:           "free",
		Brand:            "Samsung",
		PowerСonsumption: 1500,
		MinParameter:     50, //temperature
		MaxParameter:     120,
	}
	
	mockRepo.EXPECT().CreateDevice(device).Return(10, nil)

	deviceService := service.NewDeviceService(mockRepo)

	deviceID, err := deviceService.CreateDevice(device)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if deviceID != 10 {
		t.Errorf("Expected userID 10, got %d", deviceID)
	}
}

func TestDeleteDevice(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks_service.NewMockIDeviceRepo(ctrl)

	device := pkg.Devices{
		Name:             "tea",
		TypeDevice:         "kettle",
		Status:           "free",
		Brand:            "Samsung",
		PowerСonsumption: 1500,
		MinParameter:     50, //temperature
		MaxParameter:     120,
	}

	mockRepo.EXPECT().DeleteDevice(10, device).Return(nil)

	deviceService := service.NewDeviceService(mockRepo)

	err := deviceService.DeleteDevice(10, device)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestUpdateDevice(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks_service.NewMockIDeviceRepo(ctrl)

	device := pkg.Devices{
		Name:             "tea",
		TypeDevice:         "kettle",
		Status:           "free",
		Brand:            "Samsung",
		PowerСonsumption: 1500,
		MinParameter:     50, //temperature
		MaxParameter:     120,
	}

	mockRepo.EXPECT().UpdateDevice(10, device).Return(nil)

	deviceService := service.NewDeviceService(mockRepo)

	err := deviceService.UpdateDevice(10, device)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestAddHomeDevice(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks_service.NewMockIDeviceRepo(ctrl)

	device := pkg.Devices{
		Name:             "tea",
		TypeDevice:         "kettle",
		Status:           "free",
		Brand:            "Samsung",
		PowerСonsumption: 1500,
		MinParameter:     50, //temperature
		MaxParameter:     120,
	}

	mockRepo.EXPECT().AddHomeDevice(10, 1, device).Return(nil)

	deviceService := service.NewDeviceService(mockRepo)

	err := deviceService.AddHomeDevice(10, 1, device)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestDeleteHomeDevice(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks_service.NewMockIDeviceRepo(ctrl)

	device := pkg.Devices{
		Name:             "tea",
		TypeDevice:         "kettle",
		Status:           "free",
		Brand:            "Samsung",
		PowerСonsumption: 1500,
		MinParameter:     50, //temperature
		MaxParameter:     120,
	}

	mockRepo.EXPECT().DeleteHomeDevice(10, 1, device).Return(nil)

	deviceService := service.NewDeviceService(mockRepo)

	err := deviceService.DeleteHomeDevice(10, 1, device)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
