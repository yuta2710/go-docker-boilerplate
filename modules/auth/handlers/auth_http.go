package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yuta_2710/go-clean-arc-reviews/modules/auth/models"
	"github.com/yuta_2710/go-clean-arc-reviews/modules/auth/usecases"
	"github.com/yuta_2710/go-clean-arc-reviews/shared"
)

type AuthHttp struct {
	AuthUsecase usecases.AuthUsecase
}

func (ahp *AuthHttp) Login(ctx echo.Context) error {
	fmt.Println("Started trigger login")
	// Check user exist in the repo
	body := new(models.LoginRequest)

	if err := ctx.Bind(body); err != nil {
		return shared.Response(ctx, false, http.StatusBadRequest, "Invalid login request body", nil)

	}
	fmt.Println(body)

	authResp, err := ahp.AuthUsecase.Login(body)

	if err != nil {
		return shared.Response(ctx, false, http.StatusBadRequest, err.Error(), nil)
	}

	return shared.Response(ctx, true, http.StatusOK, "Login successfully", &models.AuthResponse{
		AccessToken:  authResp.AccessToken,
		RefreshToken: authResp.RefreshToken,
	})
}

func (ahp *AuthHttp) SignUp(ctx echo.Context) error {
	return nil
}
func (ahp *AuthHttp) Profile(ctx echo.Context) error {
	return nil
}
func (ahp *AuthHttp) SignOut(ctx echo.Context) error {
	return nil
}

func NewAuthHttp(auc usecases.AuthUsecase) AuthHandler {
	return &AuthHttp{
		AuthUsecase: auc,
	}
}
