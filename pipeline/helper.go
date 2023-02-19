package pipeline

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path"

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
	fmt.Println(j)
	jenkins, err := gojenkins.CreateJenkins(nil, j.HostUrl.String(), j.UserName, j.Password).Init(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return jenkins

}
