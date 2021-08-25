package server

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/ish-xyz/ssp/internal/config"
	"github.com/ish-xyz/ssp/internal/jobs"
	"github.com/ish-xyz/ssp/internal/k8s"
	"github.com/ish-xyz/ssp/internal/logger"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Start(c config.Config) {

	// Create a mux for routing incoming requests
	r := mux.NewRouter()
	_ = initRoutes(r)
	_ = initK8sClient()
	http.Handle("/", r)

	logger.InfoLogger.Printf("Web Server started on :%d ...\n", c.ServerPort)
	logger.ErrorLogger.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", c.ServerAddr, c.ServerPort), nil))
}

func initRoutes(r *mux.Router) error {
	s := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	r.PathPrefix("/static/").Handler(s)
	r.HandleFunc("/", listJobs).Methods("GET")
	r.HandleFunc("/create-job", createJobGet).Methods("GET")
	return nil
}

func initK8sClient() error {
	// Init k8s client
	// TODO: remove hard coded config path
	jobs.Client, _ = k8s.NewClient("/Users/ishamaraia/.kube/config")
	return nil
}
