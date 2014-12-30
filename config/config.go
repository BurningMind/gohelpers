package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
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

func Value(key string) interface{} {
	switch configType {
	case "json":
		return readValueJson(key)

	default:
		panic("Unhandled config type")
	}
}

func SetValue(key string, value interface{}) {
	switch configType {
	case "json":
		writeValueJson(key, value)

	default:
		panic("Unhandled config type.")
	}
}

// JSON R/W functions

func readValueJson(key string) interface{} {
	data, err := ioutil.ReadFile(ConfigName() + "." + ConfigType())
	check(err)

	config := make(map[string]interface{})
	err = json.Unmarshal(data, &config)
	check(err)

	return config[key]
}

func writeValueJson(key string, value interface{}) {
	file, err := os.OpenFile(ConfigName()+"."+ConfigType(),
		os.O_RDWR|os.O_CREATE,
		0666)
	check(err)
	defer file.Close()

	config := make(map[string]interface{})
	if stat, err := file.Stat(); stat.Size() != 0 && err == nil {
		data, err := ioutil.ReadAll(file)
		check(err)

		err = json.Unmarshal(data, &config)
		check(err)
	}

	config[key] = value

	data, err := json.Marshal(config)
	check(err)

	file.Truncate(0)
	_, err = file.WriteAt(data, 0)
	check(err)
}
