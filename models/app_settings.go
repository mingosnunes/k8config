package models

import (
	"encoding/json"
	"k8config/utils"
	"log"
	"os"
	"time"

	"github.com/AlecAivazis/survey/v2"
)

type AppSettings struct {
	CurrentConfig K8sConfig   `json:"current"`
	ConfigList    []K8sConfig `json:"configs"`
	UpdatedAt     time.Time   `json:"updatedAt"`
}

func NewAppSettings() AppSettings {
	return AppSettings{K8sConfig{}, make([]K8sConfig, 0), time.Now()}
}

func (settings *AppSettings) CheckConfigName(name string) bool {
	for _, config := range settings.ConfigList {
		if config.Name == name {
			return false
		}
	}

	return true
}

func (settings *AppSettings) AddConfig(newConfig K8sConfig) bool {

	for _, config := range settings.ConfigList {
		if config.Name == newConfig.Name {
			return false
		}
	}

	settings.ConfigList = append(settings.ConfigList, newConfig)

	settings.SaveFile()

	return true
}

func (settings *AppSettings) SaveFile() {

	prevSettings := GetSettings()

	if prevSettings.UpdatedAt != settings.UpdatedAt {
		override := false
		prompt := &survey.Confirm{
			Message: "Settings file change in the meantime. Override?",
			Default: true,
		}
		survey.AskOne(prompt, &override)

		if !override {
			return
		}
	}

	settings.UpdatedAt = time.Now()

	file, _ := json.Marshal(settings)

	errWriteFile := os.WriteFile(utils.SettingsPath, file, 0644)

	if errWriteFile != nil {
		log.Fatalln("Error: ", errWriteFile.Error())
	}
}

func CreateSettings() {
	settings := NewAppSettings()

	file, _ := json.MarshalIndent(settings, "", " ")

	errWriteFile := os.WriteFile(utils.SettingsPath, file, 0644)

	if errWriteFile != nil {
		log.Fatalln("Error: ", errWriteFile.Error())
	}
}

func GetSettings() AppSettings {
	bytesRead, err := os.ReadFile(utils.SettingsPath)

	if err != nil {
		log.Fatal("Settings file reading failed.")
	}

	var tempSettings AppSettings

	err = json.Unmarshal(bytesRead, &tempSettings)

	if err != nil {
		log.Fatal("Settings file decoding failed.")
	}

	return tempSettings
}
