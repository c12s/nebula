package main

type Deployment struct {
	Version    string   `yaml:"version"`
	Metadata   Metadata `yaml:"metadata"`
	Deploy     Deploy   `yaml:"deploy"`
}

type Metadata struct {
	Name   string            `yaml:"name"`
	Namespace string         `yaml:"namespace"`
	Labels map[string]string `yaml:"labels"`
}

type Deploy struct {
	Selector   map[string]string `yaml:"selector"`
	Containers []Container       `yaml:"containers"`
}

type Container struct {
	Name  string         `yaml:"name"`
	Image string         `yaml:"image"`
	Replicas   int       `yaml:replicas"`
	Ports map[string]int `yaml:"ports"`
}
