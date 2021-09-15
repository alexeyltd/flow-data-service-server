package graph

import "flow-data-service-server/pkg/models/common"

type DBNode struct {
	GraphObject
	common.DataUI
	DataNode
}

type DataNode struct {
	//Placement
	LocalId uint `json:"localId"`

	//Invocation
	Module    string `json:"module"`
	Function  string `json:"function"`
	Arguments string `json:"arguments"`
}
