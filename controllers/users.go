package controllers

import (
	"blogmore/models"
	"blogmore/services"
	"blogmore/utils"
	"blogmore/views"
	"fmt"
	"net/http"
)

// NewUsers represting new users view
func NewUsers(us services.UserService) *Users {
	return &Users{
		NewView:   views.New("users/new"),
		LoginView: views.New("users/login"),
		us:        us,
	}
}

// Users struct to render view and it's related methods
type Users struct {
	NewView   *views.View
	LoginView *views.View
	us        services.UserService
}

// New render
// GET /signup view
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, "index", nil)
}

// SignUpForm type to represent blogmore user register
type SignUpForm struct {
	Username string `schema:"username"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// SignUp signs up new user
// POST /signup API
func (u *Users) SignUp(w http.ResponseWriter, r *http.Request) {
	var vd views.Data
	form := new(SignUpForm)
	if err := utils.ParseForm(form, r); err != nil {
		vd.Alert = &views.Alert{
			Level:   views.AlertLvlError,
			Message: views.AlertMsgGeneric,
		}
		fmt.Println(err)
		u.NewView.Render(w, "index", vd)
		return
	}
	user := models.User{
		Username: form.Username,
		Email:    form.Email,
		Password: form.Password,
	}
	userRecord, err := u.us.Create(&user)
	if err != nil {
		vd.Alert = &views.Alert{
			Level:   views.AlertLvlError,
			Message: err.Error(),
		}
		fmt.Println(err)
		u.NewView.Render(w, "index", vd)
		return
	}
	err = u.signIn(w, userRecord)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	http.Redirect(w, r, "/profile", http.StatusFound)
}

// LoginForm type to represent blogmore user login
type LoginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// Login signs up new user
// POST /login API
func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	u.LoginView.Render(w, "index", nil)
}

// LoginAction signs up new user
// POST /login API
func (u *Users) LoginAction(w http.ResponseWriter, r *http.Request) {
	form := new(LoginForm)
	if err := utils.ParseForm(form, r); err != nil {
		panic(err)
	}
	userRecord, err := u.us.Login(form.Email, form.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = u.signIn(w, userRecord)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/profile", http.StatusFound)
	fmt.Fprint(w, userRecord)
}

func (u *Users) signIn(w http.ResponseWriter, user *models.User) error {
	if user.RememberToken == "" {
		token, err := utils.RememberToken()
		if err != nil {
			return err
		}
		user.RememberToken = token
		_, err = u.us.Update(user)
		if err != nil {
			return err
		}
	}
	cookie := http.Cookie{
		Name:     "remember_token",
		Value:    user.RememberToken,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	return nil
}
