package make

import "github.com/spf13/cobra"

var PyCommand = &cobra.Command{
	Use:   "py",
	Short: "create a project with python template",
	Long:  "create project with venv environment support. Support for data science with data sets from Kaggle coming soon",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
