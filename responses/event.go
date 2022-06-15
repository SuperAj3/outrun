package responses

import (
	"strconv"

	"github.com/Mtbcooler/outrun/enums"
	"github.com/Mtbcooler/outrun/logic"
	"github.com/Mtbcooler/outrun/netobj"
	"github.com/Mtbcooler/outrun/obj"
	"github.com/Mtbcooler/outrun/responses/responseobjs"
)

type EventListResponse struct {
	BaseResponse
	EventList []obj.Event `json:"eventList"`
}

func EventList(base responseobjs.BaseInfo, eventList []obj.Event) EventListResponse {
	baseResponse := NewBaseResponse(base)
	out := EventListResponse{
		baseResponse,
		eventList,
	}
	return out
}

func DefaultEventList(base responseobjs.BaseInfo) EventListResponse {
	return EventList(
		base,
		[]obj.Event{
			/*
			   obj.NewEvent(
			       //enums.EventIDSpecialStage+10002, // game subtracts one from number?
			       //enums.EventIDAdvert+50002, // 50002 converts to ui_event_50005_Atlas_en?
			       //enums.EventIDBGM+70002, // 70002 goes to 70007
			       enums.EventIDQuick+60002, // 60002 goes to 60006
			       0,                        // event type
			       now.BeginningOfDay().Unix(),
			       now.EndOfDay().Unix(),
			       now.EndOfDay().Unix(),
			   ),
			*/
		},
	)
}

type EventRewardListResponse struct {
	BaseResponse
	EventRewardList []obj.EventReward `json:"eventRewardList"`
}

func EventRewardList(base responseobjs.BaseInfo, eventRewardList []obj.EventReward) EventRewardListResponse {
	baseResponse := NewBaseResponse(base)
	out := EventRewardListResponse{
		baseResponse,
		eventRewardList,
	}
	return out
}

func DefaultEventRewardList(base responseobjs.BaseInfo) EventRewardListResponse {
	// TODO: Maybe *don't* hardcode this mess? We'd have to recompile Outrun to change rewards.
	return EventRewardList(
		base,
		[]obj.EventReward{
			obj.NewEventReward(
				1,
				1,
				strconv.Itoa(int(enums.ItemIDInvincible)),
				1,
			),
			obj.NewEventReward(
				2,
				1000,
				strconv.Itoa(int(enums.ItemIDTrampoline)),
				1,
			),
			obj.NewEventReward(
				3,
				3000,
				strconv.Itoa(int(enums.ItemIDDrill)),
				1,
			),
			obj.NewEventReward(
				4,
				6000,
				strconv.Itoa(int(enums.ItemIDRedRing)),
				10,
			),
			obj.NewEventReward(
				5,
				9000,
				strconv.Itoa(int(enums.ItemIDBarrier)),
				1,
			),
			obj.NewEventReward(
				6,
				12000,
				strconv.Itoa(int(enums.ItemIDCombo)),
				1,
			),
			obj.NewEventReward(
				7,
				16000,
				strconv.Itoa(int(enums.ChaoIDCarbuncle)),
				1,
			),
			obj.NewEventReward(
				8,
				20000,
				strconv.Itoa(int(enums.ItemIDEnergy)),
				1,
			),
			obj.NewEventReward(
				9,
				24000,
				strconv.Itoa(int(enums.ItemIDAsteroid)),
				1,
			),
			obj.NewEventReward(
				10,
				28000,
				strconv.Itoa(int(enums.ItemIDMagnet)),
				1,
			),
			obj.NewEventReward(
				11,
				32000,
				strconv.Itoa(int(enums.ItemIDRedRing)),
				10,
			),
			obj.NewEventReward(
				12,
				36000,
				strconv.Itoa(int(enums.ItemIDLaser)),
				1,
			),
			obj.NewEventReward(
				13,
				40000,
				strconv.Itoa(int(enums.ItemIDRing)),
				1000,
			),
			obj.NewEventReward(
				14,
				44000,
				strconv.Itoa(int(enums.ItemIDInvincible)),
				2,
			),
			obj.NewEventReward(
				15,
				48000,
				strconv.Itoa(int(enums.ItemIDTrampoline)),
				2,
			),
			obj.NewEventReward(
				16,
				52000,
				strconv.Itoa(int(enums.ChaoIDCarbuncle)),
				1,
			),
			obj.NewEventReward(
				17,
				56000,
				strconv.Itoa(int(enums.ItemIDEnergy)),
				2,
			),
			obj.NewEventReward(
				18,
				60000,
				strconv.Itoa(int(enums.ItemIDDrill)),
				2,
			),
			obj.NewEventReward(
				19,
				64000,
				strconv.Itoa(int(enums.ItemIDBarrier)),
				2,
			),
			obj.NewEventReward(
				20,
				68000,
				strconv.Itoa(int(enums.ItemIDCombo)),
				2,
			),
			obj.NewEventReward(
				21,
				72000,
				strconv.Itoa(int(enums.ItemIDRedRing)),
				20,
			),
			obj.NewEventReward(
				22,
				76000,
				strconv.Itoa(int(enums.ItemIDAsteroid)),
				2,
			),
			obj.NewEventReward(
				23,
				80000,
				strconv.Itoa(int(enums.ItemIDRing)),
				2000,
			),
			obj.NewEventReward(
				24,
				84000,
				strconv.Itoa(int(enums.ItemIDMagnet)),
				2,
			),
			obj.NewEventReward(
				25,
				88000,
				strconv.Itoa(int(enums.ItemIDLaser)),
				2,
			),
			obj.NewEventReward(
				26,
				92000,
				strconv.Itoa(int(enums.ChaoIDCarbuncle)),
				1,
			),
			obj.NewEventReward(
				27,
				97000,
				strconv.Itoa(int(enums.ItemIDEnergy)),
				3,
			),
			obj.NewEventReward(
				28,
				100000,
				strconv.Itoa(int(enums.ItemIDInvincible)),
				3,
			),
			obj.NewEventReward(
				29,
				105000,
				strconv.Itoa(int(enums.ItemIDTrampoline)),
				3,
			),
			obj.NewEventReward(
				30,
				110000,
				strconv.Itoa(int(enums.ItemIDDrill)),
				3,
			),
			obj.NewEventReward(
				31,
				115000,
				strconv.Itoa(int(enums.ItemIDRedRing)),
				30,
			),
			obj.NewEventReward(
				32,
				120000,
				strconv.Itoa(int(enums.ItemIDBarrier)),
				3,
			),
			obj.NewEventReward(
				33,
				125000,
				strconv.Itoa(int(enums.ItemIDCombo)),
				3,
			),
			obj.NewEventReward(
				34,
				130000,
				strconv.Itoa(int(enums.ItemIDRing)),
				3000,
			),
			obj.NewEventReward(
				35,
				135000,
				strconv.Itoa(int(enums.ItemIDAsteroid)),
				3,
			),
			obj.NewEventReward(
				36,
				140000,
				strconv.Itoa(int(enums.ItemIDMagnet)),
				3,
			),
			obj.NewEventReward(
				37,
				150000,
				strconv.Itoa(int(enums.ChaoIDCarbuncle)),
				1,
			),
			obj.NewEventReward(
				38,
				160000,
				strconv.Itoa(int(enums.ItemIDEnergy)),
				3,
			),
			obj.NewEventReward(
				39,
				170000,
				strconv.Itoa(int(enums.ItemIDLaser)),
				3,
			),
			obj.NewEventReward(
				40,
				180000,
				strconv.Itoa(int(enums.ItemIDInvincible)),
				4,
			),
			obj.NewEventReward(
				41,
				190000,
				strconv.Itoa(int(enums.ItemIDTrampoline)),
				4,
			),
			obj.NewEventReward(
				42,
				200000,
				strconv.Itoa(int(enums.ItemIDRedRing)),
				30,
			),
			obj.NewEventReward(
				43,
				220000,
				strconv.Itoa(int(enums.ItemIDDrill)),
				4,
			),
			obj.NewEventReward(
				44,
				240000,
				strconv.Itoa(int(enums.ItemIDBarrier)),
				4,
			),
			obj.NewEventReward(
				45,
				260000,
				strconv.Itoa(int(enums.ItemIDRing)),
				4000,
			),
			obj.NewEventReward(
				46,
				280000,
				strconv.Itoa(int(enums.ItemIDCombo)),
				4,
			),
			obj.NewEventReward(
				47,
				300000,
				strconv.Itoa(int(enums.ChaoIDCarbuncle)),
				1,
			),
			obj.NewEventReward(
				48,
				340000,
				strconv.Itoa(int(enums.ItemIDEnergy)),
				3,
			),
			obj.NewEventReward(
				49,
				380000,
				strconv.Itoa(int(enums.ItemIDAsteroid)),
				4,
			),
			obj.NewEventReward(
				50,
				420000,
				strconv.Itoa(int(enums.ItemIDMagnet)),
				4,
			),
			obj.NewEventReward(
				51,
				460000,
				strconv.Itoa(int(enums.ItemIDRedRing)),
				40,
			),
			obj.NewEventReward(
				52,
				500000,
				strconv.Itoa(int(enums.ItemIDLaser)),
				4,
			),
			obj.NewEventReward(
				53,
				550000,
				strconv.Itoa(int(enums.ItemIDRing)),
				5000,
			),
			obj.NewEventReward(
				54,
				600000,
				strconv.Itoa(int(enums.ItemIDInvincible)),
				5,
			),
			obj.NewEventReward(
				55,
				650000,
				strconv.Itoa(int(enums.ItemIDTrampoline)),
				5,
			),
			obj.NewEventReward(
				56,
				700000,
				strconv.Itoa(int(enums.ItemIDRedRing)),
				50,
			),
			obj.NewEventReward(
				57,
				800000,
				strconv.Itoa(int(enums.IDRouletteTicketPremium)),
				5,
			),
			obj.NewEventReward(
				58,
				900000,
				strconv.Itoa(int(enums.IDRouletteTicketPremium)),
				10,
			),
			obj.NewEventReward(
				59,
				1000000,
				strconv.Itoa(int(enums.CharaTypeAmitieAmy)),
				10,
			),
		},
	)
}

type EventStateResponse struct {
	BaseResponse
	netobj.EventState `json:"eventState"`
}

func EventState(base responseobjs.BaseInfo, eventState netobj.EventState) EventStateResponse {
	baseResponse := NewBaseResponse(base)
	out := EventStateResponse{
		baseResponse,
		eventState,
	}
	return out
}

/*type EventUserRaidbossStateResponse struct {
	BaseResponse
	netobj.EventUserRaidbossState `json:"eventUserRaidboss"`
}

func EventUserRaidbossState(base responseobjs.BaseInfo, userRaidbossState netobj.EventUserRaidbossState) EventUserRaidbossStateResponse {
	baseResponse := NewBaseResponse(base)
	out := EventUserRaidbossStateResponse{
		baseResponse,
		userRaidbossState,
	}
	return out
}

type EventUserRaidbossListResponse struct {
	BaseResponse
	netobj.EventUserRaidbossState `json:"eventUserRaidboss"`
	EventRaidbossStates           []netobj.EventRaidbossState `json:"eventUserRaidbossList"`
}

func DefaultEventUserRaidbossList(base responseobjs.BaseInfo, userRaidbossState netobj.EventUserRaidbossState) EventUserRaidbossListResponse {
	baseResponse := NewBaseResponse(base)
	out := EventUserRaidbossListResponse{
		baseResponse,
		userRaidbossState,
		[]netobj.EventRaidbossState{
			netobj.DefaultRaidbossState(),
			netobj.DefaultRaidbossState2(),
			netobj.DefaultRaidbossState3(),
		},
	}
	return out
}

type EventActStartResponse struct {
	ActStartBaseResponse
	netobj.EventUserRaidbossState `json:"eventUserRaidboss"`
}

func EventActStart(base responseobjs.BaseInfo, playerState netobj.PlayerState, campaignList []obj.Campaign, eventUserRaidbossState netobj.EventUserRaidbossState) EventActStartResponse {
	actStartBase := ActStartBase(base, playerState, campaignList)
	return EventActStartResponse{
		actStartBase,
		eventUserRaidbossState,
	}
}

func DefaultEventActStart(base responseobjs.BaseInfo, player netobj.Player) EventActStartResponse {
	campaignList := []obj.Campaign{}
	playerState := player.PlayerState
	eventUserRaidbossState := player.EventUserRaidbossState
	return EventActStart(
		base,
		playerState,
		campaignList,
		eventUserRaidbossState,
	)
}

type EventPostGameResultsResponse struct {
	BaseResponse
	netobj.EventUserRaidbossState `json:"eventUserRaidboss"`
}

func EventPostGameResults(base responseobjs.BaseInfo, userRaidbossState netobj.EventUserRaidbossState) EventPostGameResultsResponse {
	baseResponse := NewBaseResponse(base)
	out := EventPostGameResultsResponse{
		baseResponse,
		userRaidbossState,
	}
	return out
}

type EventUpdateGameResultsResponse struct {
	BaseResponse
	PlayerState             netobj.PlayerState    `json:"playerState"`
	ChaoState               []netobj.Chao         `json:"chaoState"`
	DailyChallengeIncentive []obj.Incentive       `json:"dailyChallengeIncentive"` // should be obj.Item, but game doesn't care
	CharacterState          []netobj.Character    `json:"characterState"`
	MessageList             []obj.Message         `json:"messageList"`
	OperatorMessageList     []obj.OperatorMessage `json:"operatorMessageList"`
	TotalMessage            int64                 `json:"totalMessage"`
	TotalOperatorMessage    int64                 `json:"totalOperatorMessage"`
	PlayCharacterState      []netobj.Character    `json:"playCharacterState"`
	EventIncentiveList      []obj.Item            `json:"eventIncentiveList"`
	WheelOptions            netobj.WheelOptions   `json:"wheelOptions"`
	EventState              netobj.EventState     `json:"eventState,omitempty"`
}

func EventUpdateGameResults(base responseobjs.BaseInfo, player netobj.Player, dci []obj.Incentive, ml []obj.Message, oml []obj.OperatorMessage, pcs []netobj.Character, eil []obj.Item, wo netobj.WheelOptions, es netobj.EventState) EventUpdateGameResultsResponse {
	baseResponse := NewBaseResponse(base)
	playerState := player.PlayerState
	chaoState := player.ChaoState
	dailyChallengeIncentive := dci
	characterState := player.CharacterState
	messageList := []obj.Message{}
	operatorMessageList := []obj.OperatorMessage{}
	totalMessage := int64(len(messageList))
	totalOperatorMessage := int64(len(operatorMessageList))
	playCharacterState := pcs
	return EventUpdateGameResultsResponse{
		baseResponse,
		playerState,
		chaoState,
		dailyChallengeIncentive,
		characterState,
		messageList,
		operatorMessageList,
		totalMessage,
		totalOperatorMessage,
		playCharacterState,
		eil,
		wo,
		es,
	}
}

func DefaultEventUpdateGameResults(base responseobjs.BaseInfo, player netobj.Player, pcs []netobj.Character, es netobj.EventState) EventUpdateGameResultsResponse {
	baseResponse := NewBaseResponse(base)
	playerState := player.PlayerState
	chaoState := player.ChaoState
	dailyChallengeIncentive := []obj.Incentive{}
	characterState := player.CharacterState
	messageList := []obj.Message{}
	operatorMessageList := []obj.OperatorMessage{}
	totalMessage := int64(len(messageList))
	totalOperatorMessage := int64(len(operatorMessageList))
	eil := []obj.Item{}
	player.LastWheelOptions = logic.WheelRefreshLogic(player, player.LastWheelOptions)
	wo := player.LastWheelOptions
	return EventUpdateGameResultsResponse{
		baseResponse,
		playerState,
		chaoState,
		dailyChallengeIncentive,
		characterState,
		messageList,
		operatorMessageList,
		totalMessage,
		totalOperatorMessage,
		pcs,
		eil,
		wo,
		es,
	}
}*/
