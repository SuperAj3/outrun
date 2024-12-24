package requests

type ItemStockNumRequest struct {
	Base
	EventID int64   `json:"eventId,string"`
	ItemIDs []int64 `json:"itemIdList"`
}

type CommitRaidbossWheelSpinRequest struct {
	EventID    int64 `json:"eventId,string"`
	SpinID     int64 `json:"id,string"`
	CostItemID int64 `json:"costItemId,string"` // Usually for the Raid Boss Roulette it's Power Rings (960000), although other roulettes such as the Event Roulette can differ in this regard.
	SpinNum    int64 `json:"num,string"`
}
