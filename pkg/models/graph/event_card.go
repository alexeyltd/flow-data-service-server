package graph

import "autoflow/pkg/common"

type DBEventCard struct {
	GraphObject

	common.DataUI
	DataEventCard
}

type DataEvent struct {
	Platform string `json:"platform,omitempty" `

	OwnerType string `json:"ownerType"`
	OwnerId   string `json:"ownerId"`

	ResourceType string `json:"resourceType,omitempty"`
	ResourceId   string `json:"resourceId,omitempty"`

	ContextType string `json:"contextType,omitempty"`
	ContextId   string `json:"contextId,omitempty"`

	InitiatorType string `json:"initiatorType,omitempty"`
	InitiatorId   string `json:"initiatorId,omitempty"`

	StaticType string `json:"staticType,omitempty"`
	StaticId   string `json:"staticId,omitempty"`
}

type DataEventCard struct {
	HttpVote uint `json:"httpVote,omitempty"`

	//Placement
	TargetId  uint   `json:"targetId"`
	SlidePort string `json:"slidePort,omitempty"`

	DataEvent
}
