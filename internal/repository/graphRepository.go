package repository

import (
	"context"
	"flow-data-service-server/pkg/models/common"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type GraphRepository interface {
	GetProjectObject(ctx context.Context, object common.ProjectObject, entity common.ProjectObject) error
	SaveProjectObject(ctx context.Context, save common.ProjectObject, entity common.ProjectObject) error
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
