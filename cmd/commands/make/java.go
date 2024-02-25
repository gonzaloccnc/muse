package make

import (
	"bytes"
	"encoding/json"
	"muse/utils"
	"muse/utils/spring"
	"net/url"
	"os"
	"reflect"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/AlecAivazis/survey/v2"
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
		qs := generateQueryString(data)

		url := "https://start.spring.io/starter.zip?" + qs

		zip := utils.Fetch(url)

		if err := os.WriteFile("data.zip", zip, 0777); err != nil {
			logrus.Fatalln(err)
		}
		// TODO this is done but catch errors
	},
}

func springInitializr(metadata spring.ResponseMetadata) spring.QueryString {
	dependenciesManager := []string{"Gradle - Groovy", "Gradle - Kotlin", "Maven"}
	languages := []string{"Java", "Kotlin", "Grovy"}
	packagings := []string{"jar", "war"}
	bootVersions := getSpringVersions(metadata.BootVersion.Values)
	javaVersions := getSpringVersions(metadata.JavaVersion.Values)
	projects := getTypeDepencies(metadata.Dependencies.Values)

	manager := mapType[selectOne("Select the dependecies manager", dependenciesManager, 0)]
	language := selectOne("Select the dependecies manager", languages, 0)
	bootVersion := selectOne("Select the dependecies manager", bootVersions, 0)
	groupId := input("Group Id: ...", "demo")
	artifactId := input("Artifact Id: ...", "demo")
	name := input("Name: ...", "demo")
	description := input("Description: ...", "Demo project for Spring Boot")
	packageName := input("Package name: ...", "com.example.demo")
	packaging := selectOne("Packaging", packagings, 0)
	javaVersion := selectOne("Java version: ...", javaVersions, 0)

	dependencies := chooseDependencies(projects, metadata.Dependencies.Values)

	dependenciesStr := deleteDuplicateAndSplit(dependencies)

	return spring.QueryString{
		Type:         manager,
		Language:     language,
		BootVersion:  bootVersion,
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

func multiSelect(label string, opts []string) []string {
	res := []string{}
	prompt := &survey.MultiSelect{
		Message: label,
		Options: opts,
	}
	survey.AskOne(prompt, &res)

	return res
}

func selectOne(label string, opts []string, pages int) string {
	var res string
	prompt := &survey.Select{
		Message:  label,
		Options:  opts,
		PageSize: pages,
	}
	survey.AskOne(prompt, &res)

	return res
}

func input(label string, defaultmg string) string {
	var res string
	promp := &survey.Input{
		Message: label,
		Default: defaultmg,
	}
	survey.AskOne(promp, &res)

	return res
}

func getSpringVersions(structs []spring.BootVersionValue) []string {
	strings := make([]string, len(structs))
	for i, v := range structs {
		strings[i] = v.Name
	}

	return strings
}

func getTypeDepencies(root []spring.DependenciesValue) []string {
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
		project := selectOne("Choose an project for add a dependency", appendProjects, 8)

		for _, v := range projectsMetadata {
			if v.Name == project {
				dependencies = append(
					dependencies,
					multiSelect("Select dependency", getDependencies(v.Values))...,
				)
			}
		}

		if project == "Exit" {
			break
		}
	}

	var ids = make([]string, len(dependencies))

	for _, v := range dependencies {
		ids = append(ids, mapDeps[v])
	}

	return ids
}

func deleteDuplicateAndSplit(arr []string) string {
	// missing delete duplicates
	return strings.Join(arr, ",")
}

func generateQueryString(structData interface{}) string {
	values := url.Values{}

	structValues := reflect.ValueOf(structData)
	types := structValues.Type()

	for i := 0; i < structValues.NumField(); i++ {
		b := []byte(types.Field(i).Name)
		b[0] = bytes.ToLower(b)[0]
		fieldName := string(b)
		fieldValue := strings.ToLower(structValues.Field(i).String())

		if fieldName == "dependencies" {
			break
		}

		values.Set(fieldName, fieldValue)
	}

	return values.Encode()
}
