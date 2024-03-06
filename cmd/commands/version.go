package commands

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "See the version number",
	Run: func(cmd *cobra.Command, args []string) {
		// call an api for get latest version
		logrus.Println("v0.1.1beta HEAD")
	},
}
