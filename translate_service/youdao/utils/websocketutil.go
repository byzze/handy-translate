package utils

import (
	"fmt"
	"github.com/gorilla/websocket"
	neturl "net/url"
	"os"
	"strconv"
	"strings"
	"sync"
)

/*
初始化websocket连接
*/
func InitConnection(url string) (*websocket.Conn, *sync.WaitGroup) {
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		fmt.Println("connection failed.", err)
		os.Exit(-1)
	}
	wg := sync.WaitGroup{}
	// 监听返回数据
	go messageHandler(ws, &wg)
	wg.Add(1)
	return ws, &wg
}

/*
初始化websocket连接, 并附带参数
*/
func InitConnectionWithParams(url string, paramsMap map[string][]string) (*websocket.Conn, *sync.WaitGroup) {
	params := neturl.Values{}
	for k, v := range paramsMap {
		params[k] = v
	}
	parseUrl, _ := neturl.Parse(url)
	parseUrl.RawQuery = params.Encode()
	return InitConnection(parseUrl.String())
}

/*
发送binary message
*/
func SendBinaryMessage(ws *websocket.Conn, message []byte) {
	ws.WriteMessage(websocket.BinaryMessage, message)
	fmt.Println("send binary message length: " + strconv.Itoa(len(message)))
}

/*
发送text message
*/
func SendTextMessage(ws *websocket.Conn, message string) {
	ws.WriteMessage(websocket.TextMessage, []byte(message))
	fmt.Println("send text message: " + message)
}

func messageHandler(ws *websocket.Conn, wg *sync.WaitGroup) {
	for {
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("message handler error ", err)
			break
		}
		switch msgType {
		case websocket.TextMessage:
			message := string(msg)
			fmt.Println("received text message: " + message)
			if !strings.Contains(message, "\"errorCode\":\"0\"") {
				wg.Done()
				os.Exit(-1)
			}
		case websocket.BinaryMessage:
			fmt.Println("received binary message length: " + string(rune(len(msg))))
		case websocket.CloseMessage:
			fmt.Println("connection closed. " + string(msg))
		}
	}
}
