package models

type K8sConfig struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

func NewK8sConfig(name string, location string) K8sConfig {
	return K8sConfig{name, location}
}
