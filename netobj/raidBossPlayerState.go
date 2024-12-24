package netobj

import "time"

type RaidBossPlayerState struct {
	NumRaidBossRings       int64   `json:"numRaidbossRings"`       // Number of Power Rings the player has
	RaidBossEnergy         int64   `json:"raidbossEnergy"`         // Number of Boss Challenge Tokens
	RaidBossEnergyBuy      int64   `json:"raidbossEnergyBuy"`      // Number of Boss Challenge Tokens purchased from shop
	NumBeatedEncounter     int64   `json:"numBeatedEncounter"`     // ?
	NumBeatedEnterprise    int64   `json:"numBeatedEnterprise"`    // Seems to be related to the event rewards?
	NumRaidBossEncountered int64   `json:"numTotalEncountered"`    // Total encounters
	EnergyRenewsAt         int64   `json:"raidbossEnergyRenewsAt"` // Next Boss Challenge Token regen
	PlayerRaidBossIDs      []int64 `json:"playerRaidBossIDs"`      // list of IDs for raid bosses
}

func DefaultRaidBossPlayerState() RaidBossPlayerState {
	numRaidbossRings := int64(0)
	raidbossEnergy := int64(20)
	raidbossEnergyBuy := int64(0)
	numBeatedEncounter := int64(0)
	numBeatedEnterprise := int64(0)
	numRaidBossEncountered := int64(0)
	energyRenewsAt := int64(time.Now().UTC().Unix() + 600)
	playerRaidbossIDs := []int64{}
	return RaidBossPlayerState{
		numRaidbossRings,
		raidbossEnergy,
		raidbossEnergyBuy,
		numBeatedEncounter,
		numBeatedEnterprise,
		numRaidBossEncountered,
		energyRenewsAt,
		playerRaidbossIDs,
	}
}

func NewRaidBossPlayerState(numRaidbossRings, raidbossEnergy, raidbossEnergyBuy, numBeatedEncounter, numBeatedEnterprise, numRaidBossEncountered, energyRenewsAt int64, playerRaidbossIDs []int64) RaidBossPlayerState {
	return RaidBossPlayerState{
		numRaidbossRings,
		raidbossEnergy,
		raidbossEnergyBuy,
		numBeatedEncounter,
		numBeatedEnterprise,
		numRaidBossEncountered,
		energyRenewsAt,
		playerRaidbossIDs,
	}
}

// Converts a RaidBossPlayerState to EventUserRaidbossState
func ConvertRaidBossPlayerState(raidbossPlayerState RaidBossPlayerState) EventUserRaidbossState {
	return EventUserRaidbossState{
		raidbossPlayerState.NumRaidBossRings,
		raidbossPlayerState.RaidBossEnergy,
		raidbossPlayerState.RaidBossEnergyBuy,
		raidbossPlayerState.NumBeatedEncounter,
		raidbossPlayerState.NumBeatedEnterprise,
		raidbossPlayerState.NumRaidBossEncountered,
		raidbossPlayerState.EnergyRenewsAt,
	}
}
