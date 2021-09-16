package service

import (
	"context"
	repository "flow-data-service-server/internal/repository/mock"
	"flow-data-service-server/pkg/models/common"
	"flow-data-service-server/pkg/models/graph"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetGraph(t *testing.T) {
	mockRepo := new(repository.MockGraphRepository)

	p := &common.ProjectModel{ProjectId: 1, Id: 1}

	mockRepo.On("GetProjectObject", context.Background(), p, new(graph.DBGraph)).
		Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(2).(*graph.DBGraph)
		arg.ProjectModel = *p
	})

	testService := NewGraphServiceImpl(mockRepo)

	gr, _ := testService.GetGraph(context.Background(), p)

	mockRepo.AssertExpectations(t)
	assert.Equal(t, uint(1), gr.GetId())
}

func TestSaveGraph(t *testing.T) {
	mockRepo := new(repository.MockGraphRepository)

	p := &common.ProjectModel{ProjectId: 1, Id: 1}
	dbGraph := &graph.DBGraph{ProjectModel: *p}

	mockRepo.On("SaveProjectObject", context.Background(), dbGraph, new(graph.DBGraph)).
		Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(2).(*graph.DBGraph)
		arg.ProjectModel = *p
	})

	testService := NewGraphServiceImpl(mockRepo)

	gr, _ := testService.SaveGraph(context.Background(), dbGraph)

	mockRepo.AssertExpectations(t)
	assert.Equal(t, uint(1), gr.GetId())
}

func TestDeleteGraph(t *testing.T) {
	mockRepo := new(repository.MockGraphRepository)

	p := &common.ProjectModel{ProjectId: 1, Id: 1}
	dbGraph := &graph.DBGraph{ProjectModel: *p}

	mockRepo.On("DeleteProjectObject", context.Background(), p, new(graph.DBGraph)).
		Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(2).(*graph.DBGraph)
		arg.ProjectModel = *p
	})

	testService := NewGraphServiceImpl(mockRepo)

	er := testService.DeleteGraph(context.Background(), p)

	mockRepo.AssertExpectations(t)
	require.NoError(t, er)
	assert.Equal(t, uint(1), dbGraph.GetId())
}
