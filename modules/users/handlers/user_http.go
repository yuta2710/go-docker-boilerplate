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

	fmt.Println()

	// Check error of binding body
	if err := ctx.Bind(body); err != nil {
		return err
	}

	// Call use case from this layer
	authId, err := u.UserUsecase.InsertNewUser(body)

	if err != nil {
		return shared.Response(ctx, false, http.StatusBadRequest, "Error inserting account", nil, nil)
	}

	return shared.Response(ctx, true, http.StatusOK, "Inserted account successfully", authId, nil)
}

func (u *UserHttp) GetUserById(ctx echo.Context) error {
	id := ctx.Param("id")

	user, err := u.UserUsecase.FindById(id)

	if err != nil {
		return shared.Response(ctx, false, http.StatusNotFound, "User not found", nil, nil)
	}

	return shared.Response(ctx, true, http.StatusOK, "Successfully fetched", user, nil)
}

func (u *UserHttp) GetUsers(ctx echo.Context) error {
	users, err := u.UserUsecase.FindAll()

	if err != nil {
		return shared.Response(ctx, false, http.StatusNotFound, "Users not found or empty", nil, nil)
	}

	return shared.Response(ctx, true, http.StatusOK, "Successfully fetched", users, nil)
}

func NewUserHttp(uc usecases.UserUseCase) UserHandler {
	return &UserHttp{
		UserUsecase: uc,
	}
}
