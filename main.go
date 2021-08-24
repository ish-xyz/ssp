package main

import (
	"fmt"

	"github.com/ish-xyz/ssp/internal/jobs"
	"github.com/ish-xyz/ssp/internal/k8s"
)

func main() {

	client, _ := k8s.NewClient("/Users/ishamaraia/.kube/config")

	job := jobs.Job{
		Client:    client,
		Name:      "test",
		Namespace: "default",
		Command:   []string{"ls", "-al"},
		Image:     "debian:latest",
	}

	job.Create()

	fmt.Println("here")
}
