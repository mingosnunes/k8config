package cmd

import (
	"errors"
	"fmt"
	"io/fs"
	"testing"
)

type InstallTestTable struct {
	Name                  string
	getHomeDir            func() (string, error)
	CheckInstallationList func() []int
	mkDirRoot             func(name string, perm fs.FileMode) error
	mkDirConfig           func(name string, perm fs.FileMode) error
	createSettings        func() error
	expectsErr            bool
}

func TestInstall(t *testing.T) {
	tests := []InstallTestTable{
		{
			Name: "Good Check",
			CheckInstallationList: func() []int {
				return []int{}
			},
			expectsErr: false,
		},
		{
			Name: "Bad Check - Bad os.UserHomeDir",
			getHomeDir: func() (string, error) {
				return "", errors.New("[unit-test] Error message")
			},
			CheckInstallationList: func() []int {
				return []int{1}
			},
			expectsErr: true,
		},
		{
			Name: "Bad Check - 1 Error",
			getHomeDir: func() (string, error) {
				return "/home/test", nil
			},
			CheckInstallationList: func() []int {
				return []int{1}
			},
			mkDirRoot: func(name string, perm fs.FileMode) error {
				return nil
			},
			mkDirConfig: func(name string, perm fs.FileMode) error {
				return nil
			},
			createSettings: func() error {
				return nil
			},
			expectsErr: false,
		},
		{
			Name: "Bad Check - 1 Error - Bad mkDirRoot",
			getHomeDir: func() (string, error) {
				return "/home/test", nil
			},
			CheckInstallationList: func() []int {
				return []int{1}
			},
			mkDirRoot: func(name string, perm fs.FileMode) error {
				return errors.New("[unit-test] Error message")
			},
			expectsErr: true,
		},
		{
			Name: "Bad Check - 1 Error - Bad mkDirConfigs",
			getHomeDir: func() (string, error) {
				return "/home/test", nil
			},
			CheckInstallationList: func() []int {
				return []int{1}
			},
			mkDirRoot: func(name string, perm fs.FileMode) error {
				return nil
			},
			mkDirConfig: func(name string, perm fs.FileMode) error {
				return errors.New("[unit-test] Error message")
			},
			createSettings: func() error {
				return nil
			},
			expectsErr: true,
		},
		{
			Name: "Bad Check - 1 Error - Bad createSettings",
			getHomeDir: func() (string, error) {
				return "/home/test", nil
			},
			CheckInstallationList: func() []int {
				return []int{1}
			},
			mkDirRoot: func(name string, perm fs.FileMode) error {
				return nil
			},
			createSettings: func() error {
				return errors.New("[unit-test] Error message")
			},
			expectsErr: true,
		},
		{
			Name: "Bad Check - 2 Error",
			getHomeDir: func() (string, error) {
				return "/home/test", nil
			},
			CheckInstallationList: func() []int {
				return []int{2}
			},
			createSettings: func() error {
				return nil
			},
			expectsErr: false,
		},
		{
			Name: "Bad Check - 2 Error - Bad createSettings",
			getHomeDir: func() (string, error) {
				return "/home/test", nil
			},
			CheckInstallationList: func() []int {
				return []int{2}
			},
			createSettings: func() error {
				return errors.New("[unit-test] Error message")
			},
			expectsErr: true,
		},
		{
			Name: "Bad Check - 3 Error",
			getHomeDir: func() (string, error) {
				return "/home/test", nil
			},
			mkDirConfig: func(name string, perm fs.FileMode) error {
				return nil
			},
			CheckInstallationList: func() []int {
				return []int{3}
			},
			expectsErr: false,
		},
		{
			Name: "Bad Check - 3 Error - Bad mkdir",
			getHomeDir: func() (string, error) {
				return "/home/test", nil
			},
			mkDirConfig: func(name string, perm fs.FileMode) error {
				return errors.New("[unit-test] Error message")
			},
			CheckInstallationList: func() []int {
				return []int{3}
			},
			expectsErr: true,
		},
		{
			Name: "Bad Check - 4 Error",
			getHomeDir: func() (string, error) {
				return "/home/test", nil
			},
			CheckInstallationList: func() []int {
				return []int{4}
			},
			expectsErr: false,
		},
	}

	for _, test := range tests {
		t.Run(
			test.Name,
			func(t *testing.T) {
				// functions to mock
				getHomeDir = test.getHomeDir
				checkInstallation = test.CheckInstallationList
				mkDirRoot = test.mkDirRoot
				mkDirConfigs = test.mkDirConfig
				createSettings = test.createSettings

				// execute de command
				_, output, err := executeCmd(rootCmd, "install")

				fmt.Println(output)
				if err != nil {
					if !test.expectsErr {
						t.Error("Expected not error and got error")
					}
				} else {
					if test.expectsErr {
						t.Error("Expected error and got not error")
					}
				}
			},
		)
	}

}
