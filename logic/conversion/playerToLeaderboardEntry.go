package conversion

import (
	"github.com/Mtbcooler/outrun/enums"
	"github.com/Mtbcooler/outrun/netobj"
	"github.com/Mtbcooler/outrun/obj"
	"github.com/jinzhu/now"
)

func PlayerToLeaderboardEntry(player netobj.Player, mode, revivalVerId int64) obj.LeaderboardEntry {
	friendID := player.ID
	name := player.Username
	url := player.ID + "_findme" // TODO: where is this used?
	grade := int64(1)
	exposeOnline := int64(0)
	rankingScore := player.PlayerState.HighScore
	rankChanged := int64(0)
	isSentEnergy := int64(0)
	expireTime := now.EndOfWeek().UTC().Unix()
	numRank := player.PlayerState.Rank
	loginTime := player.LastLogin
	mainCharaID := player.PlayerState.MainCharaID
	mainCharaLevel := int64(0)
	subCharaID := player.PlayerState.SubCharaID
	subCharaLevel := int64(0)
	mainChaoID := player.PlayerState.MainChaoID
	mainChaoLevel := int64(0)
	subChaoID := player.PlayerState.SubChaoID
	subChaoLevel := int64(0)
	if revivalVerId < 1 { // before 2.0.4
		if mainCharaID == enums.CTStrMarine {
			mainCharaID = enums.CTStrTikal
		}
		if subCharaID == enums.CTStrMarine {
			subCharaID = enums.CTStrTikal
		}
	}
	if player.IndexOfChara(mainCharaID) != -1 {
		mainCharaLevel = player.CharacterState[player.IndexOfChara(mainCharaID)].Level
	}
	if player.IndexOfChara(subCharaID) != -1 {
		subCharaLevel = player.CharacterState[player.IndexOfChara(subCharaID)].Level
	}
	if player.IndexOfChao(mainChaoID) != -1 {
		mainChaoLevel = player.ChaoState[player.IndexOfChao(mainChaoID)].Level
	}
	if player.IndexOfChao(subChaoID) != -1 {
		subChaoLevel = player.ChaoState[player.IndexOfChao(subChaoID)].Level
	}
	language := int64(enums.LangEnglish)
	league := player.PlayerState.RankingLeague
	maxScore := player.PlayerState.HighScore
	if mode == 1 { //timed mode?
		rankingScore = player.PlayerState.TimedHighScore
		league = player.PlayerState.QuickRankingLeague
		maxScore = player.PlayerState.TimedHighScore
	}
	return obj.LeaderboardEntry{
		friendID,
		name,
		url,
		grade,
		exposeOnline,
		rankingScore,
		rankChanged,
		isSentEnergy,
		expireTime,
		numRank,
		loginTime,
		mainCharaID,
		mainCharaLevel,
		subCharaID,
		subCharaLevel,
		mainChaoID,
		mainChaoLevel,
		subChaoID,
		subChaoLevel,
		language,
		league,
		maxScore,
		0,
	}
}
