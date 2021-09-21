package server

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ish-xyz/ssp/internal/logger"
	jsoniter "github.com/json-iterator/go"
	"gopkg.in/yaml.v2"
)

func jsonResponse(w http.ResponseWriter, r *http.Request, statusCode int, payload Response) {
	/*
		Standard JSON response for APIs
	*/
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		logger.DebugLogger.Printf("Can't encode payload struct to json")
		w.WriteHeader(500)
		w.Write([]byte("{\"status\": \"failed\", \"data\": \"request failed\"}"))
	}
	w.Write(jsonBytes)
}

func loadJob(filename string) (JobTemplate, error) {
	/*
		Load and return a JobTemplate object
	*/
	var job JobTemplate
	yamlFile, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", appConfig.JobTemplatesPath, filename))
	if err != nil {
		return JobTemplate{}, err
	}
	err = yaml.Unmarshal(yamlFile, &job)
	if err != nil {
		return JobTemplate{}, err
	}
	return job, nil
}
