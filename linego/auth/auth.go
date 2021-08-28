package auth

import (
	"../service"
	"../config"
	"../LineThrift"
	"net/http"
	"fmt"
	"io/ioutil"
	"time"
	"encoding/json"
	"context"
	// "crypto/rsa"
	//"strconv"
)

type InsideResults struct {
	Verifier string `json:"verifier"`
	AuthPhase string `json:"authPhase"`
}
type Results struct {
	Result InsideResults `json:"result"`
	Timestamp string `json:"timestamp"`
}


func getResult(body []byte) (*Results, error) {
	var s = new(Results)
	err := json.Unmarshal(body, &s)
	return s, err
}

func LoadService() (*LineThrift.TalkServiceClient,*LineThrift.CallServiceClient,*LineThrift.SquareServiceClient){
	if service.IsLogin != true {
		panic("[Error]Not yet logged in.")
	}
	a:=time.Now()
	talk := service.TalkService()
	fmt.Println(time.Since(a))
	call := service.CallService()
	square := service.SquareService()
	rev, err := talk.GetLastOpRevision(context.TODO())
	if err != nil {
		panic(err)
	}
	getprof, err := talk.GetProfile(context.TODO())
	if err != nil {
		panic(err)
	}
	service.Revision = rev
	service.MID = getprof.Mid
	fmt.Println("Name: ", getprof.DisplayName)
	fmt.Println("MID: ", getprof.Mid)
	fmt.Println("\nGoLang Bot")
	return talk, call, square
}

func loginRequestQR(identity LineThrift.IdentityProvider, verifier string) *LineThrift.LoginRequest{
	lreq := &LineThrift.LoginRequest{
		Type: 1,
		KeepLoggedIn: true,
		IdentityProvider: identity,
		AccessLocation: config.IP_ADDR,
		SystemName: config.SYSTEM_NAME,
		Verifier: verifier,
		E2eeVersion: 0,
	}
	return lreq
}

// not yet working, still learning...
// func LoginWithCredential(email string, password string) {
// 	fmt.Println(email, password)
// 	tauth := service.AuthService()
// 	rsaKey, err := tauth.GetRSAKeyInfo(context.TODO(), 1)
// 	fmt.Println(rsaKey, err)
// 	rsaSessionKey := rsaKey.SessionKey
// 	rsaNvalue := strconv.FormatInt(int(rsaKey.Nvalue), 16)
// 	rsaEvalue := strconv.FormatInt(int(rsaKey.Evalue), 16)
// 	message := (string(len(rsaSessionKey)) + rsaSessionKey + string(len(email)) + email + string(len(password)) + password)
// 	pubKey := &rsa.PublicKey{
// 		N: rsaNvalue,
// 		E: rsaEvalue,
// 	}
// 	fmt.Println(message)

// }


func LoginWithQrCode(keepLoggedIn bool) (*LineThrift.TalkServiceClient,*LineThrift.CallServiceClient,*LineThrift.SquareServiceClient){
	tauth := service.AuthService()
	qrCode, err := tauth.GetAuthQrcode(context.TODO(), keepLoggedIn, config.SYSTEM_NAME, true)
	if err != nil{
		panic(err)
	}

	fmt.Println("Click this link qr code for login bot.\nline://au/q/"+qrCode.Verifier)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", config.LINE_HOST_DOMAIN+config.LINE_CERTIFICATE_PATH, nil)
	req.Header.Set("User-Agent",config.USER_AGENT)
	req.Header.Set("X-Line-Application",config.LINE_APPLICATION)
	req.Header.Set("X-Line-Carrier",config.CARRIER)
	req.Header.Set("X-Line-Access",qrCode.Verifier)
	service.AuthToken = qrCode.Verifier
	res, _ := client.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	s, _ := getResult([]byte(body))
	iR := s.Result
	_verifier := iR.Verifier

	loginZ := service.LoginZService()
	loginReq := loginRequestQR(1, _verifier)
	resultz, err := loginZ.LoginZ(context.TODO(),loginReq)
	if err != nil {
		panic(err)
	}
	service.IsLogin = true
	talk, call, square := LoginWithAuthToken(resultz.AuthToken)
	return talk, call, square
}

func LoginWithAuthToken(authToken string)(*LineThrift.TalkServiceClient,*LineThrift.CallServiceClient,*LineThrift.SquareServiceClient){
	service.AuthToken = authToken
	service.IsLogin = true
	talk, call, square := LoadService()
	return talk, call, square
}
