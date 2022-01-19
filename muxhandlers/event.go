package muxhandlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/Mtbcooler/outrun/config/eventconf"
	"github.com/Mtbcooler/outrun/config/gameconf"
	"github.com/Mtbcooler/outrun/consts"
	"github.com/Mtbcooler/outrun/db"
	"github.com/Mtbcooler/outrun/db/dbaccess"
	"github.com/Mtbcooler/outrun/emess"
	"github.com/Mtbcooler/outrun/helper"
	"github.com/Mtbcooler/outrun/logic/conversion"
	"github.com/Mtbcooler/outrun/logic/gameplay"
	"github.com/Mtbcooler/outrun/netobj"
	"github.com/Mtbcooler/outrun/obj"
	"github.com/Mtbcooler/outrun/requests"
	"github.com/Mtbcooler/outrun/responses"
	"github.com/Mtbcooler/outrun/status"
	"github.com/jinzhu/now"
)

func GetEventList(helper *helper.Helper) {
	sid, _ := helper.GetSessionID()
	if !helper.CheckSession(true) {
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
				eventList = append(eventList, newEvent)
			}
		} else {
			for _, ce := range player.PersonalEvents {
				e := conversion.ConfiguredEventToEvent(ce)
				eventList = append(eventList, e)
			}
		}
	}
	helper.DebugOut("Personal event list: %v", player.PersonalEvents)
	helper.DebugOut("Global event list: %v", eventconf.CFile.CurrentEvents)
	helper.DebugOut("Event list: %v", eventList)
	response := responses.EventList(baseInfo, eventList)
	response.Seq, _ = db.BoltGetSessionIDSeq(sid)
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
	sid, _ := helper.GetSessionID()
	if !helper.CheckSession(true) {
		return
	}
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
	response.Seq, _ = db.BoltGetSessionIDSeq(sid)
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
	err = helper.SendInsecureResponse(response)
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
	// TODO: get the raid bosses from the database
	response := responses.DefaultEventUserRaidbossList(baseInfo, player.EventUserRaidbossState)
	err = helper.SendInsecureResponse(response)
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
	response := responses.DefaultEventActStart(baseInfo, player)
	err = helper.SendInsecureResponse(response)
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
	err = helper.SendInsecureResponse(response)
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
		player.PlayerState.NumRouletteTicket += request.RedRings // TODO: URGENT! Remove as soon as possible!
		player.PlayerState.Animals += request.Animals
		player.OptionUserResult.NumTakeAllRings += request.Rings
		player.OptionUserResult.NumTakeAllRedRings += request.RedRings
		playerTimedHighScore := player.PlayerState.TimedHighScore
		if request.Score > playerTimedHighScore {
			player.PlayerState.TimedHighScore = request.Score
		}

		_, leagueresettime, err := dbaccess.GetStartAndEndTimesForLeague(player.PlayerState.RankingLeague, player.PlayerState.RankingLeagueGroup)
		if err != nil {
			helper.InternalErr("Error getting league reset time (try restarting Outrun or sending the ResetRankingData RPC command)", err)
			return
		}
		if time.Now().UTC().Unix() < leagueresettime {
			playerLeagueHighScore := player.PlayerState.LeagueHighScore
			if request.Score > playerLeagueHighScore {
				player.PlayerState.LeagueHighScore = request.Score
			}
			player.PlayerState.TotalScore += request.Score
			if player.PlayerState.TotalScore > player.PlayerState.HighTotalScore {
				player.PlayerState.HighTotalScore = player.PlayerState.TotalScore
				player.OptionUserResult.TotalSumHighScore = player.PlayerState.HighTotalScore
			}
		}
		helper.DebugOut("request.DailyChallengeValue: %v", request.DailyChallengeValue)
		helper.DebugOut("request.DailyChallengeComplete: %v", request.DailyChallengeComplete)
		if player.PlayerState.DailyChallengeComplete == 0 && request.DailyChallengeComplete == 1 {
			if player.PlayerState.NextNumDailyChallenge <= 0 {
				player.PlayerState.NumDailyChallenge = int64(0)
				player.PlayerState.NextNumDailyChallenge = int64(1)
			}
			player.AddOperatorMessage(
				"A Daily Challenge reward.",
				obj.NewMessageItem(
					consts.DailyMissionRewards[player.PlayerState.NextNumDailyChallenge-1],
					consts.DailyMissionRewardCounts[player.PlayerState.NextNumDailyChallenge-1],
					0,
					0,
				),
				2592000,
			)
			player.PlayerState.NumDailyChallenge = player.PlayerState.NextNumDailyChallenge
		}
		if player.PlayerState.DailyChallengeComplete == 0 {
			player.PlayerState.DailyChallengeComplete = request.DailyChallengeComplete
		}
		player.PlayerState.DailyChallengeValue = request.DailyChallengeValue
		if time.Now().UTC().Unix() >= player.PlayerState.DailyMissionEndTime {
			if player.PlayerState.DailyChallengeComplete == 1 && player.PlayerState.DailyChalSetNum < 10 {
				helper.DebugOut("Advancing to next daily mission...")
				player.PlayerState.DailyChalSetNum++
			} else {
				player.PlayerState.DailyChalCatNum = int64(rand.Intn(5))
				player.PlayerState.DailyChalSetNum = int64(0)
			}
			if player.PlayerState.DailyChallengeComplete == 0 {
				player.PlayerState.NumDailyChallenge = int64(0)
				player.PlayerState.NextNumDailyChallenge = int64(1)
			} else {
				player.PlayerState.NextNumDailyChallenge++
				if int(player.PlayerState.NextNumDailyChallenge) > len(consts.DailyMissionRewards) {
					player.PlayerState.NumDailyChallenge = int64(0)
					player.PlayerState.NextNumDailyChallenge = int64(1) //restart from beginning
					player.PlayerState.DailyChalCatNum = int64(rand.Intn(5))
					player.PlayerState.DailyChalSetNum = int64(0)
				}
			}
			player.PlayerState.DailyChalPosNum = int64(1 + rand.Intn(2))
			player.PlayerState.DailyMissionID = int64((player.PlayerState.DailyChalCatNum * 33) + (player.PlayerState.DailyChalSetNum * 3) + player.PlayerState.DailyChalPosNum)
			player.PlayerState.DailyChallengeValue = int64(0)
			player.PlayerState.DailyChallengeComplete = int64(0)
			player.PlayerState.DailyMissionEndTime = now.EndOfDay().UTC().Unix() + 1
			helper.DebugOut("New daily mission ID: %v", player.PlayerState.DailyMissionID)
		}
		//player.PlayerState.TotalDistance += request.Distance  // We don't do this in timed mode!
		// increase character(s)'s experience
		expIncrease := request.Rings + request.FailureRings // all rings collected
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
			player.EventUserRaidbossState.NumBeatedEnterprise++
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

	err = helper.SendInsecureResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
		return
	}
}
