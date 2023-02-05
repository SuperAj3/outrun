package requests

type GetDailyBattleHistoryRequest struct {
	Count int64 `json:"count,string"`
}

type ResetDailyBattleMatchingRequest struct {
	Type int64 `json:"type,string"`
}