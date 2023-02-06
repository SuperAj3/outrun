package consts

const (
	DBFileName = "outrun.db"
	DBBucketSessionIDs = "sessionIDs"
	DBBucketPlayers    = "players"
	DBBucketAnalytics  = "analytics"
	DBSessionExpiryTime = 5600
)

const (
	BattleDBFileName = "battle.db"
	BattleDBBucketSessionIDs = "sessionIDs"
	BattleDBBucketMatched    = "matched"
	BattleDBBucketWaiting  = "waiting"
	BattleDBSessionExpiryTime = 5600
)
