package netobj

type GlobalState struct {
	TotalRuns            int64    `json:"totalRuns"`
	LastRaidBossID       int64    `json:"lastRaidBossId"`
	RaidBossPlayerIDPool []string `json:"raidBossPlayerIdPool"`
}
