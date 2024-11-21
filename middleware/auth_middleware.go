package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/yuta_2710/go-clean-arc-reviews/modules/users/entities"
	"github.com/yuta_2710/go-clean-arc-reviews/modules/users/repositories"
	"github.com/yuta_2710/go-clean-arc-reviews/shared"
)

// Protect middleware for authorization
func Protect(getUserById func(authId string) (*entities.FetchUserDto, error)) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")

			if authHeader == "" {
				return shared.Response(c, false, http.StatusUnauthorized, "Authorization token is missing", nil, nil)
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return shared.Response(c, false, http.StatusUnauthorized, "Invalid authorization format", nil, nil)
			}

			acToken := parts[1]
			acSecret := os.Getenv("ACCESS_SECRET")

			// Parse and validate token
			token, err := jwt.Parse(acToken, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
				}
				return []byte(acSecret), nil
			})

			if err != nil {
				return shared.Response(c, false, http.StatusUnauthorized, "Unauthorized", nil, nil)
			}

			// Extract claims
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				return shared.Response(c, false, http.StatusUnauthorized, "Invalid token claims", nil, nil)
			}

			authId, ok := claims["authId"].(string)
			if !ok {
				return shared.Response(c, false, http.StatusUnauthorized, "User ID not found in token", nil, nil)
			}

			// Use the provided function to fetch the user
			user, err := getUserById(authId)
			if err != nil {
				return shared.Response(c, false, http.StatusUnauthorized, "User not found", nil, nil)
			}

			// Attach user info to context
			c.Set("user", user)

			ctx := context.WithValue(c.Request().Context(), "authId", authId)
			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}

func NewProtectMiddleware(userRepo repositories.UserRepository) echo.MiddlewareFunc {
	return Protect(func(authId string) (*entities.FetchUserDto, error) {
		// Decode the authId
		decodedUID, err := shared.DecomposeUidV2(authId)
		if err != nil {
			return nil, fmt.Errorf("invalid authId: %v", err)
		}

		idStr := fmt.Sprintf("%d", decodedUID.GetLocalID())

		// Find user by ID
		user, err := userRepo.FindById(idStr)
		if err != nil {
			return nil, err
		}

		// Map user to FetchUserDto
		return &entities.FetchUserDto{
			FakeId:    authId,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		}, nil
	})
}
