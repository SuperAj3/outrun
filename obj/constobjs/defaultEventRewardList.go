package constobjs

import (
	"strconv"
	"github.com/RunnersRevival/outrun/enums"
	"github.com/RunnersRevival/outrun/obj"
)

var DefaultEventRewardList = func() []obj.EventReward {
	return []obj.EventReward{
		obj.NewEventReward(
			1,
			1,
			strconv.Itoa(int(enums.ItemIDLaser)),
			2,
		),
		obj.NewEventReward(
			2,
			1000,
			strconv.Itoa(int(enums.ItemIDBarrier)),
			2,
		),
		obj.NewEventReward(
			3,
			3000,
			strconv.Itoa(int(enums.ItemIDAsteroid)),
			2,
		),
		obj.NewEventReward(
			4,
			4000,
			strconv.Itoa(int(enums.ChaoIDGenesis)),
			1,
		),
		obj.NewEventReward(
			5,
			6000,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			15,
		),
		obj.NewEventReward(
			6,
			7500,
			strconv.Itoa(int(enums.ItemIDInvincible)),
			2,
		),
		obj.NewEventReward(
			7,
			9000,
			strconv.Itoa(int(enums.ItemIDRing)),
			15000,
		),
		obj.NewEventReward(
			8,
			10500,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			15,
		),
		obj.NewEventReward(
			9,
			12000,
			strconv.Itoa(int(enums.ChaoIDGenesis)),
			1,
		),
		obj.NewEventReward(
			10,
			16000,
			strconv.Itoa(int(enums.ItemIDMagnet)),
			2,
		),
		obj.NewEventReward(
			11,
			20000,
			strconv.Itoa(int(enums.ItemIDDrill)),
			2,
		),
		obj.NewEventReward(
			12,
			24000,
			strconv.Itoa(int(enums.ChaoIDGenesis)),
			1,
		),
		obj.NewEventReward(
			13,
			28000,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			30,
		),
		obj.NewEventReward(
			14,
			32000,
			strconv.Itoa(int(enums.ItemIDTrampoline)),
			2,
		),
		obj.NewEventReward(
			15,
			36000,
			strconv.Itoa(int(enums.ItemIDRing)),
			15000,
		),
		obj.NewEventReward(
			16,
			40000,
			strconv.Itoa(int(enums.ItemIDCombo)),
			2,
		),
		obj.NewEventReward(
			17,
			44000,
			strconv.Itoa(int(enums.ItemIDLaser)),
			4,
		),
		obj.NewEventReward(
			18,
			48000,
			strconv.Itoa(int(enums.ChaoIDGenesis)),
			1,
		),
		obj.NewEventReward(
			19,
			52000,
			strconv.Itoa(int(enums.ItemIDBarrier)),
			4,
		),
		obj.NewEventReward(
			20,
			56000,
			strconv.Itoa(int(enums.ItemIDAsteroid)),
			4,
		),
		obj.NewEventReward(
			21,
			60000,
			strconv.Itoa(int(enums.ItemIDInvincible)),
			4,
		),
		obj.NewEventReward(
			22,
			64000,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			30,
		),
		obj.NewEventReward(
			23,
			68000,
			strconv.Itoa(int(enums.ItemIDMagnet)),
			4,
		),
		obj.NewEventReward(
			24,
			72000,
			strconv.Itoa(int(enums.ChaoIDGenesis)),
			1,
		),
		obj.NewEventReward(
			25,
			76000,
			strconv.Itoa(int(enums.ItemIDDrill)),
			4,
		),
		obj.NewEventReward(
			26,
			80000,
			strconv.Itoa(int(enums.ItemIDRing)),
			30000,
		),
		obj.NewEventReward(
			27,
			84000,
			strconv.Itoa(int(enums.ItemIDTrampoline)),
			4,
		),
		obj.NewEventReward(
			28,
			88000,
			strconv.Itoa(int(enums.ItemIDCombo)),
			4,
		),
		obj.NewEventReward(
			29,
			92000,
			strconv.Itoa(int(enums.ChaoIDGenesis)),
			1,
		),
		obj.NewEventReward(
			30,
			97000,
			strconv.Itoa(int(enums.ItemIDLaser)),
			6,
		),
		obj.NewEventReward(
			31,
			100000,
			strconv.Itoa(int(enums.ItemIDBarrier)),
			6,
		),
		obj.NewEventReward(
			32,
			105000,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			45,
		),
		obj.NewEventReward(
			33,
			110000,
			strconv.Itoa(int(enums.ItemIDAsteroid)),
			6,
		),
		obj.NewEventReward(
			34,
			115000,
			strconv.Itoa(int(enums.ItemIDInvincible)),
			6,
		),
		obj.NewEventReward(
			35,
			120000,
			strconv.Itoa(int(enums.ChaoIDGenesis)),
			1,
		),
		obj.NewEventReward(
			36,
			125000,
			strconv.Itoa(int(enums.ItemIDMagnet)),
			6,
		),
		obj.NewEventReward(
			37,
			130000,
			strconv.Itoa(int(enums.ItemIDDrill)),
			6,
		),
		obj.NewEventReward(
			38,
			135000,
			strconv.Itoa(int(enums.ItemIDRing)),
			45000,
		),
		obj.NewEventReward(
			39,
			140000,
			strconv.Itoa(int(enums.ItemIDTrampoline)),
			6,
		),
		obj.NewEventReward(
			40,
			145000,
			strconv.Itoa(int(enums.ItemIDCombo)),
			6,
		),
		obj.NewEventReward(
			41,
			150000,
			strconv.Itoa(int(enums.ChaoIDGenesis)),
			1,
		),
		obj.NewEventReward(
			42,
			155000,
			strconv.Itoa(int(enums.ItemIDLaser)),
			8,
		),
		obj.NewEventReward(
			43,
			160000,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			60,
		),
		obj.NewEventReward(
			44,
			165000,
			strconv.Itoa(int(enums.ItemIDBarrier)),
			8,
		),
		obj.NewEventReward(
			45,
			170000,
			strconv.Itoa(int(enums.ItemIDAsteroid)),
			8,
		),
		obj.NewEventReward(
			46,
			175000,
			strconv.Itoa(int(enums.ItemIDInvincible)),
			8,
		),
		obj.NewEventReward(
			47,
			180000,
			strconv.Itoa(int(enums.ItemIDRing)),
			60000,
		),
		obj.NewEventReward(
			48,
			185000,
			strconv.Itoa(int(enums.ItemIDMagnet)),
			8,
		),
		obj.NewEventReward(
			49,
			190000,
			strconv.Itoa(int(enums.ItemIDDrill)),
			8,
		),
		obj.NewEventReward(
			50,
			195000,
			strconv.Itoa(int(enums.ChaoIDGenesis)),
			1,
		),
		obj.NewEventReward(
			51,
			200000,
			strconv.Itoa(int(enums.ItemIDTrampoline)),
			10,
		),
		obj.NewEventReward(
			52,
			210000,
			strconv.Itoa(int(enums.ItemIDCombo)),
			10,
		),
		obj.NewEventReward(
			53,
			220000,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			60,
		),
		obj.NewEventReward(
			54,
			230000,
			strconv.Itoa(int(enums.ItemIDLaser)),
			10,
		),
		obj.NewEventReward(
			55,
			240000,
			strconv.Itoa(int(enums.ChaoIDGenesis)),
			1,
		),
		obj.NewEventReward(
			56,
			250000,
			strconv.Itoa(int(enums.ItemIDRing)),
			60000,
		),
		obj.NewEventReward(
			57,
			260000,
			strconv.Itoa(int(enums.ItemIDBarrier)),
			10,
		),
		obj.NewEventReward(
			58,
			270000,
			strconv.Itoa(int(enums.ItemIDAsteroid)),
			10,
		),
		obj.NewEventReward(
			59,
			280000,
			strconv.Itoa(int(enums.ItemIDInvincible)),
			10,
		),
		obj.NewEventReward(
			60,
			290000,
			strconv.Itoa(int(enums.ChaoIDGenesis)),
			1,
		),
		obj.NewEventReward(
			61,
			300000,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			75,
		),
		obj.NewEventReward(
			62,
			310000,
			strconv.Itoa(int(enums.IDRouletteTicketItem)),
			10,
		),
		obj.NewEventReward(
			63,
			320000,
			strconv.Itoa(int(enums.ItemIDDrill)),
			10,
		),
		obj.NewEventReward(
			64,
			330000,
			strconv.Itoa(int(enums.ItemIDMagnet)),
			10,
		),
		obj.NewEventReward(
			65,
			340000,
			strconv.Itoa(int(enums.ItemIDRing)),
			75000,
		),
		obj.NewEventReward(
			66,
			350000,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			75,
		),
		obj.NewEventReward(
			67,
			360000,
			strconv.Itoa(int(enums.ItemIDBarrier)),
			10,
		),
		obj.NewEventReward(
			68,
			370000,
			strconv.Itoa(int(enums.ItemIDLaser)),
			10,
		),
		obj.NewEventReward(
			69,
			380000,
			strconv.Itoa(int(enums.ItemIDMagnet)),
			10,
		),
		obj.NewEventReward(
			70,
			390000,
			strconv.Itoa(int(enums.IDRouletteTicketItem)),
			10,
		),
		obj.NewEventReward(
			71,
			400000,
			strconv.Itoa(int(enums.ItemIDAsteroid)),
			10,
		),
		obj.NewEventReward(
			72,
			420000,
			strconv.Itoa(int(enums.ItemIDRing)),
			75000,
		),
		obj.NewEventReward(
			73,
			440000,
			strconv.Itoa(int(enums.ItemIDTrampoline)),
			10,
		),
		obj.NewEventReward(
			74,
			460000,
			strconv.Itoa(int(enums.ItemIDCombo)),
			10,
		),
		obj.NewEventReward(
			75,
			480000,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			75,
		),
		obj.NewEventReward(
			76,
			500000,
			strconv.Itoa(int(enums.IDRouletteTicketItem)),
			10,
		),
		obj.NewEventReward(
			77,
			520000,
			strconv.Itoa(int(enums.ItemIDDrill)),
			10,
		),
		obj.NewEventReward(
			78,
			540000,
			strconv.Itoa(int(enums.ItemIDInvincible)),
			10,
		),
		obj.NewEventReward(
			79,
			560000,
			strconv.Itoa(int(enums.IDRouletteTicketItem)),
			10,
		),
		obj.NewEventReward(
			80,
			580000,
			strconv.Itoa(int(enums.ItemIDRing)),
			75000,
		),
		obj.NewEventReward(
			81,
			600000,
			strconv.Itoa(int(enums.IDRouletteTicketPremium)),
			10,
		),
		obj.NewEventReward(
			82,
			620000,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			90,
		),
		obj.NewEventReward(
			83,
			640000,
			strconv.Itoa(int(enums.IDRouletteTicketItem)),
			10,
		),
		obj.NewEventReward(
			84,
			660000,
			strconv.Itoa(int(enums.ItemIDRing)),
			90000,
		),
		obj.NewEventReward(
			85,
			680000,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			90,
		),
		obj.NewEventReward(
			86,
			700000,
			strconv.Itoa(int(enums.ItemIDRing)),
			90000,
		),
		obj.NewEventReward(
			87,
			720000,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			90,
		),
		obj.NewEventReward(
			88,
			740000,
			strconv.Itoa(int(enums.IDRouletteTicketPremium)),
			15,
		),
		obj.NewEventReward(
			89,
			770000,
			strconv.Itoa(int(enums.IDRouletteTicketPremium)),
			15,
		),
		obj.NewEventReward(
			90,
			800000,
			strconv.Itoa(int(enums.CharaTypeClassicSonic)),
			10,
		),
	}
}

var DefaultXmasEventRewardList = func() []obj.EventReward {
	return []obj.EventReward{
		obj.NewEventReward(
			1,
			1,
			strconv.Itoa(int(enums.ItemIDLaser)),
			2,
		),
		obj.NewEventReward(
			2,
			1000,
			strconv.Itoa(int(enums.ItemIDBarrier)),
			2,
		),
		obj.NewEventReward(
			3,
			3000,
			strconv.Itoa(int(enums.ItemIDAsteroid)),
			2,
		),
		obj.NewEventReward(
			4,
			4000,
			strconv.Itoa(int(enums.ChaoIDChristmasYeti)),
			1,
		),
		obj.NewEventReward(
			5,
			6000,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			30,
		),
		obj.NewEventReward(
			6,
			7500,
			strconv.Itoa(int(enums.ItemIDInvincible)),
			2,
		),
		obj.NewEventReward(
			7,
			9000,
			strconv.Itoa(int(enums.ItemIDRing)),
			30000,
		),
		obj.NewEventReward(
			8,
			10500,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			30,
		),
		obj.NewEventReward(
			9,
			12000,
			strconv.Itoa(int(enums.ChaoIDChristmasYeti)),
			1,
		),
		obj.NewEventReward(
			10,
			16000,
			strconv.Itoa(int(enums.ItemIDMagnet)),
			2,
		),
		obj.NewEventReward(
			11,
			20000,
			strconv.Itoa(int(enums.ItemIDDrill)),
			2,
		),
		obj.NewEventReward(
			12,
			24000,
			strconv.Itoa(int(enums.ChaoIDChristmasYeti)),
			1,
		),
		obj.NewEventReward(
			13,
			28000,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			45,
		),
		obj.NewEventReward(
			14,
			32000,
			strconv.Itoa(int(enums.ItemIDTrampoline)),
			2,
		),
		obj.NewEventReward(
			15,
			36000,
			strconv.Itoa(int(enums.ItemIDRing)),
			45000,
		),
		obj.NewEventReward(
			16,
			40000,
			strconv.Itoa(int(enums.ItemIDCombo)),
			2,
		),
		obj.NewEventReward(
			17,
			44000,
			strconv.Itoa(int(enums.ItemIDLaser)),
			4,
		),
		obj.NewEventReward(
			18,
			48000,
			strconv.Itoa(int(enums.ChaoIDChristmasYeti)),
			1,
		),
		obj.NewEventReward(
			19,
			52000,
			strconv.Itoa(int(enums.ItemIDBarrier)),
			4,
		),
		obj.NewEventReward(
			20,
			56000,
			strconv.Itoa(int(enums.ItemIDAsteroid)),
			4,
		),
		obj.NewEventReward(
			21,
			60000,
			strconv.Itoa(int(enums.ItemIDInvincible)),
			4,
		),
		obj.NewEventReward(
			22,
			64000,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			60,
		),
		obj.NewEventReward(
			23,
			68000,
			strconv.Itoa(int(enums.ItemIDMagnet)),
			4,
		),
		obj.NewEventReward(
			24,
			72000,
			strconv.Itoa(int(enums.ChaoIDChristmasYeti)),
			1,
		),
		obj.NewEventReward(
			25,
			76000,
			strconv.Itoa(int(enums.ItemIDDrill)),
			4,
		),
		obj.NewEventReward(
			26,
			80000,
			strconv.Itoa(int(enums.ItemIDRing)),
			60000,
		),
		obj.NewEventReward(
			27,
			84000,
			strconv.Itoa(int(enums.ItemIDTrampoline)),
			4,
		),
		obj.NewEventReward(
			28,
			88000,
			strconv.Itoa(int(enums.ItemIDCombo)),
			4,
		),
		obj.NewEventReward(
			29,
			92000,
			strconv.Itoa(int(enums.ChaoIDChristmasYeti)),
			1,
		),
		obj.NewEventReward(
			30,
			97000,
			strconv.Itoa(int(enums.ItemIDLaser)),
			6,
		),
		obj.NewEventReward(
			31,
			100000,
			strconv.Itoa(int(enums.ItemIDBarrier)),
			6,
		),
		obj.NewEventReward(
			32,
			105000,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			75,
		),
		obj.NewEventReward(
			33,
			110000,
			strconv.Itoa(int(enums.ItemIDAsteroid)),
			6,
		),
		obj.NewEventReward(
			34,
			115000,
			strconv.Itoa(int(enums.ItemIDInvincible)),
			6,
		),
		obj.NewEventReward(
			35,
			120000,
			strconv.Itoa(int(enums.ChaoIDChristmasYeti)),
			1,
		),
		obj.NewEventReward(
			36,
			125000,
			strconv.Itoa(int(enums.ItemIDMagnet)),
			6,
		),
		obj.NewEventReward(
			37,
			130000,
			strconv.Itoa(int(enums.ItemIDDrill)),
			6,
		),
		obj.NewEventReward(
			38,
			135000,
			strconv.Itoa(int(enums.ItemIDRing)),
			75000,
		),
		obj.NewEventReward(
			39,
			140000,
			strconv.Itoa(int(enums.ItemIDTrampoline)),
			6,
		),
		obj.NewEventReward(
			40,
			145000,
			strconv.Itoa(int(enums.ItemIDCombo)),
			6,
		),
		obj.NewEventReward(
			41,
			150000,
			strconv.Itoa(int(enums.ChaoIDChristmasYeti)),
			1,
		),
		obj.NewEventReward(
			42,
			155000,
			strconv.Itoa(int(enums.ItemIDLaser)),
			8,
		),
		obj.NewEventReward(
			43,
			160000,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			90,
		),
		obj.NewEventReward(
			44,
			165000,
			strconv.Itoa(int(enums.ItemIDBarrier)),
			8,
		),
		obj.NewEventReward(
			45,
			170000,
			strconv.Itoa(int(enums.ItemIDAsteroid)),
			8,
		),
		obj.NewEventReward(
			46,
			175000,
			strconv.Itoa(int(enums.ItemIDInvincible)),
			8,
		),
		obj.NewEventReward(
			47,
			180000,
			strconv.Itoa(int(enums.ItemIDRing)),
			90000,
		),
		obj.NewEventReward(
			48,
			185000,
			strconv.Itoa(int(enums.ItemIDMagnet)),
			8,
		),
		obj.NewEventReward(
			49,
			190000,
			strconv.Itoa(int(enums.ItemIDDrill)),
			8,
		),
		obj.NewEventReward(
			50,
			195000,
			strconv.Itoa(int(enums.ChaoIDChristmasYeti)),
			1,
		),
		obj.NewEventReward(
			51,
			200000,
			strconv.Itoa(int(enums.ItemIDTrampoline)),
			10,
		),
		obj.NewEventReward(
			52,
			210000,
			strconv.Itoa(int(enums.ItemIDCombo)),
			10,
		),
		obj.NewEventReward(
			53,
			220000,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			150,
		),
		obj.NewEventReward(
			54,
			230000,
			strconv.Itoa(int(enums.ItemIDLaser)),
			10,
		),
		obj.NewEventReward(
			55,
			240000,
			strconv.Itoa(int(enums.ChaoIDChristmasYeti)),
			1,
		),
		obj.NewEventReward(
			56,
			250000,
			strconv.Itoa(int(enums.ItemIDRing)),
			150000,
		),
		obj.NewEventReward(
			57,
			260000,
			strconv.Itoa(int(enums.ItemIDBarrier)),
			10,
		),
		obj.NewEventReward(
			58,
			270000,
			strconv.Itoa(int(enums.ItemIDAsteroid)),
			10,
		),
		obj.NewEventReward(
			59,
			280000,
			strconv.Itoa(int(enums.ItemIDInvincible)),
			10,
		),
		obj.NewEventReward(
			60,
			290000,
			strconv.Itoa(int(enums.ChaoIDChristmasYeti)),
			1,
		),
		obj.NewEventReward(
			61,
			300000,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			200,
		),
		obj.NewEventReward(
			62,
			310000,
			strconv.Itoa(int(enums.IDRouletteTicketItem)),
			10,
		),
		obj.NewEventReward(
			63,
			320000,
			strconv.Itoa(int(enums.IDRouletteTicketItem)),
			10,
		),
		obj.NewEventReward(
			64,
			330000,
			strconv.Itoa(int(enums.IDRouletteTicketItem)),
			10,
		),
		obj.NewEventReward(
			65,
			340000,
			strconv.Itoa(int(enums.IDRouletteTicketItem)),
			20,
		),
		obj.NewEventReward(
			66,
			350000,
			strconv.Itoa(int(enums.IDRouletteTicketItem)),
			30,
		),
		obj.NewEventReward(
			67,
			360000,
			strconv.Itoa(int(enums.IDRouletteTicketPremium)),
			5,
		),
		obj.NewEventReward(
			68,
			370000,
			strconv.Itoa(int(enums.IDRouletteTicketPremium)),
			5,
		),
		obj.NewEventReward(
			69,
			380000,
			strconv.Itoa(int(enums.IDRouletteTicketPremium)),
			10,
		),
		obj.NewEventReward(
			70,
			400000,
			strconv.Itoa(int(enums.IDRouletteTicketPremium)),
			20,
		),
	}
}

var DefaultAnimalEventRewardList = func() []obj.EventReward {
	return []obj.EventReward{
		obj.NewEventReward(
			1,
			20,
			strconv.Itoa(int(enums.ChaoIDRCMonkey)),
			1,
		),
		obj.NewEventReward(
			2,
			40,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			15,
		),
		obj.NewEventReward(
			3,
			60,
			strconv.Itoa(int(enums.ItemIDRing)),
			15000,
		),
		obj.NewEventReward(
			4,
			80,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			15,
		),
		obj.NewEventReward(
			5,
			100,
			strconv.Itoa(int(enums.ItemIDRing)),
			15000,
		),
		obj.NewEventReward(
			6,
			120,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			15,
		),
		obj.NewEventReward(
			7,
			140,
			strconv.Itoa(int(enums.ItemIDRing)),
			15000,
		),
		obj.NewEventReward(
			8,
			160,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			15,
		),
		obj.NewEventReward(
			9,
			180,
			strconv.Itoa(int(enums.ItemIDRing)),
			15000,
		),
		obj.NewEventReward(
			10,
			200,
			strconv.Itoa(int(enums.ItemIDRing)),
			15000,
		),
		obj.NewEventReward(
			11,
			220,
			strconv.Itoa(int(enums.ChaoIDRCMonkey)),
			1,
		),
		obj.NewEventReward(
			12,
			240,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			20,
		),
		obj.NewEventReward(
			13,
			260,
			strconv.Itoa(int(enums.ItemIDRing)),
			20000,
		),
		obj.NewEventReward(
			14,
			280,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			20,
		),
		obj.NewEventReward(
			15,
			300,
			strconv.Itoa(int(enums.ItemIDRing)),
			20000,
		),
		obj.NewEventReward(
			16,
			320,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			20,
		),
		obj.NewEventReward(
			17,
			340,
			strconv.Itoa(int(enums.ItemIDRing)),
			20000,
		),
		obj.NewEventReward(
			18,
			360,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			20,
		),
		obj.NewEventReward(
			19,
			380,
			strconv.Itoa(int(enums.ItemIDRing)),
			20000,
		),
		obj.NewEventReward(
			20,
			400,
			strconv.Itoa(int(enums.ItemIDRing)),
			30000,
		),
		obj.NewEventReward(
			21,
			420,
			strconv.Itoa(int(enums.ChaoIDRCMonkey)),
			1,
		),
		obj.NewEventReward(
			22,
			440,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			25,
		),
		obj.NewEventReward(
			23,
			460,
			strconv.Itoa(int(enums.ItemIDRing)),
			35000,
		),
		obj.NewEventReward(
			24,
			480,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			25,
		),
		obj.NewEventReward(
			25,
			500,
			strconv.Itoa(int(enums.ItemIDRing)),
			35000,
		),
		obj.NewEventReward(
			26,
			520,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			25,
		),
		obj.NewEventReward(
			27,
			540,
			strconv.Itoa(int(enums.ItemIDRing)),
			35000,
		),
		obj.NewEventReward(
			28,
			560,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			25,
		),
		obj.NewEventReward(
			29,
			580,
			strconv.Itoa(int(enums.ItemIDRing)),
			35000,
		),
		obj.NewEventReward(
			30,
			600,
			strconv.Itoa(int(enums.ItemIDRing)),
			40000,
		),
		obj.NewEventReward(
			31,
			620,
			strconv.Itoa(int(enums.ChaoIDRCMonkey)),
			1,
		),
		obj.NewEventReward(
			32,
			640,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			25,
		),
		obj.NewEventReward(
			33,
			660,
			strconv.Itoa(int(enums.ItemIDRing)),
			40000,
		),
		obj.NewEventReward(
			34,
			680,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			25,
		),
		obj.NewEventReward(
			35,
			700,
			strconv.Itoa(int(enums.ItemIDRing)),
			40000,
		),
		obj.NewEventReward(
			36,
			720,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			25,
		),
		obj.NewEventReward(
			37,
			740,
			strconv.Itoa(int(enums.ItemIDRing)),
			40000,
		),
		obj.NewEventReward(
			38,
			760,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			25,
		),
		obj.NewEventReward(
			39,
			780,
			strconv.Itoa(int(enums.ItemIDRing)),
			40000,
		),
		obj.NewEventReward(
			40,
			800,
			strconv.Itoa(int(enums.ItemIDRing)),
			45000,
		),
		obj.NewEventReward(
			41,
			820,
			strconv.Itoa(int(enums.ChaoIDRCMonkey)),
			1,
		),
		obj.NewEventReward(
			42,
			840,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			30,
		),
		obj.NewEventReward(
			43,
			860,
			strconv.Itoa(int(enums.ItemIDRing)),
			50000,
		),
		obj.NewEventReward(
			44,
			880,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			30,
		),
		obj.NewEventReward(
			45,
			900,
			strconv.Itoa(int(enums.ItemIDRing)),
			50000,
		),
		obj.NewEventReward(
			46,
			920,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			30,
		),
		obj.NewEventReward(
			47,
			940,
			strconv.Itoa(int(enums.ItemIDRing)),
			50000,
		),
		obj.NewEventReward(
			48,
			960,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			30,
		),
		obj.NewEventReward(
			49,
			980,
			strconv.Itoa(int(enums.ItemIDRing)),
			50000,
		),
		obj.NewEventReward(
			50,
			1000,
			strconv.Itoa(int(enums.ItemIDRing)),
			50000,
		),
		obj.NewEventReward(
			51,
			1050,
			strconv.Itoa(int(enums.ChaoIDRCMonkey)),
			1,
		),
		obj.NewEventReward(
			52,
			1100,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			30,
		),
		obj.NewEventReward(
			53,
			1150,
			strconv.Itoa(int(enums.ItemIDRing)),
			50000,
		),
		obj.NewEventReward(
			54,
			1200,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			30,
		),
		obj.NewEventReward(
			55,
			1250,
			strconv.Itoa(int(enums.ItemIDRing)),
			50000,
		),
		obj.NewEventReward(
			56,
			1300,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			30,
		),
		obj.NewEventReward(
			57,
			1350,
			strconv.Itoa(int(enums.ItemIDRing)),
			50000,
		),
		obj.NewEventReward(
			58,
			1400,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			30,
		),
		obj.NewEventReward(
			59,
			1450,
			strconv.Itoa(int(enums.ItemIDRing)),
			50000,
		),
		obj.NewEventReward(
			60,
			5500,
			strconv.Itoa(int(enums.ItemIDRing)),
			50000,
		),
		obj.NewEventReward(
			61,
			1550,
			strconv.Itoa(int(enums.ChaoIDRCMonkey)),
			1,
		),
		obj.NewEventReward(
			62,
			1600,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			30,
		),
		obj.NewEventReward(
			63,
			1650,
			strconv.Itoa(int(enums.ItemIDRing)),
			50000,
		),
		obj.NewEventReward(
			64,
			1700,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			30,
		),
		obj.NewEventReward(
			65,
			1750,
			strconv.Itoa(int(enums.ItemIDRing)),
			50000,
		),
		obj.NewEventReward(
			66,
			1800,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			30,
		),
		obj.NewEventReward(
			67,
			1850,
			strconv.Itoa(int(enums.ItemIDRing)),
			50000,
		),
		obj.NewEventReward(
			68,
			1900,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			30,
		),
		obj.NewEventReward(
			69,
			1950,
			strconv.Itoa(int(enums.ItemIDRing)),
			50000,
		),
		obj.NewEventReward(
			70,
			2000,
			strconv.Itoa(int(enums.ItemIDRing)),
			50000,
		),
		obj.NewEventReward(
			71,
			2050,
			strconv.Itoa(int(enums.ChaoIDRCMonkey)),
			1,
		),
		obj.NewEventReward(
			72,
			2100,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			30,
		),
		obj.NewEventReward(
			73,
			2150,
			strconv.Itoa(int(enums.ItemIDRing)),
			50000,
		),
		obj.NewEventReward(
			74,
			2200,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			30,
		),
		obj.NewEventReward(
			75,
			2250,
			strconv.Itoa(int(enums.ItemIDRing)),
			50000,
		),
		obj.NewEventReward(
			76,
			2300,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			30,
		),
		obj.NewEventReward(
			77,
			2350,
			strconv.Itoa(int(enums.ItemIDRing)),
			50000,
		),
		obj.NewEventReward(
			78,
			2400,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			30,
		),
		obj.NewEventReward(
			79,
			2450,
			strconv.Itoa(int(enums.ItemIDRing)),
			50000,
		),
		obj.NewEventReward(
			80,
			7500,
			strconv.Itoa(int(enums.ItemIDRing)),
			50000,
		),
		obj.NewEventReward(
			81,
			2550,
			strconv.Itoa(int(enums.ChaoIDRCMonkey)),
			1,
		),
		obj.NewEventReward(
			82,
			2600,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			30,
		),
		obj.NewEventReward(
			83,
			2650,
			strconv.Itoa(int(enums.ItemIDRing)),
			50000,
		),
		obj.NewEventReward(
			84,
			2700,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			30,
		),
		obj.NewEventReward(
			85,
			2750,
			strconv.Itoa(int(enums.ItemIDRing)),
			50000,
		),
		obj.NewEventReward(
			86,
			2800,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			30,
		),
		obj.NewEventReward(
			87,
			2850,
			strconv.Itoa(int(enums.ItemIDRing)),
			50000,
		),
		obj.NewEventReward(
			88,
			2900,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			30,
		),
		obj.NewEventReward(
			89,
			2950,
			strconv.Itoa(int(enums.ItemIDRing)),
			50000,
		),
		obj.NewEventReward(
			90,
			3000,
			strconv.Itoa(int(enums.ItemIDRing)),
			50000,
		),
		obj.NewEventReward(
			91,
			3050,
			strconv.Itoa(int(enums.ChaoIDRCMonkey)),
			1,
		),
		obj.NewEventReward(
			92,
			3100,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			30,
		),
		obj.NewEventReward(
			93,
			3150,
			strconv.Itoa(int(enums.ItemIDRing)),
			60000,
		),
		obj.NewEventReward(
			94,
			3200,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			30,
		),
		obj.NewEventReward(
			95,
			3250,
			strconv.Itoa(int(enums.ItemIDRing)),
			60000,
		),
		obj.NewEventReward(
			96,
			3300,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			30,
		),
		obj.NewEventReward(
			97,
			3350,
			strconv.Itoa(int(enums.ItemIDRing)),
			60000,
		),
		obj.NewEventReward(
			98,
			3400,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			30,
		),
		obj.NewEventReward(
			99,
			3450,
			strconv.Itoa(int(enums.ItemIDRing)),
			60000,
		),
		obj.NewEventReward(
			100,
			9500,
			strconv.Itoa(int(enums.ItemIDRing)),
			60000,
		),
		obj.NewEventReward(
			101,
			3550,
			strconv.Itoa(int(enums.ChaoIDRCMonkey)),
			1,
		),
		obj.NewEventReward(
			102,
			3600,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			30,
		),
		obj.NewEventReward(
			103,
			3650,
			strconv.Itoa(int(enums.ItemIDRing)),
			60000,
		),
		obj.NewEventReward(
			104,
			3700,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			30,
		),
		obj.NewEventReward(
			105,
			3750,
			strconv.Itoa(int(enums.ItemIDRing)),
			60000,
		),
		obj.NewEventReward(
			106,
			3800,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			30,
		),
		obj.NewEventReward(
			107,
			3850,
			strconv.Itoa(int(enums.ItemIDRing)),
			60000,
		),
		obj.NewEventReward(
			108,
			3900,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			30,
		),
		obj.NewEventReward(
			109,
			3950,
			strconv.Itoa(int(enums.IDRouletteTicketPremium)),
			10,
		),
		obj.NewEventReward(
			110,
			4000,
			strconv.Itoa(int(enums.ItemIDRedRing)),
			55,
		),
	}
}



func GetPendingEventRewards(oldParam, newParam int64) ([]obj.EventReward, int64) {
	rewards := DefaultEventRewardList()
	pendingRewards := []obj.EventReward{}
	oldParamRewardId := int64(8008135)
	highestRewardIdSoFar := int64(-1)
	for _, reward := range rewards {
		if highestRewardIdSoFar < reward.RewardID {
			highestRewardIdSoFar = reward.RewardID
		}
		if oldParam <= reward.Param {
			// found it!
			oldParamRewardId = reward.RewardID - 1
			break
		}
	}
	if oldParamRewardId == 8008135 {
		// apparently all rewards were already obtained!
		return pendingRewards, highestRewardIdSoFar
	}
	newParamRewardId := int64(8008135)
	for _, reward := range rewards {
		if highestRewardIdSoFar < reward.RewardID {
			highestRewardIdSoFar = reward.RewardID
		}
		if newParam <= reward.Param {
			// found it!
			newParamRewardId = reward.RewardID - 1
			break
		}
	}
	if newParamRewardId == 8008135 {
		newParamRewardId = highestRewardIdSoFar
	}
	for _, reward := range rewards {
		if reward.RewardID > oldParamRewardId && reward.RewardID <= newParamRewardId {
			pendingRewards = append(pendingRewards, reward)
		}
	}
	return pendingRewards, newParamRewardId
}
