package main

import (
	"fmt"
	"time"
)

const (
	UTC_GPS_Diff = 315964800
)

type BaseEvent struct {
	Command   string
	Timestamp time.Time
	Serial    string
	Checksum  string
	Linefeed  string

	//	Command   [2]string
	//	Timestamp [8]string
	//	Serial    [1]string
	//	Checksum  [1]string
	//	Linefeed  [1]string
}

func NewBaseEvent(clickString string) *BaseEvent {
	bt := new(BaseEvent)
	// A	time    ... s-n c-s LF
	// 41 44287C70  ...  B0  E5  0A
	bt.Command = clickString[0:2]
	bt.Timestamp = convertToTime(clickString[2:10])
	bt.Serial = clickString[len(clickString)-6 : len(clickString)-4]
	bt.Checksum = clickString[len(clickString)-4 : len(clickString)-2]
	bt.Linefeed = clickString[len(clickString)-2 : len(clickString)]
	return bt
}

func (bt BaseEvent) String() string {
	return fmt.Sprintf("C:[%s]\tTimestamp:[%s]", bt.Command, bt.Timestamp)
}

func (bt BaseEvent) Diagnostic() string {
	return fmt.Sprintf("Serial:[%s]\tChecksum:[%s]\tLinefeed:[%s]", bt.Serial, bt.Checksum, bt.Linefeed)
}

// ---------- Ad --------------------
type AdEvent struct {
	*BaseEvent
	AdType string
	AdId   string
	//	AdType    [1]string
	//	AdId      [4]string
}

func NewAdEvent(clickString string) *AdEvent {
	at := new(AdEvent)
	at.BaseEvent = NewBaseEvent(clickString)
	// A	time  type  adId   s-n c-s LF
	// 41 44287C70 00 AB5ADBF2 B0  E5  0A
	at.AdType = clickString[10:12]
	at.AdId = clickString[12:20]
	return at
}

func (at AdEvent) String() string {
	return fmt.Sprintf("[%s]\tAdType:[%s]\tAdId:[%s]", at.BaseEvent, at.AdType, at.AdId)
}

// ---------- Button Config --------------------
type ButtonConfigEvent struct {
	*BaseEvent
	ButtonId      string
	ButtonType    string
	ButtonText    string
	ButtonVarData string
}

func NewButtonConfigEvent(clickString string) *ButtonConfigEvent {
	btcf := new(ButtonConfigEvent)
	btcf.BaseEvent = NewBaseEvent(clickString)
	//
	//"42 4427ABE8 00F7 0B 0C 4D6F746F722053706F727473 030B4D6F746F7273706F7274733164646464643164306464306430303164646464646464646464303030303030303130303030303030303030303030303030 60 8C 0A"
	btcf.ButtonId = clickString[10:14]
	btcf.ButtonType = clickString[14:16]

	buttonTextLenght := convertToInt(clickString[16:18])
	btcf.ButtonText = convertToString(clickString[18 : 18+buttonTextLenght*2])
	btcf.ButtonVarData = clickString[18+buttonTextLenght*2 : len(clickString)-6]

	return btcf
}

func (btcf ButtonConfigEvent) String() string {
	return fmt.Sprintf("[%s]\tButtonId:[%s]\tButtonType:[%s]\tText:[%s]\tData:[%s]",
		btcf.BaseEvent, btcf.ButtonId, btcf.ButtonType, btcf.ButtonText, btcf.ButtonVarData)
}

// ---------- Channel Change Verbose--------------------
