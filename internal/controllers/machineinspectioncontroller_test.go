package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"bytes"
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/fercarcedo/quadrant-server/internal/controllers"
	"github.com/fercarcedo/quadrant-server/internal/models"
)

type mockMachineInspectionService struct {
	mock.Mock
}

func (m *mockMachineInspectionService) GetMachineInspectionsByMachineId(machineId int, onlyObservations bool) ([]models.MachineInspection, error) {
	args := m.Called(machineId, onlyObservations)
	return args.Get(0).([]models.MachineInspection), args.Error(1)
}

func (m *mockMachineInspectionService) CreateMachineInspection(machineInspection *models.MachineInspection) error {
	args := m.Called(machineInspection)
	return args.Error(0)
}

func TestGetMachineInspectionsByMachineId(t *testing.T) {
	machineInspectionService := new(mockMachineInspectionService)
	machineInspections := []models.MachineInspection { 
		models.MachineInspection { MachineId: 2, Date: time.Now(), Observations: "obs1" }, 
		models.MachineInspection { MachineId: 2, Date: time.Now(), Observations: "obs2" }, 
	}
	machineInspectionService.On("GetMachineInspectionsByMachineId", 2, false).Return(machineInspections, nil)
	machineInspectionController := controllers.NewMachineInspectionController(machineInspectionService)
	router := controllers.NewRouter(nil, nil, machineInspectionController, nil).SetUpRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/machines/2/inspections", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	jsonBytes, _ := json.Marshal(&machineInspections)
	assert.Equal(t, string(jsonBytes)+"\n", w.Body.String())
}

func TestCreateMachineInspection(t *testing.T) {
	machineInspectionService := new(mockMachineInspectionService)
	machineInspectionService.On("CreateMachineInspection", mock.MatchedBy(func(machineInspection *models.MachineInspection) bool {
		return machineInspection.MachineId == 2 &&
			machineInspection.Observations == "obs1"
	})).Return(nil)
	machineInspectionController := controllers.NewMachineInspectionController(machineInspectionService)
	router := controllers.NewRouter(nil, nil, machineInspectionController, nil).SetUpRouter()
	w := httptest.NewRecorder()
	jsonStr, _ := json.Marshal(&models.MachineInspection { Observations: "obs1" })
	req, _ := http.NewRequest("POST", "/api/v1/machines/2/inspections", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
	assert.Equal(t, "{}\n", w.Body.String())
}