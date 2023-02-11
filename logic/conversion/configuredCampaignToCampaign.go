package conversion

import (
	"time"

	"github.com/RunnersRevival/outrun/config/campaignconf"
	"github.com/RunnersRevival/outrun/obj"
	"github.com/jinzhu/now"
)

func ConfiguredCampaignToCampaign(cc campaignconf.ConfiguredCampaign) obj.Campaign {
	// Should be used by the game as soon as possible
	startTime := cc.StartTime
	switch startTime {
	case -2:
		startTime = now.BeginningOfDay().UTC().Unix()
	case -3:
		startTime = now.EndOfDay().UTC().Unix()
	case -4:
		startTime = time.Now().UTC().Unix() - 1
	}
	endTime := cc.EndTime
	switch endTime {
	case -2:
		endTime = now.BeginningOfDay().UTC().Unix()
	case -3:
		endTime = now.EndOfDay().UTC().Unix()
	case -4:
		endTime = time.Now().UTC().Add(24 * time.Hour).Unix()
	}
	newEvent := obj.Campaign{
		cc.RealType(),
		cc.Content,
		cc.SubContent,
		startTime,
		endTime,
	}
	return newEvent
}
