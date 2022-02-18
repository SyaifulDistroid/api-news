package model

import "github.com/lib/pq"

type News struct {
	ID       int            `gorm:"primary_key" json:"id"`
	TopicsID int            `gorm:"colum:topics_id; type:int REFERENCES topics_tb(id)" json:"topics_id"`
	Tags     pq.StringArray `gorm:"type:text[]" json:"tags"`
	News     string         `gorm:"column:news" json:"news" validate:"required"`
	Status   string         `gorm:"column:status" json:"status" validate:"required"`
}

func (n News) TableName() string {
	return "news_tb"
}
