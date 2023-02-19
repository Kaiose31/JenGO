package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"

	"github.com/Kaiose31/JenGO/pipeline"
)

func main() {
	Host := flag.String("host", "localhost:8080", "URL for jenkins")
	Action := flag.String("action", "create", "Pipeline Management actions \n1. create\n2. run")
	Username := flag.String("username", "username", "username")
	Password := flag.String("password", "password", "password")
	ConfigPath := flag.String("config-path", "", "Config json path for Pipeline Creation")

	flag.Parse()
	host, err := url.Parse(*Host)
	if err != nil {
		log.Fatal("Error Parsing URL")
	}

	pipelineManager := pipeline.JenkinsConfig{
		HostUrl:  *host,
		UserName: *Username,
		Password: *Password,
	}

	switch *Action {
	case "create":
		{
			job := pipelineManager.CreatePipeline(*ConfigPath)
			fmt.Println("Created Job :", job.Raw.Name)
		}
	case "run":
		{
			pipelineManager.RunPipeline(*ConfigPath)

		}

	case "info":
		{
			// TODO
		}
	default:
		log.Fatal("Invalid Action for Pipeline:", *ConfigPath)

	}

}
