package infra

import (
	"flow-data-service-server/pkg/models/endpoint"
	graph2 "flow-data-service-server/pkg/models/graph"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strings"
)

func NewGorm(config *AppConfig) (*gorm.DB, error) {
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			NameReplacer: strings.NewReplacer("db_", ""),
		},
	}
	if config.ShowSql {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}
	db, err := gorm.Open(postgres.Open(config.DSN), gormConfig)

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&graph2.DBGraph{},
		&graph2.DBEventCard{},
		&graph2.DBNode{},
		&graph2.DBConnection{},
		&endpoint.DBEndpoint{},
		&endpoint.DBError{},
		&endpoint.DBFlowEndpoint{},
	)
	if err != nil {
		return nil, err
	}
	return db, nil
}
