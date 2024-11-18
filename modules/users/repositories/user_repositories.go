package repositories

import "github.com/yuta_2710/go-clean-arc-reviews/modules/users/entities"

type UserRepository interface {
	Insert(in *entities.InsertUserDto) error
	InsertBatch(in []*entities.InsertUserDto) error
	FindById(id string) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
}
