package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/fercarcedo/quadrant-server/internal/controllers"
	"github.com/fercarcedo/quadrant-server/internal/models"
)

type mockMachineService struct {
	mock.Mock
}

func (m *mockMachineService) GetMachinesByInspectorIdAndRevision(inspectorId int) ([]models.Machine, error) {
	args := m.Called(inspectorId)
	return args.Get(0).([]models.Machine), args.Error(1)
}

func (m *mockMachineService) CreateMachine(machine *models.Machine) error {
	args := m.Called(machine)
	return args.Error(0)
}

func (m *mockMachineService) UpdateMachine(machine *models.Machine) error {
	args := m.Called(machine)
	return args.Error(0)
}

func (m *mockMachineService) DeleteMachine(machineId int) error {
	args := m.Called(machineId)
	return args.Error(0)
}

func TestGetMachinesByInspectorIdAndRevision(t *testing.T) {
	machineService := new(mockMachineService)
	machines := []models.Machine { 
		models.Machine { Name: "machine1", Period: 7, InspectorId: 3 }, 
		models.Machine { Name: "machine2", Period: 6, InspectorId: 3 }, 
	}
	machineService.On("GetMachinesByInspectorIdAndRevision", 3).Return(machines, nil)
	machineController := controllers.NewMachineController(machineService)
	router := controllers.NewRouter(nil, machineController, nil, nil).SetUpRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/users/3/machines", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	jsonBytes, _ := json.Marshal(&machines)
	assert.Equal(t, string(jsonBytes)+"\n", w.Body.String())
}

func TestGetMachinesByInspectorIdAndRevisionBadRequest(t *testing.T) {
	machineService := new(mockMachineService)
	machines := []models.Machine { 
		models.Machine { Name: "machine1", Period: 7, InspectorId: 3 }, 
		models.Machine { Name: "machine2", Period: 6, InspectorId: 3 }, 
	}
	machineService.On("GetMachinesByInspectorIdAndRevision", 3).Return(machines, nil)
	machineController := controllers.NewMachineController(machineService)
	router := controllers.NewRouter(nil, machineController, nil, nil).SetUpRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/users/error/machines", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{}\n", w.Body.String())
}

func TestCreateMachine(t *testing.T) {
	machineService := new(mockMachineService)
	machineService.On("CreateMachine", mock.MatchedBy(func(machine *models.Machine) bool {
		return machine.Name == "testmachine" &&
			machine.Period == 7 &&
			machine.InspectorId == 3
	})).Return(nil)
	machineController := controllers.NewMachineController(machineService)
	router := controllers.NewRouter(nil, machineController, nil, nil).SetUpRouter()
	w := httptest.NewRecorder()
	jsonStr, _ := json.Marshal(&models.Machine { Name: "testmachine", Period: 7, InspectorId: 3 })
	req, _ := http.NewRequest("POST", "/api/v1/machines", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
	assert.Equal(t, "{}\n", w.Body.String())
}

func TestCreateMachineBadRequest(t *testing.T) {
	machineService := new(mockMachineService)
	machineService.On("CreateMachine", mock.MatchedBy(func(machine *models.Machine) bool {
		return machine.Name == "testmachine" &&
			machine.Period == 7 &&
			machine.InspectorId == 3
	})).Return(errors.New("Error"))
	machineController := controllers.NewMachineController(machineService)
	router := controllers.NewRouter(nil, machineController, nil, nil).SetUpRouter()
	w := httptest.NewRecorder()
	jsonStr, _ := json.Marshal(&models.Machine { Name: "testmachine", Period: 7, InspectorId: 3 })
	req, _ := http.NewRequest("POST", "/api/v1/machines", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{}\n", w.Body.String())
}

func TestUpdateMachine(t *testing.T) {
	machineService := new(mockMachineService)
	machineService.On("UpdateMachine", mock.MatchedBy(func(machine *models.Machine) bool {
		return machine.Id == 3 && 
			machine.Name == "testmachine" &&
			machine.Period == 7 &&
			machine.InspectorId == 3
	})).Return(nil)
	machineController := controllers.NewMachineController(machineService)
	router := controllers.NewRouter(nil, machineController, nil, nil).SetUpRouter()
	w := httptest.NewRecorder()
	jsonStr, _ := json.Marshal(&models.Machine { Name: "testmachine", Period: 7, InspectorId: 3 })
	req, _ := http.NewRequest("PUT", "/api/v1/machines/3", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)
	assert.Equal(t, 204, w.Code)
	assert.Equal(t, "", w.Body.String())
}

func TestUpdateMachineBadRequest(t *testing.T) {
	machineService := new(mockMachineService)
	machineService.On("UpdateMachine", mock.MatchedBy(func(machine *models.Machine) bool {
		return machine.Id == 3 && 
			machine.Name == "testmachine" &&
			machine.Period == 7 &&
			machine.InspectorId == 3
	})).Return(errors.New("Error"))
	machineController := controllers.NewMachineController(machineService)
	router := controllers.NewRouter(nil, machineController, nil, nil).SetUpRouter()
	w := httptest.NewRecorder()
	jsonStr, _ := json.Marshal(&models.Machine { Name: "testmachine", Period: 7, InspectorId: 3 })
	req, _ := http.NewRequest("PUT", "/api/v1/machines/3", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{}\n", w.Body.String())
}

func TestDeleteMachine(t *testing.T) {
	machineService := new(mockMachineService)
	machineService.On("DeleteMachine", 3).Return(nil)
	machineController := controllers.NewMachineController(machineService)
	router := controllers.NewRouter(nil, machineController, nil, nil).SetUpRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/v1/machines/3", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 204, w.Code)
	assert.Equal(t, "", w.Body.String())
}

func TestDeleteMachineBadRequest(t *testing.T) {
	machineService := new(mockMachineService)
	machineService.On("DeleteMachine", 3).Return(errors.New("error"))
	machineController := controllers.NewMachineController(machineService)
	router := controllers.NewRouter(nil, machineController, nil, nil).SetUpRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/v1/machines/3", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{}\n", w.Body.String())
}