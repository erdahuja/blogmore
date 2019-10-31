package main

import (
	"blogmore/services"
	"fmt"
	"net/http"

	"blogmore/controllers"
	"blogmore/views"

	"github.com/gorilla/mux"
)

var (
	homeView    *views.View
	profileView *views.View
	us          services.UserService
)

func homeFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	cookie, err := r.Cookie("remember_token")
	if err != nil {
		must(homeView.Render(w, "index", nil))
		return
	}
	user, err := us.ByRemember(cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	must(homeView.Render(w, "index", user))
}

func profileFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	cookie, err := r.Cookie("remember_token")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user, err := us.ByRemember(cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var vd views.Data
	vd.Yield = user
	must(profileView.Render(w, "index", vd))
}

func pageNotFoundFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Page not found")
}

func init() {
	homeView = views.New("home")
	profileView = views.New("users/profile")
	var err error
	us, err = services.NewUserService()
	must(err)
}

func main() {
	router := mux.NewRouter()
	usersC := controllers.NewUsers(us)
	us.DestructiveReset()
	us.AutoMigrate()
	router.HandleFunc("/", homeFunc)
	router.HandleFunc("/profile", profileFunc).Methods("GET")
	router.HandleFunc("/signup", usersC.New).Methods("GET")
	router.HandleFunc("/signup", usersC.SignUp).Methods("POST")
	router.HandleFunc("/login", usersC.Login).Methods("GET")
	router.HandleFunc("/login", usersC.LoginAction).Methods("POST")
	router.NotFoundHandler = http.HandlerFunc(pageNotFoundFunc)
	fmt.Println("Server listening on PORt :3000")
	http.ListenAndServe(":3000", router)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
