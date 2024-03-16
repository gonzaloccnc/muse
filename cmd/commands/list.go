package commands

import (
	"fmt"
	ck "muse/cmd/commands/make"
	"muse/db"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	templateStyle = lipgloss.NewStyle().
			Bold(true).
			Margin(1).
			Foreground(lipgloss.Color("#8934eb")).
			BorderLeft(true).
			BorderForeground(lipgloss.Color("#8934eb")).
			BorderStyle(lipgloss.NormalBorder())

	purpleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#8934eb"))
	itemStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#16cc29"))
)

var ListCmd = &cobra.Command{
	Use:   "list [language_abbr]",
	Short: "list according of the flag provided",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			cmd.Help()
			return
		}

		// get from a map
		if args[0] == "alias" {
			listAliases()
		}

		if args[0] == "ts" || args[0] == "js" {
			getViteTemplates()
		}
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

func listAliases() {
	aliases := db.GetJsonFile(db.JsonFile)

	s7 := strings.Repeat("-", 7)
	s12 := strings.Repeat("-", 12)
	s22 := strings.Repeat("-", 22)

	borderStyle := fmt.Sprintf("|%s|%s|%s|", s7, s12, s22)
	slashStyle := purpleStyle.Render("|")

	title := fmt.Sprintf("| %-5s | %-10s | %-20s |", "Order", "Alias", "Path")
	title = strings.ReplaceAll(title, "|", slashStyle)

	fmt.Println(purpleStyle.Render(borderStyle))
	fmt.Println(title)
	fmt.Println(purpleStyle.Render(borderStyle))

	for i, val := range aliases.Aliases {
		item := fmt.Sprintf("| %-5d | %-10s | %-20s |", i+1, val.Alias, val.Path)
		fmt.Println(itemStyle.Render(item))
		fmt.Println(purpleStyle.Render(borderStyle))
	}

}
