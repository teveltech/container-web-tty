package config

var SHELL_LIST = []string{"/bin/bash", "/bin/ash", "/bin/sh"}

type DockerConfig struct {
	DockerHost string // default is /var/run/docker.sock
	PsOptions  string
}

type KubeConfig struct {
	KubectlPath string // default is /usr/bin/kubectl
	ConfigPath  string // normally is $HOME/.kube/config
}

type BackendConfig struct {
	Type      string // docker or kubectl (for now)
	Docker    DockerConfig
	Kube      KubeConfig
	ExtraArgs []string // extra args pass to docker or kubectl
}

type ControlConfig struct {
	Enable  bool
	All     bool
	Start   bool
	Stop    bool
	Restart bool
}

type Config struct {
	Port    int
	Debug   bool
	Control ControlConfig
	Backend BackendConfig
	Servers []string // for proxy mode
}
