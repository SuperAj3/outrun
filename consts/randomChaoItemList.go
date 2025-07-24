package consts

import "github.com/RunnersRevival/outrun/enums"

type PrizeInfo struct {
	AppearanceChance float64 // % chance for it to be chosen to be in wheel by the server
	Type             int64   // 0 for Chao, 1 for Character
}

// A 'load' as depicted below is the chance for the server to pick
// the associated item, where chosen is if randFloat(0, 100) < load.
// IMPORTANT: This load is exclusive to the rarity of the Chao that
// is being chosen by the server.

var RandomChaoWheelCharacterPrizes = map[string]float64{
	// characterID: load
	// Hopefully this should sum up to 100 just for
	// simplicity, but it shouldn't be a requirement.
	enums.CTStrSonic:    0.71, // Initial character - all groups
	enums.CTStrTails:    0.71, // Obtained in story mode - all groups
	enums.CTStrKnuckles: 0.71, // Obtained in story mode - all groups
	enums.CTStrAmy:      0.71, //Group 2
	enums.CTStrBig:      0.71, //Group 2
	//enums.CTStrBlaze:           0.81, //Group 3
	// enums.CTStrCharmy: 0.81, //Group 1
	enums.CTStrCream:  0.71, //Group 2
	// enums.CTStrEspio:  0.81, //Group 1
	enums.CTStrMephiles:        0.71, // Group 2
	// enums.CTStrOmega: 0.81, //Group 1
	enums.CTStrPSISilver:       0.71, // Group 2
	//enums.CTStrRouge:  0.81, //Group 1
	//enums.CTStrShadow: 0.81, //Group 1
	//enums.CTStrMarine:          0.81, // Group 3
	//enums.CTStrTangle: 0.81, // Group 3
	//enums.CTStrWhisper: 0.81, // Group 3
	//enums.CTStrSticks:          0.81, // Group 3
	//enums.CTStrTikal:           0.81, // Group 3
	//enums.CTStrVector: 0.81, //Group 1
	//enums.CTStrWerehog:         0.81, // Group 3
	enums.CTStrClassicSonic: 0.71, // Group 2
	// enums.CTStrMetalSonic:      0.81, // Group 1
	enums.CTStrSilver:      0.71, // Group 2

	//enums.CTStrAmitieAmy:       1.81, // Event (Puyo Puyo Quest)
	// enums.CTStrGothicAmy:       0.0, // Revival Event
	// enums.CTStrHalloweenShadow: 1.0, // Event (Halloween)
	// enums.CTStrHalloweenRouge:  1.0, // Event (Halloween)
	enums.CTStrHalloweenOmega:  1.81, // Event (Halloween)
	// enums.CTStrXMasSonic:       1.0, // Event (Christmas)
	// enums.CTStrXMasTails:       1.81, // Event (Christmas)
	// enums.CTStrXMasKnuckles:    1.0, // Event (Christmas)
	// enums.CTStrXT:              1.2, // Revival Event (Christmas)
}

var RandomChaoWheelChaoPrizes = map[string]float64{
	// TODO: Balance these
	enums.ChaoIDStrHeroChao:             3.0, // Event (Animal Rescue event 1.0)
	enums.ChaoIDStrGoldChao:             3.0, // Event (Animal Rescue event 1.0)
	enums.ChaoIDStrDarkChao:             3.7, // Event (Animal Rescue event 1.0)
	enums.ChaoIDStrJewelChao:            3.7, // Event (Animal Rescue event 1.0)
	//enums.ChaoIDStrNormalChao:           3.0, // Event (Animal Rescue event 1.0)
	//enums.ChaoIDStrOmochao:              3.0, // Event (Animal Rescue event 1.0)
	enums.ChaoIDStrRCMonkey:             3.7, // Event (Animal Rescue event 1.0)
	enums.ChaoIDStrRCSpring:             3.7,
	enums.ChaoIDStrRCElectromagnet:      3.7,
	enums.ChaoIDStrBabyCyanWisp:         3.7,
	enums.ChaoIDStrBabyIndigoWisp:       3.7,
	enums.ChaoIDStrBabyYellowWisp:       3.7,
	// enums.ChaoIDStrRCPinwheel:           3.7,
	enums.ChaoIDStrRCPiggyBank:          3.7,
	//enums.ChaoIDStrRCBalloon:            3.0,
	// enums.ChaoIDStrEasterChao:           3.7, // Event (Easter; Increase Odds During Event)
	enums.ChaoIDStrEasterBunny:          1.73, // Event (Easter; Increase Odds During Event)
	//enums.ChaoIDStrMerlina:              1.5, // Event (Easter: Premium Roulette for Timed Mode Event Only; Obtainable through Rewards List Only for Story Event)
	//enums.ChaoIDStrPurplePapurisu:       2.5, // Event (Puyo Puyo Quest)
	//enums.ChaoIDStrSuketoudara:          2.5, // Event (Puyo Puyo Quest)
	enums.ChaoIDStrCarbuncle:            3.38, // Event (Puyo Puyo Quest: Premium Roulette for Timed Mode Event Only; Obtainable through Rewards List Only for Story Event)
	enums.ChaoIDStrEggChao:           1.73,
	enums.ChaoIDStrPumpkinChao:       1.73,
	enums.ChaoIDStrSkullChao:         1.73,
	enums.ChaoIDStrYacker:            1.73,
	enums.ChaoIDStrRCGoldenPiggyBank: 1.73,
	enums.ChaoIDStrWizardChao:        1.73,
	// enums.ChaoIDStrRCTurtle:          0.82,
	//enums.ChaoIDStrRCUFO:             1.98,
	enums.ChaoIDStrRCBomber:          1.73,
	// enums.ChaoIDStrStarShapedMissile:    0.82, // Event (Zazz Raid Boss; Increase Odds During Event)
	//enums.ChaoIDStrRCSatellite:          1.98, // Event (Zazz Raid Boss; Increase Odds During Event)
	//enums.ChaoIDStrRCMoonMech:           0.0, // Event (Zazz Raid Boss; Only Obtainable through the Raid Boss Roulette, which is currently unavailable)
	//enums.ChaoIDStrRappy:                1.0, // Event (Phantasy Star Online 2)
	enums.ChaoIDStrKuna:                 3.38, // Event (Phantasy Star Online 2)
	//enums.ChaoIDStrMagLv1:               6.8, // Event (Phantasy Star Online 2)
	enums.ChaoIDStrBlowfishTransporter:  1.73, // Event (Tropical Coast; Increase Odds During Event)
	//enums.ChaoIDStrMotherWisp:           1.2, // Event (Tropical Coast: Premium Roulette for Timed Mode Event Only; Obtainable through Rewards List Only for Story Event)
	// enums.ChaoIDStrMarineChao:           0.82, // Event (Tropical Coast; Increase Odds During Event)
	//enums.ChaoIDStrGenesis:              1.5, // Event (Birthday: Premium Roulette for Timed Mode Event Only; Obtainable through Rewards List Only for Story Event)
	//enums.ChaoIDStrCartridge:            4.0, // Event (Birthday; Increase Odds During Event)
	//enums.ChaoIDStrDeathEgg:             3.38, // Event (Birthday; Increase Odds During Event)
	enums.ChaoIDStrRCFighter:            1.73,
	//enums.ChaoIDStrRCHovercraft:         0.82,
	//enums.ChaoIDStrRCHelicopter:         1.98,
	enums.ChaoIDStrGreenCrystalMonsterS: 1.73,
	//enums.ChaoIDStrGreenCrystalMonsterL: 0.82,
	//enums.ChaoIDStrRCAirship:            1.98,
	enums.ChaoIDStrMagicLamp:            1.73, // Event (Desert Ruins and Animal Rescue 2.0; Increase Odds During Event)
	//enums.ChaoIDStrDesertChao:           0.82, // Event (Desert Ruins; Increase Odds During Event)
	//enums.ChaoIDStrErazorDjinn:          2.5, // Event (Desert Ruins: Premium Roulette for Timed Mode Event Only; Obtainable through Rewards List Only for Story Event)
	enums.ChaoIDStrNightopian:           4.0, // Event (NiGHTS)
	//enums.ChaoIDStrNiGHTS:               2.0, // Event (NiGHTS)
	//enums.ChaoIDStrReala:                2.0, // Event (NiGHTS)
	//enums.ChaoIDStrSonicOmochao:         1.98, // Event (Team Sonic Omochao)
	enums.ChaoIDStrTailsOmochao:         1.73, // Event (Team Sonic Omochao)
	//enums.ChaoIDStrKnucklesOmochao:      0.82, // Event (Team Sonic Omochao)
	//enums.ChaoIDStrKingBoomBoo:          1.5, // Event (Halloween: Premium Roulette for Timed Mode Event Only; Obtainable through Rewards List Only for Story Event)
	//enums.ChaoIDStrBoo:                  1.98, // Event (Halloween; Increase Odds During Event)
	enums.ChaoIDStrHalloweenChao:        1.73, // Event (Halloween; Increase Odds During Event)
	//enums.ChaoIDStrHeavyBomb:            1.2, // Event (Fantasy Zone)
	//enums.ChaoIDStrOPapa:                1.5, // Event (Fantasy Zone)
	//enums.ChaoIDStrOpaOpa:               1.81, // Event (Fantasy Zone)
	//enums.ChaoIDStrBlockBomb:  0.82,
	//enums.ChaoIDStrHunkofMeat: 1.98,
	// enums.ChaoIDStrYeti:                 1.5, // Event (Christmas)
	// enums.ChaoIDStrSnowChao:             2.0, // Event (Christmas)
	// enums.ChaoIDStrChristmasYeti:        1.5, // Event (Christmas)
	// enums.ChaoIDStrChristmasNiGHTS:      1.5, // Event (Christmas NiGHTS)
	// enums.ChaoIDStrIdeya:                1.5, // Event (Christmas NiGHTS)
	// enums.ChaoIDStrChristmasNightopian:  1.5, // Event (Christmas NiGHTS)
	enums.ChaoIDStrOrbot:      1.73,
	// enums.ChaoIDStrCubot:      0.82,
	enums.ChaoIDStrLightChaos: 1.18,
	enums.ChaoIDStrHeroChaos:  1.18,
	enums.ChaoIDStrDarkChaos:  1.18,
	enums.ChaoIDStrChip:       1.18,
	//enums.ChaoIDStrShahra:               4.0, // Runners' League Story Mode
	enums.ChaoIDStrCaliburn:         1.18,
	enums.ChaoIDStrKingArthursGhost: 1.18,
	enums.ChaoIDStrRCTornado:        1.18,
	// enums.ChaoIDStrRCBattleCruiser:  1.47,
	//enums.ChaoIDStrRedCrystalMonsterS: 1.8,
	enums.ChaoIDStrRedCrystalMonsterL: 1.18,
	// enums.ChaoIDStrGoldenGoose:        1.47,
	enums.ChaoIDStrRCPirateSpaceship: 1.18,
	//enums.ChaoIDStrGoldenAngel:       1.8,
	// enums.ChaoIDStrRCTornado2:           1.47, // Event (Sonic Adventure; Increase Odds During Event)
	//enums.ChaoIDStrChaos:                1.8, // Event (Sonic Adventure; Increase Odds During Event)
	//enums.ChaoIDStrOrca:                 1.98, // Event (Sonic Adventure; Increase Odds During Event)
	//enums.ChaoIDStrChaoWalker:           0.0, // Daily Battle
	// enums.ChaoIDStrDarkQueen:            1.81, // Runners' League Timed Mode
	enums.ChaoIDStrRCBlockFace: 1.18,
	//enums.ChaoIDStrDFekt:                0.0, // Revival Event (assets TBD)
	//enums.ChaoIDStrDarkChaoWalker:       1.8, // Daily Battle?
	//enums.ChaoIDStrPrideChaoL:           3.09, // Revival Event (Pride Month Celebration)
	//enums.ChaoIDStrPrideChaoG:           3.09, // Revival Event (Pride Month Celebration)
	//enums.ChaoIDStrPrideChaoB:           3.09, // Revival Event (Pride Month Celebration)
	//enums.ChaoIDStrPrideChaoT:           3.09, // Revival Event (Pride Month Celebration)
	//enums.ChaoIDStrPrideChaoP:           3.09, // Revival Event (Pride Month Celebration)
	//enums.ChaoIDStrPrideChaoA:           3.09, // Revival Event (Pride Month Celebration)
	//enums.ChaoIDStrPrideChaoNB:          3.09, // Revival Event (Pride Month Celebration)
}
