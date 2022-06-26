package muxhandlers

import (
	"encoding/json"

	"github.com/RunnersRevival/outrun/config/eventconf"
	"github.com/RunnersRevival/outrun/emess"
	"github.com/RunnersRevival/outrun/enums"
	"github.com/RunnersRevival/outrun/helper"
	"github.com/RunnersRevival/outrun/logic/conversion"
	"github.com/RunnersRevival/outrun/obj"
	"github.com/RunnersRevival/outrun/requests"
	"github.com/RunnersRevival/outrun/responses"
	"github.com/RunnersRevival/outrun/status"
)

func IsEventTypeValidForGameVersion(gameVersion string, eventType int64) bool {
	WhitelistedEventTypes := []int64{ // 1.x.x events
		enums.EventTypeSpecialStage,  // event stage, storyline, roulette, and rewards
		enums.EventTypeRaidBoss,      // unique yearly event where one of the deadly six show up
		enums.EventTypeCollectObject, // e.g. Animal Rescue Event
		enums.EventTypeGacha,         // roulette event
		enums.EventTypeAdvert,        // banner only
	}
	if gameVersion[0] == '2' {
		if gameVersion[2] == '0' {
			WhitelistedEventTypes = []int64{ // 2.0.x events
				// the first four types were technically made unused, but they can still be called despite their incompleteness
				enums.EventTypeAdvert, // banner only
				enums.EventTypeQuick,  // timed mode stage
				enums.EventTypeBGM,    // custom BGM during gameplay
			}
		} else {
			WhitelistedEventTypes = []int64{ // 2.x.x events
				enums.EventTypeSpecialStage,  // event stage, storyline, roulette, and rewards (broken in 2.0.x)
				enums.EventTypeRaidBoss,      // unique yearly event where one of the deadly six show up (broken in 2.0.x)
				enums.EventTypeCollectObject, // e.g. Animal Rescue Event
				enums.EventTypeGacha,         // roulette event
				enums.EventTypeAdvert,        // banner only
				enums.EventTypeQuick,         // timed mode stage
				enums.EventTypeBGM,           // custom BGM during gameplay
			}
		}
	}
	for _, a := range WhitelistedEventTypes {
		if eventType == a {
			return true
		}
	}
	return false
}

func GetEventList(helper *helper.Helper) {
	data := helper.GetGameRequest()
	var request requests.Base
	err := json.Unmarshal(data, &request)
	if err != nil {
		helper.InternalErr("Error unmarshalling", err)
		return
	}
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	// construct event list
	eventList := []obj.Event{}
	if eventconf.CFile.AllowEvents {
		if eventconf.CFile.EnforceGlobal || len(player.PersonalEvents) == 0 {
			for _, confEvent := range eventconf.CFile.CurrentEvents {
				newEvent := conversion.ConfiguredEventToEvent(confEvent)
				if IsEventTypeValidForGameVersion(request.Version, newEvent.Type) {
					helper.Warn("Event %v may not work on game version %s!", newEvent.ID, request.Version)
				}
				eventList = append(eventList, newEvent)
			}
		} else {
			for _, ce := range player.PersonalEvents {
				e := conversion.ConfiguredEventToEvent(ce)
				if IsEventTypeValidForGameVersion(request.Version, e.Type) {
					eventList = append(eventList, e)
				}
			}
		}
	}
	helper.DebugOut("Personal event list: %v", player.PersonalEvents)
	helper.DebugOut("Global event list: %v", eventconf.CFile.CurrentEvents)
	helper.DebugOut("Event list: %v", eventList)
	response := responses.EventList(baseInfo, eventList)
	//response.BaseResponse = responses.NewBaseResponseV(baseInfo, request.Version)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}

func GetEventReward(helper *helper.Helper) {
	recv := helper.GetGameRequest()
	var request requests.GenericEventRequest
	err := json.Unmarshal(recv, &request)
	if err != nil {
		helper.Err("Error unmarshalling", err)
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DefaultEventRewardList(baseInfo)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}

func GetEventState(helper *helper.Helper) {
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}
	recv := helper.GetGameRequest()
	var request requests.GenericEventRequest
	err = json.Unmarshal(recv, &request)
	if err != nil {
		helper.Err("Error unmarshalling", err)
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.EventState(baseInfo, player.EventState)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}
