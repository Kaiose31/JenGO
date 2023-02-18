package pipeline

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/url"
	"path"
	"time"

	"github.com/bndr/gojenkins"
)

// TODO! Update with Config Format
type ConfigFormat struct {
	Name string `json:"name"`
}

type JenkinsConfig struct {
	HostUrl url.URL
	Token   string
}

var (
	ctx = context.Background()
)

func (j *JenkinsConfig) CreatePipeline(configPath string) bool {

	configPath = path.Clean(configPath)

	conf, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	var config ConfigFormat

	if err := json.Unmarshal(conf, &config); err != nil {
		log.Fatal(err)
	}

	jenkins := gojenkins.CreateJenkins(nil, j.HostUrl.String(), j.Token)
	_, err = jenkins.Info(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//TODO! Create The Pipeline
	node, err := jenkins.CreateNode(ctx, config.Name, 1, "Description", "/var/lib/jenkins", "jdk8 docker", map[string]string{"method": "JNLPLauncher"})
	if err != nil {
		log.Fatal(err)
	}

	for true {
		time.Sleep(1 * time.Second)
		res, err := node.IsOnline(ctx)
		if err != nil {
			log.Fatal(err)
		}

		if res {
			return res
		}
	}
	return false
}
