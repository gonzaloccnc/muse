package commands

import (
	"muse/db"
	"muse/utils"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var ConfigCmd = &cobra.Command{
	Use:   "config [alias] [path]",
	Short: "Set up your context with a shortcut to create quick projects",
	Long:  "Set up aliases for quick project creation: config alias path and create projects with make -a alias. Example: config \"js\" \"/projects/javascript\" then make -a js",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			logrus.Warnln("You need provide two args for this command. use --help for more information")
			os.Exit(1)
		}

		alias := args[0]
		path := args[1]

		if !utils.ValidateAlias(alias) {
			logrus.Fatalln("You should send an valid alias that is between 2 to 10 chars")
		}

		if !utils.ValidatePath(path) {
			logrus.Fatalln("You should send a valid path in absolute or relative")
		}

		p, err := filepath.Abs(path)

		if err != nil {
			logrus.Fatalln("Ups there was an error")
		}

		db.InsertItem(db.ItemJSON{
			Alias: alias,
			Path:  p,
		})

		logrus.Infoln("Alias: ", alias)
		logrus.Infoln("Path: ", p)
	},
}
