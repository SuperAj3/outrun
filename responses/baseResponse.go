package responses

import (
	"github.com/RunnersRevival/outrun/consts"
	"github.com/RunnersRevival/outrun/meta"
	"github.com/RunnersRevival/outrun/responses/responseobjs"
)

type BaseResponse struct {
	responseobjs.BaseInfo
	AssetsVersion     string `json:"assets_version"` // doesn't necessarily have to be a number
	ClientDataVersion string `json:"client_data_version"`
	DataVersion       string `json:"data_version"`
	InfoVersion       string `json:"info_version"`
	Version           string `json:"version"`
	OutrunVersion     string `json:"ORN_version"`
}

func NewBaseResponse(base responseobjs.BaseInfo) BaseResponse {
	return BaseResponse{
		base,
		"054",
		"2.2.2",
		"15",
		"017",
		"2.2.0",
		meta.Version,
	}
}

func NewBaseResponseV(base responseobjs.BaseInfo, gameVersion string) BaseResponse {
	return BaseResponse{
		base,
		consts.DataVersionForGameVersion[gameVersion],
		gameVersion,
		"15",
		"017",
		gameVersion,
		meta.Version,
	}
}

type NextVersionResponse struct {
	responseobjs.BaseInfo
	NumRedRingsIOS        int64  `json:"numRedRingsIOS,string"` // UNCONFIRMED!
	NumRedRingsANDROID    int64  `json:"numRedRingsANDROID,string"`
	NumBuyRedRingsIOS     int64  `json:"numBuyRedRingsIOS,string"` // UNCONFIRMED!
	NumBuyRedRingsANDROID int64  `json:"numBuyRedRingsANDROID,string"`
	Username              string `json:"userName"`
	CloseMessageJP        string `json:"closeMessageJP"`
	CloseMessageEN        string `json:"closeMessageEN"`
	CloseURL              string `json:"closeUrl"`
}

func NewNextVersionResponse(base responseobjs.BaseInfo, numRedRings, numBuyRedRings int64, username, japaneseMessage, englishMessage, url string) NextVersionResponse {
	return NextVersionResponse{
		base,
		numRedRings,
		numRedRings,
		numBuyRedRings,
		numBuyRedRings,
		username,
		japaneseMessage,
		englishMessage,
		url,
	}
}
