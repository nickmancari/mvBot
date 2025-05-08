package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var File = "/etc/mvBot/settings.json"

type Setting struct {
	RemoteServer	string
	RemoteUser	string
	RemoteFolder	string
	LocalFolder	string
}

func Read() Setting {

	var setting Setting

	jsonFile, err := os.ReadFile(File)
	if err != nil {
		fmt.Printf("Cannot Read File: %v\n", err)
		return setting
	}

	err = json.Unmarshal(jsonFile, &setting)
	if err != nil {
		fmt.Printf("Cannot Unmarshal Settings: %v\n", err)
	}

	return setting

}

