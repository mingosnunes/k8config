package cmd

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mingosnunes/k8config/models"
	"github.com/stretchr/testify/mock"
)

type MockAppSettings struct {
	ConfigList    []models.K8sConfig
	CurrentConfig models.K8sConfig
	mock.Mock
}

func (m *MockAppSettings) UseConfig(configName string) error {
	err := m.Called(configName)
	return err.Error(0)
}

func (settings *MockAppSettings) GetConfigList() []models.K8sConfig {
	return settings.ConfigList
}

func (settings *MockAppSettings) GetCurrentConfig() models.K8sConfig {
	return settings.CurrentConfig
}

func (m *MockAppSettings) SaveFile() error {
	err := m.Called()
	return err.Error(0)
}

func (m *MockAppSettings) AddConfig(newConfig models.K8sConfig) bool {
	args := m.Called(newConfig)
	return args.Bool(0)
}

func (m *MockAppSettings) CheckConfigName(name string) bool {
	args := m.Called(name)
	return args.Bool(0)
}

func (m *MockAppSettings) DelConfigs(configsSelected []string) {
}

func (settings *MockAppSettings) GetUpdatedAt() time.Time {
	return time.Time{}
}

type UseTestTable struct {
	Name                  string
	CheckInstallationList func() []int
	GetSettings           func() (models.IAppSettings, error)
	SurveyAskOne          func(p survey.Prompt, response interface{}, opts ...survey.AskOpt) error
	ExpectsErr            bool
}

func TestUse(t *testing.T) {
	tests := []UseTestTable{

		{
			Name: "Good Check - Bad GetSettings",
			CheckInstallationList: func() []int {
				return []int{}
			},
			GetSettings: func() (models.IAppSettings, error) {

				return nil, errors.New("[unit-test] Expected error")
			},
			SurveyAskOne: func(p survey.Prompt, response interface{}, opts ...survey.AskOpt) error {
				// Mock Survey Response
				target := reflect.ValueOf(response)
				elem := target.Elem()
				elem.SetString("test")

				return nil
			},
			ExpectsErr: true,
		},
		{
			Name: "Good Check - Good GetSettings - Bad SurveyAskOne",
			CheckInstallationList: func() []int {
				return []int{}
			},
			GetSettings: func() (models.IAppSettings, error) {

				settings := new(MockAppSettings)

				settings.ConfigList = []models.K8sConfig{
					{
						Name:     "Test-1",
						Location: "Location-1",
					},
					{
						Name:     "Test-2",
						Location: "Location-2",
					},
				}

				return settings, nil
			},
			SurveyAskOne: func(p survey.Prompt, response interface{}, opts ...survey.AskOpt) error {

				return errors.New("[unit-test] Expected error")
			},
			ExpectsErr: true,
		},
		{
			Name: "Good Check - Good GetSettings - Good SurveyAskOne - Bad UseConfig",
			CheckInstallationList: func() []int {
				return []int{}
			},
			GetSettings: func() (models.IAppSettings, error) {

				settings := new(MockAppSettings)

				settings.ConfigList = []models.K8sConfig{
					{
						Name:     "Test-1",
						Location: "Location-1",
					},
					{
						Name:     "Test-2",
						Location: "Location-2",
					},
				}

				settings.On("UseConfig", "test").Return(errors.New("[unit-test] Expected error"))

				return settings, nil
			},
			SurveyAskOne: func(p survey.Prompt, response interface{}, opts ...survey.AskOpt) error {

				// Mock Survey Response
				target := reflect.ValueOf(response)
				elem := target.Elem()
				elem.SetString("test")

				return nil
			},
			ExpectsErr: true,
		},
		{
			Name: "Good Check - Good GetSettings - Good SurveyAskOne - Good UseConfig",
			CheckInstallationList: func() []int {
				return []int{}
			},
			GetSettings: func() (models.IAppSettings, error) {

				settings := new(MockAppSettings)

				settings.ConfigList = []models.K8sConfig{
					{
						Name:     "Test-1",
						Location: "Location-1",
					},
					{
						Name:     "Test-2",
						Location: "Location-2",
					},
				}

				settings.On("UseConfig", "test").Return(nil)

				return settings, nil
			},
			SurveyAskOne: func(p survey.Prompt, response interface{}, opts ...survey.AskOpt) error {

				// Mock Survey Response
				target := reflect.ValueOf(response)
				elem := target.Elem()
				elem.SetString("test")

				return nil
			},
			ExpectsErr: false,
		},
	}

	for _, test := range tests {
		t.Run(
			test.Name,
			func(t *testing.T) {
				// functions to mock
				checkInstallation = test.CheckInstallationList
				useGetSettings = test.GetSettings
				useSurveyAskOne = test.SurveyAskOne

				// execute de command
				_, output, err := executeCmd(rootCmd, "use")

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
