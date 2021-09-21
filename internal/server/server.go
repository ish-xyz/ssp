package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ish-xyz/ssp/internal/config"
	"github.com/ish-xyz/ssp/internal/jobs"
	"github.com/ish-xyz/ssp/internal/k8s"
	"github.com/ish-xyz/ssp/internal/logger"
)

var appConfig config.Config

func Run(c config.Config) {

	appConfig = c

	// Create a mux for routing incoming requests
	r := mux.NewRouter().StrictSlash(true)
	_ = initRoutes(r)
	_ = initK8sClient()

	srv := &http.Server{
		Handler:      r,
		Addr:         appConfig.BackendAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.InfoLogger.Printf("API Server started on %s ...\n", appConfig.BackendAddr)
	logger.ErrorLogger.Fatal(srv.ListenAndServe())
}

func initK8sClient() error {
	// Init k8s client
	// TODO: remove hard coded config path
	jobs.Client, _ = k8s.NewClient(appConfig.KubeConfigPath)
	return nil
}

func initRoutes(r *mux.Router) error {
	/*
		Init api server routes
	*/

	r.HandleFunc("/v1/list-job-templates", listJobTemplates).Methods("GET")
	r.HandleFunc("/v1/get-job-template/{name}", getJobTemplate).Methods("GET")
	r.HandleFunc("/v1/run-job/{name}", runJob).Methods("POST")
	return nil
}
