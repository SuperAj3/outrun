package netobj

import "time"

type EventRaidbossState struct {
	ID               int64  `json:"raidbossId"`
	Level            int64  `json:"raidbossLevel"`
	Rarity           int64  `json:"raidbossRarity"` // 0: Normal, 1: Rare, 2: Super-rare
	HP               int64  `json:"raidbossHitPoint"`
	MaxHP            int64  `json:"raidbossMaxHitPoint"`
	Status           int64  `json:"raidbossStatus"`   // TODO: research
	EscapeAt         int64  `json:"raidbossEscapeAt"` // the time when the raid boss expires
	EncounterName    string `json:"encounterName"`
	EncounterFlg     int64  `json:"encounterFlg"`
	CrowdedFlg       int64  `json:"crowdedFlg"`       // raid boss is full?
	ParticipateCount int64  `json:"participateCount"` // number of people taking on this raid boss?
}

func NewRaidbossState(id, level, rarity int64, encounterName string) EventRaidbossState {
	if level < 1 {
		level = 1
	}
	hpMult := int64(1)
	if rarity == 1 {
		hpMult = 2
	}
	if rarity == 2 {
		hpMult = 6
	}
	hpOffset := ((level - 1) / 10) * 10
	maxHp := int64((10 + ((level - 1) * 2) + hpOffset) * hpMult)
	hp := maxHp
	status := int64(0)
	escapeAt := int64(time.Now().UTC().Unix() + 3600) // raid boss expires after 1 hour
	encounterFlg := int64(0)
	crowdedFlg := int64(0)
	participateCount := int64(0)
	return EventRaidbossState{
		id,
		level,
		rarity,
		hp,
		maxHp,
		status,
		escapeAt,
		encounterName,
		encounterFlg,
		crowdedFlg,
		participateCount,
	}
}

// Default normal raid boss state (level 2)
func DefaultRaidbossState() EventRaidbossState {
	// TODO: establish as constants
	id := int64(0)
	level := int64(1)
	rarity := int64(0)
	encounterName := "System"
	return NewRaidbossState(
		id,
		level,
		rarity,
		encounterName,
	)
}

// Default rare raid boss state (level 3)
func DefaultRaidbossState2() EventRaidbossState {
	// TODO: establish as constants
	id := int64(1)
	level := int64(3)
	rarity := int64(1)
	encounterName := "System"
	return NewRaidbossState(
		id,
		level,
		rarity,
		encounterName,
	)
}

// Default super-rare raid boss state (level 5)
func DefaultRaidbossState3() EventRaidbossState {
	// TODO: establish as constants
	id := int64(2)
	level := int64(5)
	rarity := int64(2)
	encounterName := "System"
	return NewRaidbossState(
		id,
		level,
		rarity,
		encounterName,
	)
}
