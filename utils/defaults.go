package utils

import (
	"os"
)

var ConfigsPath string = os.Getenv("HOME") + "/.k8config/configs"
var HomePath string = os.Getenv("HOME")
var SettingsPath string = os.Getenv("HOME") + "/.k8config/settings.json"
var RootPath string = os.Getenv("HOME") + "/.k8config"
var ActualConfigPath string = os.Getenv("HOME") + "/.k8config/actual"
