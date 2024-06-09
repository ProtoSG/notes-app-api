package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func EncrypPass(password string) (string, error) {
	hass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hass), nil
}

func ComparePass(password, hass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hass), []byte(password))
	return err == nil
}
