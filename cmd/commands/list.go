package commands

import (
	"fmt"
	ck "muse/cmd/commands/make"

	"github.com/charmbracelet/lipgloss"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var templateStyle = lipgloss.NewStyle().
	Bold(true).
	Margin(1).
	Foreground(lipgloss.Color("#8934eb")).
	BorderLeft(true).
	BorderForeground(lipgloss.Color("#8934eb")).
	BorderStyle(lipgloss.NormalBorder())

var ListCmd = &cobra.Command{
	Use:   "list [language_abbr]",
	Short: "list according of the flag provided",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			cmd.Help()
			return
		}

		getViteTemplates()
	},
}

func getViteTemplates() {
	i := 1
	logrus.Infoln("listing vite templates:")

	for k := range ck.ViteTemplates {
		f := templateStyle.Render(fmt.Sprintf(" %d. %s", i, k))

		fmt.Println(f)
		i++
	}
}

func GetBackJsTemplates() {

}
