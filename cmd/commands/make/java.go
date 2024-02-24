package make

import "github.com/spf13/cobra"

var JavaCommand = &cobra.Command{
	Use:   "java",
	Short: "create a project with java template",
	Long:  "Create a project with java template. Only support for Spring framework",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
