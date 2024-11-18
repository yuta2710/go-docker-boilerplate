package entities

import "github.com/yuta_2710/go-clean-arc-reviews/shared"

type (
	User struct {
		shared.BaseSQLModel
		FirstName string `gorm:"column:firstName" json:"firstName"`
		LastName  string `gorm:"column:lastName" json:"lastName"`
		Email     string `gorm:"column:email;unique" json:"email"`
		Password  string `gorm:"column:password" json:"password"`
		Role      string `gorm:"column:role" json:"role"`
		IsActive  bool   `gorm:"column:isActive" json:"isActive"`
		IsAdmin   bool   `gorm:"column:isAdmin" json:"isAdmin"`
		IsBlocked bool   `gorm:"column:isBlocked" json:"isBlocked"`
	}

	InsertUserDto struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		Role      string `json:"role"`
		IsActive  bool   `json:"isActive"`
		IsAdmin   bool   `json:"isAdmin"`
		IsBlocked bool   `json:"isBlocked"`
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
