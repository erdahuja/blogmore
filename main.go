package main

import (
	"fmt"
	"net/http"

	"realworld-starter-kit/blogmore/views"

	"github.com/gorilla/mux"
)

var (
	homeView    *views.View
	profileView *views.View
	loginView   *views.View
	signUpView  *views.View
)

func homeFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, "index", nil))
}

func profileFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(profileView.Render(w, "index", nil))
}

func signUpFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(signUpView.Render(w, "index", nil))
}

func loginFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(loginView.Render(w, "index", nil))
}

func pageNotFoundFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Page not found")
}

func init() {
	homeView = views.New("./views/home.gohtml")
	profileView = views.New("./views/profile.gohtml")
	signUpView = views.New("./views/signUp.gohtml")
	loginView = views.New("./views/login.gohtml")
}

func main() {
	if homeView.Err != nil {
		panic(homeView.Err)
	}
	if profileView.Err != nil {
		panic(profileView.Err)
	}
	router := mux.NewRouter()
	router.HandleFunc("/", homeFunc)
	router.HandleFunc("/profile", profileFunc)
	router.HandleFunc("/signup", signUpFunc)
	router.HandleFunc("/login", loginFunc)
	router.NotFoundHandler = http.HandlerFunc(pageNotFoundFunc)
	http.ListenAndServe(":3000", router)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
