package usecases

import (
	"github.com/yuta_2710/go-clean-arc-reviews/modules/users/entities"
	"github.com/yuta_2710/go-clean-arc-reviews/modules/users/models"
)

type UserUseCase interface {
	InsertNewUser(mod *models.InsertUserRequest) error
	GetUserById(id string) (*entities.FetchUserDto, error)
}
