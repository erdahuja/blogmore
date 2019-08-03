package controllers

import (
	"blogmore/views"
	"fmt"
	"net/http"
)

// NewUsers represting new users view
func NewUsers() *Users {
	return &Users{
		NewView: views.New("views/users/new.gohtml"),
	}
}

// Users struct to render view and it's related methods
type Users struct {
	NewView *views.View
}

// New render
// GET /signup view
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, "index", nil)
}

// Create signs up new user
// POST /signup API
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("new user crreated")
	u.NewView.Render(w, "index", "i am created")
}
