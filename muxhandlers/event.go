package muxhandlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/RunnersRevival/outrun/config/eventconf"
	"github.com/RunnersRevival/outrun/config/gameconf"
	"github.com/RunnersRevival/outrun/consts"
	"github.com/RunnersRevival/outrun/db"
	"github.com/RunnersRevival/outrun/emess"
	"github.com/RunnersRevival/outrun/enums"
	"github.com/RunnersRevival/outrun/helper"
	"github.com/RunnersRevival/outrun/logic/conversion"
	"github.com/RunnersRevival/outrun/logic/gameplay"
	"github.com/RunnersRevival/outrun/netobj"
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
			// Revival since 2.2.0 supports all event types
			WhitelistedEventTypes = []int64{ // 2.x.x events
				enums.EventTypeSpecialStage,  // event stage, storyline, roulette, and rewards - Functional since 2.2.0
				enums.EventTypeRaidBoss,      // unique yearly event where one of the deadly six show up - Functional since 2.2.0
				enums.EventTypeCollectObject, // e.g. Animal Rescue Event - Functional since 2.2.0 (worked in 2.0 but missing menu UI)
				enums.EventTypeGacha,         // roulette event - Functional since 2.1.0 (worked in 2.0 but BGM was buggy)
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

func GetEventUserRaidbossState(helper *helper.Helper) {
	recv := helper.GetGameRequest()
	var request requests.GenericEventRequest
	err := json.Unmarshal(recv, &request)
	if err != nil {
		helper.Err("Error unmarshalling", err)
		return
	}
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}
	for time.Now().UTC().Unix() >= player.EventUserRaidbossState.EnergyRenewsAt && player.EventUserRaidbossState.RaidBossEnergy < 3 {
		player.EventUserRaidbossState.RaidBossEnergy++
		player.EventUserRaidbossState.EnergyRenewsAt += 1200
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.EventUserRaidbossState(baseInfo, player.EventUserRaidbossState)
	err = helper.SendCompatibleResponse(response, true)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}

func GetEventUserRaidbossList(helper *helper.Helper) {
	recv := helper.GetGameRequest()
	var request requests.GenericEventRequest
	err := json.Unmarshal(recv, &request)
	if err != nil {
		helper.Err("Error unmarshalling", err)
		return
	}
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}
	for time.Now().UTC().Unix() >= player.EventUserRaidbossState.EnergyRenewsAt && player.EventUserRaidbossState.RaidBossEnergy < 3 {
		player.EventUserRaidbossState.RaidBossEnergy++
		player.EventUserRaidbossState.EnergyRenewsAt += 1200
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DefaultEventUserRaidbossList(baseInfo, player.EventUserRaidbossState)
	err = helper.SendCompatibleResponse(response, true)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}

func EventActStart(helper *helper.Helper) {
	recv := helper.GetGameRequest()
	var request requests.EventActStartRequest
	err := json.Unmarshal(recv, &request)
	if err != nil {
		helper.Err("Error unmarshalling", err)
		return
	}
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)

	// consume items
	helper.DebugOut(fmt.Sprintf("%v", player.PlayerState.Items))
	for time.Now().UTC().Unix() >= player.EventUserRaidbossState.EnergyRenewsAt && player.EventUserRaidbossState.RaidBossEnergy < 3 {
		player.EventUserRaidbossState.RaidBossEnergy++
		player.EventUserRaidbossState.EnergyRenewsAt += 1200
	}
	if player.EventUserRaidbossState.RaidBossEnergy+player.EventUserRaidbossState.RaidBossEnergyBuy >= request.EnergyExpend {
		if gameconf.CFile.EnableEnergyConsumption {
			if player.EventUserRaidbossState.RaidBossEnergyBuy > 0 {
				player.EventUserRaidbossState.RaidBossEnergyBuy -= request.EnergyExpend
				if player.EventUserRaidbossState.RaidBossEnergyBuy < 0 { //did we go negative?
					player.EventUserRaidbossState.RaidBossEnergy += player.EventUserRaidbossState.RaidBossEnergyBuy
					player.EventUserRaidbossState.RaidBossEnergyBuy = 0
				}
			} else {
				player.EventUserRaidbossState.RaidBossEnergy -= request.EnergyExpend
				if player.EventUserRaidbossState.RaidBossEnergy < 3 {
					player.EventUserRaidbossState.EnergyRenewsAt = time.Now().UTC().Unix() + 1200
				}
			}
		}
		player.PlayerState.NumPlaying++
		if !gameconf.CFile.AllItemsFree {
			consumedRings := gameplay.GetRequiredItemPayment(request.Modifier, player)
			for _, citemID := range request.Modifier {
				if citemID[:2] == "11" { // boosts, not items
					continue
				}
				index := player.IndexOfItem(citemID)
				if index == -1 {
					helper.Uncatchable(fmt.Sprintf("Player sent bad item ID '%s', cannot continue", citemID))
					helper.InvalidRequest()
					return
				}
				if player.PlayerState.Items[index].Amount >= 1 { // can use item
					player.PlayerState.Items[index].Amount--
				}
			}
			if player.PlayerState.NumRings < consumedRings { // not enough rings
				baseInfo.StatusCode = status.NotEnoughRings
			}
			player.PlayerState.NumRings -= consumedRings
		}
	} else {
		baseInfo.StatusCode = status.NotEnoughEnergy
	}
	respPlayer := player
	if request.Version == "1.1.4" { // must send fewer characters
		// only get first 21 characters
		// TODO: enforce order 300000 to 300020?
		//cState = cState[:len(cState)-(len(cState)-10)]
		cState := respPlayer.CharacterState
		cState = cState[:16]
		helper.DebugOut("cState length: " + strconv.Itoa(len(cState)))
		helper.DebugOut("Sent character IDs: ")
		for _, char := range cState {
			helper.DebugOut(char.ID)
		}
		respPlayer.CharacterState = cState
	}
	response := responses.DefaultEventActStart(baseInfo, respPlayer)
	err = helper.SendCompatibleResponse(response, true)
	if err != nil {
		helper.InternalErr("Error sending response", err)
		return
	}
	err = db.SavePlayer(player)
	if err != nil {
		helper.InternalErr("Error saving player", err)
		return
	}
	//TODO: Add analytics for this
}

func EventPostGameResults(helper *helper.Helper) {
	recv := helper.GetGameRequest()
	var request requests.EventPostGameResultsRequest
	err := json.Unmarshal(recv, &request)
	if err != nil {
		helper.Err("Error unmarshalling", err)
		return
	}
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	for time.Now().UTC().Unix() >= player.EventUserRaidbossState.EnergyRenewsAt && player.EventUserRaidbossState.RaidBossEnergy < 3 {
		player.EventUserRaidbossState.RaidBossEnergy++
		player.EventUserRaidbossState.EnergyRenewsAt += 1200
	}
	player.EventUserRaidbossState.NumRaidbossRings += request.NumRaidbossRings
	response := responses.EventUserRaidbossState(baseInfo, player.EventUserRaidbossState)
	err = helper.SendCompatibleResponse(response, true)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
	err = db.SavePlayer(player)
	if err != nil {
		helper.InternalErr("Error saving player", err)
		return
	}
}

func EventUpdateGameResults(helper *helper.Helper) {
	recv := helper.GetGameRequest()
	var request requests.EventUpdateGameResultsRequest
	err := json.Unmarshal(recv, &request)
	if err != nil {
		helper.Err("Error unmarshalling", err)
		return
	}
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}

	//update energy counter
	for time.Now().UTC().Unix() >= player.PlayerState.EnergyRenewsAt && player.PlayerState.Energy < player.PlayerVarious.EnergyRecoveryMax {
		player.PlayerState.Energy++
		player.PlayerState.EnergyRenewsAt += player.PlayerVarious.EnergyRecoveryTime
	}
	for time.Now().UTC().Unix() >= player.EventUserRaidbossState.EnergyRenewsAt && player.EventUserRaidbossState.RaidBossEnergy < 3 {
		player.EventUserRaidbossState.RaidBossEnergy++
		player.EventUserRaidbossState.EnergyRenewsAt += 1200
	}

	hasSubCharacter := player.PlayerState.SubCharaID != "-1"
	var subC netobj.Character
	mainC, err := player.GetMainChara()
	if err != nil {
		helper.InternalErr("Error getting main character", err)
		return
	}
	playCharacters := []netobj.Character{ // assume only main character active right now
		mainC,
	}
	if hasSubCharacter {
		subC, err = player.GetSubChara()
		if err != nil {
			helper.InternalErr("Error getting sub character", err)
			return
		}
		playCharacters = []netobj.Character{ // add sub character to playCharacters
			mainC,
			subC,
		}
	}
	if request.Closed == 0 { // If the game wasn't exited out of
		player.PlayerState.NumRings += request.Rings
		player.PlayerState.NumRedRings += request.RedRings
		//player.PlayerState.NumRouletteTicket += request.RedRings
		player.PlayerState.Animals += request.Animals
		player.OptionUserResult.NumTakeAllRings += request.Rings
		player.OptionUserResult.NumTakeAllRedRings += request.RedRings
		player.PlayerState.TotalDistance += request.Distance
		// increase character(s)'s experience
		expIncrease := request.Rings + request.FailureRings // all rings collected
		// TODO: This isn't original server behavior! Expected behavior is described below:
		// Normal runs: The experience is the total score divided by 5000 (divided by 1000 prior to 2.0.0)
		// Boss battles: 2000 EXP if you clear it, nothing if you don't
		// Raid bosses: 10 EXP normal, 100 EXP rare, 500 EXP super-rare
		abilityIndex := 1
		for abilityIndex == 1 { // unused ability is at index 1
			abilityIndex = rand.Intn(len(mainC.AbilityLevel))
		}
		// check that increases exist
		_, ok := consts.UpgradeIncreases[mainC.ID]
		if !ok {
			helper.InternalErr("Error getting upgrade increase", fmt.Errorf("no key '%s' in consts.UpgradeIncreases", mainC.ID))
			return
		}
		if hasSubCharacter {
			_, ok = consts.UpgradeIncreases[subC.ID]
			if !ok {
				helper.InternalErr("Error getting upgrade increase for sub character", fmt.Errorf("no key '%s' in consts.UpgradeIncreases", subC.ID))
				return
			}
		}
		playCharacters[0].AbilityLevelUp = []int64{}
		playCharacters[0].AbilityLevelUpExp = []int64{}
		if playCharacters[0].Level < 100 {
			playCharacters[0].Exp += expIncrease
			for playCharacters[0].Exp >= playCharacters[0].Cost {
				// more exp than cost = level up
				if playCharacters[0].Level < 100 {
					abilityIndex = 1
					for abilityIndex == 1 || playCharacters[0].AbilityLevel[abilityIndex] >= 10 { // unused ability is at index 1
						abilityIndex = rand.Intn(len(mainC.AbilityLevel))
					}
					playCharacters[0].Level++                                               // increase level
					playCharacters[0].AbilityLevel[abilityIndex]++                          // increase ability level
					playCharacters[0].Exp -= playCharacters[0].Cost                         // remove cost from exp
					playCharacters[0].Cost += consts.UpgradeIncreases[playCharacters[0].ID] // increase cost
					playCharacters[0].AbilityLevelUp = append(playCharacters[0].AbilityLevelUp, int64(abilityIndex))
					playCharacters[0].AbilityLevelUpExp = append(playCharacters[0].AbilityLevelUpExp, playCharacters[0].Cost)
				} else {
					playCharacters[0].Exp -= playCharacters[0].Cost
				}
			}
		}

		if hasSubCharacter {
			playCharacters[1].AbilityLevelUp = []int64{}
			playCharacters[1].AbilityLevelUpExp = []int64{}
			if playCharacters[1].Level < 100 {
				playCharacters[1].Exp += expIncrease
				for playCharacters[1].Exp >= playCharacters[1].Cost {
					// more exp than cost = level up
					if playCharacters[1].Level < 100 {
						abilityIndex = 1
						for abilityIndex == 1 || playCharacters[1].AbilityLevel[abilityIndex] >= 10 { // unused ability is at index 1
							abilityIndex = rand.Intn(len(playCharacters[1].AbilityLevel))
						}
						playCharacters[1].Level++                                               // increase level
						playCharacters[1].AbilityLevel[abilityIndex]++                          // increase ability level
						playCharacters[1].Exp -= playCharacters[1].Cost                         // remove cost from exp
						playCharacters[1].Cost += consts.UpgradeIncreases[playCharacters[1].ID] // increase cost
						playCharacters[1].AbilityLevelUp = append(playCharacters[1].AbilityLevelUp, int64(abilityIndex))
						playCharacters[1].AbilityLevelUpExp = append(playCharacters[1].AbilityLevelUpExp, playCharacters[1].Cost)
					} else {
						playCharacters[1].Exp -= playCharacters[1].Cost
					}
				}
			}
		}

		helper.DebugOut("Old mainC Exp: %v / %v", mainC.Exp, mainC.Cost)
		helper.DebugOut("Old mainC Level: %v", mainC.Level)
		if hasSubCharacter {
			helper.DebugOut("Old subC Exp: %v / %v", subC.Exp, subC.Cost)
			helper.DebugOut("Old subC Level: %v", subC.Level)
		}
		helper.DebugOut("New mainC Exp: %v / %v", playCharacters[0].Exp, playCharacters[0].Cost)
		helper.DebugOut("New mainC Level: %v", playCharacters[0].Level)
		helper.DebugOut("mainC AbilityLevelUp: %v", playCharacters[0].AbilityLevelUp)
		helper.DebugOut("mainC AbilityLevelUpExp: %v", playCharacters[0].AbilityLevelUpExp)
		if hasSubCharacter {
			helper.DebugOut("New subC Exp: %v / %v", playCharacters[1].Exp, playCharacters[1].Cost)
			helper.DebugOut("New subC Level: %v", playCharacters[1].Level)
			helper.DebugOut("subC AbilityLevelUp: %v", playCharacters[1].AbilityLevelUp)
			helper.DebugOut("subC AbilityLevelUpExp: %v", playCharacters[1].AbilityLevelUpExp)
		}

		helper.DebugOut("Event ID: %v", request.EventID)
		helper.DebugOut("Player got %v event object(s)", request.EventValue)
		player.EventState.Param += request.EventValue

		helper.DebugOut("Raid boss ID: %v", request.RaidbossID)
		helper.DebugOut("It took %v point(s) of damage", request.RaidbossDamage)
		if request.RaidbossBeatFlg != 0 {
			helper.DebugOut("It was defeated!")
			player.EventUserRaidbossState.NumBeatedEncounter++
		}
	}

	mainCIndex := player.IndexOfChara(mainC.ID) // TODO: check if -1
	subCIndex := -1
	if hasSubCharacter {
		subCIndex = player.IndexOfChara(subC.ID) // TODO: check if -1
	}

	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DefaultEventUpdateGameResults(baseInfo, player, playCharacters, player.EventState)
	// apply the save after the response so that we don't break the leveling
	mainC = playCharacters[0]
	if hasSubCharacter {
		subC = playCharacters[1]
	}
	player.CharacterState[mainCIndex] = mainC
	if hasSubCharacter {
		player.CharacterState[subCIndex] = subC
	}
	err = db.SavePlayer(player)
	if err != nil {
		helper.InternalErr("Error saving player", err)
		return
	}
	helper.DebugOut(fmt.Sprintf("%v", player.PlayerState.Items))

	err = helper.SendCompatibleResponse(response, true)
	if err != nil {
		helper.InternalErr("Error sending response", err)
		return
	}
}
