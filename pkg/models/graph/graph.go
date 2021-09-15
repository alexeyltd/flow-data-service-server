package graph

import "flow-data-service-server/pkg/models/common"

type GraphObject struct {
	common.ProjectModel
	GraphId uint `gorm:"index" json:"graphId,omitempty"`
}

func (i *GraphObject) GetGraphId() uint {
	return i.GraphId
}

type DBGraph struct {
	common.ProjectModel

	common.DataUI
	DataGraph

	Nodes       []DBNode       `gorm:"foreignKey:GraphId;references:Id;constraint:OnDelete:CASCADE;" json:"nodes,omitempty"`
	Cards       []DBEventCard  `gorm:"foreignKey:GraphId;references:Id;constraint:OnDelete:CASCADE;" json:"cards,omitempty"`
	Connections []DBConnection `gorm:"foreignKey:GraphId;references:Id;constraint:OnDelete:CASCADE;" json:"connections,omitempty"`
}

type DataGraph struct {
	Counter uint `json:"counter"`
}

type Object interface {
	common.ProjectObject
	GetGraphId() uint
}

//TODO what is that?
var _ Object = (*GraphObject)(nil)
