package entities

import (
	"github.com/yuta_2710/go-clean-arc-reviews/modules/todo/entities"
	"github.com/yuta_2710/go-clean-arc-reviews/shared"
)

type (
	User struct {
		shared.BaseSQLModel
		FirstName string          `gorm:"column:first_name" json:"firstName"`
		LastName  string          `gorm:"column:last_name" json:"lastName"`
		Email     string          `gorm:"column:email;unique" json:"email"`
		Password  string          `gorm:"column:password" json:"password"`
		Role      string          `gorm:"column:role" json:"role"`
		IsActive  bool            `gorm:"column:is_active" json:"isActive"`
		IsAdmin   bool            `gorm:"column:is_admin" json:"isAdmin"`
		IsBlocked bool            `gorm:"column:is_blocked" json:"isBlocked"`
		IsDeleted bool            `gorm:"column:is_deleted" json:"isDeleted"`
		Todos     []entities.Todo `gorm:"foreignKey:AuthId;constraint:OnDelete:CASCADE;" json:"todos"`
	}

	InsertUserDto struct {
		FirstName string          `json:"first_name"`
		LastName  string          `json:"last_name"`
		Email     string          `json:"email"`
		Password  string          `json:"password"`
		Role      string          `json:"role"`
		IsActive  bool            `json:"is_active"`
		IsAdmin   bool            `json:"is_admin"`
		IsBlocked bool            `json:"is_blocked"`
		IsDeleted bool            `gorm:"column:is_deleted" json:"isDeleted"`
		Todos     []entities.Todo `gorm:"foreignKey:AuthId;constraint:OnDelete:CASCADE;" json:"todos"`
	}

	FetchUserDto struct {
		FakeId    string `json:"fakeId"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
	}
)

func (u *User) Mask(dbType shared.DbType) {
	u.BaseSQLModel.Mask(dbType)
}
