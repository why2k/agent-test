package agent

import "path/filepath"

import "io/ioutil"
import "fmt"

import "gopkg.in/yaml.v2"

type Job struct {
	Repo    string `yaml:"Repo"`
	Image   string `yaml:"Image"`
	Command string `yaml:"Command"`
	Status  string `yaml:"Status"`
}

func New(path string) {
	var configjob Job
	filename, _ := filepath.Abs(path + "/.build.yml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	configjob.Parse(yamlFile)

	fmt.Printf("%+v\n", configjob)
}

func (j *Job) Parse(data []byte) {
	yaml.Unmarshal(data, j)
}
