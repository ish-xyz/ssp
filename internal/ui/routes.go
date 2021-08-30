package ui

import (
	"net/http"

	"github.com/gorilla/mux"
)

func listPage(w http.ResponseWriter, r *http.Request) {
	/*
		List Jobs page
	*/
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

func lanchPage(w http.ResponseWriter, r *http.Request) {
	/*
		Launch Job page
	*/

	params := mux.Vars(r)
	req, err := apiCall(w, r, "/v1/get-job-template/"+params["name"])
	if err != nil {
		return
	}

	err = templates.ExecuteTemplate(w, "formPage", req.Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
