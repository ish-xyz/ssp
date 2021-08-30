package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/ish-xyz/ssp/internal/logger"
	"gopkg.in/yaml.v2"
)

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

func listJobTemplates(w http.ResponseWriter, r *http.Request) {
	/*
		Get a list of jobTemplates
	*/
	var jobTemplatesList []JobTemplate
	files, _ := ioutil.ReadDir(appConfig.JobTemplatesPath)

	for _, f := range files {

		job, err := loadJob(f.Name())
		if err != nil {
			logger.ErrorLogger.Printf("Can't load file %s. Skipping. Error => %v", f.Name(), err)
			continue
		}

		job.Name = strings.Split(f.Name(), ".yaml")[0]
		jobTemplatesList = append(jobTemplatesList, job)
	}

	resp := Response{
		Status: "ok",
		Data:   jobTemplatesList,
	}
	jsonResponse(w, r, 200, resp)
	return
}

func getJobTemplate(w http.ResponseWriter, r *http.Request) {
	/*
		Get a single JobTemplate
	*/
	params := mux.Vars(r)
	filename := fmt.Sprintf("%s.yaml", params["name"])

	job, err := loadJob(filename)
	if err != nil {
		logger.ErrorLogger.Printf("Can't load file %s. Skipping. Error => %v", filename, err)
		jsonResponse(w, r, 500, Response{
			Status: "failed",
			Data:   "request failed",
		})
		return
	}
	job.Name = params["name"]
	resp := Response{
		Status: "ok",
		Data:   job,
	}
	jsonResponse(w, r, 200, resp)
	return
}
