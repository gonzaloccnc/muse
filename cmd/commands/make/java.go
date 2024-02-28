package make

import (
	"bytes"
	"encoding/json"
	"muse/utils"

	"muse/utils/multiselect"
	"muse/utils/singleselect"
	"muse/utils/spring"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	ti "muse/utils/textinput"
)

var mapType = map[string]string{
	"Maven":           "maven-project",
	"Gradle - Groovy": "gradle-project",
	"Gradle - Kotlin": "gradle-project-kotlin",
}

var JavaCommand = &cobra.Command{
	Use:   "java",
	Short: "create a project with java template",
	Long:  "Create a project with java template. Only support for Spring framework",
	Run: func(cmd *cobra.Command, args []string) {

		responseData := utils.Fetch("https://start.spring.io/metadata/client")

		var metadata spring.ResponseMetadata
		json.Unmarshal(responseData, &metadata)

		data := springInitializr(metadata)
		qs := structToQueryString(data)

		url := "https://start.spring.io/starter.zip?" + qs

		zip := utils.Fetch(url)

		outDir := filepath.Join(FinalPath, Name)
		utils.CreateDir(outDir)

		outZipDir := filepath.Join(FinalPath, Name, Name+".zip")
		if err := os.WriteFile(outZipDir, zip, 0777); err != nil {
			logrus.Fatalln(err)
		}

		logrus.Println("project create successful: ", outZipDir)
	},
}

func springInitializr(metadata spring.ResponseMetadata) spring.QueryString {
	dependenciesManager := []string{"Gradle - Groovy", "Gradle - Kotlin", "Maven"}
	languages := []string{"Java", "Kotlin", "Groovy"}
	packagings := []string{"jar", "war"}
	bootVersions, bootVersionsMap := getSpringVersions(metadata.BootVersion.Values)
	javaVersions, _ := getSpringVersions(metadata.JavaVersion.Values)
	projects := getTypeDependencies(metadata.Dependencies.Values)

	manager := mapType[selectOne("Select the dependencies manager", dependenciesManager)]
	language := selectOne("Select the dependencies manager", languages)
	bootVersion := selectOne("Select the dependencies manager", bootVersions)
	groupId := input("Group Id", "com.example")
	artifactId := input("Artifact Id", "demo")
	name := input("Name", artifactId)
	description := input("Description", "Demo project for Spring Boot")
	packageName := input("Package name", groupId+"."+artifactId)
	packaging := selectOne("Packaging", packagings)
	javaVersion := selectOne("Java versions", javaVersions)

	dependencies := chooseDependencies(projects, metadata.Dependencies.Values)

	dependenciesStr := deleteDuplicateAndSplit(dependencies)

	return spring.QueryString{
		Type:         manager,
		Language:     language,
		BootVersion:  bootVersionsMap[bootVersion],
		GroupId:      groupId,
		ArtifactId:   artifactId,
		Name:         name,
		Description:  description,
		PackageName:  packageName,
		Packaging:    packaging,
		JavaVersion:  javaVersion,
		Dependencies: dependenciesStr,
	}
}

func checkboxes(label string, opts []string) []string {
	choices, err := multiselect.Run(label, opts, 6)

	if err != nil {
		logrus.Fatalln(err)
	}

	return choices
}

func selectOne(label string, opts []string) string {
	choice, err := singleselect.Run(label, opts, 6)

	if err != nil {
		logrus.Fatalln(err)
	}

	return choice
}

func input(label string, placeholder string) string {
	value, err := ti.Run(placeholder, label)

	if err != nil {
		logrus.Fatalln(err)
	}

	return value
}

func getSpringVersions(structs []spring.BootVersionValue) ([]string, map[string]string) {
	strings := make([]string, len(structs))
	mapVersion := make(map[string]string)

	for i, v := range structs {
		mapVersion[v.Name] = v.ID
		strings[i] = v.Name
	}

	return strings, mapVersion
}

func getTypeDependencies(root []spring.DependenciesValue) []string {
	strings := make([]string, len(root))
	for i, v := range root {
		strings[i] = v.Name
	}
	return strings
}

func getDependencies(dependencies []spring.ValueValue) []string {
	arrDeps := make([]string, len(dependencies))

	for i, v := range dependencies {
		arrDeps[i] = v.Name
	}

	return arrDeps
}

func convertToMap(dependenciesValues []spring.DependenciesValue) map[string]string {
	var mapDeps = map[string]string{}

	for _, v := range dependenciesValues {
		for _, dep := range v.Values {
			mapDeps[dep.Name] = dep.ID
		}
	}

	return mapDeps
}

func chooseDependencies(projects []string, projectsMetadata []spring.DependenciesValue) []string {
	var dependencies []string

	appendProjects := append(projects, "Exit")
	mapDeps := convertToMap(projectsMetadata)

	for {
		project := selectOne("Choose an project for add a dependency", appendProjects)

		for _, v := range projectsMetadata {
			if v.Name == project {
				dependencies = append(
					dependencies,
					checkboxes("Select dependency", getDependencies(v.Values))...,
				)
			}
		}

		if project == "Exit" {
			break
		}
	}

	var ids []string
	for _, v := range dependencies {
		ids = append(ids, mapDeps[v])
	}

	return ids
}

func deleteDuplicateAndSplit(arr []string) string {
	// missing delete duplicates
	return strings.Join(arr, ",")
}

func structToQueryString(structData interface{}) string {
	values := url.Values{}

	structValues := reflect.ValueOf(structData)
	types := structValues.Type()

	for i := 0; i < structValues.NumField(); i++ {
		b := []byte(types.Field(i).Name)
		b[0] = bytes.ToLower(b)[0]
		fieldName := string(b)
		fieldValue := strings.ToLower(structValues.Field(i).String())

		if fieldName == "dependencies" {
			values.Set(fieldName, fieldValue)
		}

		values.Set(fieldName, fieldValue)
	}

	return values.Encode()
}
