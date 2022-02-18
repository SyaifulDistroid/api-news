package model

type Tags struct {
	ID  int    `gorm:"primary_key" json:"id"`
	Tag string `gorm:"column:tag" json:"tag" validate:"required"`
}

func (t Tags) TableName() string {
	return "tags_tb"
}
