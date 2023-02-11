package obj

import "sort"

type BattlePair struct { // This is just used for organization within the response
	StartTime       int64      `json:"startTime"`
	EndTime         int64      `json:"endTime"`
	BattleData      BattleData `json:"battleData"`
	RivalBattleData BattleData `json:"rivalBattleData"`
}

func NewBattlePair(startTime, endTime int64, battleData, rivalBattleData BattleData) BattlePair {
	return BattlePair{
		startTime,
		endTime,
		battleData,
		rivalBattleData,
	}
}

type RewardBattlePair struct { // This is just used for organization within the response
	StartTime       int64      `json:"rewardStartTime"`
	EndTime         int64      `json:"rewardEndTime"`
	BattleData      BattleData `json:"rewardBattleData"`
	RivalBattleData BattleData `json:"rewardRivalBattleData"`
}

func NewRewardBattlePair(startTime, endTime int64, battleData, rivalBattleData BattleData) RewardBattlePair {
	return RewardBattlePair{
		startTime,
		endTime,
		battleData,
		rivalBattleData,
	}
}

// The game expects the battle pairs to be sorted by newest first. If it isn't, "Did not join" entries show up when they're not supposed to.
func SortBattlePairsByNewestFirst(pairs []BattlePair) []BattlePair {
	sort.Slice(pairs, func(i, j int) bool {
		return (int)(pairs[i].EndTime-pairs[i].StartTime) > (int)(pairs[j].EndTime-pairs[i].StartTime)
	})
	return pairs
}

func BattlePairsAreSorted(pairs []BattlePair) bool {
	return sort.SliceIsSorted(pairs, func(i, j int) bool {
		return (int)(pairs[i].EndTime-pairs[i].StartTime) > (int)(pairs[j].EndTime-pairs[i].StartTime)
	})
}
