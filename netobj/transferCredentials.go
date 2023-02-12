package netobj

type TransferCredentials struct {
	TransferID       string `json:"transferId"`
	TransferPassword string `json:"transferPassword"`
	PlayerID         string `json:"playerId"`
}

func DefaultTransferCredentials() TransferCredentials {
	return TransferCredentials{
		"",
		"",
		"0000000000",
	}
}

func PlayerToTransferCredentials(player Player) TransferCredentials {
	return TransferCredentials{
		player.MigrationPassword,
		player.UserPassword,
		player.ID,
	}
}

func VerifyTransferPasswordAndGetPlayerID(transferCreds TransferCredentials, userPassword string) (bool, string) {
	if transferCreds.TransferPassword == userPassword {
		return true, transferCreds.PlayerID
	} else {
		return false, ""
	}
}
