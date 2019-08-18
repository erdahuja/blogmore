package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

var rememberTokenBytes = 32

// CreateRandomByteSlice will generate random bytes using crypto/rand package
func CreateRandomByteSlice(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	fmt.Println(b)
	return b, nil
}

// CreateRandomString converts byte slice to a base64 encoded string
// we don't use normal string conversion because
// random bytes may not be valid UTF8
func CreateRandomString(n int) (string, error) {
	b, err := CreateRandomByteSlice(n)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// RememberToken return random generated string of predef size
func RememberToken() (string, error) {
	return CreateRandomString(rememberTokenBytes)
}
