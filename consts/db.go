package consts

const (
	DBFileName            = "outrun.db"
	DBBucketSessionIDs    = "sessionIDs"
	DBBucketPlayers       = "players"
	DBBucketAnalytics     = "analytics"
	DBBucketTransferCreds = "transferCreds"
	DBBucketGlobalParams  = "globalParams"
	DBSessionExpiryTime   = 5600
)

const (
	BattleDBFileName          = "battle.db"
	BattleDBBucketSessionIDs  = "sessionIDs"
	BattleDBBucketMatched     = "matched"
	BattleDBBucketWaiting     = "waiting"
	BattleDBSessionExpiryTime = 5600
)

const (
	RaidbossDBFileName                 = "raid.db"
	RaidbossDBBucketStandardRaidBosses = "standard"
	RaidbossDBBucketGlobalRaidBosses   = "global"
)
