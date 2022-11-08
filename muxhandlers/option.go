package muxhandlers

import (
	"github.com/RunnersRevival/outrun/consts"
	"github.com/RunnersRevival/outrun/db"
	"github.com/RunnersRevival/outrun/db/dbaccess"
	"github.com/RunnersRevival/outrun/emess"
	"github.com/RunnersRevival/outrun/helper"
	"github.com/RunnersRevival/outrun/responses"
	"github.com/RunnersRevival/outrun/status"
)

func GetOptionUserResult(helper *helper.Helper) {
	sid, _ := helper.GetSessionID()
	if !helper.CheckSession(true) {
		return
	}
	pid, err := helper.GetCallingPlayerID()
	if err != nil {
		helper.InternalErr("Error getting calling player ID", err)
		return
	}
	optionUserResult, err := dbaccess.GetOptionUserResult(consts.DBMySQLTableOptionUserResults, pid)
	if err != nil {
		helper.InternalErr("Error getting OptionUserResult data", err)
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.OptionUserResult(baseInfo, optionUserResult)
	response.Seq, _ = db.BoltGetSessionIDSeq(sid)
	helper.SendResponse(response)
}
