package netobj

import (
	"github.com/RunnersRevival/outrun/obj"
	"github.com/jinzhu/now"
)

type BattleState struct {
	ScoreRecordedToday       bool                 `json:"hasRecordedScoreToday"`
	DailyBattleHighScore     int64                `json:"maxScore"`
	PrevDailyBattleHighScore int64                `json:"lastMaxScore"`
	BattleStartsAt           int64                `json:"startTime"`
	BattleEndsAt             int64                `json:"expireTime"`
	MatchedUpWithRival       bool                 `json:"matchedUpWithRival"`
	RivalID                  string               `json:"rivalId"`
	Wins                     int64                `json:"numWin"`
	Losses                   int64                `json:"numLose"`
	Draws                    int64                `json:"numDraw"` // very rarely gets incremented, but game expects it so it's here
	Failures                 int64                `json:"numLoseByDefault"`
	WinStreak                int64                `json:"goOnWin"`
	LossStreak               int64                `json:"goOnLosses"`
	BattleHistory            []obj.BattlePair     `json:"battleDataHistory"`
	PendingReward            bool                 `json:"pendingReward"`
	PendingRewardData        obj.RewardBattlePair `json:"pendingRewardData"`
	WantsStricterMatchmaking bool                 `json:"wantsStricterMatchmaking"`
}

func NewBattleState(scoreRecordedToday bool, dailyBattleHighScore, prevDailyBattleHighScore, battleStartTime, battleEndTime int64, matchedUpWithRival bool, rivalID string, wins, losses, draws, failures, winStreak, lossStreak int64, battleHistory []obj.BattlePair, pendingReward bool, pendingRewardData obj.RewardBattlePair, wantsStricterMatchmaking bool) BattleState {
	return BattleState{
		scoreRecordedToday,
		dailyBattleHighScore,
		prevDailyBattleHighScore,
		battleStartTime,
		battleEndTime,
		matchedUpWithRival,
		rivalID,
		wins,
		losses,
		draws,
		failures,
		winStreak,
		lossStreak,
		battleHistory,
		pendingReward,
		pendingRewardData,
		wantsStricterMatchmaking,
	}
}

func DefaultBattleState() BattleState {
	scoreRecordedToday := false
	dailyBattleHighScore := int64(0)
	prevDailyBattleHighScore := int64(0)
	battleStartTime := now.BeginningOfDay().UTC().Unix()
	battleEndTime := now.EndOfDay().UTC().Unix() + 1
	matchedUpWithRival := false
	rivalID := ""
	wins := int64(0)
	losses := int64(0)
	draws := int64(0)
	failures := int64(0)
	winStreak := int64(0)
	lossStreak := int64(0)
	battleHistory := []obj.BattlePair{}
	pendingReward := false
	pendingRewardData := obj.NewRewardBattlePair(-1, -1, obj.DebugRivalBattleData(), obj.DebugRivalBattleData()) //dummy data
	wantsStricterMatchmaking := false
	return NewBattleState(
		scoreRecordedToday,
		dailyBattleHighScore,
		prevDailyBattleHighScore,
		battleStartTime,
		battleEndTime,
		matchedUpWithRival,
		rivalID,
		wins,
		losses,
		draws,
		failures,
		winStreak,
		lossStreak,
		battleHistory,
		pendingReward,
		pendingRewardData,
		wantsStricterMatchmaking,
	)
}
