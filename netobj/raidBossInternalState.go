package netobj

type RaidBossInternalState struct {
	ID            int64                         `json:"raidbossId"`
	Level         int64                         `json:"raidbossLevel"`
	Rarity        int64                         `json:"raidbossRarity"` // 0: Normal, 1: Rare, 2: Super-rare
	HP            int64                         `json:"raidbossHitPoint"`
	MaxHP         int64                         `json:"raidbossMaxHitPoint"`
	Status        int64                         `json:"raidbossStatus"`   // TODO: research
	EscapeAt      int64                         `json:"raidbossEscapeAt"` // the time when the raid boss expires
	EncounterName string                        `json:"encounterName"`
	MaxPlayers    int64                         `json:"ORN_maxPlayers"`
	PlayerStates  []RaidBossInternalPlayerState `json:"ORN_playerStates"`
}

type RaidBossInternalPlayerState struct {
	PlayerID    int64 `json:"playerId"`
	NumBattles  int64 `json:"numBattles"`  // translates to wrestleCount
	DamageTotal int64 `json:"damageTotal"` // translates to wrestleDamage
	HadFinalHit bool  `json:"hadFinalHit"` // translates to wrestleBeatFlg
	Grade       int64 `json:"grade"`
}

func DefaultRaidBossInternalState() RaidBossInternalState {
	return RaidBossInternalState{
		0,
		1,
		0,
		10,
		10,
		0,
		-1,
		"Unknown",
		10,
		[]RaidBossInternalPlayerState{},
	}
}

func GetRaidBossState(state RaidBossInternalState, encounterFlg int64) EventRaidbossState {
	crowdedFlg := int64(0)
	if len(state.PlayerStates) >= int(state.MaxPlayers) {
		crowdedFlg = 1
	}
	return EventRaidbossState{
		state.ID,
		state.Level,
		state.Rarity,
		state.HP,
		state.MaxHP,
		state.Status,
		state.EscapeAt,
		state.EncounterName,
		encounterFlg,
		crowdedFlg,
		int64(len(state.PlayerStates)),
	}
}

func ToEmptyRaidBossInternalState(state EventRaidbossState) RaidBossInternalState {
	return RaidBossInternalState{
		state.ID,
		state.Level,
		state.Rarity,
		state.HP,
		state.MaxHP,
		state.Status,
		state.EscapeAt,
		state.EncounterName,
		10,
		[]RaidBossInternalPlayerState{},
	}
}
