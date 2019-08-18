package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"hash"
	"net/http"

	schema "github.com/gorilla/Schema"
	"golang.org/x/crypto/bcrypt"
)

// ParseForm parses request body to given target
func ParseForm(dest interface{}, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return err
	}
	decoder := schema.NewDecoder()
	if err := decoder.Decode(dest, r.PostForm); err != nil {
		return err
	}
	return nil
}

// PrettyPrint variable (struct, map, array, slice) in Golang
func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("[PrettyPrint] Error while processing variable.")
		return err
	}
	fmt.Println(string(b))
	return nil
}

// GenerateHash generates bcrypt hash with default cost
func GenerateHash(byteSlice []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(byteSlice, bcrypt.DefaultCost)
}

// CompareHashAndPassword returns bool error if password and hash doesn't match
func CompareHashAndPassword(token []byte, pwd []byte) error {
	if err := bcrypt.CompareHashAndPassword(token, pwd); err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			return errors.New("invalid password for user")
		default:
			return err
		}
	}
	return nil
}

// HMAC is a wrapper around crypt/hmac pkg
type HMAC struct {
	hmac hash.Hash
}

// NewHMAC creates and returns new hmac object
func NewHMAC(key string) HMAC {
	h := hmac.New(sha256.New, []byte(key))
	return HMAC{
		hmac: h,
	}
}

// Hash generates hash for given input with secret key of hmac object
func (h HMAC) Hash(input string) string {
	h.hmac.Reset()
	h.hmac.Write([]byte(input))
	b := h.hmac.Sum(nil)
	return base64.URLEncoding.EncodeToString(b)
}
