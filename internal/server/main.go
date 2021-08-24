package server

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/ish-xyz/ssp/internal/logger"
)

func Start(port string) {

	// Create a mux for routing incoming requests
	r := mux.NewRouter()
	_ = initRoutes(r)
	http.Handle("/", r)

	logger.InfoLogger.Printf("Web Server started on :%s ...\n", port)
	logger.ErrorLogger.Fatal(http.ListenAndServe(":"+port, nil))
}

func initRoutes(r *mux.Router) error {
	s := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	r.PathPrefix("/static/").Handler(s)
	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/create-job", createJobGet).Methods("GET")
	return nil
}

func index(w http.ResponseWriter, r *http.Request) {
	// check if user is authenticated
	//   if not, redirect to login
	//   if yes, render the index page
	login(w, r)
}

func login(w http.ResponseWriter, r *http.Request) {
	// Login page
	t, _ := template.ParseFiles("templates/login.html")

	t.Execute(w, map[string]string{})
}

func createJobGet(w http.ResponseWriter, r *http.Request) {
	return
}
