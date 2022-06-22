package obj

type EventReward struct {
    Item
    RewardID int64  `json:"rewardId"` 
    Param    int64  `json:"param"`    
}

func NewEventReward(rewardId, param int64, itemId string, numItem int64) EventReward {
    return EventReward{
        NewItem(itemId, numItem),
        rewardId,
        param,
    }
}

func ItemToEventReward(i Item, rid int64, param int64) EventReward {
    return EventReward{
        i,
        rid,
        param,
    }
}