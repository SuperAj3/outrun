package netobj

import (
	"encoding/json"
	"log"
)

type PlayerInfo struct {
	Username              string      `json:"username" db:"username"`
	Password              string      `json:"password" db:"password"`
	MigrationPassword     string      `json:"migrationPassword" db:"migrate_password"` // used in migration
	UserPassword          string      `json:"userPassword" db:"user_password"`         // used in migration
	Key                   string      `json:"key" db:"player_key"`
	LastLogin             int64       `json:"lastLogin" db:"last_login"`
	Language              int64       `json:"language" db:"language"`
	CharacterState        []Character `json:"characters" db:"characters"`
	ChaoState             []Chao      `json:"chao" db:"chao"`
	SuspendedUntil        int64       `json:"suspendedUntil" db:"suspended_until"`
	SuspendReason         int64       `json:"suspendReason" db:"suspend_reason"`
	LastLoginDevice       string      `json:"lastLoginDevice" db:"last_login_device"`
	LastLoginPlatform     int64       `json:"lastLoginPlatform" db:"last_login_platform"`
	LastLoginVersionId    int64       `json:"lastLoginVersionId" db:"last_login_versionid"`
	AcceptedOpeMessageIds []int64     `json:"acceptedOpeMessageIds" db:"accepted_ope_message_ids"`
}

type StoredPlayerInfo struct {
	Username              string `json:"username" db:"username"`
	Password              string `json:"password" db:"password"`
	MigrationPassword     string `json:"migrationPassword" db:"migrate_password"` // used in migration
	UserPassword          string `json:"userPassword" db:"user_password"`         // used in migration
	Key                   string `json:"key" db:"player_key"`
	LastLogin             int64  `json:"lastLogin" db:"last_login"`
	Language              int64  `json:"language" db:"language"`
	CharacterState        []byte `json:"characters" db:"characters"`
	ChaoState             []byte `json:"chao" db:"chao"`
	SuspendedUntil        int64  `json:"suspendedUntil" db:"suspended_until"`
	SuspendReason         int64  `json:"suspendReason" db:"suspend_reason"`
	LastLoginDevice       string `json:"lastLoginDevice" db:"last_login_device"`
	LastLoginPlatform     int64  `json:"lastLoginPlatform" db:"last_login_platform"`
	LastLoginVersionId    int64  `json:"lastLoginVersionId" db:"last_login_versionid"`
	AcceptedOpeMessageIds []byte `json:"acceptedOpeMessageIds" db:"accepted_ope_message_ids"`
}

func PlayerInfoToStoredPlayerInfo(pli PlayerInfo) StoredPlayerInfo {
	characterstate, err := json.Marshal(pli.CharacterState)
	if err != nil {
		log.Printf("[WARN] Couldn't marshal character state: %s\n", err)
	}
	chaostate, err := json.Marshal(pli.ChaoState)
	if err != nil {
		log.Printf("[WARN] Couldn't marshal chao state: %s\n", err)
	}
	acceptedopemsgids, err := json.Marshal(pli.AcceptedOpeMessageIds)
	if err != nil {
		log.Printf("[WARN] Couldn't marshal accepted operator message id list: %s\n", err)
	}
	return StoredPlayerInfo{
		pli.Username,
		pli.Password,
		pli.MigrationPassword,
		pli.UserPassword,
		pli.Key,
		pli.LastLogin,
		pli.Language,
		characterstate,
		chaostate,
		pli.SuspendedUntil,
		pli.SuspendReason,
		pli.LastLoginDevice,
		pli.LastLoginPlatform,
		pli.LastLoginVersionId,
		acceptedopemsgids,
	}
}

func StoredPlayerInfoToPlayerInfo(pli StoredPlayerInfo) PlayerInfo {
	var characterstate []Character
	var chaostate []Chao
	var acceptedopemsgids []int64
	err := json.Unmarshal(pli.CharacterState, &characterstate)
	if err != nil {
		log.Printf("[WARN] Couldn't unmarshal character state: %s\n", err)
	}
	err = json.Unmarshal(pli.ChaoState, &chaostate)
	if err != nil {
		log.Printf("[WARN] Couldn't unmarshal chao state: %s\n", err)
	}
	err = json.Unmarshal(pli.AcceptedOpeMessageIds, &acceptedopemsgids)
	if err != nil {
		log.Printf("[WARN] Couldn't unmarshal accepted operator message id list: %s\n", err)
	}
	return PlayerInfo{
		pli.Username,
		pli.Password,
		pli.MigrationPassword,
		pli.UserPassword,
		pli.Key,
		pli.LastLogin,
		pli.Language,
		characterstate,
		chaostate,
		pli.SuspendedUntil,
		pli.SuspendReason,
		pli.LastLoginDevice,
		pli.LastLoginPlatform,
		pli.LastLoginVersionId,
		acceptedopemsgids,
	}
}
