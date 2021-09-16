package service

import (
	"context"
	"flow-data-service-server/internal/repository"
	"flow-data-service-server/pkg/models/common"
	"flow-data-service-server/pkg/models/graph"
	"flow-data-service-server/pkg/models/storage"
)

// go:generate mockgen -destination=./internal/service/mock/MockGraphRepository_gen.go -package=service flow-data-service-server/internal/service GraphService

type GraphService interface {
	ListGraph(c context.Context, r *storage.ListGraphRequest) (*storage.ListGraphResponse, error)
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

func (g *GraphServiceImpl) ListGraph(c context.Context, r *storage.ListGraphRequest) (*storage.ListGraphResponse, error) {
	res, err := g.graphRepository.ListGraph(c, r)
	if err != nil {
		return nil, err
	}
	return res, nil
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
