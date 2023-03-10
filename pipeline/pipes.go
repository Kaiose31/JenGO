package pipeline

import (
	"context"
	"log"
	"net/url"

	"github.com/bndr/gojenkins"
)

// TODO! Update with Config Format
type ConfigFormat struct {
	Name    string `json:"name"`
	JobName string `json:"jobname"`
}

type JenkinsConfig struct {
	HostUrl  url.URL
	UserName string
	Password string
}

var (
	ctx = context.Background()
)

// TODO! Example to Config Driven
func (j *JenkinsConfig) CreatePipeline(configPath string) *gojenkins.Job {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	configString := `<?xml version='1.0' encoding='UTF-8'?>
				<project>
				<actions/>
				<description></description>
				<keepDependencies>false</keepDependencies>
				<properties/>
				<scm class="hudson.scm.NullSCM"/>
				<canRoam>true</canRoam>
				<disabled>false</disabled>
				<blockBuildWhenDownstreamBuilding>false</blockBuildWhenDownstreamBuilding>
				<blockBuildWhenUpstreamBuilding>false</blockBuildWhenUpstreamBuilding>
				<triggers class="vector"/>
				<concurrentBuild>false</concurrentBuild>
				<builders/>
				<publishers/>
				<buildWrappers/>
				</project>`

	config := processConfig(configPath)
	jenkins := setupJenkins(ctx, j)

	pFolder, err := jenkins.CreateFolder(ctx, config.Name)

	if err != nil {
		log.Fatal(err)
	}

	job, err := jenkins.CreateJobInFolder(ctx, configString, config.JobName, pFolder.GetName())
	if err != nil {
		log.Fatal(err)
	}

	return job
}

func (j *JenkinsConfig) RunPipeline(configPath string) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	config := processConfig(configPath)
	jenkins := setupJenkins(ctx, j)

	job, err := jenkins.GetJob(ctx, config.JobName, config.Name)
	if err != nil {
		log.Fatal(err)
	}

	_, err = job.InvokeSimple(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	checkJobStatus(ctx, j, job)

}
