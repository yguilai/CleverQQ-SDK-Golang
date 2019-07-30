package main

import (
	"github.com/yguilai/CleverQQ-SDK-Golang/clvq"
	"strings"
)

func init() {
	clvq.PluginName = "demo"
	clvq.PluginVer = "0.0.1"
	clvq.PluginDesc = "这是一个demo"
	clvq.PluginAuthor = "燕归来"
	clvq.IREvent = irEvent
}

func main() {

}

func irEvent(qq string, msgType, subMsgType int, msgFrom, tigObjAct, tigObjPas, msg, msgNum, msgId, rawMsg, json string, ptrNext int) int {
	if msg[:6] == "复读" {
		str := strings.Replace(msg, "复读", "", -1)
		str = strings.TrimSpace(str)
		clvq.IRSendMsg(qq, msgType, msgFrom, tigObjAct, str, -1)
	}
	return clvq.MT_CONTINUE
}
