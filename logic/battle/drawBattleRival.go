package battle

import (
	"log"
	"math/rand"
	"time"

	"github.com/RunnersRevival/outrun/consts"
	"github.com/RunnersRevival/outrun/db"
	"github.com/RunnersRevival/outrun/db/dbaccess"
	"github.com/RunnersRevival/outrun/netobj"
)

func DrawBattleRival(player netobj.Player) netobj.BattleState {
	if !player.BattleState.MatchedUpWithRival { // are we not matched up yet?
		db.BattleSaveWaitingPlayer(player) // Save player in the waiting pool
		playerIDs := []string{}
		dbaccess.BattleDBForEachKey(consts.BattleDBBucketWaiting, func(k, v []byte) error {
			playerIDs = append(playerIDs, string(k))
			return nil
		})
		potentialRivalIDs := []string{}
		for _, pid := range playerIDs {
			potentialRival, err := db.GetPlayer(pid)
			if err != nil {
				log.Printf("[WARN] (battle.DrawBattleRival) Unable to get player '%s': %s", pid, err.Error())
			} else {
				if player.ID != pid && potentialRival.BattleState.ScoreRecordedToday && !potentialRival.BattleState.MatchedUpWithRival && time.Now().UTC().Unix() < potentialRival.BattleState.BattleEndsAt {
					potentialRivalIDs = append(potentialRivalIDs, potentialRival.ID)
				}
			}
		}
		if len(potentialRivalIDs) > 0 {
			rivalID := potentialRivalIDs[rand.Intn(len(potentialRivalIDs))]
			rival, err := db.GetPlayer(rivalID)
			if err != nil {
				log.Printf("[WARN] (battle.DrawBattleRival) Unable to get player '%s': %s", rivalID, err.Error())
			} else {
				rival.BattleState.RivalID = player.ID
				rival.BattleState.MatchedUpWithRival = true
				db.BattleDeleteWaitingPlayer(player.ID)
				db.BattleDeleteWaitingPlayer(rival.ID)
				db.BattleSaveMatchedPlayer(player)
				db.BattleSaveMatchedPlayer(rival)
				err = db.SavePlayer(rival)
				if err != nil {
					log.Printf("[WARN] (battle.DrawBattleRival) Unable to save rival data: %s", err.Error())
				} else {
					player.BattleState.RivalID = rivalID
					db.BattleDeleteWaitingPlayer(rival.ID)
					db.BattleDeleteWaitingPlayer(rival.ID)
					db.BattleSaveMatchedPlayer(player)
					db.BattleSaveMatchedPlayer(rival)
					player.BattleState.MatchedUpWithRival = true
				}
			}
		}
	} else {
		db.BattleDeleteWaitingPlayer(player.ID) // Player is matched, delete entry from the waiting pool
	}
	if player.ID == player.BattleState.RivalID && player.BattleState.MatchedUpWithRival {
		log.Printf("[WARN] (battle.DrawBattleRival) Somehow matched up with self! Removing match...")
		player.BattleState.MatchedUpWithRival = false
		db.BattleDeleteMatchedPlayer(player.ID)
		db.BattleSaveWaitingPlayer(player)
	}
	return player.BattleState
}