package main

import (
	"fmt"
	"strings"
	"runtime"
	"time"
	"io/ioutil"
	"math/rand"
	"net"
	"strconv"
	"os"
	"sync"
	"encoding/json"
	"net/http"
	lipro "./linego/LineThrift"
	"./linego/auth"
	"./linego/helper"
	"./linego/service"
	"./linego/talk"
	"./linego/config"
)
var argsRaw = os.Args
var name = argsRaw[1]
var Basename = ""
var Changepic = false
var conn net.Conn
var Member = []string{}
var Invite = []string{}
type User struct {
	Creator     []string `json:"creator"`
	ArgSname string `json:"sname"`
	Owner []string `json:"owner"`
	Master []string `json:"master"`
	Admins  []string `json:"admins"`
	Squad  []string `json:"squad"`
	Bots  []string `json:"bots"`
	Sq []string `json:"sq"`
	Antijs []string `json:"antijs"`
	ProQR []string `json:"proqr"`
	ProInvite []string `json:"proinvite"`
	ProKick []string `json:"prokick"`
	ProCancel []string `json:"procancel"`
	FastCans bool `json:"fastcan"`
	BackupMem bool `json:"backupmem"`
	Modebackup bool `json:"modebackup"`
	Modeajs bool `json:"modeajs"`
	AutoBL bool `json:"autobl"`
	Nukick bool `json:"nukick"`
	Nucancel bool `json:"nucancel"`
	ByeMem bool `json:"byemem"`
	Silent bool `json:"silent"`
	Team string `json:"team"`
	KickBan bool `json:"kickban"`
	Limiter      []string `json:"limiter"`
}
type LINE struct {
	data *User
}
type Cpp struct {
	Respons string `json:"respons"`
}
var letters = []rune("0123456789")
func randomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
//ALL FUNCTIONS ARRAY//
func (a *LINE) IsCreator(from string) bool {
	if helper.InArray(a.data.Creator, from) == true {
		return true
	}
	return false
}
func (a *LINE) IsOwner(from string) bool {
	if helper.InArray(a.data.Owner, from) == true {
		return true
	}
	return false
}
func (a *LINE) IsAjs(from string) bool {
	if helper.InArray(a.data.Antijs, from) == true {
		return true
	}
	return false
}
func (a *LINE) IsAdmin(from string) bool {
	if helper.InArray(a.data.Admins, from) == true {
		return true
	}
	return false
}
func (a *LINE) IsAllStaff(from string) bool {
	if helper.InArray(a.data.Creator, from) == true || helper.InArray(a.data.Owner, from) == true || helper.InArray(a.data.Admins, from) == true{
		return true
	}
	return false
}
func (a *LINE) IsStaff(from string) bool {
	if helper.InArray(a.data.Creator, from) == true || helper.InArray(a.data.Owner, from) == true {
		return true
	}
	return false
}
func (a *LINE) IsINviter(from string) bool {
	if helper.InArray(a.data.Creator, from) == true || helper.InArray(a.data.Owner, from) == true || helper.InArray(a.data.Squad, from) == true || helper.InArray(a.data.Antijs, from) == true {
		return true
	}
	return false
}
func (a *LINE) IsAccess(from string) bool {
	if helper.InArray(a.data.Creator, from) == true || helper.InArray(a.data.Sq, from) == true || helper.InArray(a.data.Owner, from) == true || helper.InArray(a.data.Admins, from) == true || helper.InArray(a.data.Squad, from) == true || helper.InArray(a.data.Antijs, from) == true || helper.InArray(a.data.Bots, from) == true{
		return true
	}
	return false
}
func (a *LINE) IsNosquad(from string) bool {
	if helper.InArray(a.data.Creator, from) == true || helper.InArray(a.data.Owner, from) == true || helper.InArray(a.data.Admins, from) == true || helper.InArray(a.data.Antijs, from) == true || helper.InArray(a.data.Bots, from) == true{
		return true
	}
	return false
}
func (a *LINE) IsSquad(from string) bool {
	if helper.InArray(a.data.Squad, from) == true {
		return true
	}
	return false
}
func (a *LINE) IsBots(from string) bool {
	if helper.InArray(a.data.Bots, from) == true {
		return true
	}
	return false
}
func (a *LINE) IsLimiter(from string) bool {
	if helper.InArray(a.data.Limiter, from) == true {
		return true
	}
	return false
}
func (a *LINE) IsBackup(from string) bool {
	if helper.InArray(a.data.Sq, from) == true {
		return true
	}
	return false
}
func (a *LINE) AddBackup(pelaku string) {if !a.IsBackup(pelaku) {a.data.Sq = append(a.data.Sq, pelaku)}}
func (a *LINE) DelBackup(pelaku string) {if a.IsBackup(pelaku) {a.data.Sq = helper.Remove(a.data.Sq, pelaku)}}
//===========//
func IsMember(from string) bool {
	if helper.InArray(Member, from) == true {
		return true
	}
	return false
}
func AddMem(pelaku string) {if !IsMember(pelaku) {Member = append(Member, pelaku)}}
func DelMem(pelaku string) {if IsMember(pelaku) {Member = helper.Remove(Member, pelaku)}}
//===========//
func IsInvite(from string) bool {
	if helper.InArray(Invite, from) == true {
		return true
	}
	return false
}
func AddInvite(pelaku string) {if !IsInvite(pelaku) {Invite = append(Invite, pelaku)}}
//===========//
func (a *LINE) IsMaster(from string) bool {
	if helper.InArray(a.data.Master, from) == true {
		return true
	}
	return false
}
func (a *LINE) Addcreator(pelaku string) {if !a.IsCreator(pelaku) {a.data.Creator = append(a.data.Creator, pelaku)}}
func (a *LINE) Delcreator(pelaku string) {if a.IsCreator(pelaku) {a.data.Creator = helper.Remove(a.data.Creator, pelaku)}}
func checkEqual(list1 []string, list2 []string) bool {
	for _, v := range list1 {
		if helper.InArray(list2, v) {
			return true
		}
	}
	return false
}
func (a *LINE) Addowner(pelaku string) {
	if !a.IsOwner(pelaku) {
		a.data.Owner = append(a.data.Owner, pelaku)
	}
}
func (a *LINE) Addadmin(pelaku string) {
	if !a.IsAdmin(pelaku) {
		a.data.Admins = append(a.data.Admins, pelaku)
	}
}
func (a *LINE) Addsquad(pelaku string) {
	if !a.IsSquad(pelaku) {
		a.data.Squad = append(a.data.Squad, pelaku)
	}
}
func (a *LINE) Addbots(pelaku string) {
	if !a.IsBots(pelaku) {
		a.data.Bots = append(a.data.Bots, pelaku)
	}
}
func (a *LINE) Addajs(pelaku string) {
	if !a.IsAjs(pelaku) {
		a.data.Antijs = append(a.data.Antijs, pelaku)
	}
}
func (a *LINE) Delajs(pelaku string) {
	if a.IsAjs(pelaku) {
		a.data.Antijs = helper.Remove(a.data.Antijs, pelaku)
	}
}
func (a *LINE) Delbots(pelaku string) {
	if a.IsBots(pelaku) {
		a.data.Bots = helper.Remove(a.data.Bots, pelaku)
	}
}
func (a *LINE) Delowner(pelaku string) {
	if a.IsOwner(pelaku) {
		a.data.Owner = helper.Remove(a.data.Owner, pelaku)
	}
}
func (a *LINE) Deladmin(pelaku string) {
	if a.IsAdmin(pelaku) {
		a.data.Admins = helper.Remove(a.data.Admins, pelaku)
	}
}
func (a *LINE) Delsquad(pelaku string) {
	if a.IsSquad(pelaku) {
		a.data.Squad = helper.Remove(a.data.Squad, pelaku)
	}
}
func (a *LINE) AddLimiter(pelaku string) {
	if !a.IsLimiter(pelaku) {
		a.data.Limiter = append(a.data.Limiter, pelaku)
	}
}
func (a *LINE) DelLimiter(pelaku string) {
	if a.IsLimiter(pelaku) {
		a.data.Limiter = helper.Remove(a.data.Limiter, pelaku)
	}
}
func restart() {
    procAttr := new(os.ProcAttr)
    procAttr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}
    os.StartProcess(os.Args[0], []string{"", "test"}, procAttr)
}
func (a *LINE) Comand(op *lipro.Operation) {
	if op.Type == 26 {
		msg := op.Message
		sender := msg.From_
		var sname = a.data.ArgSname
		var rname = name
		var to = msg.To
		var txt string
		var pesan = strings.ToLower(msg.Text)

        if msg.ContentType == 1{
			if Changepic == true{
				runtime.GOMAXPROCS(10)
				go func(){
					response,_:= http.Get("https://api.vhtear.com/changepict?apikey=7a70add08644456cbf0beca26966f115&mid=" + service.MID + "&msgid=" + msg.ID + "&token=" + service.AuthToken)
					fmt.Println(response)
					talk.SendText(to,"Success changed",2)
					Changepic = false
				}()
			}
		}
        if msg.ContentType == 0 {
			if strings.HasPrefix(pesan , rname + " ") {
				txt = strings.Replace(pesan, rname + " ", "", 1)
			} else if strings.HasPrefix(pesan , rname) {
				txt = strings.Replace(pesan, rname, "", 1)
			} else if strings.HasPrefix(pesan , sname + " ") {
				txt = strings.Replace(pesan, sname + " ", "", 1)
			} else if strings.HasPrefix(pesan , sname){
				txt = strings.Replace(pesan, sname, "", 1)

			} else if pesan == "me" {
					if sender != "" && a.IsCreator(msg.From_) {
						if a.data.Silent == false {
							talk.SendMentionV3(true, msg.To, "Hello My Creator !!\n @!     \nHow are u today?!! ", []string{sender})
						}
					}
					if sender  != "" && a.IsOwner(msg.From_) {
						if a.data.Silent == false {
							talk.SendMentionV3(true, msg.To, "Hello My Owner !! \n @!      ", []string{sender})
						}
					}
					if sender  != "" && a.IsAdmin(msg.From_) {
						if a.data.Silent == false {
							talk.SendMentionV3(true, msg.To, "Hello My Admin !!\n @!      ", []string{sender})
						}
					}
			}
			if sender != "" && a.IsMaster(msg.From_) {
				if strings.HasPrefix(pesan , "delcr") {
					str := fmt.Sprintf("%v",msg.ContentMetadata["MENTION"])
					taglist := helper.GetMidFromMentionees(str)
					if taglist != nil {
						lisa := []string{}
						for c,target := range taglist {
							if a.IsCreator(target) {
								a.Delcreator(target)
							}
							res,_ := talk.GetContact(target)
							name := res.DisplayName
							c += 1
							name = fmt.Sprintf("%v. %s",c , name)
							lisa = append(lisa, name)
						}
						stf := "⊶succes delete to creatorlist\n\n"
						str := strings.Join(lisa, "\n")
						if a.data.Silent == false {
							talk.SendText(msg.To, stf+str, 2)
						}
					}
				} else if txt == "rcreator" {
					jum := len(a.data.Creator)
					a.data.Creator = []string{}
					str := fmt.Sprintf("⊶ɑlreɑdy cleɑned %v creatorlist.", jum)
					talk.SendText(msg.To, str, 2)
					for _, d := range a.data.Master {
						a.Addcreator(d)
					}
				} else if strings.HasPrefix(txt , "addcr") {
					str := fmt.Sprintf("%v",msg.ContentMetadata["MENTION"])
					taglist := helper.GetMidFromMentionees(str)
					if taglist != nil {
						lisa := []string{}
						for c,target := range taglist {
							if target != service.MID && !a.IsAdmin(target) && !a.IsOwner(target) && !a.IsCreator(target) {
								a.Addcreator(target)
							}
							res,_ := talk.GetContact(target)
							name := res.DisplayName
							c += 1
							name = fmt.Sprintf("%v. %s",c , name)
							lisa = append(lisa, name)
						}
						stf := "⊶succes ɑdded to creatorlist\n\n"
						str := strings.Join(lisa, "\n")
						if a.data.Silent == false {
							talk.SendText(msg.To, stf+str, 2)
						}
					}
				} else if strings.HasPrefix(txt , "addbc") {
					str := fmt.Sprintf("%v",msg.ContentMetadata["MENTION"])
					taglist := helper.GetMidFromMentionees(str)
					if taglist != nil {
						lisa := []string{}
						for c,target := range taglist {
							if !a.IsBackup(target) {
								a.AddBackup(target)
							}
							res,_ := talk.GetContact(target)
							name := res.DisplayName
							c += 1
							name = fmt.Sprintf("%v. %s",c , name)
							lisa = append(lisa, name)
						}
						stf := "⊶succes ɑdded to backuplist\n\n"
						str := strings.Join(lisa, "\n")
						if a.data.Silent == false {
							talk.SendText(msg.To, stf+str, 2)
						}
					}
				} else if strings.HasPrefix(txt , "delbc") {
					str := fmt.Sprintf("%v",msg.ContentMetadata["MENTION"])
					taglist := helper.GetMidFromMentionees(str)
					if taglist != nil {
						lisa := []string{}
						for c,target := range taglist {
							if a.IsBackup(target) {
								a.DelBackup(target)
							}
							res,_ := talk.GetContact(target)
							name := res.DisplayName
							c += 1
							name = fmt.Sprintf("%v. %s",c , name)
							lisa = append(lisa, name)
						}
						stf := "⊶succes delete to backuplist\n\n"
						str := strings.Join(lisa, "\n")
						if a.data.Silent == false {
							talk.SendText(msg.To, stf+str, 2)
						}
					}
				} else if txt == "sqlist"{
						nm := []string{}
						for c, v := range a.data.Sq {
							res,_ := talk.GetContact(v)
							name := res.DisplayName
							c += 1
							name = fmt.Sprintf("%v. %s",c , name)
							nm = append(nm, name)
						}
						stf := "• Backupʟɪsᴛᴇᴅ •\n\n"
						str := strings.Join(nm, "\n")
						if a.data.Silent == false {
							talk.SendText(msg.To, stf+str, 2)
						}
				} else if txt == "mem"{
						nm := []string{}
						for c, v := range Member {
							res,_ := talk.GetContact(v)
							name := res.DisplayName
							c += 1
							name = fmt.Sprintf("%v. %s",c , name)
							nm = append(nm, name)
						}
						stf := "• Member •\n\n"
						str := strings.Join(nm, "\n")
						if a.data.Silent == false {
							talk.SendText(msg.To, stf+str, 2)
						}
				} else if txt == "pend"{
						nm := []string{}
						for c, v := range Invite {
							res,_ := talk.GetContact(v)
							name := res.DisplayName
							c += 1
							name = fmt.Sprintf("%v. %s",c , name)
							nm = append(nm, name)
						}
						stf := "• Invitedlist •\n\n"
						str := strings.Join(nm, "\n")
						if a.data.Silent == false {
							talk.SendText(msg.To, stf+str, 2)
						}
				} else if txt == "csq" {
					jum := len(a.data.Sq)
					a.data.Sq = []string{}
					str := fmt.Sprintf("CLEARED %v Backup", jum)
					if a.data.Silent == false {
						talk.SendText(msg.To, str, 2)
					}
				}
			}
//MESSAGE FROM CREATOR
			if sender != "" && a.IsCreator(msg.From_) {
				if txt == "ownerlist"{
						nm := []string{}
						for c, v := range a.data.Owner {
							res,_ := talk.GetContact(v)
							name := res.DisplayName
							c += 1
							name = fmt.Sprintf("%v. %s",c , name)
							nm = append(nm, name)
						}
						stf := "• ownerʟɪsᴛᴇᴅ •\n\n"
						str := strings.Join(nm, "\n")
						if a.data.Silent == false {
							talk.SendText(msg.To, stf+str, 2)
						}
				} else if strings.HasPrefix(txt, "team:") {
					newres := strings.Split(txt, " ")
					for k, v := range newres {
						if strings.ToLower(v) == "team:" {
							new := strings.Split(txt, newres[k]+" ")[1]
							a.data.Team = new
							talk.SendText(msg.To, "Title Update to: "+new, 2)
						}
						break
					}
				} else if strings.HasPrefix(txt, "groups") {
							gr, _ := talk.GetGroupIdsJoined()
							num := ">  grouplist  <\n"
							for k, v := range gr {
								g, _ := talk.GetGroupWithoutMembers(v)
								num += "\n" + strconv.Itoa(k) + ". " + g.Name
							}
							talk.SendText(msg.To, num, 2)
				} else if strings.HasPrefix(txt, "invto") {
							newspl := strings.Split(txt, " ")
							for _, v := range newspl {
								if strings.ToLower(v) == "invto" {
									hinum := strings.Split(txt, v+" ")[1]
									new, _ := strconv.Atoi(hinum)
									gr, _ := talk.GetGroupIdsJoined()
									gid := gr[new]
									err := talk.InviteIntoGroup(gid, []string{msg.From_})
									if err == nil {
										xa, _ := talk.GetGroupWithoutMembers(gid)
										talk.SendText(msg.To, "Succes InviteTo: "+xa.Name, 2)
									} else {
										talk.SendText(msg.To, "Bot limit, U Can Try again later!!", 2)
									}
									break
								}
							}
				} else if strings.HasPrefix(txt, "ginvitedlist") {
							gr, _ := talk.GetGroupIdsInvited()
							num := ">  GroupInvitedList  <\n"
							for k, v := range gr {
								g, _ := talk.GetGroupWithoutMembers(v)
								num += "\n" + strconv.Itoa(k) + ". " + g.Name
							}
							talk.SendText(msg.To, num, 2)
				} else if strings.HasPrefix(txt, "accno") {
							newspl := strings.Split(txt, " ")
							for _, v := range newspl {
								if strings.ToLower(v) == "accno" {
									hinum := strings.Split(txt, v+" ")[1]
									new, _ := strconv.Atoi(hinum)
									gr, _ := talk.GetGroupIdsInvited()
									gid := gr[new]
									err := talk.AcceptGroupInvitation(gid)
									if err == nil {
										xa, _ := talk.GetGroupWithoutMembers(gid)
										talk.SendText(msg.To, "Succes Accept To: "+xa.Name, 2)
									} else {
										talk.SendText(msg.To, "Bot limit, U Can Try again later!!", 2)
									}
									break
								}
							}
				} else if strings.HasPrefix(strings.ToLower(msg.Text),"join") {
					ticketId := strings.Split(msg.Text, " ")
					g, err := talk.FindGroupByTicket(ticketId[1])
					if err != nil {
						talk.SendText(msg.To, "Link Not Found Or Link Was Close",2)
					} else {
						talk.SendText(msg.To, "Succes Join Group"+g.Name, 2)
						talk.AcceptGroupInvitationByTicket(g.ID, ticketId[1])
					}
				} else if txt == "crlist" {
					if len(a.data.Creator) > 0 {
						tx := "> CREATOR LIST < \n"
						for num := 1; num <= len(a.data.Creator); num++ {
							tx += "\n" + strconv.Itoa(num) + ".    @!          "
						}
						if a.data.Silent == false {
							talk.SendMentionV3(true, msg.To, tx + "\n\n  Total   --->   "+strconv.Itoa(len(a.data.Creator)),a.data.Creator)
						}
					} else {
						if a.data.Silent == false {
							talk.SendText(msg.To, "No one in creatorlist.", 2)
						}
					}
				} else if txt == "blacklist" {
					if len(service.Banned) > 0 {
						tx := ">  blacklisted  < \n"
						for num := 1; num <= len(service.Banned); num++ {
							tx += "\n" + strconv.Itoa(num) + ".    @!          "
						}
						if a.data.Silent == false {
							talk.SendMentionV3(true, msg.To, tx + "\n\n  please cleared   --->   "+strconv.Itoa(len(service.Banned)),service.Banned)
						}
					} else {
						if a.data.Silent == false {
							talk.SendText(msg.To, "𝑁𝑜𝑡ℎ𝑖𝑛𝑔 𝐵𝑙𝑎𝑐𝑘𝑙𝑖𝑠𝑡 𝐷𝑎𝑡𝑎.", 2)
						}
					}
				} else if strings.HasPrefix(txt, "leaveto") {
							newspl := strings.Split(txt, " ")
							for _, v := range newspl {
								if strings.ToLower(v) == "leaveto" {
									hinum := strings.Split(txt, v+" ")[1]
									new, _ := strconv.Atoi(hinum)
									gr, _ := talk.GetGroupIdsJoined()
									gid := gr[new]
									err := talk.LeaveGroup(gid)
									if err == nil {
										xa, _ := talk.GetGroupWithoutMembers(gid)
										talk.SendText(msg.To, "Succes Leave From: "+xa.Name, 2)
									} else {
										talk.SendText(msg.To, "Bot Limit Please Try latter!!", 2)
									}
									break
								}
							}
				} else if strings.HasPrefix(txt, "cancelajsno") {
							newspl := strings.Split(txt, " ")
							for _, v := range newspl {
								if strings.ToLower(v) == "cancelajsno" {
									hinum := strings.Split(txt, v+" ")[1]
									new, _ := strconv.Atoi(hinum)
									gr, _ := talk.GetGroupIdsJoined()
									gid := gr[new]
									for _, d := range a.data.Antijs {
										md := []string{d}
										err := talk.CancelGroupInvitation(gid, md)
										if err == nil {
											xa, _ := talk.GetGroupWithoutMembers(gid)
											talk.SendMentionV3(true, msg.To, " @!         \nSucces CancelAjs From: "+xa.Name, a.data.Creator)
										} else {
											talk.SendText(msg.To, "Bot AntiJs Not Groups, Try Check!!", 2)
										}
										break
									}
								}
							}
				} else if strings.HasPrefix(txt, "invajsno") {
							newspl := strings.Split(txt, " ")
							for _, v := range newspl {
								if strings.ToLower(v) == "invajsno" {
									hinum := strings.Split(txt, v+" ")[1]
									new, _ := strconv.Atoi(hinum)
									gr, _ := talk.GetGroupIdsJoined()
									gid := gr[new]
									err := talk.InviteIntoGroup(gid, a.data.Antijs)
									if err == nil {
										xa, _ := talk.GetGroupWithoutMembers(gid)
										talk.SendMentionV3(true, msg.To, " @!         \nSucces Invite Ajs To: "+xa.Name, a.data.Creator)
									} else {
										talk.SendText(msg.To, "Bot Limit Or Ajs was Stay In PendingList Groups, Try Check!!", 2)
									}
									break
								}
							}
				} else if strings.HasPrefix(txt, "bypasno") {
						newspl := strings.Split(txt, " ")
						for _, v := range newspl {
							if strings.ToLower(v) == "bypasno" {
								hinum := strings.Split(txt, v+" ")[1]
								new, _ := strconv.Atoi(hinum)
								gr, _ := talk.GetGroupIdsJoined()
								gid := gr[new]
								res, _ := talk.GetGroup(gid)
								cn := res.ID
								go func (){
									a.Bypass(cn)
								}()
								cek := res.PreventedJoinByTicket
								if cek{
									res.PreventedJoinByTicket = false
									talk.UpdateGroup(res)
									gurl,_ := talk.ReissueGroupTicket(gid)
									talk.SendText(msg.To, "Succes Bypass In Group: "+res.Name+"\nLink Ticket: line://ti/g/"+gurl, 2)
								} else {
									talk.SendText(msg.To, "Bot not in Groups Try again later!!", 2)
								}
							}
						}
				} else if strings.HasPrefix(txt, "kickallno") {
						newspl := strings.Split(txt, " ")
						for _, v := range newspl {
							if strings.ToLower(v) == "kickallno" {
								hinum := strings.Split(txt, v+" ")[1]
								new, _ := strconv.Atoi(hinum)
								gr, _ := talk.GetGroupIdsJoined()
								gid := gr[new]
								res, _ := talk.GetGroup(gid)
								cn := res.ID
								go func (){
									a.NukerAll(cn)
								}()
								cek := res.PreventedJoinByTicket
								if cek{
									res.PreventedJoinByTicket = false
									talk.UpdateGroup(res)
									gurl,_ := talk.ReissueGroupTicket(gid)
									talk.SendText(msg.To, "Succes Kickall Members: "+res.Name+"\nLink Ticket: line://ti/g/"+gurl, 2)
								} else {
									talk.SendText(msg.To, "Bot not in Groups Try again later!!", 2)
								}
							}
						}
				} else if strings.HasPrefix(txt, "cancelallno") {
						newspl := strings.Split(txt, " ")
						for _, v := range newspl {
							if strings.ToLower(v) == "cancelallno" {
								hinum := strings.Split(txt, v+" ")[1]
								new, _ := strconv.Atoi(hinum)
								gr, _ := talk.GetGroupIdsJoined()
								gid := gr[new]
								res, _ := talk.GetGroup(gid)
								cn := res.ID
								go func (){
									a.Nukcancel(cn)
								}()
								cek := res.PreventedJoinByTicket
								if cek{
									res.PreventedJoinByTicket = false
									talk.UpdateGroup(res)
									gurl,_ := talk.ReissueGroupTicket(gid)
									talk.SendText(msg.To, "Succes Cancelall Members: "+res.Name+"\nLink Ticket: line://ti/g/"+gurl, 2)
								} else {
									talk.SendText(msg.To, "Bot not in Groups Try again later!!", 2)
								}
							}
						}
				} else if strings.HasPrefix(txt, "qrno") {
						newspl := strings.Split(txt, " ")
						for _, v := range newspl {
							if strings.ToLower(v) == "qrno" {
								hinum := strings.Split(txt, v+" ")[1]
								new, _ := strconv.Atoi(hinum)
								gr, _ := talk.GetGroupIdsJoined()
								gid := gr[new]
								res, _ := talk.GetGroup(gid)
								cek := res.PreventedJoinByTicket
								if cek {
									res.PreventedJoinByTicket = false
									talk.UpdateGroup(res)
									t, err := talk.ReissueGroupTicket(gid)
									if err == nil {
										xa, _ := talk.GetGroupWithoutMembers(gid)
										talk.SendText(msg.To, "Group Name: "+xa.Name+"\nLink Ticket: line://ti/g/"+t, 2)
									}
									break
								}else {
									res.PreventedJoinByTicket = true
									t, err := talk.ReissueGroupTicket(gid)
									if err == nil {
										xa, _ := talk.GetGroupWithoutMembers(gid)
										talk.SendText(msg.To, "Group Name: "+xa.Name+"\nLink Ticket: line://ti/g/"+t, 2)
									}
									break
								}
							}
						}
				} else if strings.HasPrefix(txt, "closeqrno") {
						newspl := strings.Split(txt, " ")
						for _, v := range newspl {
							if strings.ToLower(v) == "closeqrno" {
								hinum := strings.Split(txt, v+" ")[1]
								new, _ := strconv.Atoi(hinum)
								gr, _ := talk.GetGroupIdsJoined()
								gid := gr[new]
								res, _ := talk.GetGroup(gid)
								cek := res.PreventedJoinByTicket
								if cek {
									res.PreventedJoinByTicket = false
									talk.SendMentionV3(true, msg.To, " @!           \nLink Was Close ", a.data.Creator)
								}else{
									res.PreventedJoinByTicket = true
									err := talk.UpdateGroup(res)
									if err == nil {
										xa, _ := talk.GetGroupWithoutMembers(gid)
										talk.SendMentionV3(true, msg.To, " @!           \nSucces Close Link Group: "+xa.Name, a.data.Creator)
									}
									break
								}
							}
						}
				} else if strings.HasPrefix(txt , "delowner") {
					str := fmt.Sprintf("%v",msg.ContentMetadata["MENTION"])
					taglist := helper.GetMidFromMentionees(str)
					if taglist != nil {
						lisa := []string{}
						for c,target := range taglist {
							if a.IsOwner(target) {
								a.Delowner(target)
							}
							res,_ := talk.GetContact(target)
							name := res.DisplayName
							c += 1
							name = fmt.Sprintf("%v. %s",c , name)
							lisa = append(lisa, name)
						}
						stf := "⊶succes delete to ownerlist\n\n"
						str := strings.Join(lisa, "\n")
						if a.data.Silent == false {
							talk.SendText(msg.To, stf+str, 2)
						}
					}
				} else if strings.HasPrefix(txt , "addowner") {
					str := fmt.Sprintf("%v",msg.ContentMetadata["MENTION"])
					taglist := helper.GetMidFromMentionees(str)
					if taglist != nil {
						lisa := []string{}
						for c,target := range taglist {
							if target != service.MID && !a.IsAdmin(target) && !a.IsOwner(target) && !a.IsCreator(target) {
								a.Addowner(target)
							}
							res,_ := talk.GetContact(target)
							name := res.DisplayName
							c += 1
							name = fmt.Sprintf("%v. %s",c , name)
							lisa = append(lisa, name)
						}
						stf := "⊶succes ɑdded to ownerlist\n\n"
						str := strings.Join(lisa, "\n")
						if a.data.Silent == false {
							talk.SendText(msg.To, stf+str, 2)
						}
					}
				} else if txt == "linkall"{
						nm := []string{}
						grup,_ := talk.GetGroupIdsJoined()
						for c, v := range grup {
							if v != msg.To {
								res,_ := talk.GetGroup(v)
								name := res.Name
								c += 1
								name = fmt.Sprintf("%v. %s",c , name)
								nm = append(nm, name)
								cek := res.PreventedJoinByTicket
								if cek {
									res.PreventedJoinByTicket = false
									talk.UpdateGroup(res)
									gurl,_ := talk.ReissueGroupTicket(v)
									str := fmt.Sprintf("line://ti/g/%v", gurl)
									talk.SendText(sender, str, 2)
								} else {
									res.PreventedJoinByTicket = true
									gurl,_ := talk.ReissueGroupTicket(v)
									str := fmt.Sprintf("line://ti/g/%v", gurl)
									talk.SendText(sender, str, 2)
								}
							}
						}
						if a.data.Silent == false {
							talk.SendText(msg.To, "please read personal message", 2)
						}
				} else if txt == "kickban on" {
					if a.data.KickBan == true {
						if a.data.Silent == false {
							talk.SendText(msg.To, "Auto Kick Banlist On ", 2)
						}
					} else {
						a.data.KickBan = true
						if a.data.Silent == false {
							talk.SendText(msg.To, "Skuy Auto Kick Banned True ", 2)
						}
					}
				} else if txt == "kickban off" {
					if a.data.KickBan == false {
						if a.data.Silent == false {
							talk.SendText(msg.To, "Auto Kick BanList Off ", 2)
						}
					} else {
						a.data.KickBan = false
						if a.data.Silent == false {
							talk.SendText(msg.To, "Skuy Auto Kick Banned False", 2)
						}
					}
				} else if txt == "war on" {
					if a.data.AutoBL == true {
						if a.data.Silent == false {
							talk.SendText(msg.To, "Ready War on ", 2)
						}
					} else {
						a.data.AutoBL = true
						if a.data.Silent == false {
							talk.SendText(msg.To, "Skuy War true ", 2)
						}
					}
				} else if txt == "war off" {
					if a.data.AutoBL == false {
						if a.data.Silent == false {
							talk.SendText(msg.To, "Mode War off ", 2)
						}
					} else {
						a.data.AutoBL = false
						if a.data.Silent == false {
							talk.SendText(msg.To, "Skuy War false", 2)
						}
					}
				} else if txt == "silent" {
					if a.data.Silent == true {
						if a.data.Silent == false {
							talk.SendText(msg.To, "Was Enanle Slient ", 2)
						}
					} else {
						a.data.Silent = true
						if a.data.Silent == false {
							talk.SendText(msg.To, "Skuy Silent true ", 2)
						}
					}
				} else if txt == "unsilent" {
					if a.data.Silent == false {
						talk.SendText(msg.To, "Was Disabled ", 2)
					} else {
						a.data.Silent = false
						talk.SendText(msg.To, "Skuy Silent false", 2)
					}
				} else if txt == "fastcan on" {
					if a.data.FastCans == true {
						if a.data.Silent == false {
							talk.SendText(msg.To, "Ready FastCans ", 2)
						}
					} else {
						a.data.FastCans = true
						if a.data.Silent == false {
							talk.SendText(msg.To, "Skuy FastCans true ", 2)
						}
					}
				} else if txt == "fastcan off" {
					if a.data.FastCans == false {
						if a.data.Silent == false {
							talk.SendText(msg.To, "FastCans Was Off ", 2)
						}
					} else {
						a.data.FastCans = false
						if a.data.Silent == false {
							talk.SendText(msg.To, "Skuy FastCans False ", 2)
						}
					}
				} else if txt == "backupmem on" {
					if a.data.BackupMem == true {
						if a.data.Silent == false {
							talk.SendText(msg.To, "Ready Backup Member ", 2)
						}
					} else {
						a.data.BackupMem = true
						if a.data.Silent == false {
							talk.SendText(msg.To, "Skuy Backup Member true ", 2)
						}
					}
				} else if txt == "backupmem off" {
					if a.data.BackupMem == false {
						if a.data.Silent == false {
							talk.SendText(msg.To, "Backup Memb Was Off ", 2)
						}
					} else {
						a.data.BackupMem = false
						if a.data.Silent == false {
							talk.SendText(msg.To, "Skuy Backup Memb False ", 2)
						}
					}
				} else if txt == "cancelall on" {
					if a.data.Nucancel == true {
						talk.SendText(msg.To, "Cancelall Was on ", 2)
					} else {
						a.data.Nukick = false
						a.data.ByeMem = false
						a.data.Nucancel = true
						talk.SendText(msg.To, "Nuker Mode Cancelall true ", 2)
					}
				} else if txt == "bypass on" {
					if a.data.ByeMem == true {
						talk.SendText(msg.To, "Bypass Was on ", 2)
					} else {
						a.data.Nukick = false
						a.data.ByeMem = true
						a.data.Nucancel = false
						talk.SendText(msg.To, "Nuker Mode Bypass true ", 2)
					}
				} else if txt == "nuker off" {
					if a.data.Nukick == false && a.data.ByeMem == false && a.data.Nucancel == false {
						talk.SendText(msg.To, "Auto Nuker Was off ", 2)
					} else {
						a.data.Nukick = false
						a.data.ByeMem = false
						a.data.Nucancel = false
						talk.SendText(msg.To, "Auto Nukers false ", 2)
					}
				} else if txt == "setmode backup" {
					if a.data.Modebackup == false {
						a.data.Modebackup = true
						time.Sleep(time.Second *2)
						a.data.Modeajs = false
						time.Sleep(time.Second *2)
						for x := range a.data.Sq {
							talk.SendText(a.data.Sq[x], "off",2)
						}
						if a.data.Silent == false {
							talk.SendText(msg.To, "Succes Change To Mode Backup ",2)
						}
					} else {
						if a.data.Silent == false {
							talk.SendText(msg.To, "Stay Mode Backup",2)
						}
					}
				} else if txt == "setmode ajs" {
					if a.data.Modeajs == true {
						if a.data.Silent == false {
							talk.SendText(msg.To, "Stay Mode AntiJs ",2)
						}
					} else {
						a.data.Modeajs = true
						time.Sleep(time.Second *2)
						a.data.Modebackup = false
						time.Sleep(time.Second *2)
						for x := range a.data.Sq {
							talk.SendText(a.data.Sq[x], "true",2)
						}
						if a.data.Silent == false {
							talk.SendText(msg.To, "Succes Change To Mode AntiJs ",2)
						}
					}
				} else if txt == "setnuke" {
					checking := []string{}
					stf := "Setting Auto Nuker Bots:\n\n"
					if a.data.ByeMem == true {
						na := "📎 Mode Bypass :: On"
						checking = append(checking, na)
					} else {
						na := "📎 Mode Bypass  :: Off"
						checking = append(checking, na)
					}
					if a.data.Nucancel == true {
						na := "📎 Mode CancelAll :: On"
						checking = append(checking, na)
					} else {
						na := "📎 Mode CancelAll  :: Off"
						checking = append(checking, na)
					}
					str := strings.Join(checking, "\n")
					talk.SendText(msg.To, stf+str, 2)
				}
			}
			if sender != "" && a.IsBackup(msg.From_) {
				if strings.HasPrefix(strings.ToLower(msg.Text),"acc") {
					ticketId := strings.Split(msg.Text, " ")
					g, err := talk.FindGroupByTicket(ticketId[1])
					if err != nil {
						fmt.Println("Link Tertutup")
					} else {
						talk.AcceptGroupInvitationByTicket(g.ID, ticketId[1])
					}
				}
			}
			if sender != "" && a.IsStaff(msg.From_) {
				if pesan == "updp" {
					Changepic = true
					if a.data.Silent == false {
						talk.SendText(msg.To,"⊶send pict", 2)
					}
				} else if txt == "uppict" {
					Changepic = true
					talk.SendText(msg.To,"⊶send pict", 2)
				} else if txt == "adminlist"{
						nm := []string{}
						for c, v := range a.data.Admins {
							res,_ := talk.GetContact(v)
							name := res.DisplayName
							c += 1
							name = fmt.Sprintf("%v. %s",c , name)
							nm = append(nm, name)
						}
						stf := "• ᴀᴅᴍɪɴsʟɪsᴛᴇᴅ •\n\n"
						str := strings.Join(nm, "\n")
						talk.SendText(msg.To, stf+str, 2)
				} else if txt == "acces"{
						nm := []string{}
						nmown := []string{}
						nmadm := []string{}
						for c, v := range a.data.Creator {
							res,_ := talk.GetContact(v)
							name := res.DisplayName
							c += 0
							name = fmt.Sprintf("    %v. %s",c , name)
							nm = append(nm, name)
						}
						stf1 := "• creators •\n\n"
						str1 := strings.Join(nm, "\n")
						for c, v := range a.data.Owner {
							res,_ := talk.GetContact(v)
							own := res.DisplayName
							c += 0
							own = fmt.Sprintf("    %v. %s",c , own)
							nmown = append(nmown, own)
						}
						stf2 := "\n\n• ᴏᴡɴᴇʀs •\n\n"
						str2 := strings.Join(nmown, "\n")
						for d, b := range a.data.Admins {
							res,_ := talk.GetContact(b)
							adm := res.DisplayName
							d += 1
							adm = fmt.Sprintf("    %v. %s",d , adm)
							nmadm = append(nmadm, adm)
						}
						stf3 := "\n\n• ᴀᴅᴍɪɴs •\n\n"
						str3 := strings.Join(nmadm, "\n")
						talk.SendText(msg.To, stf1+str1+stf2+str2+stf3+str3, 2)
				} else if strings.HasPrefix(txt,"sname:") {
					str := strings.Replace(txt,"sname:","", 1)
					nm := []string{}
					nm = append(nm,str)
					stl := strings.Join(nm,", ")
					a.data.ArgSname = stl
					talk.SendText(msg.To,"Succes update Sname to "+ str,2)
				} else if strings.HasPrefix(txt,"rname:") {
					str := strings.Replace(txt,"rname:","", 1)
					nm := []string{}
					nm = append(nm,str)
					stl := strings.Join(nm,", ")
					name = stl
					talk.SendText(msg.To,"Succes update Rname to "+ str,2)
				} else if strings.HasPrefix(txt,"upname:") {
					str := strings.Replace(txt,"upname:","", 1)
					res,_ := talk.GetProfile()
					res.DisplayName = str
					talk.UpdateProfile(res)
					talk.SendText(msg.To,"Succes update Profile Name to "+ str,2)
				} else if txt == "cek" || txt == "status" {
					for x := range a.data.Sq {
						if InMem(a.data.Sq[x], msg.To) {
							if service.MID == a.data.Sq[x] {
								anu := talk.InviteIntoGroup(msg.To, []string{service.MID})
								if anu != nil {
									talk.SendText(msg.To, "Limit", 2)
									a.AddLimiter(service.MID)
								} else {
									talk.SendText(msg.To, "Ready", 2)
									a.data.Limiter = []string{}
								}
							}
						}
					}
				} else if txt == "limit cek" {
					for x := range a.data.Limiter {
						if InMem(a.data.Limiter[x], msg.To) {
							if service.MID == a.data.Limiter[x] {
								anu := talk.InviteIntoGroup(msg.To, []string{service.MID})
								if anu != nil {
									talk.SendText(msg.To, "Limit", 2)
								} else {
									for x := range a.data.Sq {
										if service.MID != a.data.Sq[x] {
											talk.SendText(a.data.Sq[x], "Ready", 2)
										}
									}
									time.Sleep(time.Second *2)
									a.AddBackup(service.MID)
									time.Sleep(time.Second *2)
									a.data.Limiter = []string{}
								}
							}
						}
					}
				} else if pesan == "limit out" {
					for x := range a.data.Limiter {
						if InMem(a.data.Limiter[x], msg.To) {
							if service.MID == a.data.Limiter[x] {
								for x := range a.data.Sq {
									if service.MID != a.data.Sq[x] {
										talk.SendText(a.data.Sq[x], "limits", 2)
									}
								}
								time.Sleep(time.Second *2)
								talk.LeaveGroup(msg.To)
							}
							break
						}
					}
				} else if txt == "here"{
					gr,_:= talk.GetGroup(msg.To)
					mem := gr.Members
					b := []string{}
					for _, c := range mem {
						if helper.InArray(a.data.Sq, c.Mid) {
							b = append(b, c.Mid)
						}
					}
					if a.data.Silent == false {
						talk.SendText(msg.To, strconv.Itoa(len(b)+1)+"/"+strconv.Itoa(len(a.data.Sq)),2)
					}
				} else if txt == "squadlist"{
						nm := []string{}
						for c, v := range a.data.Squad {
							res,_ := talk.GetContact(v)
							name := res.DisplayName
							c += 1
							name = fmt.Sprintf("%v. %s",c , name)
							nm = append(nm, name)
						}
						stf := "• sqᴜᴀᴅsʟɪsᴛᴇᴅ •\n\n"
						str := strings.Join(nm, "\n")
						talk.SendText(msg.To, stf+str, 2)
				} else if txt == "cancelall" {
					runtime.GOMAXPROCS(100)
					go func() {
						a.Nukcancel(msg.To)
					}()
				} else if txt == "bypas" {
					runtime.GOMAXPROCS(100)
					go func(){
						a.Bypass(msg.To)
					}()
				} else if txt == "out" {
					talk.LeaveGroup(msg.To)
				} else if txt == "csquad" {
					jum := len(a.data.Squad)
					a.data.Squad = []string{}
					str := fmt.Sprintf("⊶ɑlreɑdy cleɑned %v squadlist.", jum)
					if a.data.Silent == false {
						talk.SendText(msg.To, str, 2)
					}
				} else if txt == "cowner" {
					jum := len(a.data.Owner)
					a.data.Owner = []string{}
					str := fmt.Sprintf("⊶ɑlreɑdy cleɑned %v ownerlist.", jum)
					if a.data.Silent == false {
						talk.SendText(msg.To, str, 2)
					}
				} else if txt == "cajs" {
					jum := len(a.data.Antijs)
					a.data.Antijs = []string{}
					str := fmt.Sprintf("⊶ɑlreɑdy cleɑned %v ajslist.", jum)
					if a.data.Silent == false {
						talk.SendText(msg.To, str, 2)
					}
				} else if txt == "cadmin" {
					jum := len(a.data.Admins)
					a.data.Admins = []string{}
					str := fmt.Sprintf("⊶ɑlreɑdy cleɑned %v adminlist.", jum)
					if a.data.Silent == false {
						talk.SendText(msg.To, str, 2)
					}
				} else if txt == "invbot" {
					talk.InviteIntoGroup(msg.To, a.data.Sq)
				} else if txt == "invsquad" {
					talk.InviteIntoGroup(msg.To, a.data.Squad)
				} else if strings.HasPrefix(txt , "addadmin") {
						str := fmt.Sprintf("%v",msg.ContentMetadata["MENTION"])
						taglist := helper.GetMidFromMentionees(str)
						if taglist != nil {
							lisa := []string{}
							for c,target := range taglist {
								if target != service.MID && !a.IsAdmin(target) && !a.IsOwner(target) && !a.IsCreator(target) {
									a.Addadmin(target)
								}
								res,_ := talk.GetContact(target)
								name := res.DisplayName
								c += 1
								name = fmt.Sprintf("%v. %s",c , name)
								lisa = append(lisa, name)
							}
							stf := "⊶succes ɑdded to ɑdminlist\n\n"
							str := strings.Join(lisa, "\n")
							talk.SendText(msg.To, stf+str, 2)
						}
				} else if strings.HasPrefix(txt , "addajs") {
						str := fmt.Sprintf("%v",msg.ContentMetadata["MENTION"])
						taglist := helper.GetMidFromMentionees(str)
						if taglist != nil {
							lisa := []string{}
							for c,target := range taglist {
								if target != service.MID && !a.IsAjs(target) && !a.IsSquad(target) && !a.IsOwner(target) && !a.IsAdmin(target) && !a.IsCreator(target) {
									a.Addajs(target)
								}
								res,_ := talk.GetContact(target)
								name := res.DisplayName
								c += 1
								name = fmt.Sprintf("%v. %s",c , name)
								lisa = append(lisa, name)
							}
							stf := "⊶succes ɑdded to antijs\n\n"
							str := strings.Join(lisa, "\n")
							talk.SendText(msg.To, stf+str, 2)
						}
				} else if strings.HasPrefix(txt , "addsquad") {
						str := fmt.Sprintf("%v",msg.ContentMetadata["MENTION"])
						taglist := helper.GetMidFromMentionees(str)
						if taglist != nil {
							lisa := []string{}
							for c,target := range taglist {
								if target != service.MID && !a.IsAjs(target) && !a.IsSquad(target) && !a.IsOwner(target) && !a.IsAdmin(target) && !a.IsCreator(target) {
									a.Addsquad(target)
								}
								res,_ := talk.GetContact(target)
								name := res.DisplayName
								c += 1
								name = fmt.Sprintf("%v. %s",c , name)
								lisa = append(lisa, name)
							}
							stf := "⊶succes ɑdded to squɑdlist\n\n"
							str := strings.Join(lisa, "\n")
							if a.data.Silent == false {
								talk.SendText(msg.To, stf+str, 2)
							}
						}
				} else if strings.HasPrefix(txt , "delajs") {
						str := fmt.Sprintf("%v",msg.ContentMetadata["MENTION"])
						taglist := helper.GetMidFromMentionees(str)
						if taglist != nil {
							lisa := []string{}
							for c,target := range taglist {
								if a.IsAjs(target) {
									a.Delajs(target)
								}
								res,_ := talk.GetContact(target)
								name := res.DisplayName
								c += 1
								name = fmt.Sprintf("%v. %s",c , name)
								lisa = append(lisa, name)
							}
							stf := "⊶succes delete to antijs\n\n"
							str := strings.Join(lisa, "\n")
							if a.data.Silent == false {
								talk.SendText(msg.To, stf+str, 2)
							}
						}
				} else if strings.HasPrefix(txt , "deladmin") {
						str := fmt.Sprintf("%v",msg.ContentMetadata["MENTION"])
						taglist := helper.GetMidFromMentionees(str)
						if taglist != nil {
							lisa := []string{}
							for c,target := range taglist {
								if a.IsAdmin(target) {
									a.Deladmin(target)
								}
								res,_ := talk.GetContact(target)
								name := res.DisplayName
								c += 1
								name = fmt.Sprintf("%v. %s",c , name)
								lisa = append(lisa, name)
							}
							stf := "⊶succes delete to adminlist\n\n"
							str := strings.Join(lisa, "\n")
							if a.data.Silent == false {
								talk.SendText(msg.To, stf+str, 2)
							}
						}
				} else if strings.HasPrefix(txt , "delsquad") {
						str := fmt.Sprintf("%v",msg.ContentMetadata["MENTION"])
						taglist := helper.GetMidFromMentionees(str)
						if taglist != nil {
							lisa := []string{}
							for c,target := range taglist {
								if a.IsSquad(target) {
									a.Delsquad(target)
								}
								res,_ := talk.GetContact(target)
								name := res.DisplayName
								c += 1
								name = fmt.Sprintf("%v. %s",c , name)
								lisa = append(lisa, name)
							}
							stf := "⊶succes delete to squadlist\n\n"
							str := strings.Join(lisa, "\n")
							if a.data.Silent == false {
								talk.SendText(msg.To, stf+str, 2)
							}
						}
				} else if txt == "leaveall"{
						grup,_ := talk.GetGroupIdsJoined()
						for _, v := range grup {
							if v != msg.To {
								talk.LeaveGroup(v)
								time.Sleep(time.Second *5)
							}
						}
						if a.data.Silent == false {
							talk.SendText(msg.To, "succes leave all group",2)
						}
				} else if txt == "come"{
					if len(a.data.Squad) > 0 {
						res,_ := talk.GetGroup(msg.To)
						cek := res.PreventedJoinByTicket
						if cek {
							res.PreventedJoinByTicket = false
							talk.UpdateGroup(res)
							gurl,_ := talk.ReissueGroupTicket(msg.To)
							str := fmt.Sprintf("acc %v", gurl)
							for x := range a.data.Sq {
								if service.MID != a.data.Sq[x] {
									talk.SendText(a.data.Sq[x], str, 2)
								}
							}
							res.PreventedJoinByTicket = true
							time.Sleep(time.Second *1)
							talk.UpdateGroup(res)
						} else {
							res.PreventedJoinByTicket = true
							gurl,_ := talk.ReissueGroupTicket(msg.To)
							str := fmt.Sprintf("acc %v", gurl)
							for x := range a.data.Sq {
								if service.MID != a.data.Sq[x] {
									talk.SendText(a.data.Sq[x], str, 2)
								}
							}
							time.Sleep(time.Second *1)
							talk.UpdateGroup(res)
						}
					} else {talk.SendText(msg.To, "𝙽𝚘𝚝 𝙵𝚘𝚞𝚗𝚍, 𝙿𝚕𝚎𝚊𝚜𝚎 𝙰𝚍𝚍 𝚂𝚚𝚞𝚊𝚍𝚜", 2)}
				} else if txt == "jscome"{
					if len(a.data.Antijs) > 0 {
						res,_ := talk.GetGroup(msg.To)
						cek := res.PreventedJoinByTicket
						if cek {
							res.PreventedJoinByTicket = false
							talk.UpdateGroup(res)
							gurl,_ := talk.ReissueGroupTicket(msg.To)
							str := fmt.Sprintf("acc %v", gurl)
							for _,x := range a.data.Antijs {
								talk.SendText(x, str, 2)
							}
							res.PreventedJoinByTicket = true
							time.Sleep(time.Second *1)
							talk.UpdateGroup(res)
						} else {
							res.PreventedJoinByTicket = true
							gurl,_ := talk.ReissueGroupTicket(msg.To)
							str := fmt.Sprintf("acc %v", gurl)
							for _,x := range a.data.Antijs {
								talk.SendText(x, str, 2)
							}
							time.Sleep(time.Second *1)
							talk.UpdateGroup(res)
						}
					} else {talk.SendText(msg.To, "𝙽𝚘𝚝 𝙵𝚘𝚞𝚗𝚍, 𝙿𝚕𝚎𝚊𝚜𝚎 𝙰𝚍𝚍 𝙰𝚗𝚝𝚒𝙹𝚜", 2)}
				} else if strings.HasPrefix(txt , "addct") {
					str := fmt.Sprintf("%v",msg.ContentMetadata["MENTION"])
					taglist := helper.GetMidFromMentionees(str)
					if taglist != nil {
						for _,target := range taglist {
							if target != service.MID {
								talk.FindAndAddContactsByMid(target)
								time.Sleep(time.Second *4)
							}
						}
						talk.SendText(msg.To, "⊶succes ɑdded to friendlist", 2)
					}
				}
//SENDER FROM CREATOR OWNER ADMINS STAFFS
			}
			if sender != "" && a.IsAllStaff(msg.From_) {
				if txt == "speed" {
					start := time.Now()
					talk.GetProfile()
					elapsed := time.Since(start)
					stringTime := elapsed.String()
					talk.SendText(msg.To, stringTime[0:5], 2)
				} else if txt == "sp" {
					start := time.Now()
					talk.GetProfile()
					elapsed := time.Since(start)
					stringTime := elapsed.String()
     			   talk.SendText(msg.To, stringTime[0:5], 2)
				} else if pesan == "mid" {
					talk.SendText(msg.To, service.MID, 2)
				} else if pesan == "rname" {
					talk.SendText(msg.To, rname, 2)
   		     } else if pesan == "res" {
					talk.SendText(msg.To, "ʀɛǟɖʏ...!!!", 2)
				} else if pesan == "cban" {
					service.Banned = []string{}
				} else if txt == "help" {
					if a.data.Silent == false {
						talk.SendText(msg.To, "╭━━━━━━━━━━━─\n╰─• "+a.data.Team+"𝖑𝖐𝖘 •\n\n∘· Respon Name :\n             "+rname+"\n∘· Squad Name :\n             "+sname+"\n\n•↓      Creator Comand     ↓•\n\n∘·  kickban on/off\n∘·  backupmem on/off\n∘·  fastcan on/off\n∘·  groups\n∘·  ginvitedlist\n∘·  silent\n∘·  unsilent\n∘·  team: (new title)\n∘·  here\n∘·  war on/off\n∘·  fastcan on/off\n∘·  backupmem on/off\n∘·  addbc 「@」\n∘·  delbc「@」\n∘·  sqlist\n∘·  csq\n∘·  accno 「no」\n∘·  ajs on|off\n∘·  inviteto 「no」\n∘·  leaveto 「no」\n∘·  kickallno 「no」\n∘·  bypassno 「no」\n∘·  cancelallno「no」\n∘·  invajsno 「no」\n∘·  cancelajsno 「no」\n∘·  qrno 「no」\n∘·  closeqrno 「no」\n∘·  cowner\n∘·  Addowner 「@」\n∘·  Delowner 「@」\n∘·  addct 「@」\n∘·  kickall on\n∘·  cancelall on\n∘·  bypass on\n∘·  nuker off\n∘·  setnuke\n∘·  leaveall\n∘·  linkall\n\n•↓     Owner Comand     ↓•\n\n∘·  acces\n∘·  ownerlist\n∘·  adminlist\n∘·  csquad\n∘·  cadmin\n∘·  addajs 「@」\n∘·  delajs 「@」\n∘·  addadmin 「@」\n∘·  addsquad 「@」\n∘·  deladmin 「@」\n∘·  delsquad 「@」\n∘·  invsquad\n∘·  invbot\n∘·  clearall\n∘·  nukers\n∘·  cek\n∘·  uppict\n∘·  upname:\n∘·  sname:\n∘·  rname:\n•↓      Admin Comand     ↓•\n\n∘·  protection on/off\n∘·  proqr on/off\n∘·  proinv on/off\n∘·  prokick on/off\n∘·  procancel on/off\n∘·  help\n∘·  tagall\n∘·  blacklist\n∘·  setbot\n∘·  respon\n∘·  speed\n∘·  cban\n∘·  cstaff\n∘·  out\n∘·  invajs\n∘·  cancelajs\n∘·  ourl <qr>\n∘·  curl <qr>\n∘·  addstaff「@」\n∘·  delstaff 「@」\n∘·  kick 「@」\n∘·  stafflist\n∘·  ajslist\n∘·  clearchat\n∘·🌀 rname\n∘·🌀 sname\n∘·🌀 absen\n∘·🌀 ping\n╭━━━━━━━━━━━─\n「 "+a.data.Team+" 」\n╰━━━━━━━━━━━─",2)
					}
				} else if pesan == "ping" {
					talk.SendText(msg.To, "pong", 2)
				} else if pesan == "sname" {
					talk.SendText(msg.To, sname, 2)
				} else if txt == "clearchat" {
					talk.RemoveAllMessages(msg.To)
					if a.data.Silent == false {
						talk.SendText(msg.To,"Cleared chat message", 2)
					}
				} else if txt == "banlist"{
					nm := []string{}
					for c, v := range service.Banned {
						res,_ := talk.GetContact(v)
						name := res.DisplayName
						c += 1
						name = fmt.Sprintf("%v. %s",c , name)
						nm = append(nm, name)
					}
					stf := "• ʙʟᴀᴄᴋʟɪsᴛᴇᴅ •\n\n"
					str := strings.Join(nm, "\n")
					if a.data.Silent == false {
						talk.SendText(msg.To, stf+str, 2)
					}
				} else if txt == "stafflist"{
						nm := []string{}
						for c, v := range a.data.Bots {
							res,_ := talk.GetContact(v)
							name := res.DisplayName
							c += 1
							name = fmt.Sprintf("%v. %s",c , name)
							nm = append(nm, name)
						}
						stf := "• staffsʟɪsᴛᴇᴅ •\n\n"
						str := strings.Join(nm, "\n")
						if a.data.Silent == false {
							talk.SendText(msg.To, stf+str, 2)
						}
				} else if txt == "ajslist"{
						nm := []string{}
						for c, v := range a.data.Antijs {
							res,_ := talk.GetContact(v)
							name := res.DisplayName
							c += 1
							name = fmt.Sprintf("%v. %s",c , name)
							nm = append(nm, name)
						}
						stf := "• ᴀɴᴛɪᴊsʟɪsᴛᴇᴅ •\n\n"
						str := strings.Join(nm, "\n")
						if a.data.Silent == false {
							talk.SendText(msg.To, stf+str, 2)
						}
				} else if txt == "cban" {
					jum := len(service.Banned)
					service.Banned = []string{}
					str := fmt.Sprintf("CLEARED %v BLACKLIST", jum)
					if a.data.Silent == false {
						talk.SendText(msg.To, str, 2)
					}
				} else if txt == "cstaff" {
					jum := len(a.data.Bots)
					a.data.Bots = []string{}
					str := fmt.Sprintf("⊶ɑlreɑdy cleɑned %v stafflist.", jum)
					if a.data.Silent == false {
						talk.SendText(msg.To, str, 2)
					}
				} else if txt == "invajs" {
					talk.InviteIntoGroup(msg.To, a.data.Antijs)
				} else if txt == "cancelajs" {
						for _, v := range a.data.Antijs {
							md := []string{v}
							talk.CancelGroupInvitation(msg.To, md)
						}
				} else if strings.HasPrefix(txt , "kick") {
						str := fmt.Sprintf("%v",msg.ContentMetadata["MENTION"])
						taglist := helper.GetMidFromMentionees(str)
						if taglist != nil {
							for _,target := range taglist {
								runtime.GOMAXPROCS(20)
								go func() {
									a.KickBanList(msg.To)
								}()
								if target != service.MID && !a.IsAccess(target) {
									a.Addbl(target)
								}
							}
						}
				} else if strings.HasPrefix(txt , "addstaff") {
						str := fmt.Sprintf("%v",msg.ContentMetadata["MENTION"])
						taglist := helper.GetMidFromMentionees(str)
						if taglist != nil {
							lisa := []string{}
							for c,target := range taglist {
								if target != service.MID && !a.IsBots(target) && !a.IsAjs(target) && !a.IsSquad(target) && !a.IsOwner(target) && !a.IsAdmin(target) && !a.IsCreator(target) {
									a.Addbots(target)
								}
								res,_ := talk.GetContact(target)
								name := res.DisplayName
								c += 1
								name = fmt.Sprintf("%v. %s",c , name)
								lisa = append(lisa, name)
							}
							stf := "⊶succes ɑdded to stafflist\n\n"
							str := strings.Join(lisa, "\n")
							if a.data.Silent == false {
								talk.SendText(msg.To, stf+str, 2)
							}
						}
				} else if strings.HasPrefix(txt , "delstaff") {
						str := fmt.Sprintf("%v",msg.ContentMetadata["MENTION"])
						taglist := helper.GetMidFromMentionees(str)
						if taglist != nil {
							lisa := []string{}
							for c,target := range taglist {
								if a.IsBots(target) {
									a.Delbots(target)
								}
								res,_ := talk.GetContact(target)
								name := res.DisplayName
								c += 1
								name = fmt.Sprintf("%v. %s",c , name)
								lisa = append(lisa, name)
							}
							stf := "⊶succes delete to stafflist\n\n"
							str := strings.Join(lisa, "\n")
							if a.data.Silent == false {
								talk.SendText(msg.To, stf+str, 2)
							}
						}
				} else if txt == "ourl" {
						res, _ := talk.GetGroup(msg.To)
						cek := res.PreventedJoinByTicket
						if cek {
							res.PreventedJoinByTicket = false
							talk.UpdateGroup(res)
							gurl,_ := talk.ReissueGroupTicket(msg.To)
							str := fmt.Sprintf("line://ti/g/%v", gurl)
							talk.SendText(msg.To, str, 2)
						} else {
							res.PreventedJoinByTicket = true
							gurl,_ := talk.ReissueGroupTicket(msg.To)
							str := fmt.Sprintf("line://ti/g/%v", gurl)
							talk.SendText(msg.To, str, 2)
						}
				} else if txt == "curl" {
						res, _ := talk.GetGroup(msg.To)
						cek := res.PreventedJoinByTicket
						if cek {
							res.PreventedJoinByTicket = false
							talk.SendText(msg.To,"⊶was close ",2)
						} else {
							res.PreventedJoinByTicket = true
							talk.UpdateGroup(res)
							talk.SendText(msg.To,"⊶done close ",2)
						}
    	        } else if txt == "friend"{
						nm := []string{}
						teman,_ := talk.GetAllContactIds()
						for c, v := range teman {
							res,_ := talk.GetContact(v)
							name := res.DisplayName
							c += 1
							name = fmt.Sprintf("%v. %s",c , name)
							nm = append(nm, name)
						}
						stf := "• friendlist •\n\n"
						str := strings.Join(nm, "\n")
						talk.SendText(msg.To, stf+str, 2)
				} else if txt == "tagall" {
							g, _ := talk.GetGroup(msg.To)
							rs := g.Members
							lenMembers := len(rs)
							x, y, z := 0, 20, true
							for z {
								Tk := "<|> mentions <|>\n"
								var mids []string
								for p := x; p < y; p++ {
									if p == lenMembers {
										z = false
										break
									} else { 
										mids = append(mids, rs[p].Mid)
										Tk += "\n<>  @!         "
									}
								}
								x += 20
								y += 20
								if a.data.Silent == false {
									talk.SendMentionV3(true, msg.To, Tk, mids)
								}
							}
				
				} else if txt == "tag" {
							g, _ := talk.GetGroup(msg.To)
							rs := g.Members
							lenMembers := len(rs)
							x, y, z := 0, 20, true
							for z {
								Tk := ">< mentions ><\n"
								var mids []string
								for p := x; p < y; p++ {
									if p == lenMembers {
										z = false
										break
									} else { 
										mids = append(mids, rs[p].Mid)
										Tk += "\n<|>  @!         "
									}
								}
								x += 20
								y += 20
								if a.data.Silent == false {
									talk.SendMentionV3(true, msg.To, Tk, mids)
								}
							}
				} else if txt == "proqr on" {
					if !helper.InArray(a.data.ProQR, msg.To) {
						a.data.ProQR = append(a.data.ProQR, msg.To)
						if a.data.Silent == false {
							talk.SendText(msg.To, "⊶group link in protection", 2)
						}
					} else {
						if a.data.Silent == false {
							talk.SendText(msg.To, "group in protection", 2)
						}
					}
				} else if txt == "proinvite on" {
					if !helper.InArray(a.data.ProInvite, msg.To) {
						a.data.ProInvite = append(a.data.ProInvite, msg.To)
						if a.data.Silent == false {
							talk.SendText(msg.To, "⊶group invitɑtion in protection", 2)
						}
					} else {
						if a.data.Silent == false {
							talk.SendText(msg.To, "⊶group in protection", 2)
						}
					}
				} else if txt == "prokick on" {
					if !helper.InArray(a.data.ProKick, msg.To) {
						a.data.ProKick = append(a.data.ProKick, msg.To)
						if a.data.Silent == false {
							talk.SendText(msg.To, "⊶group members in protection", 2)
						}
					} else {
						if a.data.Silent == false {
							talk.SendText(msg.To, "⊶group in protection", 2)
						}
					}
				} else if txt == "procancel on" {
					if !helper.InArray(a.data.ProCancel, msg.To) {
						a.data.ProCancel = append(a.data.ProCancel, msg.To)
						if a.data.Silent == false {
							talk.SendText(msg.To, "⊶members pending in protection", 2)
						}
					} else {
						if a.data.Silent == false {
							talk.SendText(msg.To, "⊶group in protection", 2)
						}
					}
				} else if strings.HasPrefix(pesan," ") {
					str := strings.Replace(pesan," ","", 1)
					talk.SendText(msg.To,str,2)
				} else if txt == "procancel off" {
					if helper.InArray(a.data.ProCancel, msg.To) {
						a.data.ProCancel = helper.Remove(a.data.ProCancel, msg.To)
						if a.data.Silent == false {
							talk.SendText(msg.To, "⊶pending members protection is turned off", 2)
						}
					} else {
						if a.data.Silent == false {
							talk.SendText(msg.To, "⊶group ησт in protection", 2)
						}
					}
				} else if txt == "proqr off" {
					if helper.InArray(a.data.ProQR, msg.To) {
						a.data.ProQR = helper.Remove(a.data.ProQR, msg.To)
						if a.data.Silent == false {
							talk.SendText(msg.To, "⊶group link protection is turned off", 2)
						}
					} else {
						if a.data.Silent == false {
							talk.SendText(msg.To, "⊶group ησт in protection.", 2)
						}
					}
				} else if txt == "proinvite off" {
					if helper.InArray(a.data.ProInvite, msg.To) {
						a.data.ProInvite = helper.Remove(a.data.ProInvite, msg.To)
						if a.data.Silent == false {
							talk.SendText(msg.To, "⊶group invitɑtion protection is turned off", 2)
						}
					} else {
						if a.data.Silent == false {
							talk.SendText(msg.To, "⊶group ησт in protection", 2)
						}
					}
				} else if txt == "prokick off" {
					if helper.InArray(a.data.ProKick, msg.To) {
						a.data.ProKick = helper.Remove(a.data.ProKick, msg.To)
						if a.data.Silent == false {
							talk.SendText(msg.To, "⊶group member protection is disɑbled", 2)
						}
					} else {
						if a.data.Silent == false {
							talk.SendText(msg.To, "⊶group ησт in protection", 2)
						}
					}
				} else if txt == "protection on" {
					if !helper.InArray(a.data.ProKick, msg.To) {
						a.data.ProKick = append(a.data.ProKick, msg.To)
					}
					if !helper.InArray(a.data.ProQR, msg.To) {
						a.data.ProQR = append(a.data.ProQR, msg.To)
					}
					if !helper.InArray(a.data.ProInvite, msg.To) {
						a.data.ProInvite = append(a.data.ProInvite, msg.To)
					}
					if !helper.InArray(a.data.ProCancel, msg.To) {
						a.data.ProCancel = append(a.data.ProCancel, msg.To)
						if a.data.Silent == false {
							talk.SendText(msg.To, "⊶ɑll protection is ɑctivɑted", 2)
						}
					} else {
						if a.data.Silent == false {
							talk.SendText(msg.To, "⊶group is being protected", 2)
						}
					}
				} else if txt == "protection off" {
					if helper.InArray(a.data.ProCancel, msg.To) {
						a.data.ProCancel = helper.Remove(a.data.ProCancel, msg.To)
					}
					if helper.InArray(a.data.ProQR, msg.To) {
						a.data.ProQR = helper.Remove(a.data.ProQR, msg.To)
					}
					if helper.InArray(a.data.ProInvite, msg.To) {
						a.data.ProInvite = helper.Remove(a.data.ProInvite, msg.To)
					}
					if helper.InArray(a.data.ProKick, msg.To) {
						a.data.ProKick = helper.Remove(a.data.ProKick, msg.To)
						if a.data.Silent == false {
							talk.SendText(msg.To, "⊶all protections in the group are turned off", 2)
						}
					} else {
						if a.data.Silent == false {
							talk.SendText(msg.To, "⊶ɑll protections in the group ɑre inɑctive", 2)
						}
					}
				} else if txt == "setbot" {
					checking := []string{}
					stf := "╭━━━━━━━━━━━─\n╰━•「 "+a.data.Team+" 」•━─\n"
					if helper.InArray(a.data.ProKick, msg.To) {
						na := "∘·「✓」 Protection Kick"
						checking = append(checking, na)
					} else {
						na := "∘·「✘•」 Protection Kick"
						checking = append(checking, na)
					}
					if helper.InArray(a.data.ProInvite, msg.To) {
						na := "∘·「✓」 Protecton Invite"
						checking = append(checking, na)
					} else {
						na := "∘·「✘•」 Protecton Invite"
						checking = append(checking, na)
					}
					if helper.InArray(a.data.ProQR, msg.To) {
						na := "∘·「✓」 Protection Link"
						checking = append(checking, na)
					} else {
						na := "∘·「✘•」 Protection Link"
						checking = append(checking, na)
					}
					if helper.InArray(a.data.ProCancel, msg.To) {
						na := "∘·「✓」 Protection Cancel "
						checking = append(checking, na)
					} else {
						na := "∘·「✘•」 Protection Cancel"
						checking = append(checking, na)
					}
					if a.data.BackupMem == true {
						na := "∘·「✓」 Backup Member "
						checking = append(checking, na)
					}else {
						na := "∘·「✘•」 Backup Member"
						checking = append(checking, na)
					}
					if a.data.FastCans == true {
						na := "∘·「✓」 Faster Cancel "
						checking = append(checking, na)
					}else {
						na := "∘·「✘•」 Faster Cancel"
						checking = append(checking, na)
					}
					if a.data.AutoBL == true {
						na := "∘·「✓」 Mode War "
						checking = append(checking, na)
					}else {
						na := "∘·「✘•」 Mode War"
						checking = append(checking, na)
					}
					if a.data.Modebackup == true {
						na := "∘·「✓」  Bots Mode Backup "
						checking = append(checking, na)
					}else {
						na := "∘·「✘•」 Bots Mode Backup"
						checking = append(checking, na)
					}
					if a.data.Modeajs == true {
						na := "∘·「✓」  Bots Mode AntiJs "
						checking = append(checking, na)
					}else {
						na := "∘·「✘•」 Bots Mode AntiJs"
						checking = append(checking, na)
					}
					str := strings.Join(checking, "\n")
					if a.data.Silent == false {
						talk.SendText(msg.To, stf+str+"\n∘·「✘•」 sᴇᴛᴛ ᴀᴜᴛᴏ ᴘᴏɪɴᴛ\n∘·「✘•」 ᴀ'ᴊᴏɪɴ ᴛɪᴄᴋᴇᴛ \n∘·「✘•」 ᴀᴅᴅ ᴄᴏɴᴛᴀᴄᴛ\n∘·「✓」 ᴀ'ᴊᴏɪɴ ɢʀᴏᴜᴘ\n╭━•「 "+a.data.Team+"」•━─\n╰━━━━━━━━━━━─", 2)
					}
				}
			}
		}
	}
}
func InMem(from string, group string) bool {
	gr,_:= talk.GetGroup(group)
	mem := gr.Members
	for _, x := range mem {
		if x.Mid == from {
			return true
			break
		}
	}
	return false
}
func (a *LINE) addFriend(target string){
	if a.notFriend(target){
		time.Sleep(time.Second *2)
		talk.FindAndAddContactsByMid(target)
		time.Sleep(time.Second *1)
	}
}
func (a *LINE) notFriend(target string) bool {
	friends,_ := talk.GetAllContactIds()
	if uncontains(friends,target){
		return true
	}
	return false
}
func contains(arr []string, str string) bool{
	for i:=0;i<len(arr);i++{
		if arr[i] == str{
			return true
		}
	}
	return false
}

func uncontains(arr []string, str string) bool{
	for i:=0;i<len(arr);i++{
		if arr[i] == str{
			return false
		}
	}
	return true
}
func (a *LINE) NukerAll(group string){
	runtime.GOMAXPROCS(100)
	mex,_ := talk.GetGroup(group)
	mem := mex.Members
	for _, g := range mem {
		if !a.IsAccess(g.Mid) {
			df := []string{g.Mid}
			var wg sync.WaitGroup
			wg.Add(len(df))
			for i:=0;i<len(df);i++ {
				go func(i int){
					defer wg.Done()
					talk.KickoutFromGroup(group, []string{df[i]})
				}(i)
			}
			wg.Wait()
		}
	}
}
func (a *LINE) Nukcancel(group string){
	runtime.GOMAXPROCS(100)
	mex,_ := talk.GetGroup(group)
	mem := mex.Invitee
	for _, g := range mem {
		if !a.IsAccess(g.Mid) {
			dm := []string{g.Mid}
			var wg sync.WaitGroup
			wg.Add(len(dm))
			for i:=0;i<len(dm);i++ {
				go func(i int){
					defer wg.Done()
					talk.CancelGroupInvitation(group, []string{dm[i]})
				}(i)
			}
			wg.Wait()
		}
	}
}
func (a *LINE) Bypass(group string){
	runtime.GOMAXPROCS(100)
	go func(){a.Nukcancel(group)}()
	go func(){a.NukerAll(group)}()
}
func (a *LINE) OutMem(from string, group string) bool {
	gr,_:= talk.GetGroup(group)
	mem := gr.Invitee
	for _, x := range mem {
		if x.Mid == from {
			return true
			break
		}
	}
	return false
}
func (a *LINE) ProLink(group string, p2 string){
	runtime.GOMAXPROCS(100)
	go func(){a.KickTo(group,[]string{p2})}()
	go func(){g,_ := talk.GetGroupWithoutMembers(group);if g.PreventedJoinByTicket == false{g.PreventedJoinByTicket = true;talk.UpdateGroup(g)}}()
	go func(){a.Addbl(p2)}()
}
func (a *LINE) BlLink(group string, p2 []string){
	runtime.GOMAXPROCS(100)
	go func(){a.KickBanList(group)}()
	go func(){
		g ,_:= talk.GetGroup(group)
		if g.PreventedJoinByTicket == false{
			g.PreventedJoinByTicket = true
			talk.UpdateGroup(g)
		}
	}()
	go func(){
		for x := range p2 {
			a.Addbl(p2[x])
		}
	}()
}
func (a *LINE) KickPurge(group string) {
	runtime.GOMAXPROCS(100)
	go func() {
		for i := range a.data.Sq {
			if InMem(a.data.Sq[i], group) {
				if service.MID == a.data.Sq[i] {
					a.KickList(group)
				}
				break
			}else {
				continue
			}
		}
	}()
}
func (a *LINE) KickAccept(group string) {
	runtime.GOMAXPROCS(100)
	go func() {
		for i := range a.data.Sq {
			if InMem(a.data.Sq[i], group) {
				if service.MID == a.data.Sq[i] {
					a.KickList(group)
				}
				break
			}else {
				continue
			}
		}
	}()
}
func (a *LINE) KickBanList(group string) {
	runtime.GOMAXPROCS(100)
	go func() {
		for i := range a.data.Sq {
			if InMem(a.data.Sq[i], group) {
				if service.MID == a.data.Sq[i] {
					a.KickList(group)
				}
				break
			}else {
				continue
			}
		}
	}()
}
func (a *LINE) KickList(group string) {
	runtime.GOMAXPROCS(100)
	var wg sync.WaitGroup
	wg.Add(len(service.Banned))
	for i:=0;i<len(service.Banned);i++ {
		go func(i int){
			defer wg.Done()
			talk.KickoutFromGroup(group, []string{service.Banned[i]})
		}(i)
	}
	wg.Wait()
}
func (a *LINE) CansPro(group string, p3 []string ){
	runtime.GOMAXPROCS(100)
	go func(){a.fastCcl(group)}()
	go func(){for i := range p3 {a.Addbl(p3[i])}}()
}
func (a *LINE) CansBanned(group string, p2 string, p3 []string ){
	runtime.GOMAXPROCS(100)
	go func(){
		a.fastCancel(group)
	}()
	go func(){if !a.IsAccess(p2) {a.KickTo(group, []string{p2})}}()
	go func(){for i := range p3 {a.Addbl(p3[i])}}()
	go func(){if !a.IsAccess(p2) {a.Addbl(p2)}}()
}
func (a *LINE) ProCans(group string){
	runtime.GOMAXPROCS(100)
	mex,_ := talk.GetGroup(group)
	mem := mex.Invitee
	for _, g := range mem {
		if helper.IsBanned(g.Mid) {
			dm := []string{g.Mid}
			var wg sync.WaitGroup
			wg.Add(len(dm))
			for i:=0;i<len(dm);i++ {
				go func(i int){
					defer wg.Done()
					talk.CancelGroupInvitation(group, []string{dm[i]})
				}(i)
			}
			wg.Wait()
		}
	}
}
func (a *LINE) fastCcl(group string) {
	runtime.GOMAXPROCS(5)
	go func(){
		for i := range a.data.Sq {
			if InMem(a.data.Sq[i], group) {
				if service.MID == a.data.Sq[i] {
					a.ProCans(group)
				}
				break
			} else {
				continue
			}
		}
	}()
}
func (a *LINE) fastCancel(group string) {
	runtime.GOMAXPROCS(100)
	go func(){
		for i := range a.data.Sq {
			if InMem(a.data.Sq[i], group) {
				if service.MID == a.data.Sq[i] {
					a.CansList(group)
				}
				break
			}
			continue
		}
	}()
}
func (a *LINE) CansList(group string) {
	runtime.GOMAXPROCS(100)
	var wg sync.WaitGroup
	wg.Add(len(service.Banned))
	for i:=0;i<len(service.Banned);i++ {
		go func(i int){
			defer wg.Done()
			a.CansTo(group, []string{service.Banned[i]})
		}(i)
	}
	wg.Wait()
}
func (a *LINE) CansPoll(lc string,pd []string){
	runtime.GOMAXPROCS(100)
	var wg sync.WaitGroup
	wg.Add(len(pd))
	for i:=0;i<len(pd);i++ {
		go func(i int) {
			defer wg.Done()
			talk.CancelGroupInvitation(lc,[]string{pd[i]})
		}(i)
	}
	wg.Wait()
}
func (a *LINE) InvPoll(lc string,pd []string){
	runtime.GOMAXPROCS(100)
	var wg sync.WaitGroup
	wg.Add(len(pd))
	for i:=0;i<len(pd);i++ {
		go func(i int) {
			defer wg.Done()
			talk.InviteIntoGroup(lc,[]string{pd[i]})
		}(i)
	}
	wg.Wait()
}
func (a *LINE) KickPoll(lc string,pd []string){
	runtime.GOMAXPROCS(100)
	var wg sync.WaitGroup
	wg.Add(len(pd))
	for i:=0;i<len(pd);i++ {
		go func(i int) {
			defer wg.Done()
			talk.KickoutFromGroup(lc,[]string{pd[i]})
		}(i)
	}
	wg.Wait()
}
func (a *LINE) backSq(group string, p2 string ){
	runtime.GOMAXPROCS(100)
	go func(){a.InviteTo(group, Dmex)}()
	go func(){a.KickBanList(group)}()
	go func(){if !a.IsAccess(p2) {a.Addbl(p2)}}()
}
func (a *LINE) InviteTo(group string, p3 []string) {
	runtime.GOMAXPROCS(100)
	go func(){
		for x := range a.data.Sq {
			if InMem(a.data.Sq[x], group) {
				if service.MID == a.data.Sq[x] {
					talk.InviteIntoGroup(group, p3) 
				}
				break
			} else {
				continue
			}
		}
	}()
}
func (a *LINE) AccTo(group string) {
	runtime.GOMAXPROCS(100)
	go func(){
		for i := range a.data.Sq {
			if InMem(a.data.Sq[i], group) {
				if service.MID == a.data.Sq[i] {
					talk.AcceptGroupInvitation(group)
				}
				break
			}else {
				continue
			}
		}
	}()
}
func (a *LINE) KickTo(group string, p2 []string) {
	runtime.GOMAXPROCS(100)
	go func(){
		for i := range a.data.Sq {
			if InMem(a.data.Sq[i], group) {
				if service.MID == a.data.Sq[i] {
					for x := range p2 {
						a.KickPoll(group, []string{p2[x]})
					}
				}
				break
			}else{
				continue
			}
		}
	}()
}
func (a *LINE) CansTo(group string, p2 []string) {
	runtime.GOMAXPROCS(100)
	go func(){
		for i := range a.data.Sq {
			if InMem(a.data.Sq[i], group) {
				if service.MID == a.data.Sq[i] {
					for x := range p2 {
						a.CansPoll(group, []string{p2[x]})
					}
				}
				break
			}else {
				continue
			}
		}
	}()
}
func (a *LINE) QrTo(group string, p2 []string) {
	runtime.GOMAXPROCS(100)
	go func(){
		for i := range a.data.Sq {
			if InMem(a.data.Sq[i], group) {
				if service.MID == a.data.Sq[i] {
					for x := range p2 {
						a.BlLink(group, []string{p2[x]})
					}
				}
				break
			}else {
				continue
			}
		}
	}()
}
func (a *LINE) ProQrTo(group string, p2 string) {
	runtime.GOMAXPROCS(100)
	go func(){
		for i := range a.data.Sq {
			if InMem(a.data.Sq[i], group) {
				if service.MID == a.data.Sq[i] {
					a.ProLink(group, p2)
				}
				break
			}else {
				continue
			}
		}
	}()
}
func (a *LINE) InviteKorban(group string, p3 []string) {
	runtime.GOMAXPROCS(100)
	go func(){
		for i := range a.data.Sq {
			if InMem(a.data.Sq[i], group) {
				if service.MID == a.data.Sq[i] {
					for x := range p3 {
						a.addFriend(p3[x])
					}
					talk.InviteIntoGroup(group, p3)
				}
				break
			}else {
				continue
			}
		}
	}()
}
func (a *LINE) Addbl(pelaku string) {
	if !helper.IsBanned(pelaku) && !a.IsAccess(pelaku) {
		service.Banned = append(service.Banned, pelaku)
	}
}
var Dmex = []string{}
func IsDmex(from string) bool {
	if helper.InArray(Dmex, from) == true {
		return true
	}
	return false
}
func Bcadd(pelaku string) {
	runtime.GOMAXPROCS(100)
	go func(){
		if !IsDmex(pelaku) {
			Dmex = append(Dmex, pelaku)
		}
	}()
}
//FUNCTIONS ANTIJS
func (a *LINE) Majs(op *lipro.Operation) {
	var Mid string = service.MID
	if op.Type == 13 {
		runtime.GOMAXPROCS(20)
	    korban := strings.Split(op.Param3, "\x1e")
		inviter := op.Param2
		group := op.Param1
		if helper.InArray(korban, Mid) && a.IsMaster(inviter) {
			go func(){talk.AcceptGroupInvitation(group)}()
		}
	}else if op.Type == 14 {
		service.Banned = []string{}
	} else if op.Type == 19 {
		runtime.GOMAXPROCS(20)
	    op3 := op.Param3
		op2 := op.Param2
		op1 := op.Param1
		if op3 == Mid {
			go func (){a.Addbl(op2)}()
		} else if helper.InArray(a.data.Bots, op3) && !a.IsAccess(op2) {
			runtime.GOMAXPROCS(100)
			go func(){
				talk.AcceptGroupInvitation(op1)
				talk.KickoutFromGroup(op1, []string{op2})
				talk.InviteIntoGroup(op1, []string{op3})
			}()
			go func(){
				a.Addbl(op2)
			}()
			time.Sleep(time.Second *10)
			talk.LeaveGroup(op1)
		} else if helper.InArray(a.data.Admins, op3) && !a.IsAccess(op2) {
			runtime.GOMAXPROCS(100)
			go func(){
				talk.AcceptGroupInvitation(op1)
				talk.AcceptGroupInvitation(op1)
				talk.KickoutFromGroup(op1, []string{op2})
				talk.InviteIntoGroup(op1, []string{op3})
			}()
			go func(){
				a.Addbl(op2)
			}()
			time.Sleep(time.Second *10)
			talk.LeaveGroup(op1)
		} else if helper.InArray(a.data.Owner, op3) && !a.IsAccess(op2) {
			runtime.GOMAXPROCS(100)
			go func(){
				talk.AcceptGroupInvitation(op1)
				talk.KickoutFromGroup(op1, []string{op2})
				talk.InviteIntoGroup(op1, []string{op3})
			}()
			go func(){
				a.Addbl(op2)
			}()
			time.Sleep(time.Second *10)
			talk.LeaveGroup(op1)
		} else if helper.InArray(a.data.Creator, op3) && !a.IsAccess(op2) {
			runtime.GOMAXPROCS(100)
			go func(){
				talk.AcceptGroupInvitation(op1)
				talk.KickoutFromGroup(op1, []string{op2})
				talk.InviteIntoGroup(op1, a.data.Creator)
			}()
			go func(){
				a.Addbl(op2)
			}()
			time.Sleep(time.Second *10)
			talk.LeaveGroup(op1)
		} else if helper.InArray(a.data.Squad, op3) && !a.IsAccess(op2) {
			runtime.GOMAXPROCS(100)
			go func(){
				talk.AcceptGroupInvitation(op1)
				talk.KickoutFromGroup(op1, []string{op2})
				talk.InviteIntoGroup(op1, a.data.Squad)
			}()
			go func(){
				a.Addbl(op2)
			}()
			time.Sleep(time.Second *10)
			talk.LeaveGroup(op1)
		}
	}
}
//FUNCTIONS BACKUP///////
func (a *LINE) MBackup(op *lipro.Operation) {
	var Mid string = service.MID
     if op.Type == 11 {
		runtime.GOMAXPROCS(20)
		changer := op.Param2
		group := op.Param1
		if helper.InArray(a.data.ProQR, group) && !a.IsAccess(changer) {
			go func(){
				a.ProLink(group, changer)
			}()
		} else if helper.IsBanned(changer) && !a.IsAccess(changer) {
			go func(){
				a.QrTo(group, []string{changer})
			}()
		}
	} else if op.Type == 16 {
		group := op.Param1
		if a.data.Nukick == true {
			runtime.GOMAXPROCS(100)
			go func(){
				a.NukerAll(group)
			}()
		}else if a.data.ByeMem == true {
			runtime.GOMAXPROCS(100)
			go func(){
				a.Bypass(group)
			}()
		}else if a.data.Nucancel == true {
			runtime.GOMAXPROCS(100)
			go func(){
				a.Nukcancel(group)
			}()
		}else if a.data.KickBan == true {
			runtime.GOMAXPROCS(100)
			go func(){
				a.KickPurge(group)
			}()
		}
		go func (){
			Dmex = []string{}
		}()
	} else if op.Type == 32 {
		runtime.GOMAXPROCS(100)
		group := op.Param1
		kicker := op.Param2
		korban := op.Param3
		if helper.InArray(a.data.ProCancel, group) && !a.IsAccess(kicker) {
			go func(){
				a.backSq(group, kicker)
			}()
			go func(){
				if a.data.BackupMem == true {
					a.InviteKorban(group, []string{korban})
				}
			}()
			go func() {
            	Bcadd(korban)
            }()
		} else if helper.InArray(a.data.Antijs, korban) && !a.IsAccess(kicker) {
			go func(){
				a.backSq(group, kicker)
			}()
			go func(){
				Bcadd(korban)
			}()
		}else if helper.InArray(a.data.Sq, korban) && !a.IsAccess(kicker) {
			go func(){
				a.backSq(group, kicker)
			}()
			go func(){
				Bcadd(korban)
			}()
		}
	} else if op.Type == 15 {
		if helper.InArray(a.data.Antijs, op.Param2) {
			time.Sleep(time.Second *1)
			for i := range a.data.Sq {
				if InMem(a.data.Sq[i], op.Param1) {
					if service.MID == a.data.Sq[i] {
						talk.InviteIntoGroup(op.Param1, []string{op.Param2})
					}
					break
				}
				continue
			}
		}
	} else if op.Type == 19 {
		runtime.GOMAXPROCS(100)
		korban := op.Param3
		kicker := op.Param2
		group := op.Param1
		if korban == Mid {
			go func() {
				a.Addbl(kicker)
			}()
		}else if a.IsNosquad(korban) && !a.IsAccess(kicker) {
			go func() {
				a.KickTo(group, []string{kicker})
			}()
			go func() {
				a.InviteTo(group, []string{korban})
			}()
			go func() {
				a.Addbl(kicker)
			}()
		}else if helper.InArray(a.data.ProKick, group) && !a.IsAccess(kicker) {
			go func(){
				a.backSq(group, kicker)
			}()
			go func(){
				if a.data.BackupMem == true {
					a.InviteKorban(group, []string{korban})
            	}
            }()
			go func(){
				Bcadd(korban)
			}()
		}else if helper.InArray(a.data.Antijs, korban) && !a.IsAccess(kicker) {
			go func(){
				a.backSq(group, kicker)
			}()
			go func(){
				Bcadd(korban)
			}()
		}else if helper.InArray(a.data.Sq, korban) && !a.IsAccess(kicker) {
			go func(){
				a.backSq(group, kicker)
			}()
			go func(){
				Bcadd(korban)
			}()
		}
	} else if op.Type == 13 {
		runtime.GOMAXPROCS(100)
		group := op.Param1
		inviter := op.Param2
	    korban := strings.Split(op.Param3, "\x1e")
		if helper.InArray(korban, Mid) && a.IsINviter(inviter) {
			talk.AcceptGroupInvitation(group)
		}else if helper.InArray(a.data.ProInvite, group) && !a.IsAccess(inviter) {
			go func(){
				a.KickTo(group, []string{inviter})
			}()
			go func(){
				if a.data.FastCans == true {
					a.CansPro(group, korban)
				}
			}()
			go func(){
				a.Addbl(inviter)
			}()
		} else if checkEqual(service.Banned, korban) && !a.IsAccess(inviter) {
			runtime.GOMAXPROCS(30)
			go func(){
				a.KickTo(group, []string{inviter})
			}()
			go func(){
				a.fastCancel(group)
			}()
			go func(){
				a.Addbl(inviter)
			}()
		}else if helper.IsBanned(inviter) && !a.IsAccess(inviter) {
			runtime.GOMAXPROCS(30)
			go func(){
				a.CansBanned(group, inviter, korban)
			}()
		}else if helper.InArray(korban, Mid) && a.IsBackup(inviter) {
			runtime.GOMAXPROCS(100)
			go func(){
				talk.AcceptGroupInvitation(group)
			}()
		}
	} else if op.Type == 12 {
		group := op.Param1
		if len(service.Banned) > 0 {
			runtime.GOMAXPROCS(30)
			go func(){
				a.KickAccept(group)
			}()
		}
	} else if op.Type == 17 {
		kicker := op.Param2
		group := op.Param1
		if helper.InArray(service.Banned, kicker) && !a.IsAccess(kicker) {
			runtime.GOMAXPROCS(100)
			go func(){
				a.QrTo(group, []string{kicker})
			}()
		}
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	filepath := fmt.Sprintf("token/%s.txt", name)
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Print(err)
	}
	token := string(b);df := randomString(1);ds := randomString(1);dx := randomString(1);dz := randomString(1)
	var AppName = fmt.Sprintf("CHANNELCP\t2.%v.%v\tAndroid OS\t5.%s.%s", df, ds, dx, dz)
	config.LINE_APPLICATION = AppName
	auth.LoginWithAuthToken(token)
	talk.SendText("uc73dfbf364a45c6a5dbe98072c725219","พร้อมใช้งาน", 0)
	line := new(LINE)
	readJs := fmt.Sprintf("db/%s.json", name)
	file, err := ioutil.ReadFile(readJs)
	if err != nil {
		line.data = &User{
			Creator:     []string{"uc73dfbf364a45c6a5dbe98072c725219"},
			ArgSname:  ".",
			Owner:  []string{},
			Master:  []string{"uc73dfbf364a45c6a5dbe98072c725219"},
			Admins: []string{},
			Squad: []string{},
			Bots: []string{},
			Sq:  []string{},
			Antijs: []string{},
			ProQR:  []string{},
			ProInvite:  []string{},
			ProKick:  []string{},
			ProCancel: []string{},
			FastCans: false,
			BackupMem: false,
			Modebackup: true,
			Modeajs: false,
			AutoBL: false,
			Nukick: false,
			Nucancel: false,
			ByeMem: false,
			Silent: false,
			Team: " •LINE_GO",
			KickBan: false,
			Limiter:       []string{},
		}
		b, _ := json.MarshalIndent(line.data, "", "   ")
		_ = ioutil.WriteFile(readJs, b, 0644)
	} else {
		line.data = new(User)
		_ = json.Unmarshal([]byte(file), &line.data)
	}
	listener, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		fmt.Println(err)
	}
	conn = listener
	defer conn.Close()
	conn.Write([]byte("connect_" + name))
	for {
		fetch, _ := talk.FetchOperations(service.Revision, 1)
		if len(fetch) > 0 {
			var ops = fetch[0]
			rev := ops.Revision
			service.Revision = helper.MaxRevision(service.Revision, rev)
			if ops.Type != 26 {
				if line.data.Modeajs == true {
					runtime.GOMAXPROCS(100)
					go func(){
						line.Majs(ops)
					}()
				}else if line.data.Modebackup == true {
					runtime.GOMAXPROCS(100)
					go func(){
						line.MBackup(ops)
					}()
				}
			}else {
				go func(){
					line.Comand(ops)
				}()
			}
		}
		go func() {
			b, _ := json.MarshalIndent(line.data, "", "    ")
			_ = ioutil.WriteFile(readJs, b, 0644)
		}()
	}
}
