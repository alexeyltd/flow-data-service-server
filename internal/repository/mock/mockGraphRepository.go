package mock

//
//import (
//	"context"
//	"flow-data-service-server/pkg/models/common"
//	"flow-data-service-server/pkg/models/storage"
//	"github.com/stretchr/testify/mock"
//)
//
//type MockGraphRepository struct {
//	mock.Mock
//}
//
//func (u *MockGraphRepository) ListGraph(c context.Context, r *storage.ListGraphRequest) (*storage.ListGraphResponse, error) {
//	args := u.Called(c, r)
//	if args.Error(1) != nil {
//		return nil, args.Get(1).(error)
//	}
//	return args.Get(0).(*storage.ListGraphResponse), nil
//}
//
//func (u *MockGraphRepository) GetProjectObject(ctx context.Context, object common.ProjectObject, entity common.ProjectObject) error {
//	args := u.Called(ctx, object, entity)
//	if args.Error(0) != nil {
//		return args.Get(0).(error)
//	}
//	return nil
//}
//
//func (u *MockGraphRepository) SaveProjectObject(ctx context.Context, object common.ProjectObject, entity common.ProjectObject) error {
//	args := u.Called(ctx, object, entity)
//	if args.Error(0) != nil {
//		return args.Get(0).(error)
//	}
//	return nil
//}
//
//func (u *MockGraphRepository) DeleteProjectObject(ctx context.Context, object *common.ProjectModel, entity common.ProjectObject) error {
//	args := u.Called(ctx, object, entity)
//	if args.Error(0) != nil {
//		return args.Get(0).(error)
//	}
//	return nil
//}
