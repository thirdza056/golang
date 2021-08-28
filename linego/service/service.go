package service

import (
	"../LineThrift"
	"../config"
	"../thrift"
	//"fmt"
)

var IsLogin bool = false
var AuthToken string = ""
var Revision int64 = 0
var MID string = ""
var Banned = []string{}

func TalkService() *LineThrift.TalkServiceClient {
	//fmt.Println("#### TalkService Initiated. ####")
	httpClient, _ := thrift.NewTHttpClient(config.LINE_HOST_DOMAIN+config.LINE_API_QUERY_PATH_FIR)
	buffer := thrift.NewTBufferedTransportFactory(4096)
	trans := httpClient.(*thrift.THttpClient)
	trans.SetHeader("User-Agent",config.USER_AGENT)
	trans.SetHeader("X-Line-Application",config.LINE_APPLICATION)
	trans.SetHeader("X-Line-Carrier",config.CARRIER)
	trans.SetHeader("X-Line-Access",AuthToken)
	buftrans, _ := buffer.GetTransport(trans)
	compactProtocol := thrift.NewTCompactProtocolFactory()
	return LineThrift.NewTalkServiceClientFactory(buftrans, compactProtocol)
}

func PollService() *LineThrift.TalkServiceClient {
	//fmt.Println("#### TalkService Initiated. ####")
	httpClient, _ := thrift.NewTHttpClient(config.LINE_HOST_DOMAIN+config.LINE_POLL_QUERY_PATH_THI)
	buffer := thrift.NewTBufferedTransportFactory(4096)
	trans := httpClient.(*thrift.THttpClient)
	trans.SetHeader("User-Agent",config.USER_AGENT)
	trans.SetHeader("X-Line-Application",config.LINE_APPLICATION)
	trans.SetHeader("X-Line-Carrier",config.CARRIER)
	trans.SetHeader("X-Line-Access",AuthToken)
	buftrans, _ := buffer.GetTransport(trans)
	compactProtocol := thrift.NewTCompactProtocolFactory()
	return LineThrift.NewTalkServiceClientFactory(buftrans, compactProtocol)
}

func ChannelService() *LineThrift.ChannelServiceClient {
	//fmt.Println("#### TalkService Initiated. ####")
	httpClient, _ := thrift.NewTHttpClient(config.LINE_HOST_DOMAIN+config.LINE_CHAN_QUERY_PATH)
	buffer := thrift.NewTBufferedTransportFactory(4096)
	trans := httpClient.(*thrift.THttpClient)
	trans.SetHeader("User-Agent",config.USER_AGENT)
	trans.SetHeader("X-Line-Application",config.LINE_APPLICATION)
	trans.SetHeader("X-Line-Carrier",config.CARRIER)
	trans.SetHeader("X-Line-Access",AuthToken)
	buftrans, _ := buffer.GetTransport(trans)
	compactProtocol := thrift.NewTCompactProtocolFactory()
	return LineThrift.NewChannelServiceClientFactory(buftrans, compactProtocol)
}

func CallService() *LineThrift.CallServiceClient {
	//fmt.Println("#### CallService Initiated. ####")
	httpClient, _ := thrift.NewTHttpClient(config.LINE_HOST_DOMAIN+config.LINE_CALL_QUERY_PATH)
	buffer := thrift.NewTBufferedTransportFactory(4096)
	trans := httpClient.(*thrift.THttpClient)
	trans.SetHeader("User-Agent",config.USER_AGENT)
	trans.SetHeader("X-Line-Application",config.LINE_APPLICATION)
	trans.SetHeader("X-Line-Carrier",config.CARRIER)
	trans.SetHeader("X-Line-Access",AuthToken)
	buftrans, _ := buffer.GetTransport(trans)
	compactProtocol := thrift.NewTCompactProtocolFactory()
	return LineThrift.NewCallServiceClientFactory(buftrans, compactProtocol)
}

func SquareService() *LineThrift.SquareServiceClient {
	//fmt.Println("#### SquareService Initiated. ####")
	httpClient, _ := thrift.NewTHttpClient(config.LINE_HOST_DOMAIN+config.LINE_SQUARE_QUERY_PATH)
	buffer := thrift.NewTBufferedTransportFactory(4096)
	trans := httpClient.(*thrift.THttpClient)
	trans.SetHeader("User-Agent",config.USER_AGENT)
	trans.SetHeader("X-Line-Application",config.LINE_APPLICATION)
	trans.SetHeader("X-Line-Carrier",config.CARRIER)
	trans.SetHeader("X-Line-Access",AuthToken)
	buftrans, _ := buffer.GetTransport(trans)
	compactProtocol := thrift.NewTCompactProtocolFactory()
	return LineThrift.NewSquareServiceClientFactory(buftrans, compactProtocol)
}

func AuthService() *LineThrift.TalkServiceClient {
	httpClient, _ := thrift.NewTHttpClient(config.LINE_HOST_DOMAIN+config.LINE_LOGIN_QUERY_PATH)
	buffer := thrift.NewTBufferedTransportFactory(4096)
	trans := httpClient.(*thrift.THttpClient)
	trans.SetHeader("User-Agent",config.USER_AGENT)
	trans.SetHeader("X-Line-Application",config.LINE_APPLICATION)
	trans.SetHeader("X-Line-Carrier",config.CARRIER)
	buftrans, _ := buffer.GetTransport(trans)
	compactProtocol := thrift.NewTCompactProtocolFactory()
	return LineThrift.NewTalkServiceClientFactory(buftrans, compactProtocol)
}

func LoginZService() *LineThrift.AuthServiceClient {
	httpClient, _ := thrift.NewTHttpClient(config.LINE_HOST_DOMAIN+config.LINE_AUTH_QUERY_PATH)
	buffer := thrift.NewTBufferedTransportFactory(4096)
	trans := httpClient.(*thrift.THttpClient)
	trans.SetHeader("User-Agent",config.USER_AGENT)
	trans.SetHeader("X-Line-Application",config.LINE_APPLICATION)
	trans.SetHeader("X-Line-Carrier",config.CARRIER)
	trans.SetHeader("X-Line-Access",AuthToken)
	buftrans, _ := buffer.GetTransport(trans)
	compactProtocol := thrift.NewTCompactProtocolFactory()
	return LineThrift.NewAuthServiceClientFactory(buftrans, compactProtocol)
}
