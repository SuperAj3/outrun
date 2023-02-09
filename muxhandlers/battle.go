package muxhandlers

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/RunnersRevival/outrun/db"
	"github.com/RunnersRevival/outrun/emess"
	"github.com/RunnersRevival/outrun/enums"
	"github.com/RunnersRevival/outrun/helper"
	"github.com/RunnersRevival/outrun/logic/battle"
	"github.com/RunnersRevival/outrun/logic/conversion"
	"github.com/RunnersRevival/outrun/obj"
	"github.com/RunnersRevival/outrun/obj/constobjs"
	"github.com/RunnersRevival/outrun/requests"
	"github.com/RunnersRevival/outrun/responses"
	"github.com/RunnersRevival/outrun/status"
	"github.com/jinzhu/now"
)

func GetDailyBattleData(helper *helper.Helper) {
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("error getting calling player", err)
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	var response interface{}
	if player.BattleState.ScoreRecordedToday {
		if player.BattleState.MatchedUpWithRival {
			rivalPlayer, err := db.GetPlayer(player.BattleState.RivalID)
			if err != nil {
				helper.InternalErr("error getting rival player", err)
				return
			}
			response = responses.DailyBattleData(baseInfo,
				player.BattleState.BattleStartsAt,
				player.BattleState.BattleEndsAt,
				conversion.DebugPlayerToBattleData(player),
				conversion.DebugPlayerToBattleData(rivalPlayer),
			)
		} else {
			response = responses.NoRivalDailyBattleData(baseInfo,
				player.BattleState.BattleStartsAt,
				player.BattleState.BattleEndsAt,
				conversion.DebugPlayerToBattleData(player),
			)
		}
	} else {
		response = responses.NoScoreDailyBattleData(baseInfo,
			player.BattleState.BattleStartsAt,
			player.BattleState.BattleEndsAt,
		)
	}
	err = helper.SendCompatibleResponse(response, true)
	if err != nil {
		helper.InternalErr("error sending response", err)
	}
}

func UpdateDailyBattleStatus(helper *helper.Helper) {
	data := helper.GetGameRequest()
	var request requests.Base
	err := json.Unmarshal(data, &request)
	if err != nil {
		helper.InternalErr("Error unmarshalling", err)
		return
	}
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("error getting calling player", err)
		return
	}
	var rewardBattleStartTime int64
	var rewardBattleEndTime int64
	var rewardBattlePlayerData obj.BattleData
	var rewardBattleRivalData obj.BattleData
	doReward := false
	if player.BattleState.PendingReward {
		rewardBattleStartTime = player.BattleState.PendingRewardData.StartTime
		rewardBattleEndTime = player.BattleState.PendingRewardData.EndTime
		rewardBattlePlayerData = player.BattleState.PendingRewardData.BattleData
		rewardBattleRivalData = player.BattleState.PendingRewardData.RivalBattleData
		player.BattleState.PendingReward = false
		doReward = true
		if time.Now().UTC().Unix() > player.BattleState.BattleEndsAt {
			player.BattleState.BattleStartsAt = now.BeginningOfDay().UTC().Unix()
			player.BattleState.BattleEndsAt = now.EndOfDay().UTC().Unix() + 1
			player.BattleState.ScoreRecordedToday = false
			player.BattleState.MatchedUpWithRival = false
			player.BattleState.WantsStricterMatchmaking = false
		}
	} else {
		if time.Now().UTC().Unix() > player.BattleState.BattleEndsAt {
			if player.BattleState.ScoreRecordedToday {
				if player.BattleState.MatchedUpWithRival {
					rivalPlayer, err := db.GetPlayer(player.BattleState.RivalID)
					if err != nil {
						helper.InternalErr("error getting rival player", err)
						return
					}
					rewardBattleStartTime = player.BattleState.BattleStartsAt
					rewardBattleEndTime = player.BattleState.BattleEndsAt
					rewardBattlePlayerData = conversion.DebugPlayerToBattleData(player)
					rewardBattleRivalData = conversion.DebugPlayerToBattleData(rivalPlayer)
					battlePair := obj.NewBattlePair(
						rewardBattleStartTime,
						rewardBattleEndTime,
						rewardBattlePlayerData,
						rewardBattleRivalData,
					)
					rivalBattlePair := obj.NewBattlePair(
						rewardBattleStartTime,
						rewardBattleEndTime,
						rewardBattleRivalData,
						rewardBattlePlayerData,
					)
					rivalPlayer.BattleState.PendingReward = true
					rivalPlayer.BattleState.PendingRewardData = obj.NewRewardBattlePair(
						rewardBattleStartTime,
						rewardBattleEndTime,
						rewardBattleRivalData,
						rewardBattlePlayerData,
					)
					if player.BattleState.DailyBattleHighScore > rivalPlayer.BattleState.DailyBattleHighScore {
						player.BattleState.Wins++
						player.BattleState.WinStreak++
						player.BattleState.LossStreak = 0
						rivalPlayer.BattleState.Losses++
						rivalPlayer.BattleState.LossStreak++
						rivalPlayer.BattleState.WinStreak = 0
						rivalPlayer.AddOperatorMessage(
							"A compensation reward for losing a Daily Battle.",
							obj.NewMessageItem(
								enums.ItemIDRedRing,
								15,
								0,
								0,
							),
							2592000,
						)
					} else {
						if player.BattleState.DailyBattleHighScore < rivalPlayer.BattleState.DailyBattleHighScore {
							player.BattleState.Losses++
							player.BattleState.LossStreak++
							player.BattleState.WinStreak = 0
							rivalPlayer.BattleState.Wins++
							rivalPlayer.BattleState.WinStreak++
							rivalPlayer.BattleState.LossStreak = 0
							player.AddOperatorMessage(
								"A compensation reward for losing a Daily Battle.",
								obj.NewMessageItem(
									enums.ItemIDRedRing,
									15,
									0,
									0,
								),
								2592000,
							)
						} else {
							player.BattleState.Draws++
							player.BattleState.WinStreak = 0
							player.BattleState.LossStreak = 0
							rivalPlayer.BattleState.Draws++
							rivalPlayer.BattleState.WinStreak = 0
							rivalPlayer.BattleState.LossStreak = 0
						}
					}
					// TODO: Determine if this routine is working correctly for all scenarios. At some point during the original 2019-2020 test it was determined that the daily battle rewards had a few issues, and it is unknown if these were ever properly addressed.
					rewardIndex := 0
					if player.BattleState.WinStreak > 0 {
						for player.BattleState.WinStreak > constobjs.DefaultDailyBattlePrizeList[rewardIndex].Number && rewardIndex < len(constobjs.DefaultDailyBattlePrizeList) {
							rewardIndex++
						}
						for _, item := range constobjs.DefaultDailyBattlePrizeList[rewardIndex].PresentList {
							itemid, _ := strconv.Atoi(item.ID)
							player.AddOperatorMessage(
								"A reward for "+strconv.Itoa(int(player.BattleState.WinStreak))+" consecutive Daily Battle win(s).",
								obj.MessageItem{
									int64(itemid),
									item.Amount,
									0,
									0,
								},
								2592000,
							)
						}
					}
					if rivalPlayer.BattleState.WinStreak > 0 {
						rewardIndex = 0
						for player.BattleState.WinStreak > constobjs.DefaultDailyBattlePrizeList[rewardIndex].Number && rewardIndex < len(constobjs.DefaultDailyBattlePrizeList) {
							rewardIndex++
						}
						for _, item := range constobjs.DefaultDailyBattlePrizeList[rewardIndex].PresentList {
							itemid, _ := strconv.Atoi(item.ID)
							rivalPlayer.AddOperatorMessage(
								"A reward for "+strconv.Itoa(int(rivalPlayer.BattleState.WinStreak))+" consecutive Daily Battle win(s).",
								obj.MessageItem{
									int64(itemid),
									item.Amount,
									0,
									0,
								},
								2592000,
							)
						}
					}
					player.BattleState.BattleHistory = append(player.BattleState.BattleHistory, battlePair)
					rivalPlayer.BattleState.BattleHistory = append(rivalPlayer.BattleState.BattleHistory, rivalBattlePair)
					rivalPlayer.BattleState.PendingReward = true
					err = db.SavePlayer(rivalPlayer)
					if err != nil {
						helper.InternalErr("Error saving player", err)
						return
					}
					doReward = true
				} else {
					// There appears to be no reward for failures
					// TODO: Guesswork! Confirm if this is correct.
					player.BattleState.Failures++
					player.BattleState.LossStreak++
					player.BattleState.WinStreak = 0
					player.AddOperatorMessage(
						"No rival has been found, but we got something for you.",
						obj.NewMessageItem(
							enums.ItemIDRedRing,
							15,
							0,
							0,
						),
						2592000,
					)
				}
			}
			player.BattleState.BattleStartsAt = now.BeginningOfDay().UTC().Unix()
			player.BattleState.BattleEndsAt = now.EndOfDay().UTC().Unix() + 1
			player.BattleState.ScoreRecordedToday = false
			player.BattleState.MatchedUpWithRival = false
			player.BattleState.WantsStricterMatchmaking = false
		}
	}
	battleStatus := obj.BattleStatus{
		player.BattleState.Wins,
		player.BattleState.Losses,
		player.BattleState.Draws,
		player.BattleState.Failures,
		player.BattleState.WinStreak,
		player.BattleState.LossStreak,
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	var response interface{}
	if doReward {
		response = responses.UpdateDailyBattleStatusWithReward(baseInfo, player.BattleState.BattleEndsAt, battleStatus, rewardBattleStartTime, rewardBattleEndTime, rewardBattlePlayerData, rewardBattleRivalData)
	} else {
		response = responses.UpdateDailyBattleStatus(baseInfo, player.BattleState.BattleEndsAt, battleStatus)
	}
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
}

func ResetDailyBattleMatching(helper *helper.Helper) {
	data := helper.GetGameRequest()
	var request requests.ResetDailyBattleMatchingRequest
	err := json.Unmarshal(data, &request)
	if err != nil {
		helper.InternalErr("Error unmarshalling", err)
		return
	}
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("error getting calling player", err)
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	battleData := conversion.DebugPlayerToBattleData(player)
	startTime := player.BattleState.BattleStartsAt
	endTime := player.BattleState.BattleEndsAt

	helper.DebugOut("Type: %v", request.Type)
	switch request.Type {
	case 1:
		if player.PlayerState.NumRedRings < 5 {
			baseInfo.StatusCode = status.NotEnoughRedRings
			err = helper.SendResponse(responses.NewBaseResponse(baseInfo))
			if err != nil {
				helper.InternalErr("error sending response", err)
			}
			return
		}
	case 2:
		if player.PlayerState.NumRedRings < 40 {
			baseInfo.StatusCode = status.NotEnoughRedRings
			err = helper.SendResponse(responses.NewBaseResponse(baseInfo))
			if err != nil {
				helper.InternalErr("error sending response", err)
			}
			return
		}
	}
	oldRivalID := player.BattleState.RivalID
	if request.Type == 2 {
		// we don't implement this option yet
		helper.InvalidRequest()
		return
	} else {
		if request.Type != 0 {
			player.BattleState.MatchedUpWithRival = false
			player.BattleState.WantsStricterMatchmaking = false
			oldRival, err := db.GetPlayer(oldRivalID)
			if err != nil {
				helper.InternalErr("error getting rival player", err)
				return
			}
			oldRival.BattleState.MatchedUpWithRival = false
			err = db.SavePlayer(oldRival)
			if err != nil {
				helper.InternalErr("Error saving old rival", err)
				return
			}
		}
		player.BattleState = battle.DrawBattleRival(player, 200)
	}

	if player.BattleState.RivalID != oldRivalID && player.BattleState.MatchedUpWithRival {
		switch request.Type {
		case 1:
			//helper.Warn("Skip RSR deduction for now")
			player.PlayerState.NumRedRings -= 5
		case 2:
			//helper.Warn("Skip RSR deduction for now")
			player.PlayerState.NumRedRings -= 40
		}
	}

	var response interface{}
	if player.BattleState.MatchedUpWithRival {
		rivalPlayer, err := db.GetPlayer(player.BattleState.RivalID)
		if err != nil {
			helper.InternalErr("error getting rival player", err)
			return
		}
		rivalBattleData := conversion.DebugPlayerToBattleData(rivalPlayer)
		response = responses.ResetDailyBattleMatching(baseInfo, startTime, endTime, battleData, rivalBattleData, player)
	} else {
		response = responses.ResetDailyBattleMatchingNoOpponent(baseInfo, startTime, endTime, battleData, player)
	}
	err = helper.SendCompatibleResponse(response, true)
	if err != nil {
		helper.InternalErr("error sending response", err)
	}
	err = db.SavePlayer(player)
	if err != nil {
		helper.InternalErr("Error saving player", err)
		return
	}
}

func GetDailyBattleHistory(helper *helper.Helper) {
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("error getting calling player", err)
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.GetDailyBattleHistory(baseInfo, player.BattleState.BattleHistory)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("error sending response", err)
	}
}

func GetDailyBattleStatus(helper *helper.Helper) {
	data := helper.GetGameRequest()
	var request requests.Base
	err := json.Unmarshal(data, &request)
	if err != nil {
		helper.InternalErr("Error unmarshalling", err)
		return
	}
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("error getting calling player", err)
		return
	}
	battleStatus := obj.BattleStatus{
		player.BattleState.Wins,
		player.BattleState.Losses,
		player.BattleState.Draws,
		player.BattleState.Failures,
		player.BattleState.WinStreak,
		player.BattleState.LossStreak,
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)

	response := responses.GetDailyBattleStatus(baseInfo, player.BattleState.BattleEndsAt, battleStatus)
	err = helper.SendCompatibleResponse(response, true)
	if err != nil {
		helper.InternalErr("Error sending response", err)
		return
	}
}

func PostDailyBattleResult(helper *helper.Helper) {
	data := helper.GetGameRequest()
	var request requests.Base
	err := json.Unmarshal(data, &request)
	if err != nil {
		helper.InternalErr("Error unmarshalling", err)
		return
	}
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("error getting calling player", err)
		return
	}
	var rewardBattleStartTime int64
	var rewardBattleEndTime int64
	var rewardBattlePlayerData obj.BattleData
	var rewardBattleRivalData obj.BattleData
	doReward := false
	// TODO: Solve duplication requirement!
	if player.BattleState.PendingReward {
		rewardBattleStartTime = player.BattleState.PendingRewardData.StartTime
		rewardBattleEndTime = player.BattleState.PendingRewardData.EndTime
		rewardBattlePlayerData = player.BattleState.PendingRewardData.BattleData
		rewardBattleRivalData = player.BattleState.PendingRewardData.RivalBattleData
		player.BattleState.PendingReward = false
		doReward = true
		if time.Now().UTC().Unix() > player.BattleState.BattleEndsAt {
			player.BattleState.BattleStartsAt = now.BeginningOfDay().UTC().Unix()
			player.BattleState.BattleEndsAt = now.EndOfDay().UTC().Unix() + 1
			player.BattleState.ScoreRecordedToday = false
			player.BattleState.MatchedUpWithRival = false
			player.BattleState.WantsStricterMatchmaking = false
		}
	} else {
		if time.Now().UTC().Unix() > player.BattleState.BattleEndsAt {
			if player.BattleState.ScoreRecordedToday {
				if player.BattleState.MatchedUpWithRival {
					rivalPlayer, err := db.GetPlayer(player.BattleState.RivalID)
					if err != nil {
						helper.InternalErr("error getting rival player", err)
						return
					}
					rewardBattleStartTime = player.BattleState.BattleStartsAt
					rewardBattleEndTime = player.BattleState.BattleEndsAt
					rewardBattlePlayerData = conversion.DebugPlayerToBattleData(player)
					rewardBattleRivalData = conversion.DebugPlayerToBattleData(rivalPlayer)
					battlePair := obj.NewBattlePair(
						rewardBattleStartTime,
						rewardBattleEndTime,
						rewardBattlePlayerData,
						rewardBattleRivalData,
					)
					rivalBattlePair := obj.NewBattlePair(
						rewardBattleStartTime,
						rewardBattleEndTime,
						rewardBattleRivalData,
						rewardBattlePlayerData,
					)
					rivalPlayer.BattleState.PendingReward = true
					rivalPlayer.BattleState.PendingRewardData = obj.NewRewardBattlePair(
						rewardBattleStartTime,
						rewardBattleEndTime,
						rewardBattleRivalData,
						rewardBattlePlayerData,
					)
					if player.BattleState.DailyBattleHighScore > rivalPlayer.BattleState.DailyBattleHighScore {
						player.BattleState.Wins++
						player.BattleState.WinStreak++
						player.BattleState.LossStreak = 0
						rivalPlayer.BattleState.Losses++
						rivalPlayer.BattleState.LossStreak++
						rivalPlayer.BattleState.WinStreak = 0
						rivalPlayer.AddOperatorMessage(
							"A compensation reward for losing a Daily Battle.",
							obj.NewMessageItem(
								enums.ItemIDRedRing,
								15,
								0,
								0,
							),
							2592000,
						)
					} else {
						if player.BattleState.DailyBattleHighScore < rivalPlayer.BattleState.DailyBattleHighScore {
							player.BattleState.Losses++
							player.BattleState.LossStreak++
							player.BattleState.WinStreak = 0
							rivalPlayer.BattleState.Wins++
							rivalPlayer.BattleState.WinStreak++
							rivalPlayer.BattleState.LossStreak = 0
							player.AddOperatorMessage(
								"A compensation reward for losing a Daily Battle.",
								obj.NewMessageItem(
									enums.ItemIDRedRing,
									15,
									0,
									0,
								),
								2592000,
							)
						} else {
							player.BattleState.Draws++
							player.BattleState.WinStreak = 0
							player.BattleState.LossStreak = 0
							rivalPlayer.BattleState.Draws++
							rivalPlayer.BattleState.WinStreak = 0
							rivalPlayer.BattleState.LossStreak = 0
						}
					}
					// TODO: Determine if this routine is working correctly for all scenarios. At some point during the original 2019-2020 test it was determined that the daily battle rewards had a few issues, and it is unknown if these were ever properly addressed.
					rewardIndex := 0
					if player.BattleState.WinStreak > 0 {
						for player.BattleState.WinStreak > constobjs.DefaultDailyBattlePrizeList[rewardIndex].Number && rewardIndex < len(constobjs.DefaultDailyBattlePrizeList) {
							rewardIndex++
						}
						for _, item := range constobjs.DefaultDailyBattlePrizeList[rewardIndex].PresentList {
							itemid, _ := strconv.Atoi(item.ID)
							player.AddOperatorMessage(
								"A reward for "+strconv.Itoa(int(player.BattleState.WinStreak))+" consecutive Daily Battle win(s).",
								obj.MessageItem{
									int64(itemid),
									item.Amount,
									0,
									0,
								},
								2592000,
							)
						}
					}
					if rivalPlayer.BattleState.WinStreak > 0 {
						rewardIndex = 0
						for player.BattleState.WinStreak > constobjs.DefaultDailyBattlePrizeList[rewardIndex].Number && rewardIndex < len(constobjs.DefaultDailyBattlePrizeList) {
							rewardIndex++
						}
						for _, item := range constobjs.DefaultDailyBattlePrizeList[rewardIndex].PresentList {
							itemid, _ := strconv.Atoi(item.ID)
							rivalPlayer.AddOperatorMessage(
								"A reward for "+strconv.Itoa(int(rivalPlayer.BattleState.WinStreak))+" consecutive Daily Battle win(s).",
								obj.MessageItem{
									int64(itemid),
									item.Amount,
									0,
									0,
								},
								2592000,
							)
						}
					}
					player.BattleState.BattleHistory = append(player.BattleState.BattleHistory, battlePair)
					rivalPlayer.BattleState.BattleHistory = append(rivalPlayer.BattleState.BattleHistory, rivalBattlePair)
					rivalPlayer.BattleState.PendingReward = true
					err = db.SavePlayer(rivalPlayer)
					if err != nil {
						helper.InternalErr("Error saving player", err)
						return
					}
					doReward = true
				} else {
					// There appears to be no reward for failures
					// TODO: Guesswork! Confirm if this is correct.
					player.BattleState.Failures++
					player.BattleState.LossStreak++
					player.BattleState.WinStreak = 0
				}
			}
			player.BattleState.BattleStartsAt = now.BeginningOfDay().UTC().Unix()
			player.BattleState.BattleEndsAt = now.EndOfDay().UTC().Unix() + 1
			player.BattleState.ScoreRecordedToday = false
			player.BattleState.MatchedUpWithRival = false
			player.BattleState.WantsStricterMatchmaking = false
		}
	}

	battleStatus := obj.BattleStatus{
		player.BattleState.Wins,
		player.BattleState.Losses,
		player.BattleState.Draws,
		player.BattleState.Failures,
		player.BattleState.WinStreak,
		player.BattleState.LossStreak,
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	var response interface{}
	if doReward {
		response = responses.PostDailyBattleResultWithReward(baseInfo,
			player.BattleState.BattleStartsAt,
			player.BattleState.BattleEndsAt,
			battleStatus,
			rewardBattleStartTime,
			rewardBattleEndTime,
			rewardBattlePlayerData,
			rewardBattleRivalData,
		)
	} else {
		if player.BattleState.ScoreRecordedToday {
			if player.BattleState.MatchedUpWithRival {
				rivalPlayerData, err := db.GetPlayer(player.BattleState.RivalID)
				if err != nil {
					helper.InternalErr("error getting rival player", err)
					return
				}
				response = responses.PostDailyBattleResult(baseInfo,
					player.BattleState.BattleStartsAt,
					player.BattleState.BattleEndsAt,
					conversion.DebugPlayerToBattleData(player),
					conversion.DebugPlayerToBattleData(rivalPlayerData),
					battleStatus,
				)
			} else {
				helper.DebugOut("No rival")
				response = responses.PostDailyBattleResultNoRival(baseInfo,
					player.BattleState.BattleStartsAt,
					player.BattleState.BattleEndsAt,
					conversion.DebugPlayerToBattleData(player),
					battleStatus,
				)
			}
		} else {
			helper.DebugOut("No score recorded")
			response = responses.PostDailyBattleResultNoData(baseInfo,
				player.BattleState.BattleStartsAt,
				player.BattleState.BattleEndsAt,
				battleStatus,
			)
		}
	}

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
}

func GetPrizeDailyBattle(helper *helper.Helper) {
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DefaultGetPrizeDailyBattle(baseInfo)
	err := helper.SendCompatibleResponse(response, true)
	if err != nil {
		helper.InternalErr("Error sending response", err)
		return
	}
}
