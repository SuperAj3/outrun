package rpcobj

import (
	"github.com/RunnersRevival/outrun/db"
	"github.com/RunnersRevival/outrun/netobj"
	"github.com/RunnersRevival/outrun/obj"
)

type Toolbox struct {
}

func (t *Toolbox) RegisterPlayerWithID(uid string, reply *ToolboxReply) error {
	player := db.NewAccountWithID(uid, 0)
	err := db.SavePlayer(player)
	if err != nil {
		reply.Status = StatusOtherError
		reply.Info = "unable to save player: " + err.Error()
		return err
	}
	reply.Status = StatusOK
	reply.Info = "OK"
	return nil
}

func (t *Toolbox) FetchPlayer(uid string, reply *netobj.Player) error {
	player, err := db.GetPlayer(uid)
	if err != nil {
		return err
	}
	*reply = player
	return nil
}

type ToolboxReply struct {
	Status uint
	Info   string
}

type ToolboxValueReply struct {
	Status uint
	Result interface{}
}

type ToolboxPlayerCountReply struct {
	Status               uint
	ActiveCount          int64
	TotalRegisteredCount int64
	AdditionalInfo       string
}

type ChangeValueArgs struct {
	UID   string
	Value interface{}
}

type SendOperatorMessageToAllArgs struct {
	MessageContents string
	Item            obj.MessageItem
	ExpiresAfter    int64
}

type SendOperatorMessageArgs struct {
	UID             string
	MessageContents string
	Item            obj.MessageItem
	ExpiresAfter    int64
}

type SendOperatorMessageToSomeArgs struct {
	UIDs            string
	MessageContents string
	Item            obj.MessageItem
	ExpiresAfter    int64
}

type ChangeCharacter struct {
	ID            string `json:"characterId"`
	Cost          int64  `json:"numRings"`         // interestingly, is used for both buying the character and for levelling up...
	NumRedRings   int64  `json:"numRedRings"`      // ?
	Price         int64  `json:"priceNumRings"`    // used to limit break, as far as I can tell?
	PriceRedRings int64  `json:"priceNumRedRings"` // ?
	LockCondition int64
	UIDs          string
}
