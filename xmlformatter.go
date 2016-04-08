package main

import (
	"encoding/xml"
	"io/ioutil"
)

func generateXml(events []interface{}) ([]byte, error) {
	b, err := xml.MarshalIndent(events, "", "    ")
	if err != nil {
		return []byte{}, err
	}

	return b, nil
}

func saveXmlToFile(fileName string, xmlStr []byte) error {
	err := ioutil.WriteFile(fileName, xmlStr, 0644)
	return err
}
