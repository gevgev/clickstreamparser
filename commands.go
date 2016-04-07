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
	bt.Command = convertToString(clickString[0:2])
	bt.Timestamp = convertToTime(clickString[2:10])
	bt.Serial = clickString[len(clickString)-6 : len(clickString)-4]
	bt.Checksum = clickString[len(clickString)-4 : len(clickString)-2]
	bt.Linefeed = clickString[len(clickString)-2 : len(clickString)]
	return bt
}

func (bt BaseEvent) String() string {
	return fmt.Sprintf("Command:[%s]\tTimestamp:[%s]", bt.Command, bt.Timestamp)
}

func (bt BaseEvent) Diagnostic() string {
	return fmt.Sprintf("Serial:[%s]\tChecksum:[%s]\tLinefeed:[%s]", bt.Serial, bt.Checksum, bt.Linefeed)
}

// ---------- Ad: A, 41 --------------------
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

// ---------- Button Config: B, 42 --------------------
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

// ---------- Channel Change Verbose: C, 43 --------------------

type ChannelChangeVerboseEvent struct {
	*BaseEvent
	Channel       string
	SourseId      string
	ProgramId     string
	Auth          string
	TunerInfo     string
	PreviousState string
	LastKey       string
}

func NewChannelChangeVerboseEvent(clickString string) *ChannelChangeVerboseEvent {
	channelchange := new(ChannelChangeVerboseEvent)
	channelchange.BaseEvent = NewBaseEvent(clickString)
	// C   time     chN  src prgmId A  TI PS LK
	// 43 442878E2 01F8 2B57 42AE47 41 00 07 13 AF 3B 0A

	channelchange.Channel = clickString[10:14]
	channelchange.SourseId = clickString[14:18]
	channelchange.ProgramId = clickString[18:24]
	channelchange.Auth = convertToString(clickString[24:26])
	channelchange.TunerInfo = clickString[26:28]
	channelchange.PreviousState = clickString[28:30]
	channelchange.LastKey = clickString[30:32]

	return channelchange
}

func (channelchange ChannelChangeVerboseEvent) String() string {
	return fmt.Sprintf("[%s]\tChannel:[%s]\tSourseId:[%s]\tProgramId:[%s]\tAuth:[%s]\tTuner Info:[%s]\tPrevious State:[%s]\tLast Key:[%s]",
		channelchange.BaseEvent,
		channelchange.Channel,
		channelchange.SourseId,
		channelchange.ProgramId,
		channelchange.Auth,
		channelchange.TunerInfo,
		channelchange.PreviousState,
		channelchange.LastKey)

}

// ---------- State: S, 53 --------------------

type StateEvent struct {
	*BaseEvent
	State         string
	PreviousState string
	LastKey       string
}

func NewStateEvent(clickString string) *StateEvent {
	stateEvent := new(StateEvent)
	stateEvent.BaseEvent = NewBaseEvent(clickString)
	//  S   time    S PS LK
	// 53 44287C58 F8 E2 11 EF 93 0A
	stateEvent.State = clickString[10:12]
	stateEvent.PreviousState = clickString[12:14]
	stateEvent.LastKey = clickString[14:16]

	return stateEvent
}

func (statechange StateEvent) String() string {
	return fmt.Sprintf("[%s]\tState:[%s]\tPrevious State:[%s]\tLast Key:[%s]",
		statechange.BaseEvent,
		statechange.State,
		statechange.PreviousState,
		statechange.LastKey)

}
