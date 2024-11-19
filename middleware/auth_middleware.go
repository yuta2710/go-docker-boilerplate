package custom_middleware

import (
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

func Protect(userRepo repositories.UserRepository) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")

			if authHeader == "" {
				return shared.Response(c, false, http.StatusBadRequest, "[No authorization]: authorization token is missing", nil, nil)
			}

			parts := strings.Split(authHeader, " ")
			fmt.Println(parts)

			if len(parts) != 2 || parts[0] != "Bearer" {
				return shared.Response(c, false, http.StatusBadRequest, "[No authorization]: Invalid authorization format", nil, nil)
			}

			acToken := parts[1]
			acSecret := os.Getenv("ACCESS_SECRET")

			token, err := jwt.Parse(acToken, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
				}

				return []byte(acSecret), nil
			})

			if err != nil {
				return shared.Response(c, false, http.StatusBadRequest, "[No authorization]: Unauthorized", nil, nil)

			}

			// Extract claims and retrieve the users ID
			claims, ok := token.Claims.(jwt.MapClaims)

			if !ok && !token.Valid {
				return shared.Response(c, false, http.StatusBadRequest, "[No authorization]: Unable to parse claims and invalid token", nil, nil)
			}

			fmt.Println(claims)

			authId, ok := claims["authId"].(string)

			if !ok {
				return shared.Response(c, false, http.StatusBadRequest, "[No authorization]: User ID not found in the token", nil, nil)
			}

			decodedUID, err := shared.DecomposeUidV2(authId)
			if err != nil {
				fmt.Println("DecomposeUID error:", err)
			} else {
				fmt.Printf("Decoded UID: LocalID=%d, ObjectType=%d, SharedID=%d\n",
					decodedUID.GetLocalID(), decodedUID.GetObjectType(), decodedUID.GetShardID())
			}

			if err != nil {

			}

			fmt.Println(decodedUID.GetLocalID())

			idStr := fmt.Sprintf("%d", decodedUID.GetLocalID())

			// Check user exist or not
			user, err := userRepo.FindById(idStr)

			if err != nil {
				return shared.Response(c, false, http.StatusBadRequest, "[No authorization]: User ID not found", nil, nil)
			}

			profileDto := &entities.FetchUserDto{
				FakeId:    authId,
				FirstName: user.FirstName,
				LastName:  user.LastName,
				Email:     user.Email,
			}

			c.Set("user", profileDto)

			fmt.Println(c.Get("user"))

			return next(c)
		}
	}
}
