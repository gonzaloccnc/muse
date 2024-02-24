package utils

import (
	"fmt"
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
