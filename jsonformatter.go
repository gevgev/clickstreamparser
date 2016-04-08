package main

import (
	"encoding/json"
	"io/ioutil"
)

func generateJson(events []interface{}) ([]byte, error) {
	b, err := json.MarshalIndent(events, "", "    ")
	if err != nil {
		return []byte{}, err
	}

	return b, nil
}

func saveJsonToFile(jsonStr []byte) error {
	err := ioutil.WriteFile("events.json", jsonStr, 0644)
	return err
}
