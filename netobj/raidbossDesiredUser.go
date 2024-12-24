package netobj

type RaidBossDesiredUser struct {
	DesireID            string `json:"desireId"`
	UserName            string `json:"name"`
	NumRank             int64  `json:"numRank,string"`
	LoginTime           int64  `json:"loginTime,string"`
	CharaID             string `json:"charaId"`
	CharaLevel          int64  `json:"charaLevel,string"`
	SubCharaID          string `json:"subCharaId"`
	SubCharaLevel       int64  `json:"subCharaLevel,string"`
	MainChaoID          string `json:"mainChaoId"`
	MainChaoLevel       int64  `json:"mainChaoLevel,string"`
	SubChaoID           string `json:"subChaoId"`
	SubChaoLevel        int64  `json:"subChaoLevel,string"`
	Language            int64  `json:"language,string"`
	League              int64  `json:"league,string"`
	NumBeatedEnterprise int64  `json:"numBeatedEnterprise,string"`
}

func GetDesiredUserFromPlayer(player Player, desireId string) RaidBossDesiredUser {
	return RaidBossDesiredUser{
		desireId,
		player.Username,
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
		player.RaidBossPlayerState.NumBeatedEnterprise,
	}
}
