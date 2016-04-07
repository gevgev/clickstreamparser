package main

import (
	"fmt"
)

const (
	test_KEY_A = "4144287C7000AB5ADBF2B0E50A"
	test_KEY_B = "424427ABE800F70B0C4D6F746F722053706F727473030B4D6F746F7273706F7274733164646464643164306464306430303164646464646464646464303030303030303130303030303030303030303030303030608C0A"
	test_KEY_C = "43442878E201F82B5742AE4741000713AF3B0A"
	test_KEY_D = "44287C448C4D04DA04D60000000000DDA20A"
	test_KEY_E = "45442852144444535800585E0F44277150FF7536E8B2C3EF58585300000064680A"
	//test_KEY_F = "46?"
	test_KEY_G = "4744287C560C5050503A504149442D504944D8450A"
	test_KEY_H = "4844287C6B47486A7926D244286060FAD50A"
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
	R_AD      Command = "41"
	R_BtnCnfg Command = "42"
	R_Chan    Command = "43"
)

func GetNextCommand() string {
	//return "4100112233445566778899AABBCCDDEEFF"
	return test_KEY_B
}

func CheckCommand(clickString string) Command {
	return Command(clickString[0:2])
}

func main() {
	clickString := GetNextCommand()
	fmt.Println("Got: ", clickString)
	switch CheckCommand(clickString) {
	case R_AD:
		adEvent := NewAdEvent(clickString)
		fmt.Printf("Ad event: [%s]\n", adEvent)
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
		fmt.Printf("Button Config event: [%s]\n", btcnfgEvent)
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
}
