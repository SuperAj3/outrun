package netobj

import (
	"github.com/RunnersRevival/outrun/config/gameconf"
)

type PlayerVarious struct {
	CmSkipCount        int64 `json:"cmSkipCount"` // no clear purpose
	EnergyRecoveryMax  int64 `json:"energyRecoveryMax"`
	EnergyRecoveryTime int64 `json:"energyRecveryTime"`
	OnePlayCmCount     int64 `json:"onePlayCmCount"` // number of ad continues per run??? no effect in 2.x
	//OnePlayDailyFreeCount int64 `json:"onePlayDailyFreeCount"` // number of daily free continues (not currently implemented)
	OnePlayContinueCount int64 `json:"onePlayContinueCount"` // number of RSR continues per run
	IsPurchased          int64 `json:"isPurchased"`          // removes ads? doesn't seem to do anything in 2.x
}

func DefaultPlayerVarious() PlayerVarious {
	cmSkipCount := int64(5)
	energyRecoveryMax := gameconf.CFile.EnergyRecoveryMax
	energyRecoveryTime := gameconf.CFile.EnergyRecoveryTime
	onePlayCmCount := int64(0)
	//onePlayDailyFreeCount := int64(0)
	onePlayContinueCount := int64(2)
	isPurchased := int64(0)
	return PlayerVarious{
		cmSkipCount,
		energyRecoveryMax,
		energyRecoveryTime,
		onePlayCmCount,
		//onePlayDailyFreeCount,
		onePlayContinueCount,
		isPurchased,
	}
}
