package service

//import (
//	"context"
//	"flow-data-service-server/pkg/models/common"
//	"flow-data-service-server/pkg/models/graph"
//	"flow-data-service-server/pkg/models/storage"
//	"github.com/stretchr/testify/mock"
//)
//
//type MockGraphService struct {
//	mock.Mock
//}
//
//func (g *MockGraphService) ListGraph(c context.Context, r *storage.ListGraphRequest) (*storage.ListGraphResponse, error) {
//	args := g.Called(c, r)
//	if args.Error(1) != nil {
//		return nil, args.Get(1).(error)
//	}
//	return args.Get(0).(*storage.ListGraphResponse), nil
//}
//
//func (g *MockGraphService) GetGraph(c context.Context, r *common.ProjectModel) (*graph.DBGraph, error) {
//	args := g.Called(mock.Anything, r)
//	if args.Error(1) != nil {
//		return nil, args.Get(1).(error)
//	}
//	return args.Get(0).(*graph.DBGraph), nil
//}
//
//func (g *MockGraphService) SaveGraph(c context.Context, data *graph.DBGraph) (*graph.DBGraph, error) {
//	args := g.Called(c, data)
//	if args.Error(0) != nil {
//		return nil, args.Get(1).(error)
//	}
//	return args.Get(0).(*graph.DBGraph), nil
//}
//
//func (g *MockGraphService) DeleteGraph(c context.Context, request *common.ProjectModel) error {
//	args := g.Called(c, request)
//	if args.Error(0) != nil {
//		return args.Get(1).(error)
//	}
//	return nil
//}
