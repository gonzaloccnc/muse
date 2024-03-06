package commands

import (
	ck "muse/cmd/commands/make"
	"muse/db"
	"muse/utils"
	"path/filepath"

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
	Long:  "make project with this command, This take the alias path or output concatenate with the name alias. /output_or_alias/projects + myAppName = /output_or_alias/projects/myAppName/",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		if ck.Name == "" {
			logrus.Fatalln("you need provide the name flag. see --help")
		}

		item := db.GetAlias(ck.Alias)
		// validate path output

		if item == nil {
			ck.FinalPath = ck.Output
			logrus.Infoln("the alias is not found or not provided", ck.Alias)
		} else {
			ck.FinalPath = item.Path
		}

		logrus.Infoln("the output path is: ", filepath.Join(ck.FinalPath, ck.Name))
	},
}

func init() {
	MakeCmd.AddCommand(ck.JavaCommand)
	MakeCmd.AddCommand(ck.JsCommand)
	MakeCmd.AddCommand(ck.PyCommand)
	MakeCmd.PersistentFlags().StringVarP(&ck.Alias, "alias", "a", "", "use this instead absolute or relative path with (--output, -o)")
	MakeCmd.PersistentFlags().StringVarP(&ck.Output, "output", "o", homeDir, "output where your project will be located")
	MakeCmd.PersistentFlags().StringVarP(&ck.Name, "name", "n", "", "name of your project dir")
}
