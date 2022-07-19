/*
Copyright © 2022 Domingos Nunes mingosnunes94@gmail.com

*/
package utils

import (
	"os"
)

var ConfigsPath string = os.Getenv("HOME") + "/.k8config/configs"
var SettingsPath string = os.Getenv("HOME") + "/.k8config/settings.json"
var RootPath string = os.Getenv("HOME") + "/.k8config"
var ActualConfigPath string = os.Getenv("HOME") + "/.k8config/actual"
