package cmd

import (
	cm "muse/cmd/commands"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "muse",
	Short: "Muse is very fast project generator",
	Long:  "A quick project builder to create projects with python, java, javascript, typescript with our available templates",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true,
		ForceColors:      true,
	})

	rootCmd.AddCommand(cm.VersionCmd)
	rootCmd.AddCommand(cm.MakeCmd)
	rootCmd.AddCommand(cm.ConfigCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
