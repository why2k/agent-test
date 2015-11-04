package agent

import (
	"os"

	"fmt"

	marathon "github.com/why2k/agent-test/agent/Godeps/_workspace/src/github.com/gambol99/go-marathon"
	"github.com/why2k/agent-test/agent/Godeps/_workspace/src/github.com/golang/glog"
)

func Testme() {
	marathon_url := "http://192.168.33.10:8080/"
	config := marathon.NewDefaultConfig()
	config.URL = marathon_url
	config.LogOutput = os.Stdout
	if client, err := marathon.NewClient(config); err != nil {
		glog.Fatalf("Failed to create a client for marathon, error: %s", err)
	} else {
		fmt.Printf("Success %s", client)

	}
}
