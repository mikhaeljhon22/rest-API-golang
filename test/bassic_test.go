package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"restGolang/controller"
	"restGolang/model"
)

// Mock service
type MockUserService struct{}

func (m *MockUserService) CreateAcc(user *model.UserNews) error {
	// Simulate success tanpa DB
	return nil
}

func (m *MockUserService) GenerateJwt(username string) (string, error) {
	return "mocked-jwt-token", nil
}

func TestCreateAcc(t *testing.T){
	mockService := &MockUserService{}
	controller.SetUserService(mockService)

	router := gin.Default()
	router.POST("/api/service/create/acc", controller.CreateAcc)

	jsonBody := []byte(`{
	    "Name": "mikhael chan",
		"Username": "mikhael123",
		"Email": "mikhael@example.com",
		"Password": "secret123"
	}`)

	req, _ := http.NewRequest("POST","/api/service/create/acc", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w,req)

	t.Log("Response body:", w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
}
