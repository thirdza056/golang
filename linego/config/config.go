package config

import (
	"../LineThrift"
)

/* server */

var LINE_HOST_DOMAIN = "https://legy-jp-addr-long.line.naver.jp"
var LINE_OBS_DOMAIN = "http://obs.line-apps.com"
var LINE_TIMELINE_API = "https://legy-jp-addr-long.line.naver.jp/mh/api"
var LINE_TIMELINE_MH = "https://legy-jp-addr-long.line.naver.jp/mh"

var LINE_LOGIN_QUERY_PATH = "/api/v4/TalkService.do"
var LINE_AUTH_QUERY_PATH = "/api/v4p/rs"

var LINE_API_QUERY_PATH_FIR = "/S4"
var LINE_POLL_QUERY_PATH_FIR = "/P4"
var LINE_POLL_QUERY_PATH_SEC = "/NP4"
var LINE_POLL_QUERY_PATH_THI = "/P4"
var LINE_CALL_QUERY_PATH = "/V4"
var LINE_CERTIFICATE_PATH = "/Q"
var LINE_CHAN_QUERY_PATH = "/CH4"
var LINE_SQUARE_QUERY_PATH = "/SQS1"
var LINE_SHOP_QUERY_PATH = "/SHOP4"

var APP_TYPE = func (ap LineThrift.ApplicationType) (string) { return ap.String() } (144)
var APP_VER = "\t2.9.1"
var CARRIER = "51089, 1-0"
var SYSTEM_NAME = "\tAndroid OS"
var SYSTEM_VER = "\t5.1.1"
var IP_ADDR = "127.0.0.1"
var LINE_APPLICATION string
var USER_AGENT = "LLA/2.9.1 SM-J320G 5.1.1"

var TIMELINE_CHANNEL_ID = "1341209950"
var WEBTOON_CHANNEL_ID = "1401600689"
var TODAY_CHANNEL_ID = "1518712866"
var STORE_CHANNEL_ID = "1376922440"
var MUSIC_CHANNEL_ID = "1381425814"
var SERVICES_CHANNEL_ID = "1459630796"

/* additional stuff */

var AutoCancel = false
