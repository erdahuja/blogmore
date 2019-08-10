package controllers

import (
	"encoding/json"
	"fmt"
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
