package netobj

type EventState struct {
	Param    int64 `json:"param"`
	RewardID int64 `json:"rewardId"` // record-keeping, to see which reward ID we're currently on so we don't double-reward gifts
}

func DefaultEventState() EventState {
	param := int64(0)
	rewardID := int64(0)
	return NewEventState(param, rewardID)
}

func NewEventState(param, rewardID int64) EventState {
	return EventState{
		param,
		rewardID,
	}
}
