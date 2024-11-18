package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yuta_2710/go-clean-arc-reviews/modules/users/models"
	"github.com/yuta_2710/go-clean-arc-reviews/modules/users/usecases"
	"github.com/yuta_2710/go-clean-arc-reviews/shared"
)

type UserHttp struct {
	UserUsecase usecases.UserUseCase
}

func (u *UserHttp) CreateNewUser(ctx echo.Context) error {
	fmt.Println("CreateNewUser called") // Add a log here
	// Extract body
	body := new(models.InsertUserRequest)

	// Check error of binding body
	if err := ctx.Bind(body); err != nil {
		return err
	}

	// Call use case from this layer
	if err := u.UserUsecase.InsertNewUser(body); err != nil {
		return shared.Response(ctx, false, http.StatusBadRequest, "Error inserting account", nil)
	}

	return shared.Response(ctx, true, http.StatusOK, "Inserted account successfully", nil)
}

func (u *UserHttp) GetUserById(ctx echo.Context) error {
	fmt.Println("Get user by ID called")
	id := ctx.Param("id")

	user, err := u.UserUsecase.GetUserById(id)

	if err != nil {
		return shared.Response(ctx, false, http.StatusNotFound, "User not found", nil)
	}

	return shared.Response(ctx, true, http.StatusOK, "Successfully fetched", user)
}

func NewUserHttp(uc usecases.UserUseCase) UserHandler {
	return &UserHttp{
		UserUsecase: uc,
	}
}
