package conversion

import (
	"github.com/RunnersRevival/outrun/enums"
	"github.com/RunnersRevival/outrun/netobj"
	"github.com/RunnersRevival/outrun/obj"
)

func DebugPlayerToBattleData(player netobj.Player) obj.BattleData {
	uid := player.ID
	username := player.Username
	maxScore := player.BattleState.DailyBattleHighScore // TODO: Should this be the daily high score?
	league := player.PlayerState.RankingLeague
	loginTime := player.LastLogin
	mainChaoID := player.PlayerState.MainChaoID
	mainChaoLevel := player.PlayerState.MainChaoLevel // TODO: this may be problematic if the game does checks
	subChaoID := player.PlayerState.SubChaoID
	subChaoLevel := player.PlayerState.SubChaoLevel // TODO: this may be problematic if the game does checks
	rank := player.PlayerState.Rank
	mainCharaID := player.PlayerState.MainCharaID
	mainCharaLevel := player.PlayerState.MainCharaLevel // TODO: this may be problematic if the game does checks
	subCharaID := player.PlayerState.SubCharaID
	subCharaLevel := int64(0) // TODO: this may be problematic if the game does checks
	goOnWin := player.BattleState.WinStreak
	isSentEnergy := int64(0)
	language := int64(enums.LangEnglish)
	return obj.BattleData{
		uid,
		username,
		maxScore,
		league,
		loginTime,
		mainChaoID,
		mainChaoLevel,
		subChaoID,
		subChaoLevel,
		rank,
		mainCharaID,
		mainCharaLevel,
		subCharaID,
		subCharaLevel,
		goOnWin,
		isSentEnergy,
		language,
	}
}
