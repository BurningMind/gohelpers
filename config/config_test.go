package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

// Tests config.readConfigJson() by comparing a map to
// the output of readConfigJson after writing the map as json
func TestReadConfigJson(t *testing.T) {
	// We set the config file name to be unique
	SetConfigName("TestReadConfigJson")
	SetConfigType("json")

	// We build an object to test the different JSON types
	var test = map[string]interface{}{ 
		"string":   "test",
		"number":   float64(42),
		"boolean":  true,
		"nulltype": nil,
		"array":    []interface{}{
		                "test",
		                float64(42),
		            },
		"object":   map[string]interface{}{
		                "string": "test",
		                "number": float64(42),
		            },
	}

	// We JSON-encode the test config
	data, err := json.Marshal(test)
	check(err)

	// We write the JSON test config
	err = ioutil.WriteFile(ConfigName() + "." + ConfigType(),
		                   data,
		                   0666)
	check(err)
	// We don't forget to remove the file at the end
	defer os.Remove(ConfigName() + "." + ConfigType())

	// We check the result from readConfigJson against our test object
	res := readConfigJson()
	if !reflect.DeepEqual(res, test) {
		t.Error("Got ", res, ", wanted ", test)
	}
}

// Benchmarks config.readConfigJson() by reading a simple string from json
func BenchmarkReadConfigJson(b *testing.B) {
	// We set the config file name to be unique
	SetConfigName("BenchmarkReadConfigJson")
	SetConfigType("json")

	// We write a basic JSON config file
	err := ioutil.WriteFile(ConfigName() + "." + ConfigType(),
	                        []byte(`{"string": "test"}`),
	                        0666)
	check(err)
	// We don't forget to remove the file at the end
	defer os.Remove(ConfigName() + "." + ConfigType())

	// We start the benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = readConfigJson()
	}
	b.StopTimer()
}

// Tests config.writeConfigJson() by comparing a map to
// the json map obtained from the output of writeConfigJson)
func TestWriteConfigJson(t *testing.T) {
	// We set the config file name to be unique
	SetConfigName("TestWriteConfigJson")
	SetConfigType("json")

	// We build an object to test the different JSON types
	var test = map[string]interface{}{
		"string":   "test",
		"number":   float64(42),
		"boolean":  true,
		"nulltype": nil,
		"array":    []interface{}{
		                "test",
		                float64(42),
		            },
		"object":   map[string]interface{}{
		                "string": "test",
		                "number": float64(42),
		            },
	}

	// We write the config using writeConfigJson
	writeConfigJson(key, value)
	// We don't forget to remove the file at the end
	defer os.Remove(ConfigName() + "." + ConfigType())

	// We read the JSON config file
	data, err := ioutil.ReadFile(ConfigName() + "." + ConfigType())
	check(err)

	// We JSON-decode the JSON config file
	config := make(map[string]interface{})
	err = json.Unmarshal(data, &config)
	check(err)

	// We check the object obtained via writeConfigJson against our test object
	if !reflect.DeepEqual(config, test) {
		t.Error("Got ", config, ", wanted ", test)
	}
}

// Benchmarks config.writeConfigJson() by writing a simple
// key/value couple to an already existing config file
func BenchmarkWriteConfigJson(b *testing.B) {
	// We set the config file name to be unique
	SetConfigName("BenchmarkWriteValueJson")
	SetConfigType("json")

	// We start the benchmark
	b.StopTimer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// We initialize the config file for it to be non-empty
		writeConfigJson(map[string]interface{}{ "string": "test" })

		// We write the config
		b.StartTimer()
		writeConfigJson()
		b.StopTimer()

		// We remove the config file
		os.Remove(ConfigName() + "." + ConfigType())
	}
}
