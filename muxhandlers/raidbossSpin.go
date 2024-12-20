package muxhandlers

import (
	"github.com/RunnersRevival/outrun/emess"
	"github.com/RunnersRevival/outrun/enums"
	"github.com/RunnersRevival/outrun/helper"
	"github.com/RunnersRevival/outrun/netobj"
	"github.com/RunnersRevival/outrun/responses"
	"github.com/RunnersRevival/outrun/status"
)

func GetItemStockNum(helper *helper.Helper) {
	// TODO: Flesh out properly! The game responds with
	// [IDRouletteTicketPremium, IDRouletteTicketItem, IDSpecialEgg]
	// for item IDs, along with an event ID, likely for event characters.
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DefaultItemStockNum(baseInfo)
	err := helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}

func GetRaidbossWheelOptions(helper *helper.Helper) {
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	wheelOptions := netobj.DefaultRaidbossWheelOptions(0, player.PlayerState.ChaoEggs, 0, enums.WheelRankNormal, 0)
	response := responses.RaidbossWheelOptions(baseInfo, wheelOptions)
	err = helper.SendCompatibleResponse(response, true) // For some reason the game expects an unencrypted response for all raid boss endpoints.
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}

func GetPrizeRaidbossWheelSpin(helper *helper.Helper) {
	// agnostic
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DefaultPrizeRaidbossWheel(baseInfo)
	err := helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}
