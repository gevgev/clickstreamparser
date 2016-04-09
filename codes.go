package main

type EventCode struct {
	Code string
	Text string
}

var eventCodes = [...]EventCode{
	{"41", "R_AD"},
	{"42", "R_BUTTONCONFIG"},
	{"43", "R_CHANGECHANNEL"},
	{"45", "R_PROGRAMEVENT"},
	{"53", "R_STATE"},
	{"54", "R_TURBO"},
	{"48", "R_HIGHLIGHT"},
	{"49", "R_INFO"},
	{"4F", "R_OPTION"},
	{"56", "R_VIDEO"},
	{"4B", "R_KEY"},
	{"55", "R_UNIT"},
}

var EventCodes map[string]string

func init() {
	EventCodes = make(map[string]string)

	for _, event := range eventCodes {
		EventCodes[event.Code] = event.Text
	}
}
