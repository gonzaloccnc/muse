package utils_test

import (
	"muse/utils"
	"os"
	"testing"
)

func TestValidPaths(t *testing.T) {
	paths := []string{
		"as93-.@/",
		"/c/route",
		"c/PATH123@/",
		"c/c123",
		"c\\s/fg/",
		"c-/route",
		"c/route",
		"c/sap/pel/ads\\AS\\sa",
		"c\\route\\",
		"/a-\\k",
	}

	for _, path := range paths {
		t.Run(path, func(t *testing.T) {
			want := utils.ValidatePath(path)
			if !want {
				t.Errorf("path is not valid: %s", path)
			}
		})
	}
}

func TestValidAlias(t *testing.T) {
	aliases := []string{
		"alias_",
		"alias_vali",
		"alias",
		"ali",
	}

	for _, alias := range aliases {
		t.Run(alias, func(t *testing.T) {
			want := utils.ValidateAlias(alias)
			if !want {
				t.Errorf("path is not valid: %s", alias)
			}
		})
	}
}

func TestCreateDir(t *testing.T) {
	outDir := "/valid/path"
	utils.CreateDir(outDir)

	info, err := os.Stat(outDir)

	if err != nil {
		t.Error(err)
	}

	if !info.IsDir() {
		t.Errorf("this not create a directory is a file")
	}
}

func TestValidateFetch(t *testing.T) {
	// invalidUrl := "https://start.spring.io/starter.zip?type=gradle-project&language=java&bootVersion=3.2.3&baseDir=demo&groupId=com.example&artifactId=demo&name=demo&description=Demo%20project%20for%20Spring%20Boot&packageName=com.example.demo&packaging=jar&javaVersion=17&dependencies=lombok,,,"
	validUrl := "https://start.spring.io/starter.zip?type=gradle-project&language=java&bootVersion=3.2.3&baseDir=demo&groupId=com.example&artifactId=demo&name=demo&description=Demo%20project%20for%20Spring%20Boot&packageName=com.example.demo&packaging=jar&javaVersion=17&dependencies=lombok"
	utils.Fetch(validUrl)
}
