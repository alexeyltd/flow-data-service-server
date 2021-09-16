package service

import (
	"context"
	"flow-data-service-server/internal/repository"
	"flow-data-service-server/pkg/models/common"
	"flow-data-service-server/pkg/models/graph"
)

type GraphService interface {
	GetGraph(c context.Context, r *common.ProjectModel) (*graph.DBGraph, error)
	SaveGraph(c context.Context, data *graph.DBGraph) (*graph.DBGraph, error)
	DeleteGraph(c context.Context, request *common.ProjectModel) error
}

type GraphServiceImpl struct {
	graphRepository repository.GraphRepository
}

func NewGraphServiceImpl(graphRepository repository.GraphRepository) *GraphServiceImpl {
	return &GraphServiceImpl{
		graphRepository: graphRepository,
	}
}

func (g *GraphServiceImpl) GetGraph(c context.Context, r *common.ProjectModel) (*graph.DBGraph, error) {
	gr := new(graph.DBGraph)
	err := g.graphRepository.GetProjectObject(c, r, gr)
	if err != nil {
		return nil, err
	}
	return gr, nil
}

func (g *GraphServiceImpl) SaveGraph(c context.Context, data *graph.DBGraph) (*graph.DBGraph, error) {
	entity := &graph.DBGraph{}
	err := g.graphRepository.SaveProjectObject(c, data, entity)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (g *GraphServiceImpl) DeleteGraph(c context.Context, request *common.ProjectModel) error {
	return g.graphRepository.DeleteProjectObject(c, request, &graph.DBGraph{})
}
