package controller

import (
	"flow-data-service-server/internal/service"
	"flow-data-service-server/pkg/models/common"
	"flow-data-service-server/pkg/models/graph"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GraphController struct {
	GraphService service.GraphService
}

func NewGraphController(graphService service.GraphService) *GraphController {
	return &GraphController{
		GraphService: graphService,
	}
}

func (c *GraphController) GetGraph(g *gin.Context) {
	obj := new(common.ProjectModel)
	e := g.BindJSON(obj)
	if e != nil {
		// TODO Validation error
		_ = g.Error(e)
		g.JSON(http.StatusBadRequest, gin.H{
			"error": e.Error(),
		})
		return
	}
	res, e := c.GraphService.GetGraph(g, obj)
	if e != nil {
		_ = g.Error(e)
		g.JSON(http.StatusBadRequest, gin.H{
			"error": e.Error(),
		})
		return
	}
	g.JSON(http.StatusOK, res)
}

func (c *GraphController) SaveGraph(g *gin.Context) {
	obj := new(graph.DBGraph)
	err := g.BindJSON(obj)
	if err != nil {
		_ = g.Error(err)
		return
	}
	res, err := c.GraphService.SaveGraph(g, obj)
	if err != nil {
		_ = g.Error(err)
		return
	}
	g.JSON(http.StatusOK, res)
}
