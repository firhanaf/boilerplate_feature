package helpers

import gonanoid "github.com/matoous/go-nanoid/v2"

func NanoIDGenerator() (string, error) {
	str := "abxdefghijklmopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	uuid, err := gonanoid.Generate(str, 32)
	if err != nil {
		return "", err
	}
	return uuid, nil
}
