package controller

import (
	mock "flow-data-service-server/internal/service/mock"
	"flow-data-service-server/pkg/models/common"
	"flow-data-service-server/pkg/models/graph"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type graphControllerSuite struct {
	suite.Suite

	mockService     *mock.MockGraphService
	graphController *GraphController
}

func (g *graphControllerSuite) SetupSuite() {
	g.mockService = new(mock.MockGraphService)
	g.graphController = NewGraphController(g.mockService)
}

func (g *graphControllerSuite) TestGetGraph() {
	recorder := httptest.NewRecorder()
	c, engine := gin.CreateTestContext(recorder)
	gin.SetMode(gin.TestMode)

	bodyReader := strings.NewReader(`{
    "projectId": 1,
    "id": 1
	}`)

	p := &common.ProjectModel{ProjectId: 1, Id: 1}

	g.mockService.On("GetGraph", engine, p).
		Return(&graph.DBGraph{
			ProjectModel: *p,
		}, nil)

	request := httptest.NewRequest(http.MethodGet, "/graph", bodyReader)
	engine.GET("/graph", g.graphController.GetGraph)
	engine.ServeHTTP(recorder, request)
	assert.Equal(g.T(), 200, c.Writer.Status())
}

func TestGraphControllerSuite(t *testing.T) {
	suite.Run(t, new(graphControllerSuite))
}
