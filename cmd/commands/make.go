package commands

import (
	ck "muse/cmd/commands/make"
	"muse/db"
	"muse/utils"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var homeDir = utils.GetHomeDir("/projects")
var Templates = map[string]string{
	"java": "java",
	"js":   "javascript",
	"ts":   "typescript",
	"py":   "python",
}

var MakeCmd = &cobra.Command{
	Use:   "make",
	Short: "make project with this command with our support",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var finalPath string

		if ck.Name == "" {
			logrus.Fatalln("you need provide the name flag. see --help")
		}

		item := db.GetAlias(ck.Alias)

		if item == nil {
			finalPath = ck.Output
			logrus.Infoln("the alias is not found or not provided", ck.Alias)
		} else {
			finalPath = item.Path
		}

		logrus.Infoln("the output path is: ", finalPath)
	},
}

func init() {
	MakeCmd.AddCommand(ck.JavaCommand)
	MakeCmd.AddCommand(ck.JsCommand)
	MakeCmd.AddCommand(ck.PyCommand)
	MakeCmd.AddCommand(ck.TsCommand)
	MakeCmd.PersistentFlags().StringVarP(&ck.Alias, "alias", "a", "", "use this instead absolute or relative path with (--output, -o)")
	MakeCmd.PersistentFlags().StringVarP(&ck.Output, "output", "o", homeDir, "output where your project will be located")
	MakeCmd.PersistentFlags().StringVarP(&ck.Name, "name", "n", "", "name of your project dir")
}
