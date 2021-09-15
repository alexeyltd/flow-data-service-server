package controller

import (
	"github.com/gin-gonic/gin"
)

func NewGin(g *GraphController) (*gin.Engine, error) {
	engine := gin.Default()
	engine.GET("/graph", g.GetGraph)
	engine.POST("/graph", g.SaveGraph)
	return engine, nil
}
