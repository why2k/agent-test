package agent

import "path/filepath"

import "io/ioutil"

import "gopkg.in/yaml.v2"
import "github.com/fsouza/go-dockerclient"
import "fmt"

const Docker_url = "tcp://192.168.59.103:2375"

type Job struct {
	Repo      string `yaml:"Repo"`
	Image     string `yaml:"Image"`
	Command   string `yaml:"Command"`
	Status    string `yaml:"Status"`
	Container *docker.Container
}

func New(path string) Job {
	var configjob Job
	filename, _ := filepath.Abs(path + "/.build.yml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	configjob.Parse(yamlFile)

	return configjob
}

func (j *Job) Parse(data []byte) {
	yaml.Unmarshal(data, j)
}

func (j *Job) Create() {
	volumes := make(map[string]struct{})
	volumes["/Users/212403630/Documents/github-maven-example"] = struct{}{}
	client, _ := docker.NewClient(Docker_url)
	config := docker.Config{
		AttachStdout: true,
		AttachStdin:  true,
		Image:        j.Image,
		Volumes:      volumes,
		Cmd:          []string{j.Command},
		WorkingDir:   "/usr/src",
	}
	opts := docker.CreateContainerOptions{Name: "MAVENTEST", Config: &config}
	container, _ := client.CreateContainer(opts)
	j.Container = container
	fmt.Println(j)
}

func (j *Job) Start() {
	client, _ := docker.NewClient(Docker_url)
	binds := []string{}
	binds = append(binds, "/Users/212403630/Documents/github-maven-example"+":"+"/usr/src/mymaven")
	client.StartContainer(j.Container.ID, &docker.HostConfig{
		Binds: binds,
	})
}
