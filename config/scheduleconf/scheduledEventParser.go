package scheduleconf

import "github.com/RunnersRevival/outrun/enums"

// A Scheduled Event is a type of file which has all the necessary data for an event or campaign in a given timeslot.

var TimeslotTypeMapping = map[string]int64{
	"event":            0,
	"campaign":         1,
	"rouletteOddsOnly": 2,
}

var EventTypes = map[string]int64{
	"specialStage":  enums.EventIDSpecialStage,
	"raidBoss":      enums.EventIDRaidBoss,
	"collectObject": enums.EventIDCollectObject,
	"gacha":         enums.EventIDGacha,
	"advert":        enums.EventIDAdvert,
	"quick":         enums.EventIDQuick,
	"bgm":           enums.EventIDBGM,
}

type ScheduledEventFile struct {
	IsActive     bool   `json:"isActive,omitempty"`
	TimeslotType string `json:"timeslotType,omitempty"`
}
