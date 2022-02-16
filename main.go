package main

import (
	"crypto/tls"
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {
	//11111
	url := "ws://81.70.243.31:7345/ws"
	NewWebSocket(url, "test", "test11", "1213")
	/*
		url := "http://81.70.243.31:7145"
		function := "CreateAccount"
		var req = new(msg.CreateAccountReq)
		req.Account = "test1111112"
		req.AccountType = 1
		req.Password = "test"
		var info, _ = json.Marshal(req)
		HttpPost(url, function, string(info))

	*/
}

func NewWebSocket(url, token, account, roleId string) {
	url += "?token=" + token
	url += "&account=" + account
	url += "&roleId=" + roleId
	dialer := websocket.Dialer{}
	dialer.Proxy = http.ProxyFromEnvironment
	dialer.HandshakeTimeout = 15 * time.Second
	conn, res, err := dialer.Dial(url, nil)
	if nil != err {
		fmt.Printf("err:%v", err.Error())
	}
	fmt.Printf("addr:%v\n, res:%v", conn, res)
}

func HttpPost(url, function, info string) (string, error) {
	urlpath := url + "/v2/rpc/" + function + "?http_key=defaulthttpkey"
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	//	dd := bytes.NewReader([]byte(info))
	fmt.Printf("info :%v\n", info)
	reqs := strings.NewReader(info)
	req, err := http.NewRequest("POST", urlpath, reqs)

	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "name=anny")
	var header = "Bearer " + "defaulthttpkey"
	req.Header.Set("Authorization", header)
	response, err := client.Do(req)
	if nil == response {
		return "", err
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	fmt.Printf("res:%v\n, err:%v", string(data), err)
	return string(data), nil
}
