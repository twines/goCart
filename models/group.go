package models

import "github.com/jinzhu/gorm"

type Group struct {
	gorm.Model
	Members   []*Admin `gorm:"many2many:member_groups;"`
	Title     string   `form:"title" validate:"required,gt=3"`
	CreaterID uint     //创建这ID
	Status    int8
	Roles     []*Role `gorm:"many2many:group_roles;"`
}

func (g Group) GetByTitle() Group {
	group := Group{}
	DB().First(&group, "title=?", g.Title)
	return group
}
func AllGroups() []Group {
	groups := []Group{}
	DB().Find(&groups)
	return groups
}
