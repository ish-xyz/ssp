package main

import (
	"fmt"

	"github.com/ish-xyz/ssp/internal/jobs"
	"github.com/ish-xyz/ssp/internal/k8s"
	"github.com/ish-xyz/ssp/internal/server"
)

func main() {

	client, _ := k8s.NewClient("/Users/ishamaraia/.kube/config")

	_ = jobs.Job{
		Client:    client,
		Name:      "test",
		Namespace: "default",
		Command:   []string{"ls", "-al"},
		Image:     "debian:latest",
	}

	//job.Create()

	server.Start("7000")

	fmt.Println("here")
}
