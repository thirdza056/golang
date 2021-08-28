package call

import (
	"../LineThrift"
	"../service"
	"context"
	// "fmt"
)

var err error

func AcquireCallRoute(to string) (r []string, err error) {
	TS := service.TalkService()
	res, err := TS.AcquireCallRoute(context.TODO(), to)
	return res, err
}

func AcquireGroupCallRoute(chatMid string, mediaType LineThrift.GroupCallMediaType) (r *LineThrift.GroupCallRoute, err error) {
	CS := service.CallService()
	res, err := CS.AcquireGroupCallRoute(context.TODO(), chatMid, mediaType)
	return res, err
}

func GetGroupCall(chatMid string) (r *LineThrift.GroupCall, err error) {
	CS := service.CallService()
	res, err := CS.GetGroupCall(context.TODO(), chatMid)
	return res, err
}

func InviteIntoGroupCall(chatMid string, memberMids []string, mediaType LineThrift.GroupCallMediaType) (err error) {
	CS := service.CallService()
	err = CS.InviteIntoGroupCall(context.TODO(), chatMid, memberMids, mediaType)
	return err
}