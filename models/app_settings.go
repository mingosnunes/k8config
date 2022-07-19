/*
Copyright © 2022 Domingos Nunes mingosnunes94@gmail.com

*/
package models

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"time"

	"github.com/mingosnunes/k8config/utils"

	"github.com/AlecAivazis/survey/v2"
)

type IAppSettings interface {
	CheckConfigName(name string) bool
	AddConfig(newConfig K8sConfig) bool
	DelConfigs(configsSelected []string)
	UseConfig(configName string) error
	SaveFile() error
	GetConfigList() []K8sConfig
	GetUpdatedAt() time.Time
	GetCurrentConfig() K8sConfig
}

type AppSettings struct {
	CurrentConfig K8sConfig   `json:"current"`
	UpdatedAt     time.Time   `json:"updatedAt"`
	ConfigList    []K8sConfig `json:"configs"`
}

func NewAppSettings() AppSettings {
	return AppSettings{K8sConfig{}, time.Now(), make([]K8sConfig, 0)}
}

func (settings *AppSettings) GetCurrentConfig() K8sConfig {
	return settings.CurrentConfig
}

func (settings *AppSettings) GetConfigList() []K8sConfig {
	return settings.ConfigList
}

func (settings *AppSettings) GetUpdatedAt() time.Time {
	return settings.UpdatedAt
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

func (settings *AppSettings) DelConfigs(configsSelected []string) {

	for index, config := range settings.ConfigList {
		for _, selected := range configsSelected {
			if config.Name == selected {
				settings.ConfigList = utils.RemoveFromList(settings.ConfigList, index)
				os.Remove(config.Location)
			}
		}
	}

	settings.SaveFile()
}

func (settings *AppSettings) UseConfig(configName string) error {

	for _, config := range settings.ConfigList {

		if config.Name == configName {
			bytesRead, err := os.ReadFile(config.Location)
			if err != nil {
				log.Fatal(err)
			}

			err = os.WriteFile(utils.ActualConfigPath, bytesRead, 0644)
			if err != nil {
				log.Fatal(err)
			}

			settings.CurrentConfig = config

			settings.SaveFile()
		}
	}

	return nil

}

func (settings *AppSettings) SaveFile() error {

	prevSettings, err := GetSettings()

	if err != nil {
		return err
	}

	if prevSettings.GetUpdatedAt() != settings.GetUpdatedAt() {
		override := false
		prompt := &survey.Confirm{
			Message: "Settings file change in the meantime. Override?",
			Default: true,
		}
		err := survey.AskOne(prompt, &override)

		if err != nil {
			return err
		}

		if !override {
			return nil
		}
	}

	settings.UpdatedAt = time.Now()

	file, _ := json.Marshal(settings)

	errWriteFile := os.WriteFile(utils.SettingsPath, file, 0644)

	if errWriteFile != nil {
		return err
	}

	return nil
}

func CreateSettings() error {
	settings := NewAppSettings()

	file, _ := json.MarshalIndent(settings, "", " ")

	errWriteFile := os.WriteFile(utils.SettingsPath, file, 0644)

	return errWriteFile
}

func GetSettings() (IAppSettings, error) {
	bytesRead, err := os.ReadFile(utils.SettingsPath)

	if err != nil {
		return &AppSettings{}, errors.New("Settings file reading failed.")
	}

	var tempSettings AppSettings

	err = json.Unmarshal(bytesRead, &tempSettings)

	if err != nil {
		return &AppSettings{}, errors.New("Settings file decoding failed.")
	}

	return &tempSettings, nil
}
