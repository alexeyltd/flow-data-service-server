package repository

import (
	"context"
	"flow-data-service-server/pkg/models/common"
	"flow-data-service-server/pkg/models/graph"
	"flow-data-service-server/pkg/models/storage"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type GraphRepository interface {
	ListGraph(c context.Context, r *storage.ListGraphRequest) (*storage.ListGraphResponse, error)
	GetProjectObject(ctx context.Context, object common.ProjectObject, entity common.ProjectObject) error
	SaveProjectObject(ctx context.Context, save common.ProjectObject, entity common.ProjectObject) error
	DeleteProjectObject(ctx context.Context, id *common.ProjectModel, entity common.ProjectObject) error
}

type GraphRepositoryImpl struct {
	db     *gorm.DB
	logger *zap.SugaredLogger
}

func NewGraphRepositoryImpl(db *gorm.DB, logger *zap.SugaredLogger) *GraphRepositoryImpl {
	return &GraphRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (g *GraphRepositoryImpl) ListGraph(c context.Context, r *storage.ListGraphRequest) (*storage.ListGraphResponse, error) {
	var graphs []graph.DBGraph
	err := g.db.
		Session(&gorm.Session{Context: c}).
		Transaction(func(tx *gorm.DB) error {
			res := tx.
				Where("project_id in ?", r.ProjectId).
				Find(&graphs)

			if res.Error != nil {
				return res.Error
			}
			return nil
		})

	if err != nil {
		return nil, err
	}

	return &storage.ListGraphResponse{
		Graphs: graphs,
	}, nil
}

func (g *GraphRepositoryImpl) GetProjectObject(ctx context.Context, object common.ProjectObject, entity common.ProjectObject) error {
	err := g.db.
		Session(&gorm.Session{
			Context: ctx,
		}).
		Transaction(func(tx *gorm.DB) error {
			r := tx.Where(
				"project_id = ? and id = ?",
				object.GetProjectId(),
				object.GetId(),
			).First(entity)
			if r.Error != nil {
				return r.Error
			}
			return nil
		})
	if err != nil {
		g.logger.Error("getting project object failed", zap.Error(err))
		return err
	}
	return nil
}

func (g *GraphRepositoryImpl) SaveProjectObject(ctx context.Context, save common.ProjectObject, entity common.ProjectObject) error {
	err := g.db.
		Session(&gorm.Session{
			Context: ctx,
		}).
		Transaction(func(tx *gorm.DB) error {
			r := tx.Where(
				"project_id = ? and id = ?",
				save.GetProjectId(),
				save.GetId(),
			).
				Assign(save).
				FirstOrCreate(entity)
			if r.Error != nil {
				return r.Error
			}
			return nil
		})
	if err != nil {
		g.logger.Error("saving project object failed", zap.Error(err), zap.Any("save", save))
		return err
	}
	return nil
}

func (g *GraphRepositoryImpl) DeleteProjectObject(ctx context.Context, id *common.ProjectModel, entity common.ProjectObject) error {
	err := g.db.
		Session(&gorm.Session{
			Context: ctx,
		}).
		Transaction(func(tx *gorm.DB) error {
			r := tx.Where(
				"project_id = ? and id = ?",
				id.GetProjectId(),
				id.GetId(),
			).Delete(entity)

			if r.Error != nil {
				return r.Error
			}
			return nil
		})
	if err != nil {
		g.logger.Error("deleting project object failed", zap.Error(err))
		return err
	}
	return nil
}
