package mock

import (
	database "StudentUniverse/StudentUniverseApp/Facade/Database"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockUser struct {
	mock.Mock
}

// NewMockUser creates a new mock instance.
func NewMockUser(t *testing.T) *MockUser {
	t.Helper()
	return &MockUser{}
}

// GetProduct mocks base method.
func (m *MockUser) signUp(email string, password string) (*database.Users, error) {
	args := m.Called(email, password)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*database.Users), nil
}

// CreateProduct mocks base method.
// func (m *MockUser) CreateProduct(product *database.Users, role string, email string) (*database.Users, error) {
// 	args := m.Called(product, role, email)
// 	if args.Error(1) != nil {
// 		return nil, args.Error(1)
// 	}
// 	return args.Get(0).(*models.Product), nil
// }

// // GetAllProducts mocks base method.
// func (m *MockUser) GetAllProducts(email string) ([]primitive.M, error) {
// 	args := m.Called(email)
// 	if args.Error(1) != nil {
// 		return nil, args.Error(1)
// 	}
// 	return args.Get(0).([]primitive.M), nil
// }
