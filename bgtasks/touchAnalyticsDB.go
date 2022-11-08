package bgtasks

import (
	"log"

	"github.com/RunnersRevival/outrun/consts"
	"github.com/RunnersRevival/outrun/db/dbaccess"
)

func TouchAnalyticsDB() {
	err := dbaccess.SetAnalyticsEntry(consts.DBMySQLTableAnalytics, "touch", []byte{})
	if err != nil {
		log.Println("[ERR] Unable to touch " + consts.DBMySQLTableAnalytics + ": " + err.Error())
	}
}
