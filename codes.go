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
	{"47", "R_VODCAT"},
	{"48", "R_HIGHLIGHT"},
	{"49", "R_INFO"},
	{"4B", "R_KEY"},
	{"4D", "R_MISSING"},
	{"4F", "R_OPTION"},
	{"50", "R_PULSE"},
	{"52", "R_RESET"},
	{"53", "R_STATE"},
	{"54", "R_TURBO"},
	{"56", "R_VIDEO"},
	{"55", "R_UNIT"},
}

var EventCodes map[string]string

func init() {
	EventCodes = make(map[string]string)

	for _, event := range eventCodes {
		EventCodes[event.Code] = event.Text
	}
}
