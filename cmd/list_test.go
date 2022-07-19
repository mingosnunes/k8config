package cmd

import (
	"errors"
	"fmt"
	"testing"

	"github.com/mingosnunes/k8config/models"
)

type ListTestTable struct {
	Name                  string
	CheckInstallationList func() []int
	ExpectsErr            bool
	GetSettings           func() (models.IAppSettings, error)
}

func TestList(t *testing.T) {
	tests := []ListTestTable{
		{
			Name: "Good GetSettings",
			CheckInstallationList: func() []int {
				return []int{}
			},
			GetSettings: func() (models.IAppSettings, error) {
				return &models.AppSettings{
						ConfigList: []models.K8sConfig{
							{
								Name:     "Test-1",
								Location: "Location-1",
							},
							{
								Name:     "Test-2",
								Location: "Location-2",
							},
						},
					},
					nil
			},
			ExpectsErr: false,
		},
		{
			Name: "Bad GetSettings",
			CheckInstallationList: func() []int {
				return []int{}
			},
			GetSettings: func() (models.IAppSettings, error) {
				return &models.AppSettings{},
					errors.New("[unit-test] Expected Error")
			},
			ExpectsErr: true,
		},
	}

	for _, test := range tests {
		t.Run(
			test.Name,
			func(t *testing.T) {
				// functions to mock
				checkInstallation = test.CheckInstallationList
				listGetSettings = test.GetSettings

				// execute de command
				_, output, err := executeCmd(rootCmd, "list")

				fmt.Println(output)
				if err != nil {
					if !test.ExpectsErr {
						t.Error("Expected not error and got error")
					}
				} else {
					if test.ExpectsErr {
						t.Error("Expected error and got not error")
					}
				}
			},
		)
	}
}
