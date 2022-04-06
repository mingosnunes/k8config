package utils

import (
	"os"
)

func CheckInstallation() []int {
	homePath := os.Getenv("HOME")

	rootDir, errRoot := os.Stat(RootPath)

	errorList := make([]int, 0)

	// Root Directory check
	if errRoot != nil || (errRoot == nil && !rootDir.IsDir()) {
		errorList = append(errorList, 1)
		return errorList
	}

	// Settings file check
	_, errSettings := os.Stat(homePath + "/.k8config/settings.json")

	if errSettings != nil {
		errorList = append(errorList, 2)
	}

	// Config directory check
	configDir, errConfig := os.Stat(ConfigsPath)

	if errConfig != nil || (errConfig == nil && !configDir.IsDir()) {
		errorList = append(errorList, 3)
	}

	return errorList
}
