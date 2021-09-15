package repository

import (
	"context"
	"database/sql"
	"flow-data-service-server/pkg/models/common"
	"flow-data-service-server/pkg/models/graph"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"regexp"
	"testing"
)

type graphRepositorySuite struct {
	suite.Suite

	db              *gorm.DB
	mock            sqlmock.Sqlmock
	graphRepository GraphRepository
}

func (g *graphRepositorySuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, g.mock, err = sqlmock.New()
	require.NoError(g.T(), err)

	g.db, err = gorm.Open(postgres.Dialector{Config: &postgres.Config{Conn: db}}, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	require.NoError(g.T(), err)

	g.graphRepository = NewGraphRepositoryImpl(g.db, zap.L().Sugar())
}

func (g *graphRepositorySuite) TestRepositoryGet() {
	p := &common.ProjectModel{ProjectId: 1, Id: 1}
	g.mock.ExpectBegin()
	g.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "db_graphs" WHERE project_id = $1 and id = $2 ORDER BY "db_graphs"."id" LIMIT 1`)).
		WithArgs(p.GetProjectId(), p.GetId()).
		WillReturnRows(sqlmock.NewRows([]string{"project_id", "id", "name", "ui", "counter"}).
			AddRow(1, 1, "Telegram Test Postman", nil, 0))
	g.mock.ExpectCommit()

	gr := new(graph.DBGraph)
	err := g.graphRepository.GetProjectObject(context.Background(), p, gr)
	require.NoError(g.T(), err)
	require.Equal(g.T(), p.GetId(), gr.Id)
}

func (g *graphRepositorySuite) AfterTest(_, _ string) {
	require.NoError(g.T(), g.mock.ExpectationsWereMet())
}

func TestUserServiceSuite(t *testing.T) {
	suite.Run(t, new(graphRepositorySuite))
}
