package server

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

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

func renderValues(valueFile string, data map[string]interface{}) (string, error) {

	filePath := fmt.Sprintf("%s/%s", appConfig.JobTemplatesPath, valueFile)
	t := template.Must(template.New("values").ParseFiles(filePath))

	if err := t.ExecuteTemplate(os.Stdout, "values.yaml", data); err != nil {
		return "", err
	}

	return "", nil
}
