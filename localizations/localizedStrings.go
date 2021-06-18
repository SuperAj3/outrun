package localizations

import (
	"github.com/Mtbcooler/outrun/enums"
)

var LocalizedStrings = map[string]map[string]string{
	"en": {
		"CheatWarningNotice":              "Your previous run was flagged by our automatic cheat detection system, and as such, did not count towards your high score, total score, or story progression.\nKeep in mind that we are cracking down on cheaters, so this incident has been reported. Repeated cheating may result in your account being suspended, reset, or deleted.\n\nThis notice is only shown once, so you can now close the game and reopen it to continue playing.",
		"DailyBattleWinRewardLabel":       "Daily Battle Reward.",
		"DailyBattleWinStreakRewardLabel": "Daily Battle %v-Win Streak Reward.",
		"DailyChallengeRewardLabel":       "A Daily Challenge Reward.",
		"DefaultAnnouncementMessage":      "Welcome to Sonic Runners Revival!",
		"DefaultLoginRouletteMessage":     "Earn some items to help you get a high score, and maybe even a top place in the rankings!",
		"DefaultMaintenanceMessage": "Sonic Runners Revival is currently in maintenance mode!\nPlease check our social media for more information!\n\n" +
			"Our website: https://www.sonicrunners.com/\n" +
			"Twitter: https://twitter.com/runnersrevival\n" +
			"Discord: https://discord.gg/T5ytR6T",
		"DefaultRewardLabel": "A gift from the Runners Revival Team.",
		"DeviceSuspensionNotice": "This device has been blocked from accessing the Sonic Runners Revival game server.\n\n" +
			"If you feel this is in error, please get in touch!\n" +
			"Twitter: https://twitter.com/runnersrevival\n" +
			"Discord: https://discord.gg/T5ytR6T",
		"FirstBillionPointRunRewardLabel": "A reward for performing your first billion point run.",
		"FirstLoginBonusRewardLabel":      "A Start Dash Login Bonus.",
		"InactiveAccountNotice": "The account on this device has been marked for deletion due to inactivity. If you would like to reclaim your account, please contact us as soon as possible, as your account will be fully deleted at the beginning of next month.\n" +
			"Twitter: https://twitter.com/runnersrevival\n" +
			"Discord: https://discord.gg/T5ytR6T",
		"LeagueHighRankingRewardLabel":  "A reward for getting the following position in the Runners' League High Score Ranking: %v.",
		"LeaguePromotionRewardLabel":    "Runners' League Promotion Reward. Story Mode.",
		"LeagueTotalRankingRewardLabel": "A reward for getting the following position in the Runners' League Total Score Ranking: %v.",
		"LoginBonusRewardLabel":         "A Login Bonus.",
		"NewAccountsDisabledNotice": "Registration of new accounts is disabled at the moment. For more information, please visit our Twitter or Discord.\n" +
			"Twitter: https://twitter.com/runnersrevival\n" +
			"Discord: https://discord.gg/T5ytR6T",
		"NewAccountsDisabledBetaNotice": "Registration is not possible with this server at this time. If you have accessed this server before and wish to get back to playing on this server, you will need to have a transfer ID and password on hand. If you've lost your transfer ID or your password, feel free to let us know in the beta channel on our Discord.\n" +
			"Twitter: https://twitter.com/runnersrevival\n" +
			"Discord: https://discord.gg/T5ytR6T",
		"QuickLeagueHighRankingRewardLabel":  "A reward for getting the following position in the Runners' League Timed Mode High Score Ranking: %v.",
		"QuickLeaguePromotionRewardLabel":    "Runners' League Promotion Reward. Timed Mode.",
		"QuickLeagueTotalRankingRewardLabel": "A reward for getting the following position in the Runners' League Timed Mode Total Score Ranking: %v.",
		"Rank999RewardLabel":                 "A reward for reaching Rank 999. Very few players are dedicated enough to reach this rank!",
		"SuspensionNotice_Temporary": "The account on this device has been temporarily suspended from Sonic Runners Revival for %s\n" +
			"You will be able to play the game again at: %s\n\n" +
			"If you feel this is in error, please get in touch!\n" +
			"Twitter: https://twitter.com/runnersrevival\n" +
			"Discord: https://discord.gg/T5ytR6T",
		"SuspensionNotice_Permanent": "The account on this device has been permanently banned from Sonic Runners Revival for %s\n\n" +
			"If you feel this is in error, please get in touch!\n" +
			"Twitter: https://twitter.com/runnersrevival\n" +
			"Discord: https://discord.gg/T5ytR6T",
		"SuspensionReason_0": "an unspecified reason. Please contact a Runners Revival Team member (preferably from the Yacker modmail bot on our Discord server) for more information about this suspension.",
		"SuspensionReason_1": "repeatedly cheating during runs.",
		"SuspensionReason_2": "packet manipulation or exploiting a server glitch.",
		"SuspensionReason_3": "attempting to mess up the economy through various methods.",
		"SuspensionReason_4": "leaking details about a prerelease build of Sonic Runners Revival.",
		"SuspensionReason_5": "setting an inappropriate nickname for yourself. When you are unsuspended, you will need to pick a new nickname for yourself.",
		"SuspensionReason_6": "repeatedly setting inappropriate nicknames for yourself.",
		"SuspensionReason_7": "intentionally attacking the Runners Revival service in one form or another.",
		"SuspensionReason_8": "an unknown reason (RID 8).",
		"SuspensionReason_9": "an unknown reason (RID 9).",
		"UpdateGameNotice": "Please update your game; this version is no longer supported on Sonic Runners Revival!\n\n" +
			"Our website: https://www.sonicrunners.com/\n" +
			"Twitter: https://twitter.com/runnersrevival\n" +
			"Discord: https://discord.gg/T5ytR6T",
		"VersionTooOldForAccountNotice": "The account on this device has signed in with a newer version of Sonic Runners Revival. To prevent any instabilities, we prevent older versions of Sonic Runners Revival from signing into our server once they've logged in with a newer version of Sonic Runners Revival.\n\n" +
			"Our website: https://www.sonicrunners.com/\n" +
			"Twitter: https://twitter.com/runnersrevival\n" +
			"Discord: https://discord.gg/T5ytR6T",
		"WatermarkTicker_1": "This server is powered by [ff0000]Outrun for Revival!",
		"WatermarkTicker_2": "User ID: [0000ff]%s",
		"WatermarkTicker_3": "High score (Timed Mode): [0000ff]%v",
		"WatermarkTicker_4": "High score (Story Mode): [0000ff]%v",
		"WatermarkTicker_5": "Total distance ran (Story Mode): [003fff]%v",
	},
}

var LanguageEnumToLanguageCodeTable = map[int64]string{
	enums.LangJapanese:   "ja",
	enums.LangEnglish:    "en",
	enums.LangChineseZH:  "zh",
	enums.LangChineseZHJ: "zhj",
	enums.LangKorean:     "ko",
	enums.LangFrench:     "fr",
	enums.LangGerman:     "de",
	enums.LangSpanish:    "es",
	enums.LangPortuguese: "pt",
	enums.LangItalian:    "it",
	enums.LangRussian:    "ru",
}

var LanguageEnumToLanguageNameTable = map[int64]string{
	enums.LangJapanese:   "Japanese",
	enums.LangEnglish:    "English",
	enums.LangChineseZH:  "Chinese",
	enums.LangChineseZHJ: "Chinese",
	enums.LangKorean:     "Korean",
	enums.LangFrench:     "French",
	enums.LangGerman:     "German",
	enums.LangSpanish:    "Spanish",
	enums.LangPortuguese: "Portuguese",
	enums.LangItalian:    "Italian",
	enums.LangRussian:    "Russian",
}

func GetStringByLanguage(language int64, key string, fallBackToEnglish bool) string {
	result := LocalizedStrings[LanguageEnumToLanguageCodeTable[language]][key]
	if result == "" {
		if fallBackToEnglish {
			result = LocalizedStrings["en"][key]
			if result == "" {
				result = "ERR: Key \"" + key + "\" does not exist or is empty!"
			}
		} else {
			result = "ERR: Key \"" + key + "\" does not exist in " + LanguageEnumToLanguageNameTable[language] + "!"
		}
	}
	return result
}
