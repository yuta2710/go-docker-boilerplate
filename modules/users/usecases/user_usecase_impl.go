package usecases

import (
	"fmt"

	"github.com/yuta_2710/go-clean-arc-reviews/modules/users/entities"
	"github.com/yuta_2710/go-clean-arc-reviews/modules/users/models"
	"github.com/yuta_2710/go-clean-arc-reviews/modules/users/repositories"
	"github.com/yuta_2710/go-clean-arc-reviews/shared"
)

type UserUsecaseImpl struct {
	userRepo repositories.UserRepository
}

func (uui *UserUsecaseImpl) InsertNewUser(mod *models.InsertUserRequest) error {
	insertData := &entities.InsertUserDto{
		FirstName: mod.FirstName,
		LastName:  mod.LastName,
		Email:     mod.Email,
		Password:  mod.Password,
		Role:      "user",
		IsActive:  true,
		IsAdmin:   false,
		IsBlocked: false,
	}

	if err := uui.userRepo.Insert(insertData); err != nil {
		return err
	}

	fmt.Println("[CREATE ACCOUNT FROM USECASE LAYER SUCCESSFULLY]")
	return nil
}

func (uui *UserUsecaseImpl) GetUserById(id string) (*entities.FetchUserDto, error) {
	// Get user from repo
	user, err := uui.userRepo.GetUserById(id)

	// check err is not nil
	if err != nil {
		panic("User not found")
		// return nil, nil
	}

	uid := shared.NewUID(uint32(user.BaseSQLModel.Id), int(shared.DbTypeUser), 1)
	// convert to dto
	userDto := &entities.FetchUserDto{
		FakeId:    uid.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}

	return userDto, nil
}

func NewUserUsecaseImpl(userRepo repositories.UserRepository) UserUseCase {
	return &UserUsecaseImpl{
		userRepo: userRepo,
	}
}
