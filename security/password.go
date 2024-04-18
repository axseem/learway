// Package security implements user data privacy features.
package security

import (
	"errors"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 10)
}

func CompareHashAndPassword(password string, hash []byte) error {
	return bcrypt.CompareHashAndPassword(hash, []byte(password))
}

func ValidatePassword(password string) error {
	if len(password) < 8 || len(password) > 128 {
		return errors.New("password length must be between 8 and 128 characters")
	}

	var hasNumber, hasUpper, hasLower bool
	for _, r := range password {
		hasNumber = hasNumber || unicode.IsNumber(r)
		hasUpper = hasUpper || unicode.IsUpper(r)
		hasLower = hasLower || unicode.IsLower(r)

		// ASCII characters from 33 to 126: !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijklmnopqrstuvwxyz{|}~
		if r < 33 || r > 126 {
			return errors.New("password contains an invalid character: " + string(r))
		}
	}

	if !hasNumber || !hasUpper || !hasLower {
		return errors.New("password must contain at least one number, one uppercase letter, and one lowercase letter")
	}

	return nil
}
