package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"

	"github.com/sirupsen/logrus"
)

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
