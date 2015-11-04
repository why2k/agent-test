package main

import "gopkg.in/libgit2/git2go.v22"

import "os"

import "github.build.ge.com/Dockerci/agent/agent"

func main() {

	//	git.Clone("https://github.com/shannonrdunn/github-maven-example", "/tmp/workspace", &git.CloneOptions{})
	git.Clone(os.Getenv("GIT_URL"), "/tmp/workspace", &git.CloneOptions{})
	agent.Testme()

}
