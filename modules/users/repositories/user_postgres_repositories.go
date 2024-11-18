package repositories

import (
	"fmt"
	"log"
	"strings"

	"github.com/yuta_2710/go-clean-arc-reviews/database"
	"github.com/yuta_2710/go-clean-arc-reviews/modules/users/entities"
	"github.com/yuta_2710/go-clean-arc-reviews/shared"
)

type UserPostgresRepository struct {
	db database.Database
}

func (ur *UserPostgresRepository) Insert(dto *entities.InsertUserDto) error {
	insert := &entities.User{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
		Password:  dto.Password, // Ensure password is hashed before calling this
		Role:      dto.Role,
	}

	hash, err := shared.HashPassword(insert.Password)

	if err != nil {
		log.Fatal("Error hashing password")
	}
	if hash != "" {
		insert.Password = hash
	}

	result := ur.db.GetDb().Create(insert)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "unique constraint") {
			fmt.Println("This email already exists")
			panic(result.Error.Error())
		}
	}

	fmt.Println("[INSERTED DATA SUCCESSFULLY]")
	insert.Mask(shared.DbTypeUser)

	return nil
}

func (ur *UserPostgresRepository) InsertBatch(dtos []*entities.InsertUserDto) error {
	return nil
}

func (ur *UserPostgresRepository) FindById(id string) (*entities.User, error) {
	// Get the id
	var u *entities.User
	result := ur.db.GetDb().Where("id = ?", id).First(&u)

	if result.Error != nil {
		return nil, result.Error
	}
	return u, nil
}

func (ur *UserPostgresRepository) FindByEmail(email string) (*entities.User, error) {
	var u *entities.User
	result := ur.db.GetDb().Where("email = ?", email).First(&u)

	if result.Error != nil {
		return nil, result.Error
	}
	return u, nil
}

func NewUserPostgresRepository(db database.Database) UserRepository {
	return &UserPostgresRepository{
		db: db,
	}
}
