package cmd

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
)

func executeCmd(cmd *cobra.Command, args ...string) (c *cobra.Command, output string, err error) {
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs(args)

	c, err = cmd.ExecuteC()

	return c, buf.String(), err
}

type RootTestTable struct {
	Name                  string
	CheckInstallationList func() []int
	expectsErr            bool
}

func TestRoot(t *testing.T) {
	tests := []RootTestTable{
		{
			Name: "Good Check",
			CheckInstallationList: func() []int {
				return []int{}
			},
			expectsErr: false,
		},
		{
			Name: "Bad Check - 1 Error",
			CheckInstallationList: func() []int {
				return []int{1}
			},
			expectsErr: true,
		},
		{
			Name: "Bad Check - 2 Error",
			CheckInstallationList: func() []int {
				return []int{2, 3}
			},
			expectsErr: true,
		},
	}

	for _, test := range tests {
		t.Run(
			test.Name,
			func(t *testing.T) {
				checkInstallation = test.CheckInstallationList

				_, _, err := executeCmd(rootCmd)

				// t.Log("CMD", cmd)
				// t.Log("OUTPUT", output)

				t.Log(err)
				t.Log()
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
