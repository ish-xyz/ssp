package ui

import (
	"net/http"
)

func listPage(w http.ResponseWriter, r *http.Request) {

	req, err := apiCall(w, r, "v1/list-job-templates")
	if err != nil {
		return
	}

	err = templates.ExecuteTemplate(w, "listPage", req.Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func formPage(w http.ResponseWriter, r *http.Request) {

	/*req, err := apiCall(w, r, "v1/list-job-templates")
	if err != nil {
		return
	}
	*/

	//var inputs []map[string]interface{}

	inputs := []map[string]interface{}{
		map[string]interface{}{
			"key": "options",
			"value": []string{
				"value1",
				"value2",
				"value3",
			},
			"description": "my job options",
		},
		map[string]interface{}{
			"key":         "command",
			"value":       "my-custom-command",
			"description": "my-custom-command description",
		},
		map[string]interface{}{
			"key":         "dry-run",
			"value":       false,
			"description": "Dry-run option",
		},
	}

	err := templates.ExecuteTemplate(w, "formPage", inputs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
