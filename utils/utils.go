package utils

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"log"
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

	return responseData
}

func ReadZip(body []byte) {
	zipReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		log.Fatal(err)
	}

	// Read all the files from zip archive
	for _, zipFile := range zipReader.File {
		fmt.Println("Reading file:", zipFile.Name)
		unzippedFileBytes := readZipFile(zipFile)
		if err != nil {
			log.Println(err)
			continue
		}

		_ = unzippedFileBytes // this is unzipped file bytes
	}
}

func readZipFile(zf *zip.File) []byte {

	f, err := zf.Open()
	if err != nil {
		logrus.Fatalln(err)
	}
	defer f.Close()

	fileBytes, err := io.ReadAll(f)

	if err != nil {
		logrus.Fatalln(err)
	}

	return fileBytes
}
