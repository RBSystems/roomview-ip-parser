package helpers

import (
	"encoding/json"
	"io/ioutil"
)

func ImportConfig(file string) Configuration {
	b, err := ioutil.ReadFile(file)

	if err != nil {
		panic(err)
	}

	var config Configuration

	err = json.Unmarshal(b, &config)

	if err != nil {
		panic(err)
	}

	return config
}
