package netobj

import (
	"strconv"
	"time"
)

type RaidBossUserState struct {
	WrestleID         string `json:"friendId"`
	Name              string `json:"name"`
	Grade             int64  `json:"grade,string"`
	NumRank           int64  `json:"numRank,string"`
	LoginTime         int64  `json:"loginTime,string"`
	CharacterID       string `json:"charaId"`
	CharacterLevel    int64  `json:"characterLevel,string"`
	SubcharacterID    string `json:"subCharaId"`
	SubcharacterLevel int64  `json:"subCharaLevel,string"`
	MainChaoID        string `json:"mainChaoId"`
	MainChaoLevel     int64  `json:"mainChaoLevel,string"`
	SubChaoID         string `json:"subChaoId"`
	SubChaoLevel      int64  `json:"subChaoLevel,string"`
	Language          int64  `json:"language,string"` // enums.Lang*
	League            int64  `json:"league,string"`
	WrestleCount      int64  `json:"wrestleCount,string"`
	WrestleDamage     int64  `json:"wrestleDamage,string"`
	WrestleBeatFlg    int64  `json:"wrestleBeatFlg,string"`
}

func NewRaidBossUserState(wid, n string, g, nr, lt, cid, cl, schid, schl, mcid, mcl, scid, scl, lang, league, wc, wd, wbf int64) RaidBossUserState {
	return RaidBossUserState{
		wid,
		n,
		g,
		nr,
		lt,
		strconv.Itoa(int(cid)),
		cl,
		strconv.Itoa(int(schid)),
		schl,
		strconv.Itoa(int(mcid)),
		mcl,
		strconv.Itoa(int(scid)),
		scl,
		lang,
		league,
		wc,
		wd,
		wbf,
	}
}

func DefaultRaidBossUserState(uid string) RaidBossUserState {
	return NewRaidBossUserState(
		uid,
		"",
		1,
		0,
		time.Now().UTC().Unix(),
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		1,
		0,
		0,
		0,
		0,
	)
}

func GetRaidBossUserFromPlayer(player Player, grade, wrestleCount, wrestleDamage, wrestleBeatFlg int64) RaidBossUserState {
	return RaidBossUserState{
		player.ID,
		player.Username,
		grade,
		player.PlayerState.Rank,
		player.LastLogin,
		player.PlayerState.MainCharaID,
		0,
		player.PlayerState.SubCharaID,
		0,
		player.PlayerState.MainChaoID,
		player.PlayerState.MainChaoLevel,
		player.PlayerState.SubChaoID,
		player.PlayerState.SubChaoLevel,
		0,
		player.PlayerState.RankingLeague,
		wrestleCount,
		wrestleDamage,
		wrestleBeatFlg,
	}
}
