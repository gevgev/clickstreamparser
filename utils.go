package main

import (
	"encoding/hex"
	"path/filepath"
	"strconv"
	"time"
)

func convertToTime(timestampS string) time.Time {
	timestamp, err := strconv.ParseInt(timestampS, 16, 64)
	//fmt.Println(timestampS, timestamp)
	if err == nil {
		timestamp += UTC_GPS_Diff
		//fmt.Println(timestampS, timestamp)

		t := time.Unix(timestamp, 0)
		return t
	}
	//else {
	//fmt.Println("Error:", err)
	//}
	return time.Time{}
}

func convertToInt(intS string) int64 {
	val, err := strconv.ParseInt(intS, 16, 64)
	//fmt.Println(timestampS, timestamp)
	if err == nil {
		return val
	}

	return 0
}

func convertToString(str string) string {
	bytes, err := hex.DecodeString(str)
	if err == nil {
		return string(bytes)
	}
	return ""
}

func lookUpKeyName(keyCode int) string {
	return KeyName[keyCode]
}

func lookUpEventName(code string) string {
	return EventCodes[code]
}

func addProperExtension(fileName string) string {
	switch outputFormat {
	case xmlOutput:
		fileName = fileName + "." + xmlOutput
	case jsonOutput:
		fileName = fileName + "." + jsonOutput
	case txtOutput:
		fileName = fileName + "." + txtOutput
	}
	return fileName
}

func validateOutFileName(fileName string) string {
	// Check if it has extension
	// If not, add the default extension
	ext := filepath.Ext(fileName)
	if ext != "" {
		if isRawFile(fileName) {
			fileName = fileName[:len(fileName)-len(ext)]
			fileName = addProperExtension(fileName)
		}
	} else if ext == "" {
		fileName = addProperExtension(fileName)
	}

	return fileName
}

func isRawFile(fileName string) bool {
	return filepath.Ext(fileName) == "."+inExtension
}

func getFileType(fileName string) FileType {
	ext := filepath.Ext(fileName)
	switch ext {
	case "." + csExt:
		return FT_CS
	case "." + rawExt:
		return FT_RAW
	case "." +csPayload:
		return FT_PAYLOAD
	}
	return FT_WRONG
}
