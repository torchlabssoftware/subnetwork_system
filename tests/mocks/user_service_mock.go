package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	models "github.com/torchlabssoftware/subnetwork_system/internal/server/models"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(ctx context.Context, user *models.CreateUserRequest) (*models.CreateUserResponce, int, string, error) {
	args := m.Called(ctx, user)

	var resp *models.CreateUserResponce
	if args.Get(0) != nil {
		resp = args.Get(0).(*models.CreateUserResponce)
	}
	return resp, args.Int(1), args.String(2), args.Error(3)
}
