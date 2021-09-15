package endpoint

import (
	"flow-data-service-server/pkg/models/common"
	"gorm.io/datatypes"
	"time"
)

type DBEndpoint struct {
	common.ProjectModel
	DataEndpoint
}

type DBFlowEndpoint struct {
	ID uint `gorm:"primaryKey"`
	DataEndpoint
}

type DataEndpoint struct {
	Uri     string `json:"uri"`
	Module  string `gorm:"uniq" json:"module"`
	Headers string `json:"headers"`
	Query   string `json:"query"`
}

type DBError struct {
	common.ProjectModel
	DataError
}

type DataError struct {
	Request    datatypes.JSON `json:"request"`
	Response   datatypes.JSON `json:"response"`
	Error      string
	StatusCode int       `json:"statusCode"`
	CreatedAt  time.Time `json:"createdAt"`
}
