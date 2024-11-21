package common

import (
	"fmt"
	"log"

	"github.com/yuta_2710/go-clean-arc-reviews/database"
	TodoSchema "github.com/yuta_2710/go-clean-arc-reviews/modules/todo/entities"
	TokenProviderSchema "github.com/yuta_2710/go-clean-arc-reviews/modules/token/entities"
	UserSchema "github.com/yuta_2710/go-clean-arc-reviews/modules/users/entities"
)

func LoadRelations(repo database.Database) {
	fmt.Println("Load called")
	if err := repo.GetDb().AutoMigrate(&UserSchema.User{}, &TokenProviderSchema.TokenProvider{}, &TodoSchema.Todo{}, &TodoSchema.TodoMember{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
