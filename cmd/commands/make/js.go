package make

import "github.com/spf13/cobra"

var JsCommand = &cobra.Command{
	Use:   "js",
	Short: "create a project with js template",
	Long:  "create a project with js template. Support for vite for frontend projects. Backend JS comming soon",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
