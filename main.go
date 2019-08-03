package main

import (
	"fmt"
	"net/http"

	"blogmore/controllers"
	"blogmore/views"

	"github.com/gorilla/mux"
)

var (
	homeView    *views.View
	profileView *views.View
)

func homeFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, "index", nil))
}

func profileFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(profileView.Render(w, "index", nil))
}

func pageNotFoundFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Page not found")
}

func init() {
	homeView = views.New("home")
	profileView = views.New("users/profile")
}

func main() {
	router := mux.NewRouter()
	usersC := controllers.NewUsers()
	router.HandleFunc("/", homeFunc)
	router.HandleFunc("/profile", profileFunc)
	router.HandleFunc("/signup", usersC.New).Methods("GET")
	router.HandleFunc("/signup", usersC.Create).Methods("POST")
	router.HandleFunc("/login", usersC.Login).Methods("GET")
	router.HandleFunc("/login", usersC.LoginAction).Methods("POST")
	router.NotFoundHandler = http.HandlerFunc(pageNotFoundFunc)
	http.ListenAndServe(":3000", router)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
