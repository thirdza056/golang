package talk

import (
	"../LineThrift"
	"../service"
	"context"
	"strings"
	"strconv"
	"fmt"
    "encoding/json"
	"unicode/utf8"
)

var err error
// im initiating a new TalkService every function call
// because the bytes.Buffer is not thread-safe,
// and i don't know how to do that yet.

/* User Functions */

func GetProfile() (r *LineThrift.Profile, err error) {
	TS := service.TalkService()
	res, err := TS.GetProfile(context.TODO())
	return res, err
}

func GetSettings() (r *LineThrift.Settings, err error) {
	TS := service.TalkService()
	res, err := TS.GetSettings(context.TODO())
	return res, err
}

func GetUserTicket() (r *LineThrift.Ticket, err error) {
	TS := service.TalkService()
	res, err := TS.GetUserTicket(context.TODO())
	return res, err
}

func UpdateProfile(profile *LineThrift.Profile) (err error) {
	TS := service.TalkService()
	err = TS.UpdateProfile(context.TODO(), int32(0), profile)
	return err
}

func UpdateSettings(settings *LineThrift.Settings) (err error) {
	TS := service.TalkService()
	err = TS.UpdateSettings(context.TODO(), int32(0), settings)
	return err
}

func UpdateProfileAttribute(attr LineThrift.ProfileAttribute, value string) (err error) {
	TS := service.TalkService()
	err = TS.UpdateProfileAttribute(context.TODO(), int32(0), attr, value)
	return err
}

/* Fetch Functions */

func FetchOperations(rev int64, count int32) (r []*LineThrift.Operation, err error) {
	TS := service.PollService()
	res, err := TS.FetchOperations(context.TODO(), rev, count)
	return res, err
}

//func FetchOps(rev int64, count int32) (r []*LineThrift.Operation, err error) {
    //TS := service.TalkService()
    //res, err := TS.FetchOps(context.TODO(), rev, count,0,0)
    //return res, err
//}
/* Message Functions */

func SendText(to string, text string, toType LineThrift.MIDType) (*LineThrift.Message, error) {
	TS := service.TalkService()
	M := &LineThrift.Message{
		From_: service.MID,
		To: to,
		ToType: toType,
		Text: text,
		ContentType: 0,
		ContentMetadata: nil,
		RelatedMessageId: "0", // to be honest, i don't know what this is for, and if i don't throw something it wouldn't send the message
	}
	res, err := TS.SendMessage(context.TODO(), int32(0), M)
	return res, err
}
type mention struct {
	S string `json:"S"`
	E string `json:"E"`
	M string `json:"M"`
}
type DataMention struct {
	S string `json:"S"`
	E string `json:"E"`
	M string `json:"M"`
}

type mentionMsg struct {
	MENTIONEES []DataMention `json:"MENTIONEES"`
}

func SendMention(to string, text string, mids []string) (*LineThrift.Message, error) {
	TS := service.TalkService()
	arr := []*mention{}
	mentionee := "@Z"
	texts := strings.Split(text, "@!")
	if len(mids) == 0 || len(texts) < len(mids) { return &LineThrift.Message{}, fmt.Errorf("Invalid mids") }
	textx := ""
	for i := 0; i < len(mids); i++ {
		textx += texts[i]
        arr = append(arr, &mention{S: strconv.Itoa(utf8.RuneCountInString(textx)), E: strconv.Itoa(utf8.RuneCountInString(textx) + 11), M:mids[i]})
        textx += mentionee
	}
	textx += texts[len(texts)-1]
	arrData,_ := json.Marshal(arr)
	w := make(map[string]string)
	w["MENTION"] = string(arrData)
	M := &LineThrift.Message{
		To:               to,
		ContentType:      0,
		ContentMetadata:  w,
		Text:             textx,
		RelatedMessageId: "0",
	}
	e , err := TS.SendMessage(context.TODO(), int32(0), M)
	return e, err
}
func SendMentionV3(send bool, to string, text string, mids []string) {
	TS := service.TalkService()
	if send == false {
		return
	}
	var arr []DataMention
	var textS string
	mention := "@Lipro "
	splAt := strings.Split(text, "@!")
	if len(mids) == 0 || len(splAt) < len(mids) { return}
	textS = ""
	for s := 0; s < len(mids); s++ {
		textS += splAt[s]
		slen := len(textS)
		elen := len(textS) + 15
		arrData := DataMention{
			S: strconv.Itoa(slen),
			E: strconv.Itoa(elen - 4),
			M: mids[s],
		}
		arr = append(arr, arrData)
		textS += mention
	}
	ggh := &mentionMsg{
		MENTIONEES: arr,
	}
	e, _ := json.Marshal(ggh)
	textS += splAt[len(mids)]
	w := make(map[string]string)
	w["MENTION"] = string(e)
	M := &LineThrift.Message{
		To:               to,
		ContentType:      0,
		ContentMetadata:  w,
		Text:             textS,
		RelatedMessageId: "0",
	}
	TS.SendMessage(context.TODO(), int32(0), M)
}
func SendMentionV2(to string, text string, mids []string) {
	TS := service.TalkService()
	var arr []DataMention
	var textS string
	mention := "@lipro "
	splAt := strings.Split(text, "@!")
	if len(mids) == 0 || len(splAt) < len(mids) { return}
	textS = ""
	for s := 0; s < len(mids); s++ {
		textS += splAt[s]
		slen := len(textS)
		elen := len(textS) + 15
		arrData := DataMention{
			S: strconv.Itoa(slen),
			E: strconv.Itoa(elen - 4),
			M: mids[s],
		}
		arr = append(arr, arrData)
		textS += mention
	}
	ggh := &mentionMsg{
		MENTIONEES: arr,
	}
	e, _ := json.Marshal(ggh)
	textS += splAt[len(mids)]
	w := make(map[string]string)
	w["MENTION"] = string(e)
	M := &LineThrift.Message{
		To:               to,
		ContentType:      0,
		ContentMetadata:  w,
		Text:             textS,
		RelatedMessageId: "0",
	}
	TS.SendMessage(context.TODO(), int32(0), M)
}
func UnsendMessage(messageId string) (err error) {
	TS := service.TalkService()
	err = TS.UnsendMessage(context.TODO(), int32(0), messageId)
	return err
}

func RequestResendMessage(senderMid string, messageId string) (err error) {
	TS := service.TalkService()
	err = TS.RequestResendMessage(context.TODO(), int32(0), senderMid, messageId)
	return err
}

func RespondResendMessage(receiverMid string, originalMessageId string, resendMessage *LineThrift.Message, errorCode LineThrift.ErrorCode) (err error) {
	TS := service.TalkService()
	err = TS.RespondResendMessage(context.TODO(), int32(0), receiverMid, originalMessageId, resendMessage, errorCode)
	return err
}

func RemoveMessage(messageId string) (r bool, err error) {
	TS := service.TalkService()
	res, err := TS.RemoveMessage(context.TODO(), messageId)
	return res, err
}

func RemoveAllMessages(lastMessageId string) (err error) {
	TS := service.TalkService()
	err = TS.RemoveAllMessages(context.TODO(), int32(0), lastMessageId)
	return err
}

func RemoveMessageFromMyHome(ctx context.Context, messageId string) (r bool, err error) {
	TS := service.TalkService()
	res, err := TS.RemoveMessageFromMyHome(context.TODO(), messageId)
	return res, err
}

func SendChatChecked(consumer string, lastMessageId string, sessionId int8) (err error) {
	TS := service.TalkService()
	err = TS.SendChatChecked(context.TODO(), int32(0), consumer, lastMessageId, sessionId)
	return err
}

func SendEvent(message *LineThrift.Message) (r *LineThrift.Message, err error) {
	TS := service.TalkService()
	res, err := TS.SendEvent(context.TODO(), int32(0), message)
	return res, err
}

func GetPreviousMessagesV2WithReadCount(messageBoxId string, endMessageId *LineThrift.MessageBoxV2MessageId, messagesCount int32) (r []*LineThrift.Message, err error) {
	TS := service.TalkService()
	res, err := TS.GetPreviousMessagesV2WithReadCount(context.TODO(), messageBoxId, endMessageId, messagesCount)
	return res, err
}

/* Contact Functions */

func BlockContact(id string) (err error) {
	TS := service.TalkService()
	err = TS.BlockContact(context.TODO(), int32(0), id)
	return err
}

func FindAndAddContactByMetaTag(userid string, reference string) (r *LineThrift.Contact, err error) {
	TS := service.TalkService()
	res, err := TS.FindAndAddContactByMetaTag(context.TODO(), int32(0), userid, reference)
	return res, err
}

func FindAndAddContactsByMid(mid string) (r map[string]*LineThrift.Contact, err error) {
	TS := service.TalkService()
	res, err := TS.FindAndAddContactsByMid(context.TODO(), int32(0), mid)
	return res, err
}

func FindAndAddContactsByEmail(emails []string) (r map[string]*LineThrift.Contact, err error) {
	TS := service.TalkService()
	res, err := TS.FindAndAddContactsByEmail(context.TODO(), int32(0), emails)
	return res, err
}

func FindAndAddContactsByUserid(userid string) (r map[string]*LineThrift.Contact, err error) {
	TS := service.TalkService()
	res, err := TS.FindAndAddContactsByUserid(context.TODO(), int32(0), userid)
	return res, err
}
func GetAllContactIds() (r []string, err error) {
	TS := service.TalkService()
	res, err := TS.GetAllContactIds(context.TODO())
	return res, err
}

func GetBlockedContactIds() (r []string, err error) {
	TS := service.TalkService()
	res, err := TS.GetBlockedContactIds(context.TODO())
	return res, err
}

func GetContact(id string) (r *LineThrift.Contact, err error) {
	TS := service.TalkService()
	res, err := TS.GetContact(context.TODO(), id)
	return res, err
}

func GetContacts(ids []string) (r []*LineThrift.Contact, err error) {
	TS := service.TalkService()
	res, err := TS.GetContacts(context.TODO(), ids)
	return res, err
}

func GetFavoriteMids() (r []string, err error) {
	TS := service.TalkService()
	res, err := TS.GetFavoriteMids(context.TODO())
	return res, err
}

func GetHiddenContactMids() (r []string, err error) {
	TS := service.TalkService()
	res, err := TS.GetHiddenContactMids(context.TODO())
	return res, err
}

/* Group Functions */

func CancelGroupInvitation(groupId string, contactIds []string) (err error) {
	TS := service.TalkService()
	err = TS.CancelGroupInvitation(context.TODO(), int32(0), groupId, contactIds)
	if err != nil{
		fmt.Println(err)
	}
	return err
}

func KickoutFromGroup(groupId string, contactIds []string) (err error) {
	TS := service.TalkService()
	err = TS.KickoutFromGroup(context.TODO(), int32(0), groupId, contactIds)
	if err != nil{
		fmt.Println(err)
	}
	return err
}

func InviteIntoGroup(groupId string, contactIds []string) (err error) {
	TS := service.TalkService()
	err = TS.InviteIntoGroup(context.TODO(), int32(0), groupId, contactIds)
	if err != nil{
		fmt.Println(err)
	}
	return err
}

func AcceptGroupInvitation(groupId string) (err error) {
	TS := service.TalkService()
	err = TS.AcceptGroupInvitation(context.TODO(), int32(0), groupId)
	return err
}

func LeaveGroup(groupId string) (err error) {
	TS := service.TalkService()
	err = TS.LeaveGroup(context.TODO(), int32(0), groupId)
	return err
}

func GetGroup(groupId string) (r *LineThrift.Group, err error) {
	TS := service.TalkService()
	res, err := TS.GetGroup(context.TODO(), groupId)
	if err != nil{
		fmt.Println(err)
	}
	return res, err
}

func UpdateGroup(group *LineThrift.Group) (err error) {
	TS := service.TalkService()
	err = TS.UpdateGroup(context.TODO(), int32(0), group)
	return err
}

func ReissueGroupTicket(groupMid string) (r string, err error) {
	TS := service.TalkService()
	res, err := TS.ReissueGroupTicket(context.TODO(), groupMid)
	if err != nil{
		fmt.Println(err)
	}
	return res, err
}

func FindGroupByTicket(ticketId string) (r *LineThrift.Group, err error) {
	TS := service.TalkService()
	res, err := TS.FindGroupByTicket(context.TODO(), ticketId)
	if err != nil{
		fmt.Println(err)
	}
	return res, err
}
func AcceptGroupInvitationByTicket(groupMid string, ticketId string) (err error) {
	TS := service.TalkService()
	err = TS.AcceptGroupInvitationByTicket(context.TODO(), int32(0), groupMid, ticketId)
	return err
}

func GetGroupIdsJoined() (r []string, err error) {
	TS := service.TalkService()
	res, err := TS.GetGroupIdsJoined(context.TODO())
	return res, err
}
func GetCompactGroup(groupId string) (r *LineThrift.Group, err error) {
	TS := service.TalkService()
	res, err := TS.GetCompactGroup(context.TODO(), groupId)
	if err != nil{
		fmt.Println(err)
	}
	return res, err
}
func CreateGroupV2(name string, contactIds []string) (r *LineThrift.Group, err error) {
	TS := service.TalkService()
	res, err := TS.CreateGroupV2(context.TODO(), int32(0), name, contactIds)
	if err != nil{
		fmt.Println(err)
	}
	return res, err
}
func GetGroupIdsInvited() (r []string, err error) {
	TS := service.TalkService()
	res, err := TS.GetGroupIdsInvited(context.TODO())
	return res, err
}
func GetGroupWithoutMembers(id string) (r *LineThrift.Group, err error) {
	TS := service.TalkService()
	r, err = TS.GetGroupWithoutMembers(context.TODO(), id)
	return r, err
}

