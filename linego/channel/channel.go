package channel

import (
	"../LineThrift"
	"../service"
	"../helper"
	"context"
	// "fmt"
)

var ChannelResult *LineThrift.ChannelToken
var err error

func LoginChannel(channelId string) {
	if service.IsLogin != true {
		panic("[Error]Not yet logged in.")
	}
	CS := service.ChannelService()
	ChannelResult, err = CS.ApproveChannelAndIssueChannelToken(context.TODO(), channelId)
	if err != nil {
		panic(err)
	}
	channelInfo, erro := GetChannelInfo(channelId)
	if erro != nil {
		panic(erro)
	}
	helper.Log(6969, "LOGS", "Succesfully logged in to: "+channelInfo.Name)
}

func GetChannelInfo(channelId string) (*LineThrift.ChannelInfo, error) {
	CS := service.ChannelService()
	res, err := CS.GetChannelInfo(context.TODO(), channelId, "EN")
	return res, err
}

func RevokeChannel(channelId string) (error) {
	CS := service.ChannelService()
	err := CS.RevokeChannel(context.TODO(), channelId)
	return err
}