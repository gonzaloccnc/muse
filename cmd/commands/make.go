package commands

import (
	"fmt"

	ck "muse/cmd/commands/make"
	"muse/utils"

	"github.com/spf13/cobra"
)

var Alias string
var Output string
var Templates = map[string]string{
	"java": "java",
	"js":   "javascript",
	"ts":   "typescript",
	"py":   "python",
}

var MakeCmd = &cobra.Command{
	Use:   "make",
	Short: "make project with this command with our support",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO show help command if the user not pass args or flags
		fmt.Println(cmd.Flags().NFlag(), cmd.Flags().NArg())
		fmt.Println(args)
	},
}

func init() {
	homeDir := utils.GetHomeDir("/projects")

	MakeCmd.AddCommand(ck.JavaCommand)
	MakeCmd.AddCommand(ck.JsCommand)
	MakeCmd.AddCommand(ck.PyCommand)
	MakeCmd.AddCommand(ck.TsCommand)
	MakeCmd.PersistentFlags().StringVarP(&Alias, "alias", "a", "p", "use this instead absolute or relative path with (--output, -o)")
	MakeCmd.PersistentFlags().StringVarP(&Output, "output", "o", homeDir, "output where your project will be located")
}
