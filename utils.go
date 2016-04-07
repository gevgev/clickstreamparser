package main

import (
	"encoding/hex"
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
