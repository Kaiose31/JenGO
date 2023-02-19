package pipeline

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"time"

	"github.com/bndr/gojenkins"
)

func processConfig(configPath string) ConfigFormat {

	configPath = path.Clean(configPath)

	conf, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	var config ConfigFormat

	if err := json.Unmarshal(conf, &config); err != nil {
		log.Fatal(err)
	}

	return config

}

func setupJenkins(ctx context.Context, j *JenkinsConfig) *gojenkins.Jenkins {
	jenkins, err := gojenkins.CreateJenkins(nil, j.HostUrl.String(), j.UserName, j.Password).Init(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return jenkins

}

func checkJobStatus(ctx context.Context, j *JenkinsConfig, job *gojenkins.Job) {
	for true {
		job.Poll(ctx)

		run, err := job.IsRunning(ctx)
		if err != nil {
			log.Fatal(err)
		}

		if run {
			fmt.Println("Job Running", job.Raw.Name)

		} else {
			build, err := job.GetLastCompletedBuild(ctx)
			if err != nil {
				log.Fatal(err)
			}

			if build.IsGood(ctx) {
				fmt.Println("Job Ran Successfully", job.Raw.Name)
				break
			}

		}
		time.Sleep(2 * time.Second)
	}
}
