package database

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	ID uint `gorm:"primaryKey;autoIncrement" json:"id,omitempty"`
	Name string `gorm:"not null;unique" json:"name"`
	Description string `gorm:"" json:"description,omitempty"`
}

func (g *Group) Out() GroupOut {
	return GroupOut{
		ID:          g.ID,
		Name:        g.Name,
		Description: g.Description,
	}
}

func (g *GroupIn) In() Group {
	return Group{
		Name:        g.Name,
		Description: g.Description,
	}
}

func GroupArrayOut(groupModels []Group) []GroupOut {
	outArr:=make([]GroupOut,len(groupModels))
	for i,group := range groupModels {
		outArr[i]=group.Out()
	}
	return outArr
}

type GroupIn struct {
	Name string `json:"name"`
	Description string `json:"description,omitempty"`
}

type GroupOut struct {
	ID uint `json:"id,omitempty"`
	Name string `json:"name"`
	Description string `json:"description,omitempty"`
}


