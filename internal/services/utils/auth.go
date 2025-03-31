package utils

import "crypto/rand"

func GenerateNewSecret() ([]byte, error) {
	secret := make([]byte, 64)

	_, err := rand.Read(secret)
	if err != nil {
		return nil, err
	}

	return secret, nil
}
