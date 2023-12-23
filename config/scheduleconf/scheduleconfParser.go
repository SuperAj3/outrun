package scheduleconf

import (
	"encoding/json"
	"log"
	"os"
)

var Defaults = map[string]interface{}{
	"DAllowScheduledEvents": true,
	"DEventDataFolderName":  "events",
	"DSchedule":             []ConfiguredTimeslot{},
}

type ConfiguredTimeslot struct {
	StartTime     int64  `json:"startTime"`
	EndTime       int64  `json:"endTime"`
	EventFilename string `json:"eventFilename"`
}

var CFile ConfigFile

type ConfigFile struct {
	AllowScheduledEvents bool                 `json:"enableScheduledEvents,omitempty"`
	EventDataFolderName  string               `json:"eventDataFolderName,omitempty"`
	Schedule             []ConfiguredTimeslot `json:"schedule,omitempty"`
}

func Parse(filename string) error {
	CFile = ConfigFile{
		Defaults["DAllowScheduledEvents"].(bool),
		Defaults["DEventDataFolderName"].(string),
		Defaults["DSchedule"].([]ConfiguredTimeslot),
	}
	file, err := loadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &CFile)
	if err != nil {
		return err
	}
	scheduledTimeslots := []ConfiguredTimeslot{}
	for _, ts := range CFile.Schedule {
		// TODO: Add parsing of scheduled events!
		scheduledTimeslots = append(scheduledTimeslots, ts)
	}
	log.Printf("[INFO] Loaded %v timeslot(s).\n", len(scheduledTimeslots))
	CFile.Schedule = scheduledTimeslots
	return nil
}

func loadFile(filename string) ([]byte, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}
	return b, err
}
