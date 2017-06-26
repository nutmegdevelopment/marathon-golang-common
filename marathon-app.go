package models

import (
	"encoding/json"
	"log"
)

// MarathonApp represents the bulk of the response from /v2/apps/<APP_ID>
// This is the JSON that is posted to the destination Marathon.
type MarathonApp struct {
	AcceptedResourceRoles []string   `yaml:"acceptedResourceRoles" json:"acceptedResourceRoles,omitempty"`
	Args                  []string   `yaml:"args" json:"args,omitempty"`
	BackoffFactor         float64    `yaml:"backoffFactor" json:"backoffFactor,omitempty"`
	BackoffSeconds        int        `yaml:"backoffSeconds" json:"backoffSeconds,omitempty"`
	Command               string     `yaml:"cmd" json:"cmd,omitempty"`
	Constraints           [][]string `yaml:"constraints" json:"constraints,omitempty"`
	Container             struct {
		ContainerType string `yaml:"type" json:"type"`
		Docker        struct {
			Image        string `yaml:"image" json:"image"`
			Network      string `yaml:"network" json:"network,omitempty"`
			PortMappings []struct {
				ContainerPort int    `yaml:"containerPort" json:"containerPort,omitempty"`
				HostPort      int    `yaml:"hostPort" json:"hostPort"`
				Protocol      string `yaml:"protocol" json:"protocol,omitempty"`
				ServicePort   int    `yaml:"servicePort" json:"servicePort,omitempty"`
			} `yaml:"portMappings" json:"portMappings"`
			Privileged bool `yaml:"privileged" json:"privileged"`
			Parameters []struct {
				Key   string `yaml:"key" json:"key,omitempty"`
				Value string `yaml:"value" json:"value,omitempty"`
			} `yaml:"parameters" json:"parameters"`
		} `yaml:"docker" json:"docker"`
		ForcePullImage bool `yaml:"forcePullImage" json:"forcePullImage"`
		Volumes        []struct {
			ContainerPath string `yaml:"containerPath" json:"containerPath"`
			HostPath      string `yaml:"hostPath" json:"hostPath"`
			Mode          string `yaml:"mode" json:"mode"`
		} `yaml:"volumes" json:"volumes"`
	} `yaml:"container" json:"container"`
	CPUs         float64           `yaml:"cpus" json:"cpus,omitempty"`
	Dependencies []string          `yaml:"dependencies" json:"dependencies,omitempty"`
	Environment  map[string]string `yaml:"env" json:"env,omitempty"`
	Executor     string            `yaml:"executor" json:"executor,omitempty"`
	Fetch        []struct {
		URI        string `yaml:"uri" json:"uri"`
		Extract    bool   `yaml:"extract" json:"extract"`
		Executable bool   `yaml:"executable" json:"executable"`
		Cache      bool   `yaml:"cache" json:"cache"`
	} `yaml:"fetch" json:"fetch,omitempty"`
	HealthChecks []struct {
		Command                map[string]string `yaml:"command" json:"command,omitempty"`
		GracePeriodSeconds     int               `yaml:"gracePeriodSeconds" json:"gracePeriodSeconds,omitempty"`
		IntervalSeconds        int               `yaml:"intervalSeconds" json:"intervalSeconds,omitempty"`
		MaxConsecutiveFailures int               `yaml:"maxConsecutiveFailures" json:"maxConsecutiveFailures,omitempty"`
		Path                   string            `yaml:"path" json:"path,omitempty"`
		PortIndex              int               `yaml:"portIndex" json:"portIndex,omitempty"`
		Protocol               string            `yaml:"protocol" json:"protocol,omitempty"`
		TimeoutSeconds         int               `yaml:"timeoutSeconds" json:"timeoutSeconds,omitempty"`
	} `yaml:"healthChecks" json:"healthChecks,omitempty"`
	ID                    string            `yaml:"id" json:"id"`
	OriginalID            string            `json:"originalID,omitempty"` // This should not really be here.  TODO: remove!
	Instances             int               `yaml:"instances" json:"instances"`
	Labels                map[string]string `yaml:"labels" json:"labels,omitempty"`
	MaxLaunchDelaySeconds int               `yaml:"maxLaunchDelaySeconds" json:"maxLaunchDelaySeconds,omitempty"`
	Memory                float64           `yaml:"mem" json:"mem,omitempty"`
	Ports                 []int             `yaml:"ports" json:"ports,omitempty"`
	PortDefinitions       []struct {
		Port     int               `yaml:"port" json:"port"`
		Protocol string            `yaml:"protocol" json:"protocol"`
		Labels   map[string]string `yaml:"labels" json:"labels"`
	} `yaml:"portDefinitions" json:"portDefinitions,omitempty"`
	RequirePorts    bool `yaml:"requirePorts,omitempty" json:"requirePorts"`
	UpgradeStrategy struct {
		MaximumOverCapacity   *float64 `yaml:"maximumOverCapacity" json:"maximumOverCapacity,omitempty"`
		MinimumHealthCapacity *float64 `yaml:"minimumHealthCapacity" json:"minimumHealthCapacity,omitempty"`
	} `yaml:"upgradeStrategy" json:"upgradeStrategy,omitempty"`
	URIs []string `yaml:"uris" json:"uris,omitempty"`
}

// ToJSON returns a JSON string representation of itself.
func (t *MarathonApp) ToJSON() []byte {
	jsonString, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling JSON: %s", err.Error())
	}
	return jsonString
}
