package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

// Tests config.readValueJson() by comparing a map to
// the output of readValueJson after writing the map as json
func TestReadValueJson(t *testing.T) {
	var tests = map[string]interface{}{
		"string":   interface{}(string("test")),
		"number":   interface{}(float64(42)),
		"boolean":  interface{}(bool(true)),
		"nulltype": interface{}(nil),
		"array": interface{}([]interface{}{
			interface{}(string("test")),
			interface{}(float64(42)),
		}),
		"object": interface{}(map[string]interface{}{
			"string": interface{}(string("test")),
			"number": interface{}(float64(42)),
		}),
	}

	SetConfigName("TestReadValueJson")
	SetConfigType("json")

	data, err := json.Marshal(tests)
	check(err)
	err = ioutil.WriteFile(ConfigName()+"."+ConfigType(),
		data,
		0666)
	check(err)
	defer os.Remove(ConfigName() + "." + ConfigType())

	for key, value := range tests {
		res := readValueJson(key)
		if !reflect.DeepEqual(res, value) {
			t.Error("Used ", key, ", got ", res, ", wanted ", value)
		}
	}
}

// Benchmarks config.readValueJson() by reading a simple string from json
func BenchmarkReadValueJson(b *testing.B) {
	SetConfigName("BenchmarkReadValueJson")
	SetConfigType("json")

	err := ioutil.WriteFile(ConfigName()+"."+ConfigType(),
		[]byte(`{"string": "test"}`),
		0666)
	check(err)
	defer os.Remove(ConfigName() + "." + ConfigType())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = readValueJson("string")
	}
	b.StopTimer()
}

// Tests config.writeValueJson() by comparing a map to
// the json map obtained from the output of writeValueJson)
func TestWriteValueJson(t *testing.T) {
	var tests = map[string]interface{}{
		"string":   interface{}(string("test")),
		"number":   interface{}(float64(42)),
		"boolean":  interface{}(bool(true)),
		"nulltype": interface{}(nil),
		"array": interface{}([]interface{}{
			interface{}(string("test")),
			interface{}(float64(42)),
		}),
		"object": interface{}(map[string]interface{}{
			"string": interface{}(string("test")),
			"number": interface{}(float64(42)),
		}),
	}

	SetConfigName("TestWriteValueJson")
	SetConfigType("json")

	for key, value := range tests {
		writeValueJson(key, value)
	}
	defer os.Remove(ConfigName() + "." + ConfigType())

	data, err := ioutil.ReadFile(ConfigName() + "." + ConfigType())
	check(err)

	config := make(map[string]interface{})
	err = json.Unmarshal(data, &config)
	check(err)

	for key, value := range tests {
		if !reflect.DeepEqual(config[key], value) {
			t.Error("Used ", key, ", got ", config[key], ", wanted ", value)
		}
	}
}

// Benchmarks config.writeValueJson() by writing a simple
// key/value couple to an already existing config file
func BenchmarkWriteValueJson(b *testing.B) {
	SetConfigName("BenchmarkWriteValueJson")
	SetConfigType("json")

	writeValueJson("stringInit", "test")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		writeValueJson("string", "test")
		b.StopTimer()
		os.Remove(ConfigName() + "." + ConfigType())
	}
}
