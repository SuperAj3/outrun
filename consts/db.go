package consts

const (
	DBFileName            = "outrun.db"
	DBBucketSessionIDs    = "sessionIDs"
	DBBucketPlayers       = "players"
	DBBucketAnalytics     = "analytics"
	DBBucketTransferCreds = "transferCreds"
	DBSessionExpiryTime   = 5600
)

const (
	BattleDBFileName          = "battle.db"
	BattleDBBucketSessionIDs  = "sessionIDs"
	BattleDBBucketMatched     = "matched"
	BattleDBBucketWaiting     = "waiting"
	BattleDBSessionExpiryTime = 5600
)
