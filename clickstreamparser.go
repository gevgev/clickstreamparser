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
	//test_KEY_K = array of samples below
	//test_KEY_L = "4C?"
	test_KEY_M  = "4D4428629E4500140161EBA10A"
	test_KEY_P  = "50442877A600000002250A"
	test_KEY_S  = "5344287C58F8E211EF930A"
	test_KEY_U  = "55442877A600059CAA293233322E343400020000000000000000000000000000000000000000000000007F0A"
	test_KEY_U1 = "55442841ED000038FD833233322E343400020000000000000000000000000000000000000000000000002A0A"
	test_KEY_U2 = "554428839D00008B8D723233322E34340002000000000000000000000000000000000000000000000000660A"
	test_KEY_V  = "5644287C5600000000EBE8220656FFFFD7460A"
)

var test_KEY_K = [...]string{"4B44287C4D00DE6D0A",
	"4B44287C4D01DE6D0A",
	"4B44287C4D02DE6D0A",
	"4B44287C4D0ADE6D0A",
	"4B44287C4D0BDE6D0A",
	"4B44287C4D0CDE6D0A",
	"4B44287C4D0DDE6D0A",
	"4B44287C4D0EDE6D0A",
	"4B44287C4D0FDE6D0A",
	"4B44287C4D10DE6D0A",
	"4B44287C4D11DE6D0A",
	"4B44287C4D12DE6D0A",
	"4B44287C4D13DE6D0A",
	"4B44287C4D14DE6D0A",
	"4B44287C4D15DE6D0A",
	"4B44287C4D16DE6D0A",
	"4B44287C4D17DE6D0A",
	"4B44287C4D18DE6D0A",
	"4B44287C4D19DE6D0A",
	"4B44287C4D1ADE6D0A",
	"4B44287C4D1BDE6D0A",
	"4B44287C4D1CDE6D0A",
	"4B44287C4D1DDE6D0A",
	"4B44287C4D1EDE6D0A",
	"4B44287C4D1FDE6D0A",
	"4B44287C4D20DE6D0A",
	"4B44287C4D21DE6D0A",
	"4B44287C4D22DE6D0A",
	"4B44287C4D23DE6D0A",
	"4B44287C4D24DE6D0A",
	"4B44287C4D25DE6D0A",
	"4B44287C4D26DE6D0A",
	"4B44287C4D27DE6D0A",
	"4B44287C4D28DE6D0A",
	"4B44287C4D29DE6D0A",
	"4B44287C4D2ADE6D0A",
	"4B44287C4D2BDE6D0A",
	"4B44287C4D2CDE6D0A",
	"4B44287C4D2DDE6D0A",
	"4B44287C4D2EDE6D0A",
	"4B44287C4D2FDE6D0A",
	"4B44287C4D30DE6D0A",
	"4B44287C4D31DE6D0A",
	"4B44287C4D32DE6D0A",
	"4B44287C4D33DE6D0A",
	"4B44287C4D34DE6D0A",
	"4B44287C4D35DE6D0A",
	"4B44287C4D36DE6D0A",
	"4B44287C4D37DE6D0A",
	"4B44287C4D38DE6D0A",
	"4B44287C4D39DE6D0A",
	"4B44287C4D3ADE6D0A",
	"4B44287C4D3BDE6D0A",
	"4B44287C4D3CDE6D0A",
	"4B44287C4D3DDE6D0A",
	"4B44287C4D3EDE6D0A",
	"4B44287C4D3FDE6D0A",
	"4B44287C4D40DE6D0A",
	"4B44287C4D41DE6D0A",
	"4B44287C4D42DE6D0A",
	"4B44287C4D43DE6D0A",
	"4B44287C4D44DE6D0A",
	"4B44287C4D45DE6D0A",
	"4B44287C4D46DE6D0A",
	"4B44287C4D47DE6D0A",
	"4B44287C4D48DE6D0A",
	"4B44287C4D49DE6D0A",
	"4B44287C4D4ADE6D0A",
	"4B44287C4D4BDE6D0A",
	"4B44287C4D4CDE6D0A",
	"4B44287C4D4DDE6D0A",
	"4B44287C4D4EDE6D0A",
	"4B44287C4D4FDE6D0A",
}

type Command string

const (
	R_AD        Command = "41"
	R_BtnCnfg   Command = "42"
	R_ChanVrb   Command = "43"
	R_STATE     Command = "53"
	R_HIGHLIGHT Command = "48"
	R_INFO      Command = "49"
	R_VIDEO     Command = "56"
	R_KEY       Command = "4B"
	R_UNIT      Command = "55"
)

var answers = []string{
	test_KEY_A,
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
	test_KEY_H6,
	test_KEY_V,
	test_KEY_U,
	test_KEY_U1,
	test_KEY_U2,
}

func GetNextCommand() string {
	//return "4100112233445566778899AABBCCDDEEFF"
	if (rand.Intn(2)) == 1 {
		return answers[rand.Intn(len(answers))]
	} else {
		return test_KEY_K[rand.Intn(len(test_KEY_K))]
	}
}

func CheckCommand(clickString string) Command {
	return Command(clickString[0:2])
}

func main() {
	rand.Seed(int64(time.Now().Second()))

	diagnostics := false

	eventsCollection := []interface{}{}

	for i := 0; i < 100; i++ {
		clickString := GetNextCommand()
		//fmt.Println("-----------------------------------------------")
		//fmt.Println("Got: ", clickString)
		fmt.Printf("Device Id: 0000008B8D72 ")
		switch CheckCommand(clickString) {
		case R_AD:
			adEvent := NewAdEvent(clickString)
			fmt.Println(adEvent)
			eventsCollection = append(eventsCollection, adEvent)
			if diagnostics {
				fmt.Println("Diagnostics: ", adEvent.BaseEvent.Diagnostic())

				fmt.Println(adEvent.Command,
					adEvent.Timestamp,
					adEvent.AdType,
					adEvent.AdId,
					adEvent.Serial,
					adEvent.Checksum,
					adEvent.Linefeed)
			}
		case R_BtnCnfg:
			btcnfgEvent := NewButtonConfigEvent(clickString)
			fmt.Println(btcnfgEvent)
			eventsCollection = append(eventsCollection, btcnfgEvent)
			if diagnostics {
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
			}
		case R_ChanVrb:
			channelchange := NewChannelChangeVerboseEvent(clickString)
			fmt.Println(channelchange)
			eventsCollection = append(eventsCollection, channelchange)
			if diagnostics {
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
			}
		case R_STATE:
			statechange := NewStateEvent(clickString)
			fmt.Println(statechange)
			eventsCollection = append(eventsCollection, statechange)
			if diagnostics {
				fmt.Println("Diagnostics: ", statechange.BaseEvent.Diagnostic())
				fmt.Println(statechange.Command,
					statechange.State,
					statechange.PreviousState,
					statechange.LastKey)
			}
		case R_INFO:
			info := NewInfoScreenEvent(clickString)
			fmt.Println(info)
			eventsCollection = append(eventsCollection, info)
			if diagnostics {

				fmt.Println("Diagnostics: ", info.BaseEvent.Diagnostic())
				fmt.Println(info.Command,
					info.Type,
					info.Id)
			}
		case R_KEY:
			key := NewKeyPressEvent(clickString)
			fmt.Println(key)
			eventsCollection = append(eventsCollection, key)
			if diagnostics {

				fmt.Println("Diagnostics: ", key.BaseEvent.Diagnostic())
				fmt.Println(key.Command,
					key.KeyCode)
			}
		case R_HIGHLIGHT:
			hilit := NewHighlightEvent(clickString)
			fmt.Println(hilit)
			eventsCollection = append(eventsCollection, hilit)
			if diagnostics {

				fmt.Println("Diagnostics: ", hilit.BaseEvent.Diagnostic())
				fmt.Println(hilit.Command,
					hilit.Type,
					hilit.IdFieldsStr)
			}
		case R_VIDEO:
			video := NewVideoPlaybackEvent(clickString)
			fmt.Println(video)
			eventsCollection = append(eventsCollection, video)
			if diagnostics {

				fmt.Println("Diagnostics: ", video.BaseEvent.Diagnostic())
				fmt.Println(video.Id,
					video.VodPlaybackMode,
					video.Source,
					video.PlayBackPosition)
			}
		case R_UNIT:
			unit := NewUnitIdentificationEvent(clickString)
			fmt.Println(unit)
			eventsCollection = append(eventsCollection, unit)
			if diagnostics {

				fmt.Println("Diagnostics: ", unit.BaseEvent.Diagnostic())
				fmt.Println(unit.PeriodicReports,
					unit.PollingReports,
					unit.HighWaterMarkReports,
					unit.BlackoutOverflowReports,
					unit.ExceededMaxReportsPerHour,
					unit.UsedBufferSize,
					unit.GuideState,
					unit.TunerInfo,
					unit.SourceIdTuner0,
					unit.SourceIdTuner1)
			}

		}
	}

	jsonString, err := generateJson(eventsCollection)
	if diagnostics {
		fmt.Println(string(jsonString))
	}
	if err == nil {
		err = saveToFile(jsonString)
		if err != nil {
			fmt.Println("Error writing Json file:", err)
		}
	} else {
		fmt.Println(err)
	}
}
