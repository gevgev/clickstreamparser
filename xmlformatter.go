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

func saveXmlToFile(xmlStr []byte) error {
	err := ioutil.WriteFile("events.xml", xmlStr, 0644)
	return err
}
