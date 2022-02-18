package model

type Topics struct {
	ID     int    `gorm:"primary_key" json:"id"`
	Topics string `gorm:"column:topics" json:"topics" validate:"required"`
}

func (t Topics) TableName() string {
	return "topics_tb"
}
