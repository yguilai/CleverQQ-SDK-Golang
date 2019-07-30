package main

import (
	"CleverQQ-SDK-Golang/clvq"
)

func init() {
	clvq.PluginName = "复读"
	clvq.PluginVer = "1.0.0"
	clvq.PluginDesc = "复读机一枚"
	clvq.PluginAuthor = "燕归来"
	clvq.IREvent = irEvent
}

// 此函数不写代码, 仅为编译需要
func main() {

}

func irEvent(qq string, msgType, subMsgType int, msgFrom, tigObjAct, tigObjPas, msg, msgNum, msgId, rawMsg, json string, ptrNext int) int {
	return clvq.MT_CONTINUE
}
