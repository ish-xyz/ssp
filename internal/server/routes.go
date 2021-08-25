package server

import (
	"net/http"
	"text/template"
)

func listJobs(w http.ResponseWriter, r *http.Request) {
	// ls jobs from c.JobTemplatesPath
	// load jobs metadata

	err := templates.ExecuteTemplate(w, "listPage", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	// Login page
	t, _ := template.ParseFiles("templates/login.html")

	t.Execute(w, map[string]string{})
}

func createJobGet(w http.ResponseWriter, r *http.Request) {
	return
}
