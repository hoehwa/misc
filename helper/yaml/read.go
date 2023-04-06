package yaml

import (
	"os"

	"gopkg.in/yaml.v2"
)

func ReadOneDemensionalYaml(path string, filename string) ([]string, []string) {
	// If file or directory doesn't exist, make one.
	if _, err := os.Stat(path + filename); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
		os.Create(path + filename)
	}
  // read the YAML file
  yamlFile, err := os.ReadFile(path + filename)
  if err != nil {
    panic(err)
  }

  // define a map to store the key-value pairs
  data := make(map[string]string)

  // unmarshal the YAML data into the map
  err = yaml.Unmarshal(yamlFile, &data)
  if err != nil {
    panic(err)
  }

	var keys []string
	var values []string
  // loop through the map and print the key-value pairs
  for key, value := range data {
    // fmt.Printf("Key: %s, Value: %s\n", key, value)
		keys = append(keys, key)
		values = append(values, value)
  }
	return keys, values
}