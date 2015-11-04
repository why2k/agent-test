package main

import "github.com/why2k/agent-test/Godeps/_workspace/src/gopkg.in/libgit2/git2go.v22"

import "os"

import "github.com/why2k/agent-test/agent"

func main() {

	//	git.Clone("https://github.com/shannonrdunn/github-maven-example", "/tmp/workspace", &git.CloneOptions{})
	git.Clone(os.Getenv("GIT_URL"), "./", &git.CloneOptions{})
	agent.Testme()

}
