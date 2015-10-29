package main

import "gopkg.in/libgit2/git2go.v22"

import "github.build.ge.com/DockerCI/agent/agent"

func main() {
	git.Clone("https://github.com/shannonrdunn/github-maven-example", "/tmp/workspace", &git.CloneOptions{})
	j := agent.New("/Users/212403630/Documents/github-maven-example")
	j.Create()
	j.Start()
}
