package make

import "github.com/spf13/cobra"

var TsCommand = &cobra.Command{
	Use:   "ts",
	Short: "create a project with ts template",
	Long:  "create a project with ts template. Support for vite for frontend projects. Backend TS coming soon",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
