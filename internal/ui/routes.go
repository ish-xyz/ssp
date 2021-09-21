package ui

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func listPage(w http.ResponseWriter, r *http.Request) {
	/*
		List Jobs page
	*/
	req, err := backendGet(w, r, "v1/list-job-templates")
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
	req, err := backendGet(w, r, "v1/get-job-template/"+params["name"])
	if err != nil {
		return
	}

	err = templates.ExecuteTemplate(w, "formPage", req.Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func runResultPage(w http.ResponseWriter, r *http.Request) {
	/*
		Job Creation result page
	*/
	//read form inputs
	params := mux.Vars(r)
	r.ParseForm()

	r.Form["_name"] = []string{params["_name"]}
	fmt.Println(r.Form)
	// send r.Form as payload to the backend api
	// read response and render right template
	w.Write([]byte("Job created successfully"))
}
