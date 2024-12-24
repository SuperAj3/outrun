package netobj

type RaidBossBonus struct {
	RaidBossEncounterBonus  int64 `json:"eventRaidbossEncounterBonus,string"`
	RaidBossWrestleBonus    int64 `json:"eventRaidbossWrestleBonus,string"`
	RaidBossDamageRateBonus int64 `json:"eventRaidbossDamageRateBonus,string"`
	RaidBossDamageTopBonus  int64 `json:"eventRaidbossDamageTopBonus,string"`
	RaidBossBeatBonus       int64 `json:"eventRaidbossBeatBonus,string"`
}
