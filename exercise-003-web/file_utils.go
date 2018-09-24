package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func saveState(users *map[string]int, dbPath string) {
	bytes, err := json.Marshal(*users)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Can't serialize")
	}
	ioutil.WriteFile(dbPath, bytes, 0644)
}

func restoreState(users *map[string]int, dbPath string) {
	if !exists(DbPath) {
		return
	}
	data, err := ioutil.ReadFile(dbPath)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Can't read")
		return
	}

	err = json.Unmarshal(data, users)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Can't deserislize")
	}
}

func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
