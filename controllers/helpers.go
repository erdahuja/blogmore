package controllers

import (
	"net/http"

	schema "github.com/gorilla/Schema"
)

func parseForm(dest interface{}, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return err
	}
	decoder := schema.NewDecoder()
	if err := decoder.Decode(dest, r.PostForm); err != nil {
		return err
	}
	return nil
}
