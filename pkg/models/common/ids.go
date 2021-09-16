package common

type ProjectModel struct {
	ProjectId uint `json:"projectId,omitempty" form:"projectId"`
	Id        uint `gorm:"primaryKey,omitempty" json:"id,omitempty" form:"id"`
}

type ProjectSpace struct {
	ProjectId uint `json:"projectId" form:"projectId"`
}

type ProjectObject interface {
	SpacedObject
	GetId() uint
}

type SpacedObject interface {
	GetProjectId() uint
}

var _ ProjectObject = (*ProjectModel)(nil)
var _ SpacedObject = (*ProjectSpace)(nil)

func (p *ProjectModel) GetId() uint {
	return p.Id
}

func (p *ProjectModel) GetProjectId() uint {
	return p.ProjectId
}

func (b *ProjectSpace) GetProjectId() uint {
	return b.ProjectId
}
