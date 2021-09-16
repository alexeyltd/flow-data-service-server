package service

import (
	"context"
	"flow-data-service-server/internal/repository/mock"
	"flow-data-service-server/pkg/models/common"
	"flow-data-service-server/pkg/models/graph"
	"flow-data-service-server/pkg/models/storage"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestListGraph(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock.NewMockGraphRepository(mockCtrl)

	r := &storage.ListGraphRequest{
		ProjectId: []uint{1},
	}

	p := &common.ProjectModel{ProjectId: 1, Id: 12}
	dbGraph := graph.DBGraph{ProjectModel: *p}
	dbGraphs := []graph.DBGraph{dbGraph}
	mockRepo.EXPECT().
		ListGraph(context.Background(), r).
		Return(&storage.ListGraphResponse{
			Graphs: dbGraphs,
		}, nil).
		Times(1)

	testService := NewGraphServiceImpl(mockRepo)

	gr, _ := testService.ListGraph(context.Background(), r)

	assert.Equal(t, uint(12), gr.Graphs[0].Id)

}

func TestGetGraph(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock.NewMockGraphRepository(mockCtrl)
	p := &common.ProjectModel{ProjectId: 1, Id: 1}

	mockRepo.EXPECT().
		GetProjectObject(context.Background(), p, new(graph.DBGraph)).
		Do(func(ctx context.Context, object common.ProjectObject, entity common.ProjectObject) {
			dbGraph := entity.(*graph.DBGraph)
			dbGraph.ProjectModel = *p
		}).
		Return(nil)

	testService := NewGraphServiceImpl(mockRepo)
	gr, _ := testService.GetGraph(context.Background(), p)
	assert.Equal(t, uint(1), gr.GetId())
}

func TestSaveGraph(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock.NewMockGraphRepository(mockCtrl)

	p := &common.ProjectModel{ProjectId: 1, Id: 1}
	dbGraph := &graph.DBGraph{ProjectModel: *p}

	mockRepo.EXPECT().
		SaveProjectObject(context.Background(), dbGraph, new(graph.DBGraph)).
		Do(func(ctx context.Context, object common.ProjectObject, entity common.ProjectObject) {
			dbGraph := entity.(*graph.DBGraph)
			dbGraph.ProjectModel = *p
		}).
		Return(nil)

	testService := NewGraphServiceImpl(mockRepo)
	gr, _ := testService.SaveGraph(context.Background(), dbGraph)
	assert.Equal(t, uint(1), gr.GetId())
}

func TestDeleteGraph(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock.NewMockGraphRepository(mockCtrl)

	p := &common.ProjectModel{ProjectId: 1, Id: 1}
	dbGraph := &graph.DBGraph{ProjectModel: *p}

	mockRepo.EXPECT().
		DeleteProjectObject(context.Background(), p, new(graph.DBGraph)).
		Do(func(ctx context.Context, object common.ProjectObject, entity common.ProjectObject) {
			dbGraph := entity.(*graph.DBGraph)
			dbGraph.ProjectModel = *p
		}).
		Return(nil)

	testService := NewGraphServiceImpl(mockRepo)
	er := testService.DeleteGraph(context.Background(), p)
	require.NoError(t, er)
	assert.Equal(t, uint(1), dbGraph.GetId())
}

// OLD WAY
// ------------------

//func TestListGraph(t *testing.T) {
//	mockRepo := new(repository.MockGraphRepository)
//
//	r := &storage.ListGraphRequest{
//		ProjectId: []uint{1},
//	}
//
//	p := &common.ProjectModel{ProjectId: 1, Id: 1}
//	dbGraph := graph.DBGraph{ProjectModel: *p}
//	dbGraphs := []graph.DBGraph{dbGraph}
//	mockRepo.On("ListGraph", context.Background(), r).Return(&storage.ListGraphResponse{
//		Graphs: dbGraphs,
//	}, nil)
//
//	testService := NewGraphServiceImpl(mockRepo)
//
//	gr, _ := testService.ListGraph(context.Background(), r)
//
//	mockRepo.AssertExpectations(t)
//	assert.Equal(t, p.GetProjectId(), gr.Graphs[0].Id)
//}
//
//func TestGetGraph(t *testing.T) {
//	mockRepo := new(repository.MockGraphRepository)
//
//	p := &common.ProjectModel{ProjectId: 1, Id: 1}
//
//	mockRepo.On("GetProjectObject", context.Background(), p, new(graph.DBGraph)).
//		Return(nil).Run(func(args mock.Arguments) {
//		arg := args.Get(2).(*graph.DBGraph)
//		arg.ProjectModel = *p
//	})
//
//	testService := NewGraphServiceImpl(mockRepo)
//
//	gr, _ := testService.GetGraph(context.Background(), p)
//
//	mockRepo.AssertExpectations(t)
//	assert.Equal(t, uint(1), gr.GetId())
//}
//
//func TestSaveGraph(t *testing.T) {
//	mockRepo := new(repository.MockGraphRepository)
//
//	p := &common.ProjectModel{ProjectId: 1, Id: 1}
//	dbGraph := &graph.DBGraph{ProjectModel: *p}
//
//	mockRepo.On("SaveProjectObject", context.Background(), dbGraph, new(graph.DBGraph)).
//		Return(nil).Run(func(args mock.Arguments) {
//		arg := args.Get(2).(*graph.DBGraph)
//		arg.ProjectModel = *p
//	})
//
//	testService := NewGraphServiceImpl(mockRepo)
//
//	gr, _ := testService.SaveGraph(context.Background(), dbGraph)
//
//	mockRepo.AssertExpectations(t)
//	assert.Equal(t, uint(1), gr.GetId())
//}
//
//func TestDeleteGraph(t *testing.T) {
//	mockRepo := new(repository.MockGraphRepository)
//
//	p := &common.ProjectModel{ProjectId: 1, Id: 1}
//	dbGraph := &graph.DBGraph{ProjectModel: *p}
//
//	mockRepo.On("DeleteProjectObject", context.Background(), p, new(graph.DBGraph)).
//		Return(nil).Run(func(args mock.Arguments) {
//		arg := args.Get(2).(*graph.DBGraph)
//		arg.ProjectModel = *p
//	})
//
//	testService := NewGraphServiceImpl(mockRepo)
//
//	er := testService.DeleteGraph(context.Background(), p)
//
//	mockRepo.AssertExpectations(t)
//	require.NoError(t, er)
//	assert.Equal(t, uint(1), dbGraph.GetId())
//}
