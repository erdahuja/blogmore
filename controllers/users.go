package controllers

import (
	"blogmore/views"
	"fmt"
	"net/http"
)

// NewUsers represting new users view
func NewUsers() *Users {
	return &Users{
		NewView:   views.New("views/users/new.gohtml"),
		LoginView: views.New("views/users/new.gohtml"),
	}
}

// Users struct to render view and it's related methods
type Users struct {
	NewView   *views.View
	LoginView *views.View
}

// New render
// GET /signup view
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, "index", nil)
}

// SignUpForm type to represent blogmore user register
type SignUpForm struct {
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// Create signs up new user
// POST /signup API
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	form := new(SignUpForm)
	if err := parseForm(form, r); err != nil {
		panic(err)
	}
	fmt.Fprint(w, form)
}

// LoginForm type to represent blogmore user login
type LoginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// Login signs up new user
// POST /login API
func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	form := new(LoginForm)
	if err := parseForm(form, r); err != nil {
		panic(err)
	}
	fmt.Fprint(w, form)
}
