/*
Copyright 2014 Rohith All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package marathon

// Info is the detailed stats returned from marathon info
type Info struct {
	EventSubscriber struct {
		HttpEndpoints []string `json:"http_endpoints"`
		Type          string   `json:"type"`
	} `json:"event_subscriber"`
	FrameworkId string `json:"frameworkId"`
	HttpConfig  struct {
		AssetsPath interface{} `json:"assets_path"`
		HttpPort   float64     `json:"http_port"`
		HTTPSPort  float64     `json:"https_port"`
	} `json:"http_config"`
	Leader         string `json:"leader"`
	MarathonConfig struct {
		Checkpoint                 bool    `json:"checkpoint"`
		Executor                   string  `json:"executor"`
		FailoverTimeout            float64 `json:"failover_timeout"`
		Ha                         bool    `json:"ha"`
		Hostname                   string  `json:"hostname"`
		LocalPortMax               float64 `json:"local_port_max"`
		LocalPortMin               float64 `json:"local_port_min"`
		Master                     string  `json:"master"`
		MesosRole                  string  `json:"mesos_role"`
		MesosUser                  string  `json:"mesos_user"`
		ReconciliationInitialDelay float64 `json:"reconciliation_initial_delay"`
		ReconciliationInterval     float64 `json:"reconciliation_interval"`
		TaskLaunchTimeout          float64 `json:"task_launch_timeout"`
	} `json:"marathon_config"`
	Name            string `json:"name"`
	Version         string `json:"version"`
	ZookeeperConfig struct {
		Zk              string `json:"zk"`
		ZkFutureTimeout struct {
			Duration float64 `json:"duration"`
		} `json:"zk_future_timeout"`
		ZkHosts   string  `json:"zk_hosts"`
		ZkPath    string  `json:"zk_path"`
		ZkState   string  `json:"zk_state"`
		ZkTimeout float64 `json:"zk_timeout"`
	} `json:"zookeeper_config"`
}

// Info retrieves the info stats from marathon
func (r *marathonClient) Info() (*Info, error) {
	info := new(Info)
	if err := r.apiGet(MARATHON_API_INFO, nil, info); err != nil {
		return nil, err
	}

	return info, nil
}

// Leader retrieves the current marathon leader node
func (r *marathonClient) Leader() (string, error) {
	var leader struct {
		Leader string `json:"leader"`
	}
	if err := r.apiGet(MARATHON_API_LEADER, nil, &leader); err != nil {
		return "", err
	}

	return leader.Leader, nil
}

// AbdicateLeader abdicates the marathon leadership
func (r *marathonClient) AbdicateLeader() (string, error) {
	var message struct {
		Message string `json:"message"`
	}

	if err := r.apiDelete(MARATHON_API_LEADER, nil, &message); err != nil {
		return "", err
	}

	return message.Message, nil
}
