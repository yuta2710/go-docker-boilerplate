package repositories

import (
	"time"

	"github.com/yuta_2710/go-clean-arc-reviews/database"
	"github.com/yuta_2710/go-clean-arc-reviews/modules/token/entities"
)

type TokenPostgresRepository struct {
	db database.Database
}

func (tpr *TokenPostgresRepository) CreateTokens(authId string, accessToken string, refreshToken string, expiredAt time.Time) error {
	tokenProvider := entities.TokenProvider{
		AuthId:       authId,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiredAt:    expiredAt,
	}

	result := tpr.db.GetDb().Create(&tokenProvider)

	return result.Error
}
func (tpr *TokenPostgresRepository) ValidateRefreshToken(refreshToken string) error {
	return nil
}
func (tpr *TokenPostgresRepository) DeleteTokens(authId string) error {
	return nil
}

func NewTokenPostgresRepository(db database.Database) TokenRepository {
	return &TokenPostgresRepository{
		db: db,
	}
}
