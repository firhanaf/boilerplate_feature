package helpers

import gonanoid "github.com/matoous/go-nanoid/v2"

func UsernameGenerator() (string, error) {
	str := "abxdefghijklmopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	uuid, err := gonanoid.Generate(str, 12)
	if err != nil {
		return "", err
	}
	return uuid, nil
}
