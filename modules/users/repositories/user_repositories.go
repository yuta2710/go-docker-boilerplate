package repositories

import "github.com/yuta_2710/go-clean-arc-reviews/modules/users/entities"

type UserRepository interface {
	Insert(in *entities.InsertUserDto) error
	InsertBatch(in []*entities.InsertUserDto) error
	GetUserById(id string) (*entities.User, error)
}
