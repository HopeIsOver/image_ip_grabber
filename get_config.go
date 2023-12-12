package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func Get_config() map[string]interface{} {
	jsonFile, err := os.Open("./config.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}

	json.Unmarshal([]byte(byteValue), &result)

	return result
}
