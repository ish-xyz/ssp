package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/ish-xyz/ssp/internal/logger"
	"gopkg.in/yaml.v2"
)

func listJobTemplates(w http.ResponseWriter, r *http.Request) {

	// return a list of jobTemplates
	var jobTemplatesList []JobTemplateData
	files, _ := ioutil.ReadDir(appConfig.JobTemplatesPath)

	for _, f := range files {

		var it JobTemplateData
		yamlFile, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", appConfig.JobTemplatesPath, f.Name()))
		if err != nil {
			logger.ErrorLogger.Printf("Can't read file %s. Skipping. Error => %v ", f.Name(), err)
			continue
		}
		err = yaml.Unmarshal(yamlFile, &it)
		if err != nil {
			logger.ErrorLogger.Printf("Can't unmarshal file %s. Skipping. Error => %v ", f.Name(), err)
			continue
		}

		it.Name = strings.Split(f.Name(), ".yaml")[0]
		jobTemplatesList = append(jobTemplatesList, it)
	}

	resp := Response{
		Status: "ok",
		Data:   jobTemplatesList,
	}

	jsonResponse(w, r, resp)
	return
}
