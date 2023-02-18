package main

import (
	"flag"
	"log"
	"net/url"

	"github.com/Kaiose31/JenGO/pipeline"
)

func main() {
	Host := flag.String("host", "localhost:8080", "URL for jenkins")
	Token := flag.String("token", "", "API Token for jenkins")
	Action := flag.String("action", "create", "Pipeline Management actions \n1. create\n2. run")
	Config := flag.String("config-path", "", "Config json path for Pipeline Creation")

	flag.Parse()
	host, err := url.Parse(*Host)
	if err != nil {
		log.Fatal("Error Parsing URL")
	}

	pipelineManager := pipeline.JenkinsConfig{
		HostUrl: *host,
		Token:   *Token,
	}

	switch *Action {
	case "create":
		{
			pipelineManager.CreatePipeline(*Config)

		}
	case "run":
		{
			// TODO
		}

	case "info":
		{
			// TODO
		}
	default:
		log.Fatal("Invalid Action for Pipeline:", *Config)

	}

}
