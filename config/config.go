package config

import (
	"encoding/json"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var configName string = "config"
var configType string = "json"

// Getters and Setters

func ConfigName() string {
	return configName
}

func SetConfigName(newConfigName string) {
	configName = newConfigName
}

func ConfigType() string {
	return configType
}

func SetConfigType(newConfigType string) {
	configType = newConfigType
}

// General R/W functions

// Simple wrapper to get a specific value given a key
func Value(key string) interface{} {
	return Config()[key]
}

// Returns the config object depending on configType
func Config() map[string]interface{} {
	// We handle the different config types
	switch configType {
	case "json":
		return readConfigJson()

	default:
		panic("Unhandled config type.")
	}
}

// Simple wrapper to assign a specific value to a key
func SetValue(key string, value interface{}) {
	config := Config()
	config[key] = value
	SetConfig(config)
}

// Sets the config object depending on configType
func SetConfig(config map[string]interface{}) {
	// We handle the different config types
	switch configType {
	case "json":
		writeConfigJson(config)

	default:
		panic("Unhandled config type.")
	}
}

// JSON R/W functions

func readConfigJson() map[string]interface{} {
	// We read the JSON config file
	data, err := ioutil.ReadFile(ConfigName() + "." + ConfigType())

	// We JSON-decode the config if it exists, else we use an empty config
	config := make(map[string]interface{})
	if err == nil && len(data) != 0 {
		err = json.Unmarshal(data, &config)
		check(err)
	}
	return config
}

func writeConfigJson(config map[string]interface{}) {
	// We JSON-encode the config
	data, err := json.Marshal(config)
	check(err)

	// We write the JSON config file
	ioutil.WriteFile(ConfigName() + "." + ConfigType(),
	                 data,
	                 0666)
	check(err)
}
