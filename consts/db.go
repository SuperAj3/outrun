package consts

const (
	DBFileName = "outrun.db"

	DBBucketSessionIDs = "sessionIDs"
	DBBucketPlayers    = "players"
	DBBucketAnalytics  = "analytics"

	DBSessionExpiryTime = 3600

	// TODO: Add more tables as needed.
	DBMySQLTableAnalytics          = "analytics"
	DBMySQLTableGameResults        = "game_results"
	DBMySQLTableCorePlayerInfo     = "player_info"
	DBMySQLTableUserRaidbossStates = "player_raidboss_states"
	DBMySQLTablePlayerStates       = "player_states"
	DBMySQLTableMileageMapStates   = "player_mileage"
	DBMySQLTableOptionUserResults  = "player_user_results"
	DBMySQLTableLastWheelOptions   = "player_roulette_options"
	DBMySQLTableRouletteInfos      = "player_item_roulette_data"
	DBMySQLTableLoginBonusStates   = "player_login_bonus_states"
	DBMySQLTablePersonalEvents     = "player_personal_events"
	DBMySQLTableMessages           = "player_messages"
	DBMySQLTableOperatorMessages   = "player_operator_messages"
	DBMySQLTableOperatorInfos      = "player_operator_infos"
	DBMySQLTableRaidBosses         = "raidbosses"
	DBMySQLTableQuickGameResults   = "quick_game_results"
	DBMySQLTableRankingLeagueData  = "ranking_league_data"
	DBMySQLTableSessionIDs         = "session_ids"

	SQLAnalyticsSchema = `
	CREATE TABLE IF NOT EXISTS ` + DBMySQLTableAnalytics + ` (
		pid VARCHAR(20) NOT NULL,
		param JSON,
		PRIMARY KEY (pid)
	) ENGINE = InnoDB;`
	SQLCorePlayerInfoSchema = `
	CREATE TABLE IF NOT EXISTS ` + DBMySQLTableCorePlayerInfo + ` (
		id BIGINT UNSIGNED NOT NULL,
		username VARCHAR(12) NOT NULL,
		password VARCHAR(10) NOT NULL,
		migrate_password VARCHAR(12) NOT NULL,
		user_password TEXT NOT NULL,
		player_key VARCHAR(32) NOT NULL,
		last_login BIGINT NOT NULL,
		language INTEGER NOT NULL,
		characters JSON,
		chao JSON,
		suspended_until BIGINT NOT NULL,
		suspend_reason INTEGER NOT NULL,
		last_login_device TEXT NOT NULL,
		last_login_platform INTEGER NOT NULL,
		last_login_versionid INTEGER NOT NULL,
		accepted_ope_message_ids JSON,
		PRIMARY KEY (id)
	) ENGINE = InnoDB;`
	SQLPlayerStatesSchema = `
	CREATE TABLE IF NOT EXISTS ` + DBMySQLTablePlayerStates + ` (
		id BIGINT UNSIGNED NOT NULL,
		items JSON NOT NULL,
		equipped_items JSON NOT NULL,
		mainchara_id TEXT NOT NULL,
		subchara_id TEXT NOT NULL,
		mainchao_id TEXT NOT NULL,
		subchao_id TEXT NOT NULL,
		num_rings INTEGER NOT NULL DEFAULT '0',
		num_buy_rings INTEGER NOT NULL DEFAULT '0',
		num_red_rings INTEGER NOT NULL DEFAULT '0',
		num_buy_red_rings INTEGER NOT NULL DEFAULT '0',
		energy INTEGER NOT NULL DEFAULT '0',
		energy_buy INTEGER NOT NULL DEFAULT '0',
		energy_renews_at BIGINT NOT NULL DEFAULT '0',
		num_messages INTEGER NOT NULL DEFAULT '0',
		ranking_league INTEGER NOT NULL DEFAULT '0',
		quick_ranking_league INTEGER NOT NULL DEFAULT '0',
		num_roulette_ticket INTEGER NOT NULL DEFAULT '0',
		num_chao_roulette_ticket INTEGER NOT NULL DEFAULT '0',
		chao_eggs INTEGER NOT NULL DEFAULT '0',
		high_score BIGINT NOT NULL DEFAULT '0',
		quick_high_score BIGINT NOT NULL DEFAULT '0',
		total_distance BIGINT NOT NULL,
		best_distance BIGINT NOT NULL,
		daily_mission_id INTEGER UNSIGNED NOT NULL DEFAULT '0',
		daily_mission_end_time BIGINT NOT NULL DEFAULT '0',
		daily_challenge_value INTEGER,
		daily_challenge_complete TINYINT UNSIGNED NOT NULL DEFAULT '0',
		num_daily_chal_cont INTEGER NOT NULL DEFAULT '0',
		num_plays INTEGER NOT NULL DEFAULT '0',
		num_animals INTEGER NOT NULL DEFAULT '0',
		rank INTEGER UNSIGNED NOT NULL DEFAULT '0',
		dm_cat INTEGER NOT NULL DEFAULT '0',
		dm_set INTEGER NOT NULL DEFAULT '0',
		dm_pos INTEGER NOT NULL DEFAULT '0',
		dm_nextcont INTEGER NOT NULL DEFAULT '0',
		league_high_score BIGINT NOT NULL DEFAULT '0',
		quick_league_high_score BIGINT NOT NULL DEFAULT '0',
		league_start_time BIGINT NOT NULL,
		league_reset_time BIGINT NOT NULL,
		ranking_league_group INTEGER NOT NULL,
		quick_ranking_league_group INTEGER NOT NULL,
		total_score BIGINT NOT NULL DEFAULT '0',
		quick_total_score BIGINT NOT NULL DEFAULT '0',
		best_total_score BIGINT NOT NULL DEFAULT '0',
		best_quick_total_score BIGINT NOT NULL DEFAULT '0',
		event_param BIGINT NOT NULL DEFAULT '0',
		PRIMARY KEY (id)
	) ENGINE = InnoDB;`
	SQLMileageMapStatesSchema = `
	CREATE TABLE IF NOT EXISTS ` + DBMySQLTableMileageMapStates + ` (
		id BIGINT UNSIGNED NOT NULL,
		map_distance INTEGER,
		num_boss_attack INTEGER,
		stage_distance INTEGER,
		stage_max_score INTEGER,
		episode INTEGER,
		chapter INTEGER,
		point INTEGER,
		stage_total_score BIGINT,
		chapter_start_time BIGINT NOT NULL,
		PRIMARY KEY (id)
	) ENGINE = InnoDB;`
	SQLOptionUserResultsSchema = `
	CREATE TABLE IF NOT EXISTS ` + DBMySQLTableOptionUserResults + ` (
		id BIGINT UNSIGNED NOT NULL,
		high_total_score INTEGER,
		high_quick_total_score INTEGER,
		total_rings INTEGER,
		total_red_rings INTEGER,
		chao_roulette_spin_count INTEGER,
		roulette_spin_count INTEGER,
		num_jackpots INTEGER,
		best_jackpot INTEGER,
		support INTEGER,
		PRIMARY KEY (id)
	) ENGINE = InnoDB;`
	SQLRouletteInfosSchema = `
	CREATE TABLE IF NOT EXISTS ` + DBMySQLTableRouletteInfos + ` (
		id BIGINT UNSIGNED NOT NULL,
		login_roulette_id INTEGER,
		roulette_period_end BIGINT NOT NULL,
		roulette_count_in_period INTEGER,
		got_jackpot_this_period INTEGER,
		PRIMARY KEY (id)
	) ENGINE = InnoDB;`
	SQLLoginBonusStatesSchema = `
	CREATE TABLE IF NOT EXISTS ` + DBMySQLTableLoginBonusStates + ` (
		id BIGINT UNSIGNED NOT NULL,
		current_start_dash_bonus_day INTEGER,
		current_login_bonus_day INTEGER,
		last_login_bonus_time BIGINT NOT NULL,
		next_login_bonus_time BIGINT NOT NULL,
		login_bonus_start_time BIGINT NOT NULL,
		login_bonus_end_time BIGINT NOT NULL,
		PRIMARY KEY (id)
	) ENGINE = InnoDB;`
	SQLOperatorMessagesSchema = `
	CREATE TABLE IF NOT EXISTS ` + DBMySQLTableOperatorMessages + ` (
		id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
		for_all TINYINT UNSIGNED NOT NULL DEFAULT '0',
		userid BIGINT UNSIGNED NOT NULL,
		contents TEXT,
		item JSON,
		expire_time BIGINT NOT NULL,
		PRIMARY KEY (id)
	) ENGINE = InnoDB;`
	SQLRankingLeagueDataSchema = `
	CREATE TABLE IF NOT EXISTS ` + DBMySQLTableRankingLeagueData + ` (
		league_id INT UNSIGNED NOT NULL,
		group_id INT UNSIGNED NOT NULL,
		start_time BIGINT NOT NULL,
		reset_time BIGINT NOT NULL,
		league_player_count INTEGER,
		group_player_count INTEGER,
		PRIMARY KEY (league_id, group_id)
	) ENGINE = InnoDB;`
	SQLSessionIDsSchema = `
	CREATE TABLE IF NOT EXISTS ` + DBMySQLTableSessionIDs + ` (
		sid VARCHAR(48) NOT NULL,
		uid BIGINT UNSIGNED NOT NULL,
		assigned_at_time BIGINT NOT NULL,
		PRIMARY KEY (sid)
	) ENGINE = InnoDB;`
	SQLOperatorInfosSchema = `
	CREATE TABLE IF NOT EXISTS ` + DBMySQLTableOperatorInfos + ` (
		uid BIGINT UNSIGNED NOT NULL,
		id INTEGER NOT NULL,
		param TEXT,
		PRIMARY KEY (uid, id)
	) ENGINE = InnoDB;`
	SQLUserRaidbossStatesSchema = `
	CREATE TABLE IF NOT EXISTS ` + DBMySQLTableUserRaidbossStates + ` (
		uid BIGINT UNSIGNED NOT NULL,
		raidboss_rings INTEGER NOT NULL,
		raid_energy INTEGER NOT NULL,
		raid_energy_max INTEGER NOT NULL DEFAULT '10',
		num_beated_encounter INTEGER NOT NULL,
		num_beated_enterprise INTEGER NOT NULL,
		energy_renews_at BIGINT NOT NULL,
		score_until_next_raidboss BIGINT NOT NULL,
		PRIMARY KEY (uid)
	) ENGINE = InnoDB;`
	SQLGameResultsSchema = `
	CREATE TABLE IF NOT EXISTS ` + DBMySQLTableGameResults + ` (
		gameid BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
		uid BIGINT UNSIGNED NOT NULL,
		score BIGINT NOT NULL,
		rings INTEGER NOT NULL,
		failure_rings INTEGER NOT NULL,
		red_rings INTEGER NOT NULL,
		distance BIGINT NOT NULL,
		daily_challenge_value INTEGER NOT NULL,
		daily_challenge_complete TINYINT NOT NULL,
		animals INTEGER NOT NULL,
		max_combo INTEGER NOT NULL,
		closed TINYINT NOT NULL,
		boss_destroyed TINYINT NOT NULL,
		chapter_clear TINYINT NOT NULL,
		get_chao_egg TINYINT NOT NULL,
		boss_hits INTEGER NOT NULL,
		reach_point INTEGER NOT NULL,
		event_id BIGINT,
		event_value INTEGER,
		cheat_result VARCHAR(8) NOT NULL,
		PRIMARY KEY (gameid)
	) ENGINE = InnoDB;`
	SQLQuickGameResultsSchema = `
	CREATE TABLE IF NOT EXISTS ` + DBMySQLTableQuickGameResults + ` (
		gameid BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
		uid BIGINT UNSIGNED NOT NULL,
		score BIGINT NOT NULL,
		rings INTEGER NOT NULL,
		failure_rings INTEGER NOT NULL,
		red_rings INTEGER NOT NULL,
		distance BIGINT NOT NULL,
		daily_challenge_value INTEGER NOT NULL,
		daily_challenge_complete TINYINT NOT NULL,
		animals INTEGER NOT NULL,
		max_combo INTEGER NOT NULL,
		closed TINYINT NOT NULL,
		cheat_result VARCHAR(8) NOT NULL,
		PRIMARY KEY (gameid)
	) ENGINE = InnoDB;`
	SQLLastWheelOptionsSchema = `
	CREATE TABLE IF NOT EXISTS ` + DBMySQLTableLastWheelOptions + ` (
		id BIGINT UNSIGNED NOT NULL,
		items JSON,
		item_count JSON,
		item_weight JSON,
		item_won INTEGER,
		next_free_spin BIGINT NOT NULL,
		spin_id BIGINT,
		spin_cost INTEGER,
		roulette_rank INTEGER,
		num_roulette_token INTEGER,
		num_jackpot_ring INTEGER,
		num_remaining_roulette INTEGER,
		item_list JSON,
		PRIMARY KEY (id)
	) ENGINE = InnoDB;`
	SQLPlayerStatesInsertTypeSchema = `(
		id,
		items,
		equipped_items,
		mainchara_id,
		subchara_id,
		mainchao_id,
		subchao_id,
		num_rings,
		num_buy_rings,
		num_red_rings,
		num_buy_red_rings,
		energy,
		energy_buy,
		energy_renews_at,
		num_messages,
		ranking_league,
		quick_ranking_league,
		num_roulette_ticket,
		num_chao_roulette_ticket,
		chao_eggs,
		high_score,
		quick_high_score,
		total_distance,
		best_distance,
		daily_mission_id,
		daily_mission_end_time,
		daily_challenge_value,
		daily_challenge_complete,
		num_daily_chal_cont,
		num_plays,
		num_animals,
		rank,
		dm_cat,
		dm_set,
		dm_pos,
		dm_nextcont,
		league_high_score,
		quick_league_high_score,
		league_start_time,
		league_reset_time,
		ranking_league_group,
		quick_ranking_league_group,
		total_score,
		quick_total_score,
		best_total_score,
		best_quick_total_score,
		event_param
	)
	VALUES (
		:id,
		:items,
		:equipped_items,
		:mainchara_id,
		:subchara_id,
		:mainchao_id,
		:subchao_id,
		:num_rings,
		:num_buy_rings,
		:num_red_rings,
		:num_buy_red_rings,
		:energy,
		:energy_buy,
		:energy_renews_at,
		:num_messages,
		:ranking_league,
		:quick_ranking_league,
		:num_roulette_ticket,
		:num_chao_roulette_ticket,
		:chao_eggs,
		:high_score,
		:quick_high_score,
		:total_distance,
		:best_distance,
		:daily_mission_id,
		:daily_mission_end_time,
		:daily_challenge_value,
		:daily_challenge_complete,
		:num_daily_chal_cont,
		:num_plays,
		:num_animals,
		:rank,
		:dm_cat,
		:dm_set,
		:dm_pos,
		:dm_nextcont,
		:league_high_score,
		:quick_league_high_score,
		:league_start_time,
		:league_reset_time,
		:ranking_league_group,
		:quick_ranking_league_group,
		:total_score,
		:quick_total_score,
		:best_total_score,
		:best_quick_total_score,
		:event_param
	)`
)
