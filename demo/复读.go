package main

import (
	"github.com/yguilai/CleverQQ-SDK-Golang/clvq"
	"strings"
)

func init() {
	clvq.PluginName = "复读"
	clvq.PluginVer = "1.0.0"
	clvq.PluginDesc = "复读机一枚"
	clvq.PluginAuthor = "燕归来"
	clvq.IREvent = irEvent
}

// main函数不写代码, 仅为编译需要
func main() {

}

func irEvent(qq string, msgType, subMsgType int, msgFrom, tigObjAct, tigObjPas, msg, msgNum, msgId, rawMsg, json string, ptrNext int) int {
	if msgType == clvq.MT_GROUP || msgType == clvq.MT_FRIEND {
		// 注意: go语言中中文占3字符
		if msg[:6] == "复读" {
			str := strings.Replace(msg, "复读", "", -1)
			str = strings.TrimSpace(str)
			clvq.IRSendMsg(qq, msgType, msgFrom, tigObjAct, msg, -1)
			return clvq.MT_CONTINUE
		}
	}
	return clvq.MT_CONTINUE
}
