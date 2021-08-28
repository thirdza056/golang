package server

import (
	"../config"
	"net/http"
	"io"
)

var client = &http.Client{}

func getNewRequest(method string, url string, body io.Reader) (res *http.Response, err error) {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("User-Agent",config.USER_AGENT)
	req.Header.Add("X-Line-Application",config.LINE_APPLICATION)
	req.Header.Add("X-Line-Carrier",config.CARRIER)
	req.Header.Add("X-Line-Access",AuthToken)
	req.Header.Add("Connection","Keep-Alive")
	res, err := client.Do(req)
	return res, err
}

func PostContent(url string, body io.Reader) (res *http.Response, err error) {
	res, err := getNewRequest("POST", url, body)
	return res, err
}
func GetContent(url string, body io.Reader) (res *http.Response, err error) {
	res, err := getNewRequest("GET", url, body)
	return res, err
}
func DeleteContent(url string, body io.Reader) (res *http.Response, err error) {
	res, err := getNewRequest("DELETE", url, body)
	return res, err
}
func PutContent(url string, body io.Reader) (res *http.Response, err error) {
	res, err := getNewRequest("PUT", url, body)
	return res, err
}