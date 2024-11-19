package usecases

import (
	"fmt"
	"time"

	"github.com/yuta_2710/go-clean-arc-reviews/modules/auth/models"
	TokenRepo "github.com/yuta_2710/go-clean-arc-reviews/modules/token/repositories"
	"github.com/yuta_2710/go-clean-arc-reviews/modules/users/entities"
	"github.com/yuta_2710/go-clean-arc-reviews/modules/users/repositories"
	UserRepo "github.com/yuta_2710/go-clean-arc-reviews/modules/users/repositories"
	"github.com/yuta_2710/go-clean-arc-reviews/shared"
)

type AuthUsecaseImpl struct {
	UserRepo  UserRepo.UserRepository
	TokenRepo TokenRepo.TokenRepository
}

func (aui *AuthUsecaseImpl) Login(mod *models.LoginRequest) (*models.AuthResponse, error) {
	u, err := aui.UserRepo.FindByEmail(mod.Email)

	// fmt.Println("Mod email ", mod.Email)

	if err != nil {
		return nil, fmt.Errorf("[Login failed]: Email is not valid")
	}

	isValidPassword := shared.CheckPasswordHash(mod.Password, u.Password)

	if !isValidPassword {
		return nil, fmt.Errorf("[Login failed]: Password is not correct")
	}

	// fmt.Println(isValidPassword)

	u.Mask(shared.DbTypeUser)

	authId := u.FakeId.String()

	fmt.Println("Generated authId:", authId)
	// fmt.Println(authId)

	// Generate access token
	// accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"authId": authId,
	// 	"exp":    time.Now().Add(15 * time.Minute).Unix(),
	// })

	// accessSecret := os.Getenv("ACCESS_SECRET")
	// accessTokenString, err := accessToken.SignedString([]byte(accessSecret))

	// if err != nil {
	// 	return nil, fmt.Errorf("Login failed, something wrong due to processing access token to String type")
	// }

	// // Generate refresh token
	// refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"authId": authId,
	// 	"exp":    time.Now().Add(7 * 24 * time.Hour).Unix(),
	// })
	// refreshSecret := os.Getenv("REFRESH_SECRET")
	// refreshTokenString, err := refreshToken.SignedString([]byte(refreshSecret))
	accessTokenString, refreshTokenString, err := shared.TokenProvider(authId)

	if err != nil {
		return nil, fmt.Errorf("Login failed, something wrong due to processing refresh token to String type")
	}

	// fmt.Println("Hahahahahaha")
	// fmt.Println(accessTokenString, refreshTokenString)

	// Save to db
	err = aui.TokenRepo.CreateTokens(authId, accessTokenString, refreshTokenString, time.Now().Add(7*24*time.Hour))

	if err != nil {
		return nil, fmt.Errorf("Login failed, something wrong due to saving token to database")
	}

	return &models.AuthResponse{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}
func (aui *AuthUsecaseImpl) SignUp(mod *models.SignUpRequest) (*models.AuthResponse, error) {
	insertData := &entities.InsertUserDto{
		FirstName: mod.FirstName,
		LastName:  mod.LastName,
		Email:     mod.Email,
		Password:  mod.Password,
		Role:      "user",
		IsActive:  true,
		IsAdmin:   false,
		IsBlocked: false,
	}

	authId, err := aui.UserRepo.Insert(insertData)
	if err != nil {
		return nil, err
	}

	accessTokenString, refreshTokenString, err := shared.TokenProvider(authId)

	if err != nil {
		return nil, fmt.Errorf("Login failed, something wrong due to processing refresh token to String type")
	}
	// fmt.Println("Hahahahahaha")
	// fmt.Println(accessTokenString, refreshTokenString)

	// Save to db
	err = aui.TokenRepo.CreateTokens(authId, accessTokenString, refreshTokenString, time.Now().Add(7*24*time.Hour))

	if err != nil {
		return nil, fmt.Errorf("Login failed, something wrong due to saving token to database")
	}

	return &models.AuthResponse{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}
func (aui *AuthUsecaseImpl) Profile() (*entities.FetchUserDto, error) {
	return nil, nil
}
func (aui *AuthUsecaseImpl) SignOut() error {
	return nil
}

func NewAuthUsecaseImpl(userRepo repositories.UserRepository, tokenRepo TokenRepo.TokenRepository) AuthUsecase {
	return &AuthUsecaseImpl{
		UserRepo:  userRepo,
		TokenRepo: tokenRepo,
	}
}
