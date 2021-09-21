package main

import (
	"github.com/ish-xyz/ssp/internal/config"
	"github.com/ish-xyz/ssp/internal/server"
	"github.com/ish-xyz/ssp/internal/ui"
)

func main() {

	c := config.Config{
		KubeConfigPath:   "/Users/ishamaraia/.kube/config",
		JobTemplatesPath: "/Users/ishamaraia/repos/ssp/debug/charts",
		BackendAddr:      "localhost:7000",
		FrontendAddr:     "localhost:8000",
	}
	go server.Run(c)
	ui.Run(c)
}
