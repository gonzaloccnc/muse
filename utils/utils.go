package utils

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/sirupsen/logrus"
)

var yesOrNoStyle = lipgloss.NewStyle().Bold(true)

func Todo(msg string) {
	logrus.Fatalf("TODO!: %s\n", msg)
}

func ValidatePath(path string) bool {
	regex := regexp.MustCompile(`^(\\|\/)?([\w-._@]+(\\|\/)?)+$`)
	return regex.MatchString(path)
}

func ValidateAlias(alias string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z_]{2,10}$`)
	return regex.MatchString(alias)
}

func GetHomeDir(or string) string {
	var defaultPath string
	homeDir, err := os.UserHomeDir()

	if err != nil {
		defaultPath, err = filepath.Abs(or)

		if err != nil {
			defaultPath = fmt.Sprintf("/%s", or)
		}
	} else {
		defaultPath = homeDir
	}

	return defaultPath
}

func Fetch(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		logrus.Fatalln(err)
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		logrus.Fatalln(err)
	}

	if response.StatusCode != 200 {
		logrus.Fatalf(
			"an error occurred while get the data, code: %d\nBody: %s\n",
			response.StatusCode,
			string(responseData),
		)
	}

	return responseData
}

func CreateDir(path string) string {
	if err := os.MkdirAll(path, 0777); err != nil {
		logrus.Fatalln(err)
	}

	return ""
}

func YesOrNo(label string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(yesOrNoStyle.Render(label))
	fmt.Print(yesOrNoStyle.Render("Do you want to continue? (yes/no): ... "))
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(strings.ToLower(answer))

	if answer == "yes" || answer == "y" {
		return true
	} else if answer == "no" || answer == "n" {
		return false
	} else {
		return false
	}
}

func RemoveDir(path string) {
	if err := os.RemoveAll(path); err != nil {
		logrus.Fatalln(err)
	}
}
