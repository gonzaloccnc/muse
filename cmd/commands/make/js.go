package make

import (
	"fmt"
	"muse/utils"
	"os"
	"os/exec"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var packageManager string
var template string
var ViteTemplates = map[string]string{
	"vanilla":    "template-vanilla",
	"vanilla-ts": "template-vanilla-ts",
	"vue":        "template-vue",
	"vue-ts":     "template-vue-ts",
	"react":      "template-react",
	"react-ts":   "template-react-ts",
	"preact":     "template-preact",
	"preact-ts":  "template-preact-ts",
	"lit":        "template-lit",
	"lit-ts":     "template-lit-ts",
	"svelte":     "template-svelte",
	"svelte-ts":  "template-svelte-ts",
	"solid":      "template-solid",
	"solid-ts":   "template-solid-ts",
	"qwik":       "template-qwik",
	"qwik-ts":    "template-qwik-ts",
}

var JsCommand = &cobra.Command{
	Use:   "js [template]",
	Short: "create a project with js template",
	Long:  "create a project with js template. Support for Vite for frontend projects. Backend JS coming soon",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			logrus.Fatalln("you need provide one arg for a template")
		}

		template = args[0]

		_, ok := ViteTemplates[template]

		if !ok {
			logrus.Fatalf(
				"unknown template for '%s' use --list for see available templates",
				template,
			)
		}

		utils.CreateDir(FinalPath)

		if err := os.Chdir(FinalPath); err != nil {
			logrus.Fatalln(err)
		}

		dirName := path.Join(FinalPath, Name)

		if existDir(dirName) {
			con := utils.YesOrNo("the directory exists, this action will replace it")

			if con {
				utils.RemoveDir(dirName)
				execCommand(template)
			}

			return
		}

		execCommand(template)
	},
}

func init() {
	JsCommand.Flags().StringVarP(&packageManager, "package", "p", "npm", "package manager")
}

func execCommand(template string) {
	// if exist directory remove if the user say 'yes'
	templateSlashes := "--template"
	slashes := ""

	if packageManager == "npm" {
		slashes = "--"
	}

	c := exec.Command(
		packageManager,
		"create",
		"vite@latest",
		Name,
		slashes,
		templateSlashes,
		template,
	)

	if runtime.GOOS == "windows" {
		c = exec.Command(
			"cmd",
			"/C",
			packageManager,
			"create",
			"vite@latest",
			Name,
			slashes,
			templateSlashes,
			template,
		)
	}

	ou, err := c.Output()

	if err != nil {
		logrus.Fatalln(err)
	}

	fmt.Println(string(ou))
}

func existDir(path string) bool {
	info, err := os.Stat(path)

	if err != nil && os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}
