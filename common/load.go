package common

import (
	"fmt"

	"github.com/yuta_2710/go-clean-arc-reviews/database"
	TokenProviderSchema "github.com/yuta_2710/go-clean-arc-reviews/modules/token/entities"
	UserSchema "github.com/yuta_2710/go-clean-arc-reviews/modules/users/entities"
)

func LoadRelations(repo database.Database) {
	fmt.Println("Load called")
	if !repo.GetDb().Migrator().HasTable("users") {
		repo.GetDb().Migrator().CreateTable(&UserSchema.User{})
	}

	if !repo.GetDb().Migrator().HasTable("token_providers") {
		repo.GetDb().Migrator().CreateTable(&TokenProviderSchema.TokenProvider{})
	}
}
