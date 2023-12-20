package scheduleconf

var Defaults = map[string]interface{}{
	"DEnableScheduledEvents": true,
	"DSchedule":              []ConfiguredTimeslot{},
}

type ConfiguredTimeslot struct {
	StartTime     int64  `json:"startTime"`
	EndTime       int64  `json:"endTime"`
	EventFilename string `json:"eventFilename"`
}

var CFile ConfigFile

type ConfigFile struct {
	EnableScheduledEvents bool                 `json:"enableScheduledEvents,omitempty"`
	Schedule              []ConfiguredTimeslot `json:"schedule,omitempty"`
}
