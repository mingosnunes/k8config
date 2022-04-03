package models

type AppSettings struct {
	currentConfig K8sConfig
	configList    []K8sConfig
}

func NewAppSettings() AppSettings {
	return AppSettings{K8sConfig{}, make([]K8sConfig, 0)}
}
