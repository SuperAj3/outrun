package muxhandlers

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"time"

	"github.com/RunnersRevival/outrun/analytics"
	"github.com/RunnersRevival/outrun/analytics/factors"
	"github.com/RunnersRevival/outrun/config"
	"github.com/RunnersRevival/outrun/config/gameconf"
	"github.com/RunnersRevival/outrun/config/infoconf"
	"github.com/RunnersRevival/outrun/db"
	"github.com/RunnersRevival/outrun/emess"
	"github.com/RunnersRevival/outrun/helper"
	"github.com/RunnersRevival/outrun/logic"
	"github.com/RunnersRevival/outrun/logic/conversion"
	"github.com/RunnersRevival/outrun/netobj"
	"github.com/RunnersRevival/outrun/obj"
	"github.com/RunnersRevival/outrun/obj/constobjs"
	"github.com/RunnersRevival/outrun/requests"
	"github.com/RunnersRevival/outrun/responses"
	"github.com/RunnersRevival/outrun/status"
	"github.com/jinzhu/now"
)

func Login(helper *helper.Helper) {
	recv := helper.GetGameRequest()
	var request requests.LoginRequest
	err := json.Unmarshal(recv, &request)
	if err != nil {
		helper.Err("Error unmarshalling", err)
		return
	}
	uid := request.LineAuth.UserID
	password := request.LineAuth.Password

	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	helper.Out("User logging in with Revival Version ID %v (%s)", request.RevivalVerID, request.Version)
	if request.Version != "2.2.0" && request.Version != "2.0.3" && !config.CFile.LegacyCompatibilityMode {
		helper.Out("Client version too old or too new for this version of Outrun!")
		baseInfo.StatusCode = status.VersionDifference
		response := responses.NewBaseResponse(baseInfo)
		err := helper.SendResponse(response)
		if err != nil {
			helper.InternalErr("Error sending response", err)
		}
		return
	}
	if uid == "0" && password == "" {
		helper.Out("Entering LoginAlpha")
		if request.Version == "2.0.3" && !config.CFile.LegacyCompatibilityMode {
			// message that only shows on 2.0.3
			baseInfo.StatusCode = status.ServerNextVersion
			err = helper.SendResponse(responses.NewNextVersionResponse(baseInfo,
				0,
				0,
				"",
				"Sonic Runners Revival has been updated! This version of Revival is no longer supported. Since you seem to be a new player, you shouldn't need to do much to update to the new version.\n\nWe semi-regularly update Sonic Runners Revival with new features and content that wouldn't be possible without app updates. Newer updates contain security features that aren't present in 2.0.3, so to protect our players, older versions of the game are no longer playable on Sonic Runners Revival.\n\nDownload the new version here: https://www.sonicrunners.com/#/download\nOur Twitter: https://twitter.com/runnersrevival\nOur Discord: https://discord.gg/T5ytR6T",
				"Sonic Runners Revival has been updated! This version of Revival is no longer supported. Since you seem to be a new player, you shouldn't need to do much to update to the new version.\n\nWe semi-regularly update Sonic Runners Revival with new features and content that wouldn't be possible without app updates. Newer updates contain security features that aren't present in 2.0.3, so to protect our players, older versions of the game are no longer playable on Sonic Runners Revival.\n\nDownload the new version here: https://www.sonicrunners.com/#/download\nOur Twitter: https://twitter.com/runnersrevival\nOur Discord: https://discord.gg/T5ytR6T",
				"https://sonicrunners.com/",
			))
			if err != nil {
				helper.InternalErr("Error sending response", err)
			}
			return
		}
		newPlayer := db.NewAccount()
		err = db.SavePlayer(newPlayer)
		if err != nil {
			helper.InternalErr("Error saving player", err)
			return
		}
		baseInfo.StatusCode = status.InvalidPassword
		baseInfo.SetErrorMessage(emess.BadPassword)
		response := responses.LoginRegister(
			baseInfo,
			newPlayer.ID,
			newPlayer.Password,
			newPlayer.Key,
		)
		err = helper.SendResponse(response)
		if err != nil {
			helper.InternalErr("Error responding", err)
		}
		return
	} else if uid == "0" && password != "" {
		helper.Out("Entering LoginBravo")
		// invalid request
		helper.InvalidRequest()
		return
	} else if uid != "0" && password == "" {
		helper.Out("Entering LoginCharlie")
		// game wants to log in
		baseInfo.StatusCode = status.InvalidPassword
		baseInfo.SetErrorMessage(emess.BadPassword)
		player, err := db.GetPlayer(uid)
		if err != nil {
			// likely account that wasn't found, so let's tell them that:
			response := responses.LoginCheckKey(baseInfo, "")
			baseInfo.StatusCode = status.MissingPlayer
			//response.BaseResponse = responses.NewBaseResponseV(baseInfo, request.Version)
			helper.SendResponse(response)
			helper.InternalErr("Error getting player", err)
			return
		}
		if request.Version == "2.0.3" && !config.CFile.LegacyCompatibilityMode {
			// message that only shows on 2.0.3
			baseInfo.StatusCode = status.ServerNextVersion
			err = helper.SendResponse(responses.NewNextVersionResponse(baseInfo,
				player.PlayerState.NumRedRings,
				player.PlayerState.NumBuyRedRings,
				player.Username,
				"Sonic Runners Revival has been updated! This version of Revival is no longer supported. Don't worry about your progress, it should be brought over to the new version automatically when you first start the new version up, JUST AS LONG AS YOU HAVE THIS VERSION OF RUNNERS REVIVAL STILL INSTALLED.\n\nIn the event your data isn't transferred automatically, please take note of your user ID ("+player.ID+"), as you'll need it to ensure that your save data is correct after performing the manual save data transfer steps outlined in our Discord.\n\nWe semi-regularly update Sonic Runners Revival with new features and content that wouldn't be possible without app updates. Newer updates contain security features that aren't present in 2.0.3, so to protect our players, older versions of the game are no longer playable on Sonic Runners Revival.\n\nDownload the new version here: https://www.sonicrunners.com/#/download\nOur Twitter: https://twitter.com/runnersrevival\nOur Discord: https://discord.gg/T5ytR6T",
				"Sonic Runners Revival has been updated! This version of Revival is no longer supported. Don't worry about your progress, it should be brought over to the new version automatically when you first start the new version up, JUST AS LONG AS YOU HAVE THIS VERSION OF RUNNERS REVIVAL STILL INSTALLED.\n\nIn the event your data isn't transferred automatically, please take note of your user ID ("+player.ID+"), as you'll need it to ensure that your save data is correct after performing the manual save data transfer steps outlined in our Discord.\n\nWe semi-regularly update Sonic Runners Revival with new features and content that wouldn't be possible without app updates. Newer updates contain security features that aren't present in 2.0.3, so to protect our players, older versions of the game are no longer playable on Sonic Runners Revival.\n\nDownload the new version here: https://www.sonicrunners.com/#/download\nOur Twitter: https://twitter.com/runnersrevival\nOur Discord: https://discord.gg/T5ytR6T",
				"https://sonicrunners.com/",
			))
			if err != nil {
				helper.InternalErr("Error sending response", err)
			}
			return
		}
		response := responses.LoginCheckKey(baseInfo, player.Key)
		//response.BaseResponse = responses.NewBaseResponseV(baseInfo, request.Version)
		err = helper.SendResponse(response)
		if err != nil {
			helper.InternalErr("Error sending response", err)
			return
		}
		return
	} else if uid != "0" && password != "" {
		helper.Out("Entering LoginDelta")
		// game is attempting to log in using key
		player, err := db.GetPlayer(uid)
		if err != nil {
			// player might not exist
			response := responses.LoginCheckKey(baseInfo, "")
			baseInfo.StatusCode = status.MissingPlayer
			//response.BaseResponse = responses.NewBaseResponseV(baseInfo, request.Version)
			helper.SendResponse(response)
			helper.InternalErr("Error getting player", err)
			return
		}
		if request.Password == logic.GenerateLoginPasskey(player) {
			baseInfo.StatusCode = status.OK
			baseInfo.SetErrorMessage(emess.OK)
			if request.Version == "2.0.3" && !config.CFile.LegacyCompatibilityMode {
				// message that only shows on 2.0.3
				baseInfo.StatusCode = status.ServerNextVersion
				err = helper.SendResponse(responses.NewNextVersionResponse(baseInfo,
					player.PlayerState.NumRedRings,
					player.PlayerState.NumBuyRedRings,
					player.Username,
					"Sonic Runners Revival has been updated! This version of Revival is no longer supported. Don't worry about your progress, it should be brought over to the new version automatically when you first start the new version up (only applies if you're on Android - this feature is not available on the iOS version of Revival), JUST AS LONG AS YOU HAVE THIS VERSION OF RUNNERS REVIVAL STILL INSTALLED.\n\nIn the event your data isn't transferred automatically, please take note of your user ID ("+player.ID+"), as you'll need it to ensure that your save data is correct after performing the manual save data transfer steps outlined in our Discord.\n\nWe semi-regularly update Sonic Runners Revival with new features and content that wouldn't be possible without app updates. Newer updates contain security features that aren't present in 2.0.3, so to protect our players, older versions of the game are no longer playable on Sonic Runners Revival.\n\nDownload the new version here: https://www.sonicrunners.com/#/download\nOur Twitter: https://twitter.com/runnersrevival\nOur Discord: https://discord.gg/T5ytR6T",
					"Sonic Runners Revival has been updated! This version of Revival is no longer supported. Don't worry about your progress, it should be brought over to the new version automatically when you first start the new version up (only applies if you're on Android - this feature is not available on the iOS version of Revival), JUST AS LONG AS YOU HAVE THIS VERSION OF RUNNERS REVIVAL STILL INSTALLED.\n\nIn the event your data isn't transferred automatically, please take note of your user ID ("+player.ID+"), as you'll need it to ensure that your save data is correct after performing the manual save data transfer steps outlined in our Discord.\n\nWe semi-regularly update Sonic Runners Revival with new features and content that wouldn't be possible without app updates. Newer updates contain security features that aren't present in 2.0.3, so to protect our players, older versions of the game are no longer playable on Sonic Runners Revival.\n\nDownload the new version here: https://www.sonicrunners.com/#/download\nOur Twitter: https://twitter.com/runnersrevival\nOur Discord: https://discord.gg/T5ytR6T",
					"https://sonicrunners.com/",
				))
				if err != nil {
					helper.InternalErr("Error sending response", err)
				}
				return
			}
			sid, err := db.AssignSessionID(uid)
			if err != nil {
				helper.InternalErr("Error assigning session ID", err)
				return
			}
			player.LastLogin = time.Now().UTC().Unix()
			player.PlayerVarious.EnergyRecoveryMax = gameconf.CFile.EnergyRecoveryMax
			player.PlayerVarious.EnergyRecoveryTime = gameconf.CFile.EnergyRecoveryTime
			err = db.SavePlayer(player)
			if err != nil {
				helper.InternalErr("Error saving player", err)
				return
			}
			response := responses.LoginSuccess(baseInfo, sid, player.Username, player.PlayerVarious.EnergyRecoveryTime, player.PlayerVarious.EnergyRecoveryMax)
			//response.BaseResponse = responses.NewBaseResponseV(baseInfo, request.Version)
			err = helper.SendResponse(response)
			if err != nil {
				helper.InternalErr("Error sending response", err)
				return
			}
			analytics.Store(player.ID, factors.AnalyticTypeLogins)
		} else {
			baseInfo.StatusCode = status.InvalidPassword
			baseInfo.SetErrorMessage(emess.BadPassword)
			helper.DebugOut("Incorrect passkey sent: \"%s\"", request.Password)
			response := responses.LoginCheckKey(baseInfo, player.Key)
			//response.BaseResponse = responses.NewBaseResponseV(baseInfo, request.Version)
			err = helper.SendResponse(response)
			if err != nil {
				helper.InternalErr("Error sending response", err)
				return
			}
		}
	}
}

func LoginNextVersion(helper *helper.Helper) {
	baseInfo := helper.BaseInfo(emess.OK, status.ServerNextVersion)
	// TODO: can we please not have this message hardcoded?
	err := helper.SendResponse(responses.NewNextVersionResponse(baseInfo,
		0,
		0,
		"",
		"Sonic Runners Revival 2.2.0 is almost here! This update will provide restored functionality for Special Stage and Gacha events, as well as the Offline Practice mode. This version will also bring in a few bugfixes and minor enhancements.\n\nIn preparation for the update's release, we are undergoing maintenance to prepare for the update.\n\nMaintenance Start Date: June 5th, 2022\nMaintenance End Date (estimate): To be announced\n\nDuring that time, Sonic Runners Revival will not be playable. We apologize for any inconvenience this may cause.\n\n- Runners Revival Team\n\nOur website: https://www.sonicrunners.com/\nYouTube: https://www.youtube.com/c/SonicRunnersRevival\nTwitter: https://twitter.com/runnersrevival\nDiscord: https://discord.gg/T5ytR6T",
		"Sonic Runners Revival 2.2.0 is almost here! This update will provide restored functionality for Special Stage and Gacha events, as well as the Offline Practice mode. This version will also bring in a few bugfixes and minor enhancements.\n\nIn preparation for the update's release, we are undergoing maintenance to prepare for the update.\n\nMaintenance Start Date: June 5th, 2022\nMaintenance End Date (estimate): To be announced\n\nDuring that time, Sonic Runners Revival will not be playable. We apologize for any inconvenience this may cause.\n\n- Runners Revival Team\n\nOur website: https://www.sonicrunners.com/\nYouTube: https://www.youtube.com/c/SonicRunnersRevival\nTwitter: https://twitter.com/runnersrevival\nDiscord: https://discord.gg/T5ytR6T",
		"https://sonicrunners.com/",
	))
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}

func LoginMaintenance(helper *helper.Helper) {
	baseInfo := helper.BaseInfo(emess.OK, status.ServerMaintenance)
	err := helper.SendResponse(responses.NewBaseResponse(baseInfo))
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}

func GetVariousParameter(helper *helper.Helper) {
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.VariousParameter(baseInfo, player)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
		return
	}
}

func GetInformation(helper *helper.Helper) {
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	infos := []obj.Information{}
	helper.DebugOut("%v", infoconf.CFile.EnableInfos)
	if infoconf.CFile.EnableInfos {
		for _, ci := range infoconf.CFile.Infos {
			newInfo := conversion.ConfiguredInfoToInformation(ci)
			infos = append(infos, newInfo)
			helper.DebugOut(newInfo.Param)
		}
	}
	operatorInfos := []obj.OperatorInformation{}
	numOpUnread := int64(len(operatorInfos))
	response := responses.Information(baseInfo, infos, operatorInfos, numOpUnread)
	err := helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}

func GetTicker(helper *helper.Helper) {
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DefaultTicker(baseInfo, player)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}

func LoginBonus(helper *helper.Helper) {
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}
	if time.Now().UTC().Unix() > player.LoginBonusState.LoginBonusEndTime {
		player.LoginBonusState = netobj.DefaultLoginBonusState(player.LoginBonusState.CurrentFirstLoginBonusDay)
	}
	doLoginBonus := false
	if time.Now().UTC().Unix() > player.LoginBonusState.NextLoginBonusTime {
		doLoginBonus = true
		player.LoginBonusState.LastLoginBonusTime = time.Now().UTC().Unix()
		player.LoginBonusState.NextLoginBonusTime = now.EndOfDay().UTC().Unix()
		player.LoginBonusState.CurrentLoginBonusDay++
		if gameconf.CFile.EnableStartDashLoginBonus {
			player.LoginBonusState.CurrentFirstLoginBonusDay++
		}
	}
	err = db.SavePlayer(player)
	if err != nil {
		helper.InternalErr("Error saving player", err)
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DefaultLoginBonus(baseInfo, player, doLoginBonus)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}

func LoginBonusSelect(helper *helper.Helper) {
	recv := helper.GetGameRequest()
	var request requests.LoginBonusSelectRequest
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

	loginBonusRewardList := constobjs.DefaultLoginBonusRewardList

	rewardList := []obj.Item{}
	firstRewardList := []obj.Item{}
	if request.FirstRewardDays != -1 && int(request.FirstRewardDays) < len(constobjs.DefaultFirstLoginBonusRewardList) {
		firstRewardList = constobjs.DefaultFirstLoginBonusRewardList[request.FirstRewardDays].SelectRewardList[request.FirstRewardSelect].ItemList
	}
	if request.RewardDays != -1 && int(request.RewardDays) < len(loginBonusRewardList) {
		rewardList = loginBonusRewardList[request.RewardDays].SelectRewardList[request.RewardSelect].ItemList
	}
	for _, item := range rewardList {
		itemid, _ := strconv.Atoi(item.ID)
		player.AddOperatorMessage(
			"A Login Bonus.",
			obj.MessageItem{
				int64(itemid),
				item.Amount,
				0,
				0,
			},
			2592000,
		)
		helper.DebugOut("Sent %s x %v to gift box (Login Bonus)", item.ID, item.Amount)
	}
	for _, item := range firstRewardList {
		itemid, _ := strconv.Atoi(item.ID)
		player.AddOperatorMessage(
			"A Debut Dash Login Bonus.",
			obj.MessageItem{
				int64(itemid),
				item.Amount,
				0,
				0,
			},
			2592000,
		)
		helper.DebugOut("Sent %s x %v to gift box (Start Dash Login Bonus)", item.ID, item.Amount)
	}
	err = db.SavePlayer(player)
	if err != nil {
		helper.InternalErr("Error saving player", err)
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.LoginBonusSelect(baseInfo, rewardList, firstRewardList)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}

func GetCountry(helper *helper.Helper) {
	// TODO: Should get correct country code!
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DefaultGetCountry(baseInfo)
	err := helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}

func GetMigrationPassword(helper *helper.Helper) {
	randChar := func(charset string, length int64) string {
		runes := []rune(charset)
		final := make([]rune, 10)
		for i := range final {
			final[i] = runes[rand.Intn(len(runes))]
		}
		return string(final)
	}
	recv := helper.GetGameRequest()
	var request requests.GetMigrationPasswordRequest
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
	if len(player.MigrationPassword) != 10 {
		player.MigrationPassword = randChar("abcdefghijklmnopqrstuvwxyz1234567890", 10)
	}
	player.UserPassword = request.UserPassword
	db.SavePlayer(player)
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.MigrationPassword(baseInfo, player)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}

func Migration(helper *helper.Helper) {
	randChar := func(charset string, length int64) string {
		runes := []rune(charset)
		final := make([]rune, 10)
		for i := range final {
			final[i] = runes[rand.Intn(len(runes))]
		}
		return string(final)
	}

	recv := helper.GetGameRequest()
	var request requests.LoginRequest
	err := json.Unmarshal(recv, &request)
	if err != nil {
		helper.Err("Error unmarshalling", err)
		return
	}
	password := request.LineAuth.MigrationPassword
	migrationUserPassword := request.LineAuth.MigrationUserPassword

	baseInfo := helper.BaseInfo(emess.OK, status.OK)

	helper.DebugOut("Transfer ID: %s", password)
	foundPlayers, err := logic.FindPlayersByMigrationPassword(password, false)
	if err != nil {
		helper.Err("Error finding players by password", err)
		return
	}
	playerIDs := []string{}
	for _, player := range foundPlayers {
		playerIDs = append(playerIDs, player.ID)
	}
	if len(playerIDs) > 0 {
		migratePlayer, err := db.GetPlayer(playerIDs[0])
		if err != nil {
			helper.InternalErr("Error getting player", err)
			return
		}
		if migrationUserPassword == migratePlayer.UserPassword {
			baseInfo.StatusCode = status.OK
			baseInfo.SetErrorMessage(emess.OK)
			migratePlayer.MigrationPassword = randChar("abcdefghijklmnopqrstuvwxyz1234567890", 10) //generate a brand new transfer ID
			migratePlayer.UserPassword = ""                                                        //clear user password
			migratePlayer.LastLogin = time.Now().UTC().Unix()
			err = db.SavePlayer(migratePlayer)
			if err != nil {
				helper.InternalErr("Error saving player", err)
				return
			}
			sid, err := db.AssignSessionID(migratePlayer.ID)
			if err != nil {
				helper.InternalErr("Error assigning session ID", err)
				return
			}
			helper.DebugOut("User ID: %s", migratePlayer.ID)
			helper.DebugOut("Username: %s", migratePlayer.Username)
			helper.DebugOut("New Transfer ID: %s", migratePlayer.MigrationPassword)
			response := responses.MigrationSuccess(baseInfo, sid, migratePlayer.ID, migratePlayer.Username, migratePlayer.Password, migratePlayer.PlayerVarious.EnergyRecoveryTime, migratePlayer.PlayerVarious.EnergyRecoveryMax)
			helper.SendResponse(response)
		} else {
			baseInfo.StatusCode = status.InvalidPassword
			baseInfo.SetErrorMessage(emess.BadPassword)
			helper.DebugOut("Incorrect password for user ID %s", migratePlayer.ID)
			response := responses.NewBaseResponseV(baseInfo, request.Version)
			helper.SendResponse(response)
		}
	} else {
		helper.DebugOut("Failed to find player")
		baseInfo.StatusCode = status.InvalidPassword
		response := responses.NewBaseResponseV(baseInfo, request.Version)
		helper.SendResponse(response)
	}
}
