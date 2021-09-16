package storage

import (
	"flow-data-service-server/pkg/models/graph"
)

type ListGraphRequest struct {
	ProjectId []uint `json:"projectID" form:"projectId"`
}

type ListGraphResponse struct {
	Graphs []graph.DBGraph `json:"graphs"`
}
