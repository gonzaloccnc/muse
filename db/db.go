package db

import (
	"encoding/json"
	"fmt"
	"muse/utils"
	"os"

	"github.com/sirupsen/logrus"
)

var DB = "db.json"

type ItemJSON struct {
	Id    string `json:"id"`
	Alias string `json:"alias"`
	Path  string `json:"path"`
}

type AliasesJSON = struct {
	Aliases []ItemJSON `json:"aliases"`
}

func CreateDB() {
	_, err := os.Stat(DB)

	if err == nil {
		return
	}

	saveInJsonFile(AliasesJSON{Aliases: []ItemJSON{}})
}

func InsertItem(item ItemJSON) {
	aliasesJSON := getJsonFile(DB)

	isDuplicated := isDuplicateAlias(item.Alias, aliasesJSON)

	if isDuplicated {
		logrus.Fatalf("the item with alias %s exist in the aliases", item.Alias)
	}

	nextId := len(aliasesJSON.Aliases) + 1

	item.Id = fmt.Sprint(nextId)

	aliasesJSON.Aliases = append(aliasesJSON.Aliases, item)

	saveInJsonFile(aliasesJSON)
}

func DeleteItem(alias string) {
	utils.Todo("delete item")
}

func UpdateItem(alias string, item ItemJSON) {
	utils.Todo("update item")
}

func getJsonFile(path string) AliasesJSON {
	var aliasesJSON AliasesJSON

	jsonBytes, err := os.ReadFile(path)

	if err != nil {
		logrus.Fatal(err)
	}

	if err := json.Unmarshal(jsonBytes, &aliasesJSON); err != nil {
		logrus.Fatal(err)
	}

	return aliasesJSON
}

func isDuplicateAlias(alias string, aliases AliasesJSON) bool {
	for _, val := range aliases.Aliases {
		if val.Alias == alias {
			return true
		}
	}

	return false
}

func saveInJsonFile(aliases AliasesJSON) {
	aliasesJSON, err := json.Marshal(aliases)

	if err != nil {
		logrus.Fatal(err)
	}

	if err := os.WriteFile(DB, aliasesJSON, 0644); err != nil {
		logrus.Fatal(err)
	}
}