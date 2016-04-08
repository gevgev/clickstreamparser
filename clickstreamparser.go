package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	test_KEY_A = "4144287C7000AB5ADBF2B0E50A"
	test_KEY_B = "424427ABE800F70B0C4D6F746F722053706F727473030B4D6F746F7273706F7274733164646464643164306464306430303164646464646464646464303030303030303130303030303030303030303030303030608C0A"
	test_KEY_C = "43442878E201F82B5742AE4741000713AF3B0A"
	test_KEY_D = "44287C448C4D04DA04D60000000000DDA20A"
	test_KEY_E = "45442852144444535800585E0F44277150FF7536E8B2C3EF58585300000064680A"
	//test_KEY_F = "46?"
	test_KEY_G = "4744287C560C5050503A504149442D504944D8450A"
	//						^^
	test_KEY_H  = "4844287C6B47486A7926D244286060FAD50A"
	test_KEY_H1 = "4844287C4C4D04DA04D60000000000DDA20A"
	test_KEY_H2 = "4844287C7C51000003A20000000000F36B0A"
	test_KEY_H3 = "4844287C7044AB5ADF000000000000AE8A0A"
	test_KEY_H4 = "48442878D74F007200000000000000A3990A"
	test_KEY_H5 = "4844287C554234E74E417A0000000069AC0A"
	test_KEY_H6 = "4844287C5D4100000100000000000077BA0A"

	test_KEY_I = "4944287C545600EBE822D55B0A"
	//test_KEY_J = "4A?"
	test_KEY_K = "4B44287C4D35DE6D0A"
	//test_KEY_L = "4C?"
	test_KEY_M = "4D4428629E4500140161EBA10A"
	test_KEY_P = "50442877A600000002250A"
	test_KEY_S = "5344287C58F8E211EF930A"
	test_KEY_V = "5644287C5600000000EBE8220656FFFFD7460A"
)

type Command string

const (
	R_AD        Command = "41"
	R_BtnCnfg   Command = "42"
	R_ChanVrb   Command = "43"
	R_STATE     Command = "53"
	R_HIGHLIGHT Command = "48"
	R_INFO      Command = "49"
	R_VIDEO     Command = "56"
)

var answers = []string{
	/*	test_KEY_A,
		test_KEY_B,
		test_KEY_C,
		test_KEY_S,
		test_KEY_I,
		test_KEY_H,
		test_KEY_H1,
		test_KEY_H2,
		test_KEY_H3,
		test_KEY_H4,
		test_KEY_H5,
		test_KEY_H6, */
	test_KEY_V,
}

func GetNextCommand() string {
	//return "4100112233445566778899AABBCCDDEEFF"
	return answers[rand.Intn(len(answers))]
}

func CheckCommand(clickString string) Command {
	return Command(clickString[0:2])
}

func main() {
	rand.Seed(int64(time.Now().Second()))

	for i := 0; i < 20; i++ {
		clickString := GetNextCommand()
		fmt.Println("-----------------------------------------------")
		fmt.Println("Got: ", clickString)
		switch CheckCommand(clickString) {
		case R_AD:
			adEvent := NewAdEvent(clickString)
			fmt.Printf("Ad event: %s\n", adEvent)
			fmt.Println("Diagnostics: ", adEvent.BaseEvent.Diagnostic())
			fmt.Println(adEvent.Command,
				adEvent.Timestamp,
				adEvent.AdType,
				adEvent.AdId,
				adEvent.Serial,
				adEvent.Checksum,
				adEvent.Linefeed)
		case R_BtnCnfg:
			btcnfgEvent := NewButtonConfigEvent(clickString)
			fmt.Printf("Button Config event: %s\n", btcnfgEvent)
			fmt.Println("Diagnostics: ", btcnfgEvent.BaseEvent.Diagnostic())
			fmt.Println(btcnfgEvent.Command,
				btcnfgEvent.Timestamp,
				btcnfgEvent.ButtonId,
				btcnfgEvent.ButtonType,
				btcnfgEvent.ButtonText,
				btcnfgEvent.ButtonVarData,
				btcnfgEvent.Serial,
				btcnfgEvent.Checksum,
				btcnfgEvent.Linefeed)
		case R_ChanVrb:
			channelchange := NewChannelChangeVerboseEvent(clickString)
			fmt.Printf("Channel change event: %s\n", channelchange)
			fmt.Println("Diagnostics: ", channelchange.BaseEvent.Diagnostic())
			fmt.Println(channelchange.Command,
				channelchange.Timestamp,
				channelchange.Channel,
				channelchange.SourseId,
				channelchange.ProgramId,
				channelchange.Auth,
				channelchange.TunerInfo,
				channelchange.PreviousState,
				channelchange.LastKey,
				channelchange.Serial,
				channelchange.Checksum,
				channelchange.Linefeed)
		case R_STATE:
			statechange := NewStateEvent(clickString)
			fmt.Printf("State event: %s\n", statechange)
			fmt.Println("Diagnostics: ", statechange.BaseEvent.Diagnostic())
			fmt.Println(statechange.Command,
				statechange.State,
				statechange.PreviousState,
				statechange.LastKey)
		case R_INFO:
			info := NewInfoScreenEvent(clickString)
			fmt.Printf("State event: %s\n", info)
			fmt.Println("Diagnostics: ", info.BaseEvent.Diagnostic())
			fmt.Println(info.Command,
				info.Type,
				info.Id)
		case R_HIGHLIGHT:
			hilit := NewHighlightEvent(clickString)
			fmt.Printf("Highlight event: %s\n", hilit)
			fmt.Println("Diagnostics: ", hilit.BaseEvent.Diagnostic())
			fmt.Println(hilit.Command,
				hilit.Type,
				hilit.IdFieldsStr)
		case R_VIDEO:
			video := NewVideoPlaybackEvent(clickString)
			fmt.Printf("Video event: %s\n", video)
			fmt.Println("Diagnostics: ", video.BaseEvent.Diagnostic())
			fmt.Println(video.Id,
				video.VodPlaybackMode,
				video.Source,
				video.PlayBackPosition)
		}
	}
}
