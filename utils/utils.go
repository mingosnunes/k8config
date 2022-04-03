package utils

import (
	"encoding/json"
	"k8config/models"
	"log"
	"os"
)

func CreateSettings(homePath string) {
	appSettings := models.NewAppSettings()

	file, _ := json.MarshalIndent(appSettings, "", " ")

	errWriteFile := os.WriteFile(homePath+"/.k8config/settings.json", file, 0666)

	if errWriteFile != nil {
		log.Fatalln("Error: ", errWriteFile.Error())
	}
}
