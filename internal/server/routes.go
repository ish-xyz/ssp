package server

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ish-xyz/ssp/internal/logger"
)

var metadataFile = "metadata.yaml"

func listJobTemplates(w http.ResponseWriter, r *http.Request) {
	/*
		Get a list of jobTemplates
	*/
	var jobTemplatesList []JobTemplate
	files, _ := ioutil.ReadDir(appConfig.JobTemplatesPath)

	for _, f := range files {

		job, err := loadJob(f.Name() + "/" + metadataFile)
		if err != nil {
			logger.ErrorLogger.Printf("Can't load file %s. Skipping. Error => %v", f.Name(), err)
			continue
		}

		job.Name = f.Name()
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
	filename := fmt.Sprintf("%s/%s", params["name"], metadataFile)

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

func runJob(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	nameSuffix := params["name"]
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(r.FormValue("name"))
	fmt.Println(string(body))
	//{inputs, name}

	fmt.Println(nameSuffix)
	//get payload with inputs
	//serialize payload
	//generate job name
	//create a k8s job
	//set job for auto-deletion
	return
}
