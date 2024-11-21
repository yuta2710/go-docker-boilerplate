package entities

import (
	"time"

	"github.com/yuta_2710/go-clean-arc-reviews/shared"
)

type Priority int

const (
	Low Priority = iota + 1
	Medium
	High
)

func (p Priority) String() string {
	switch p {
	case Low:
		return "Low"
	case Medium:
		return "Medium"
	case High:
		return "High"
	default:
		return "Unknown"
	}
}

type Todo struct {
	shared.BaseSQLModel
	AuthId      string    `gorm:"column:auth_id;type:VARCHAR(255);index;not null" json:"authId"`
	Title       string    `gorm:"column:title;not null" json:"title"`
	Description string    `gorm:"column:description" json:"description"`
	IsCompleted bool      `gorm:"column:is_completed" json:"is_completed"`
	DueDate     time.Time `gorm:"column:due_date" json:"due_date"`
	Priority    Priority  `gorm:"column:priority" json:"priority"`
}

func (td *Todo) Mask(dbType shared.DbType) {
	td.BaseSQLModel.Mask(dbType)
}
