package entities

import "time"

type TokenProvider struct {
	Id           int       `gorm:"column:id;primaryKey"`
	AuthId       string    `gorm:"column:auth_id"`
	AccessToken  string    `gorm:"column:access_token"`
	RefreshToken string    `gorm:"column:access_token"`
	ExpiredAt    time.Time `gorm:"column:expired_at"`
	CreatedAt    time.Time `gorm:"column:created_at"`
}
