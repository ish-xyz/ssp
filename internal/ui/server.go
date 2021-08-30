package ui

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/gorilla/mux"
	"github.com/ish-xyz/ssp/internal/config"
	"github.com/ish-xyz/ssp/internal/jobs"
	"github.com/ish-xyz/ssp/internal/k8s"
	"github.com/ish-xyz/ssp/internal/logger"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))
var appConfig config.Config

func Run(c config.Config) {

	appConfig = c

	// Create a mux for routing incoming requests
	r := mux.NewRouter()
	_ = initRoutes(r)
	_ = initK8sClient()

	srv := &http.Server{
		Handler:      r,
		Addr:         appConfig.FrontendAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.InfoLogger.Printf("Web Server started on %s ...\n", appConfig.FrontendAddr)
	logger.ErrorLogger.Fatal(srv.ListenAndServe())
}

func initRoutes(r *mux.Router) error {
	s := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	r.PathPrefix("/static/").Handler(s)
	r.HandleFunc("/", listPage).Methods("GET")
	r.HandleFunc("/launch/{name}", lanchPage).Methods("GET")
	return nil
}

func initK8sClient() error {
	// Init k8s client
	// TODO: remove hard coded config path
	jobs.Client, _ = k8s.NewClient(appConfig.KubeConfigPath)
	return nil
}

func apiCall(w http.ResponseWriter, r *http.Request, endpoint string) (*Request, error) {
	var req *Request

	logger.DebugLogger.Printf("Frontend calling backend at %s", endpoint)

	reqURL := fmt.Sprintf("http://%s/%s", appConfig.BackendAddr, endpoint)
	resp, err := http.Get(reqURL)
	if err != nil {
		logger.ErrorLogger.Printf("Can't request backend %s. Error => %v ", endpoint, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return nil, err
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&req)

	return req, nil
}
