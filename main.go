package main

import (
	"github.com/ish-xyz/ssp/internal/config"
	"github.com/ish-xyz/ssp/internal/server"
)

func main() {

	c := config.Config{
		KubeConfigPath:   "/Users/ishamaraia/.kube/config",
		JobTemplatesPath: "/Users/ishamaraia/repos/ssp/debug/jobs",
		ServerAddr:       "localhost",
		ServerPort:       7000,
	}
	server.Start(c)
}
