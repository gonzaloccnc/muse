package commands

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var template bool

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "list according of the flag provided",
	Run: func(cmd *cobra.Command, args []string) {
		// now only template is available becase ->

		if cmd.Flags().NFlag() == 0 {
			logrus.Fatalln("you nedd provide some flag:")
		}

		// listing templates

		logrus.Infoln("listing tempaltes:")
	},
}

func init() {
	ListCmd.Flags().BoolVarP(&template, "tempalte", "t", false, "list all project templates availables")
}
