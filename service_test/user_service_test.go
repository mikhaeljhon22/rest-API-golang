package service_test

import (
	"errors"
	"restGolang/model"
	"restGolang/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository untuk mock repository
type MockUserRepository struct{
	mock.Mock
}
// Implementasi mock untuk Create
func (m *MockUserRepository) Create (user *model.Users) error {
	args := m.Called(user)
	return args.Error(0)
}
// Implementasi metode lain jika diperlukan, tapi kita hanya membutuhkan Create
func (m *MockUserRepository) FindAll() ([]model.Users, error) {
	return nil, nil
}

func (m *MockUserRepository) FindByUsername(username string) (*model.Users, error) {
	return nil, nil
}

func (m *MockUserRepository) UpdateUsername(id uint, username string) error {
	return nil
}

func (m *MockUserRepository) DeleteByUsername(username string) error {
	return nil
}

// Test untuk CreateUser (Success)
func TestCreateUser_Success(t *testing.T) {
	// Membuat mock repository
	mockRepo := new(MockUserRepository)

	// Membuat service dengan mock repository
	userSvc := service.NewUserService(mockRepo)

	// Data yang akan digunakan untuk test
	testUser := &model.Users{Username: "tester"}

	// Ekspektasi: Create dipanggil dengan testUser dan tidak error
	mockRepo.On("Create", testUser).Return(nil)

	// Memanggil CreateUser
	err := userSvc.CreateUser(testUser)

	// Assert jika tidak ada error
	assert.NoError(t, err)

	// Verifikasi bahwa mock repo dipanggil sesuai ekspektasi
	mockRepo.AssertExpectations(t)
}

func TestCreateUser_Error(t *testing.T) {
	// Membuat mock repository
	mockRepo := new(MockUserRepository)

	// Membuat service dengan mock repository
	userSvc := service.NewUserService(mockRepo)

	// Data yang akan digunakan untuk test
	testUser := &model.Users{Username: "tester"}

	// Ekspektasi: Create dipanggil dengan testUser dan menghasilkan error
	mockRepo.On("Create", testUser).Return(errors.New("insert failed"))

	// Memanggil CreateUser
	err := userSvc.CreateUser(testUser)

	// Assert jika error yang dihasilkan sesuai
	assert.Error(t, err)
	assert.EqualError(t, err, "insert failed")

	// Verifikasi bahwa mock repo dipanggil  ekspektasi
	mockRepo.AssertExpectations(t)
}
