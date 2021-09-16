package repository

import (
	"context"
	"database/sql"
	"flow-data-service-server/pkg/models/common"
	"flow-data-service-server/pkg/models/graph"
	"flow-data-service-server/pkg/models/storage"
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

func (g *graphRepositorySuite) TestRepositoryList() {
	r := &storage.ListGraphRequest{
		ProjectId: []uint{1},
	}
	g.mock.ExpectBegin()
	g.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "db_graphs"`)).
		WithArgs(r.ProjectId[0]).
		WillReturnRows(sqlmock.NewRows([]string{"project_id", "id", "name", "ui", "counter"}).
			AddRow(1, 1, "Telegram Test Postman", nil, 0))
	g.mock.ExpectCommit()
	res, err := g.graphRepository.ListGraph(context.Background(), r)
	require.NoError(g.T(), err)
	require.Equal(g.T(), r.ProjectId[0], res.Graphs[0].GetId())
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

func (g *graphRepositorySuite) TestRepositorySave() {
	p := &common.ProjectModel{ProjectId: 1, Id: 1}
	gr := &graph.DBGraph{
		ProjectModel: *p,
		DataUI:       common.DataUI{Name: "Telegram Test Postman"},
	}
	g.mock.ExpectBegin()
	g.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "db_graphs"`)).
		WithArgs(p.GetProjectId(), p.GetId(), p.GetId()).
		WillReturnRows(sqlmock.NewRows([]string{"project_id", "id", "name", "ui", "counter"}).
			AddRow(1, 1, "Telegram Test Postman", nil, 0))
	g.mock.ExpectExec(regexp.QuoteMeta(
		`UPDATE "db_graphs"`)).
		WithArgs(p.GetId(), p.GetProjectId(), p.GetProjectId(), p.GetId(), p.GetId(), p.GetId()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	g.mock.ExpectCommit()

	err := g.graphRepository.SaveProjectObject(context.Background(), p, gr)
	require.NoError(g.T(), err)
	require.Equal(g.T(), p.GetId(), gr.Id)
}

func (g *graphRepositorySuite) TestRepositoryDelete() {
	p := &common.ProjectModel{ProjectId: 1, Id: 1}
	gr := &graph.DBGraph{
		ProjectModel: *p,
		DataUI:       common.DataUI{Name: "Telegram Test Postman"},
	}
	g.mock.ExpectBegin()
	g.mock.ExpectExec(regexp.QuoteMeta(
		`DELETE FROM "db_graphs"`)).
		WithArgs(p.GetProjectId(), p.GetId(), p.GetId()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	g.mock.ExpectCommit()

	err := g.graphRepository.DeleteProjectObject(context.Background(), p, gr)
	require.NoError(g.T(), err)
	require.Equal(g.T(), p.GetId(), gr.Id)
}

func (g *graphRepositorySuite) AfterTest(_, _ string) {
	require.NoError(g.T(), g.mock.ExpectationsWereMet())
}

func TestUserServiceSuite(t *testing.T) {
	suite.Run(t, new(graphRepositorySuite))
}
