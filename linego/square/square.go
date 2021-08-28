package square

import (
	"../LineThrift"
	"../service"
	"context"
	"strings"
)

var SquareObsToken string
var Squares *LineThrift.GetJoinedSquaresResponse

func AcquireEncryptedAccessToken() (string, error) {
	TS := service.TalkService()
	res, err := TS.AcquireEncryptedAccessToken(context.TODO(), 2)
	token := strings.Split(res, "\x1e")[1]
	SquareObsToken = token
	Squares, _ = GetJoinedSquares()
	return token, err
}

func FetchMyEvents(syncToken string, limit int32, continuationToken string) (*LineThrift.FetchMyEventsResponse, error) {
	SS := service.SquareService()
	req := &LineThrift.FetchMyEventsRequest{
		SyncToken: syncToken,
		Limit: limit,
		ContinuationToken: continuationToken,
	}
	res, err := SS.FetchMyEvents(context.TODO(), req)
	return res, err
}

func GetJoinedSquares() (*LineThrift.GetJoinedSquaresResponse, error) {
	SS := service.SquareService()
	req := &LineThrift.GetJoinedSquaresRequest{
		ContinuationToken: "",
		Limit: 50,
	}
	res, err := SS.GetJoinedSquares(context.TODO(), req)
	return res, err
}

func GetJoinedSquareChats() (*LineThrift.GetJoinedSquareChatsResponse, error) {
	SS := service.SquareService()
	req := &LineThrift.GetJoinedSquareChatsRequest{
		ContinuationToken: "",
		Limit: 50,
	}
	res, err := SS.GetJoinedSquareChats(context.TODO(), req)
	return res, err
}

func GetSquare(squareMid string) (*LineThrift.GetSquareResponse, error) {
	SS := service.SquareService()
	req := &LineThrift.GetSquareRequest{
		Mid: squareMid,
	}
	res, err := SS.GetSquare(context.TODO(), req)
	return res, err
}

func GetSquareChat(squareMid string) (*LineThrift.GetSquareChatResponse, error) {
	SS := service.SquareService()
	req := &LineThrift.GetSquareChatRequest{
		SquareChatMid: squareMid,
	}
	res, err := SS.GetSquareChat(context.TODO(), req)
	return res, err
}

func SendSquareMessage(squareChatMid string, text string) (*LineThrift.SendMessageResponse, error) {
	SS := service.SquareService()
	M := &LineThrift.Message{
		To: squareChatMid,
		Text: text,
		ContentType: 0,
		ContentMetadata: nil,
		RelatedMessageId: "0", // to be honest, i don't know what this is for, and if i don't throw something it wouldn't send the message
	}
	sqM := &LineThrift.SquareMessage{
		Message: M,
	}
	req := &LineThrift.SendMessageRequest{
		SquareChatMid: squareChatMid,
		SquareMessage: sqM,
	}
	res, err := SS.SendMessage(context.TODO(), req)
	return res, err
}