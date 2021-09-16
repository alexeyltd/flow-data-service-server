package controller

import (
	mock "flow-data-service-server/internal/service/mock"
	"flow-data-service-server/pkg/models/common"
	"flow-data-service-server/pkg/models/graph"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetGraph(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockGraphService := mock.NewMockGraphService(mockCtrl)
	g := NewGraphController(mockGraphService)

	recorder := httptest.NewRecorder()
	c, engine := gin.CreateTestContext(recorder)
	gin.SetMode(gin.TestMode)

	bodyReader := strings.NewReader(`{
    	"projectId": 1,
    	"id": 1
	}`)

	p := &common.ProjectModel{ProjectId: 1, Id: 1}

	mockGraphService.EXPECT().
		GetGraph(gomock.Any(), p).
		Return(&graph.DBGraph{
			ProjectModel: *p,
		}, nil)

	request := httptest.NewRequest(http.MethodGet, "/graph", bodyReader)
	engine.GET("/graph", g.GetGraph)
	engine.ServeHTTP(recorder, request)
	assert.Equal(t, 200, c.Writer.Status())
}
