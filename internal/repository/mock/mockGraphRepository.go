package repository

import (
	"context"
	"flow-data-service-server/pkg/models/common"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockGraphRepository struct {
	mock.Mock
}

func (u *MockGraphRepository) GetProjectObject(ctx context.Context, object common.ProjectObject, entity common.ProjectObject) error {
	args := u.Called(ctx, object, entity)
	if args.Error(0) != nil {
		return gorm.ErrInvalidData
	}
	return nil
}

func (u *MockGraphRepository) SaveProjectObject(ctx context.Context, object common.ProjectObject, entity common.ProjectObject) error {
	args := u.Called(ctx, object, entity)
	if args.Error(0) != nil {
		return gorm.ErrInvalidData
	}
	entity = &common.ProjectModel{ProjectId: 1, Id: 1}
	return nil
}
