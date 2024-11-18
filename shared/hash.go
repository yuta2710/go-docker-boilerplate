package shared

import "golang.org/x/crypto/bcrypt"

func HashPassword(pw string) (string, error) {
	hashedPw, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPw), nil
}
