package main

import (
	"fmt"
	"time"
)

const (
	UTC_GPS_Diff = 315964800
)

// ---------- Common Base: Command Code, Timestamp, Serial, Checksum, Linefeed --------------------
type BaseEvent struct {
	Command   string
	Timestamp time.Time
	Serial    string
	Checksum  string
	Linefeed  string
	DeviceId  string

	//	Command   [2]string
	//	Timestamp [8]string
	//	Serial    [1]string
	//	Checksum  [1]string
	//	Linefeed  [1]string
}

func NewBaseEvent(deviceId, clickString string) *BaseEvent {
	bt := new(BaseEvent)
	// A	time    ... s-n c-s LF
	// 41 44287C70  ...  B0  E5  0A
	bt.Command = lookUpEventName(clickString[0:2])
	bt.Timestamp = convertToTime(clickString[2:10])
	bt.Serial = clickString[len(clickString)-6 : len(clickString)-4]
	bt.Checksum = clickString[len(clickString)-4 : len(clickString)-2]
	bt.Linefeed = clickString[len(clickString)-2 : len(clickString)]
	bt.DeviceId = deviceId
	return bt
}

func (bt BaseEvent) String() string {
	return fmt.Sprintf("DeviceID: [%s]\tTimestamp:[%s]\tEvent:[%16s] ", bt.DeviceId, bt.Timestamp, bt.Command)
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

func NewAdEvent(deviceId, clickString string) *AdEvent {
	at := new(AdEvent)
	at.BaseEvent = NewBaseEvent(deviceId, clickString)
	// A	time  type  adId   s-n c-s LF
	// 41 44287C70 00 AB5ADBF2 B0  E5  0A
	at.AdType = clickString[10:12]
	at.AdId = clickString[12:20]
	return at
}

func (at AdEvent) String() string {
	return fmt.Sprintf("%s\tAdType:[%s]\tAdId:[%s]", at.BaseEvent, at.AdType, at.AdId)
}

// ---------- Button Config: B, 42 --------------------
type ButtonConfigEvent struct {
	*BaseEvent
	ButtonId      string
	ButtonType    string
	ButtonText    string
	ButtonVarData string
}

func NewButtonConfigEvent(deviceId, clickString string) *ButtonConfigEvent {
	btcf := new(ButtonConfigEvent)

	btcf.BaseEvent = NewBaseEvent(deviceId, clickString)
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
	return fmt.Sprintf("%s\tButtonId:[%s]\tButtonType:[%s]\tText:[%s]\tData:[%s]",
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

func NewChannelChangeVerboseEvent(deviceId, clickString string) *ChannelChangeVerboseEvent {
	channelchange := new(ChannelChangeVerboseEvent)
	channelchange.BaseEvent = NewBaseEvent(deviceId, clickString)
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
	return fmt.Sprintf("%s\tChannel:[%s]\tSourseId:[%s]\tProgramId:[%s]\tAuth:[%s]\tTuner Info:[%s]\tPrevious State:[%s]\tLast Key:[%s]",
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

func NewStateEvent(deviceId, clickString string) *StateEvent {
	stateEvent := new(StateEvent)
	stateEvent.BaseEvent = NewBaseEvent(deviceId, clickString)
	//  S   time    S PS LK
	// 53 44287C58 F8 E2 11 EF 93 0A
	stateEvent.State = clickString[10:12]
	stateEvent.PreviousState = clickString[12:14]
	stateEvent.LastKey = clickString[14:16]

	return stateEvent
}

func (statechange StateEvent) String() string {
	return fmt.Sprintf("%s\tState:[%s]\tPrevious State:[%s]\tLast Key:[%s]",
		statechange.BaseEvent,
		statechange.State,
		statechange.PreviousState,
		statechange.LastKey)
}

// ---------- State: I, 49 --------------------

type InfoScreenEvent struct {
	*BaseEvent
	Type string
	Id   string
}

func NewInfoScreenEvent(deviceId, clickString string) *InfoScreenEvent {
	info := new(InfoScreenEvent)
	info.BaseEvent = NewBaseEvent(deviceId, clickString)
	//
	// 49 44287C54 56 00EBE822 D5 5B 0A
	info.Type = convertToString(clickString[10:12])
	info.Id = clickString[12:20]
	return info
}

func (info InfoScreenEvent) String() string {
	return fmt.Sprintf("%s\tType:[%s]\tId:[%s]",
		info.BaseEvent,
		info.Type,
		info.Id)

}

// ---------- Highlight: H, 48 --------------------

// -----------IdFields struct for Highlight -------
type IdFields struct {
	HighlightType string
	// L or B or G
	ProgramId string
	SourceId  string
	// G
	GridTime time.Time
	// M or Q
	MenuId   string
	ButtonId string
	// A
	Position string
	// A  or O
	FunctionCode string
	// S
	OptionCode  string
	OptionValue string
	// K
	KeyCode string
	// D
	AdId string
	// V
	AssetOrSourceId string
	TemplateId      string
	ObjectId        string
	// Common
	Filler string
}

func NewIdFields(hiType, clickString string) *IdFields {
	idFields := new(IdFields)
	idFields.HighlightType = hiType

	switch idFields.HighlightType {
	case "L", "B":
		idFields.ProgramId = clickString[0:6]
		idFields.SourceId = clickString[6:10]
		idFields.Filler = clickString[10:18]
	case "G":
		idFields.ProgramId = clickString[0:6]
		idFields.SourceId = clickString[6:10]
		idFields.GridTime = convertToTime(clickString[10:18])
	case "M", "Q":
		idFields.MenuId = clickString[0:4]
		idFields.ButtonId = clickString[4:8]
		idFields.Filler = clickString[8:18]
	case "A":
		idFields.ProgramId = clickString[0:2]
		idFields.FunctionCode = clickString[2:6]
		idFields.Filler = clickString[6:18]
	case "O":
		idFields.FunctionCode = clickString[0:4]
		idFields.Filler = clickString[4:18]
	case "S":
		idFields.OptionCode = clickString[0:2]
		idFields.OptionValue = clickString[2:4]
		idFields.Filler = clickString[4:18]
	case "K":
		idFields.KeyCode = clickString[0:2]
		idFields.Filler = clickString[2:18]
	case "D":
		idFields.AdId = clickString[0:6]
		idFields.Filler = clickString[6:18]
	case "V":
		idFields.AssetOrSourceId = clickString[0:8]
		idFields.TemplateId = clickString[8:12]
		idFields.ObjectId = clickString[12:16]
	}
	return idFields
}

func (idFields IdFields) String() string {
	var str string

	switch idFields.HighlightType {
	case "L", "B":
		str = fmt.Sprintf("Program Id: %s\t SourceId %s",
			idFields.ProgramId,
			idFields.SourceId)
	case "G":
		str = fmt.Sprintf("Program Id: %s\t SourceId %s \t Grid Time: %s ",
			idFields.ProgramId,
			idFields.SourceId,
			idFields.GridTime)
	case "M", "Q":
		str = fmt.Sprintf("Menu Id: %s \t Button Id: %s ",
			idFields.MenuId,
			idFields.ButtonId)
	case "A":
		str = fmt.Sprintf("Program Id: %s\t Function Code: %s",
			idFields.ProgramId,
			idFields.FunctionCode)
	case "O":
		str = fmt.Sprintf("Function Code: %s",
			idFields.FunctionCode)
	case "S":
		str = fmt.Sprintf("Option Code: %s\t Option Value: %s",
			idFields.OptionCode,
			idFields.OptionValue)
	case "K":
		str = fmt.Sprintf("Key Code: %s\t",
			idFields.KeyCode)
	case "D":
		str = fmt.Sprintf("Ad Id: %s\t",
			idFields.AdId)
	case "V":
		str = fmt.Sprintf("Asset/Source Id: %s\t Template Id: %s\t Object Id: %s",
			idFields.AssetOrSourceId,
			idFields.TemplateId,
			idFields.ObjectId)
	}
	return str
}

// ---------- Highlight: H, 48 --------------------
type HighlightEvent struct {
	*BaseEvent
	Type        string
	IdFieldsStr string
	*IdFields
}

func NewHighlightEvent(deviceId, clickString string) *HighlightEvent {
	hilit := new(HighlightEvent)
	hilit.BaseEvent = NewBaseEvent(deviceId, clickString)
	//
	// 48 44287C6B 47 486A7926D244286060 FA D5 0A
	hilit.Type = convertToString(clickString[10:12])
	hilit.IdFieldsStr = clickString[12:30]
	hilit.IdFields = NewIdFields(hilit.Type, hilit.IdFieldsStr)
	return hilit
}

func (hilit HighlightEvent) String() string {
	return fmt.Sprintf("%s\tType:[%s]\tIdFileds:[%s]",
		hilit.BaseEvent,
		hilit.Type,
		hilit.IdFields)
}

// ---------- Video Playback: V, 56 --------------------

type VideoPlaybackEvent struct {
	*BaseEvent
	Id               string
	VodPlaybackMode  string
	Source           string
	PlayBackPosition string
}

func NewVideoPlaybackEvent(deviceId, clickString string) *VideoPlaybackEvent {
	video := new(VideoPlaybackEvent)
	video.BaseEvent = NewBaseEvent(deviceId, clickString)

	video.Id = clickString[10:24]
	video.VodPlaybackMode = clickString[24:26]
	video.Source = convertToString(clickString[26:28])
	video.PlayBackPosition = clickString[28:32]

	return video
}

func (video VideoPlaybackEvent) String() string {
	return fmt.Sprintf("%s\tId:[%s]\tVOD Playback Mode:[%s]\t Source [%s]\t Playback Position: [%s]",
		video.BaseEvent,
		video.Id,
		video.VodPlaybackMode,
		video.Source,
		video.PlayBackPosition)
}

// ---------- Key Press: K, 4B --------------------

type KeyPressEvent struct {
	*BaseEvent
	Key     string
	KeyCode int
}

func NewKeyPressEvent(deviceId, clickString string) *KeyPressEvent {
	key := new(KeyPressEvent)
	key.BaseEvent = NewBaseEvent(deviceId, clickString)

	key.KeyCode = int(convertToInt(clickString[10:12]))
	key.Key = lookUpKeyName(key.KeyCode)

	return key
}

func (key KeyPressEvent) String() string {
	return fmt.Sprintf("%s\tKey:[%s]",
		key.BaseEvent,
		key.Key)
}

// ---------- Unit Identification: U, 55 --------------------

type UnitIdentificationEvent struct {
	*BaseEvent
	UnitAddress               string
	HardwareType              string
	SoftwareVersion           string
	Profile                   string
	PeriodicReports           string
	PollingReports            string
	HighWaterMarkReports      string
	BlackoutOverflowReports   string
	ExceededMaxReportsPerHour string
	UsedBufferSize            string
	GuideState                string
	TunerInfo                 string
	SourceIdTuner0            string
	SourceIdTuner1            string
}

func NewUnitIdentificationEvent(deviceId, clickString string) *UnitIdentificationEvent {
	unit := new(UnitIdentificationEvent)
	unit.BaseEvent = NewBaseEvent(deviceId, clickString)
	//																					????
	// 55 442877A6 00059CAA29 32 33322E343400 0200 00 00 00 00 00 00 00 00 0000 0000 00000000000000000000 00 7F 0A
	// 55 4428839D 00008B8D72 32 33322E343400 0200 00 00 00 00 00 00 00 00 0000 0000 00000000000000000000 00 66 0A
	unit.UnitAddress = clickString[10:20]
	unit.HardwareType = clickString[20:22]
	unit.SoftwareVersion = convertToString(clickString[22:34])
	unit.Profile = clickString[34:38]
	unit.PeriodicReports = clickString[38:40]
	unit.PollingReports = clickString[40:42]
	unit.HighWaterMarkReports = clickString[42:44]
	unit.BlackoutOverflowReports = clickString[44:46]
	unit.ExceededMaxReportsPerHour = clickString[46:48]
	unit.UsedBufferSize = clickString[48:50]
	unit.GuideState = clickString[50:52]
	unit.TunerInfo = clickString[52:54]
	unit.SourceIdTuner0 = clickString[54:58]
	unit.SourceIdTuner1 = clickString[58:62]

	return unit
}

func (unit UnitIdentificationEvent) String() string {
	return fmt.Sprintf("%s\tAddr:[%s]\tHardware:[%s]\tSoftware:[%s]\tProfile:[%s]",
		unit.BaseEvent,
		unit.UnitAddress,
		unit.HardwareType,
		unit.SoftwareVersion,
		unit.Profile)
}

// ---------- VOD Category: G, 47 --------------------

type VodCategoryEvent struct {
	*BaseEvent
	StringLength int
	Str          string
}

func NewVodCategoryEvent(deviceId, clickString string) *VodCategoryEvent {
	vodCat := new(VodCategoryEvent)
	vodCat.BaseEvent = NewBaseEvent(deviceId, clickString)

	vodCat.StringLength = int(convertToInt(clickString[10:12]))
	vodCat.Str = clickString[12 : 12+vodCat.StringLength*2]

	return vodCat
}

func (vodCat VodCategoryEvent) String() string {
	return fmt.Sprintf("%s\tVOD Category Name:[%s]",
		vodCat.BaseEvent,
		vodCat.Str)
}

// -------------- Event E, 45 --------------------
type ProgramEventEvent struct {
	*BaseEvent
	EventType       string
	DataSource      string
	EventRecurrence string
	EventAction     string
	EventTuner      int
	TunerSelection  string
	SourceID        string
	EventDateTime   time.Time
	EventDays       string
	EventProgramID  string
	EventSeriesID   string
	EpisodeType     string
	SaveNoMoreThan  string
	SaveUntil       string
	StartOffset     int
	EndOffset       int
	Length          int
	SearchString    string
}

func NewProgramEventEvent(deviceId, clickString string) *ProgramEventEvent {
	event := new(ProgramEventEvent)
	event.BaseEvent = NewBaseEvent(deviceId, clickString)

	event.EventType = convertToString(clickString[10:12])
	event.DataSource = convertToString(clickString[12:14])
	event.EventRecurrence = convertToString(clickString[14:16])
	event.EventAction = convertToString(clickString[16:18])
	event.EventTuner = int(convertToInt(clickString[18:20]))
	event.TunerSelection = convertToString(clickString[20:22])
	event.SourceID = clickString[22:26]
	event.EventDateTime = convertToTime(clickString[26:34])
	event.EventDays = clickString[34:36]
	event.EventProgramID = clickString[36:42]
	event.EventSeriesID = clickString[42:48]
	event.EpisodeType = convertToString(clickString[48:50])
	event.SaveNoMoreThan = clickString[50:52] // complex definition
	event.SaveUntil = convertToString(clickString[52:54])
	event.StartOffset = int(convertToInt(clickString[54:56]))
	event.EndOffset = int(convertToInt(clickString[56:58]))
	event.Length = int(convertToInt(clickString[58:60]))
	event.SearchString = convertToString(clickString[60 : 60+event.Length*2])

	return event
}

func (event ProgramEventEvent) String() string {
	return fmt.Sprintf("%s\t EventType[%s] DataSource[%s] Recurrence[%s] Action[%s] Tuner[%d] TunerSelection[%s] SourceID[%s] DateTime[%v] Days[%s] ProgramID[%s] SeriesID[%s] Type[%s] SaveNoMoreThan[%s] SaveUntil[%s] StartOffset[%d] EndOffset[%d] SearchString[%s]",
		event.BaseEvent,
		event.EventType,
		event.DataSource,
		event.EventRecurrence,
		event.EventAction,
		event.EventTuner,
		event.TunerSelection,
		event.SourceID,
		event.EventDateTime,
		event.EventDays,
		event.EventProgramID,
		event.EventSeriesID,
		event.EpisodeType,
		event.SaveNoMoreThan,
		event.SaveUntil,
		event.StartOffset,
		event.EndOffset,
		event.SearchString)
}

// ---------- Turbo Key: T, 54 --------------------

type TurboKeyEvent struct {
	*BaseEvent
	Key     string
	KeyCode int
}

func NewTurboKeyEvent(deviceId, clickString string) *TurboKeyEvent {
	key := new(TurboKeyEvent)
	key.BaseEvent = NewBaseEvent(deviceId, clickString)

	key.KeyCode = int(convertToInt(clickString[10:12]))
	key.Key = lookUpKeyName(key.KeyCode)

	return key
}

func (key TurboKeyEvent) String() string {
	return fmt.Sprintf("%s\tTurboKey:[%s]",
		key.BaseEvent,
		key.Key)
}
