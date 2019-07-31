package main

import (
	"github.com/yguilai/CleverQQ-SDK-Golang/clvq"
)

func init() {
	clvq.PluginName = "demo"
	clvq.PluginVer = "0.0.1"
	clvq.PluginDesc = "这是一个demo"
	clvq.PluginAuthor = "燕归来"
	clvq.IREvent = irEvent
}

func main() {}

// 在此函数下处理所有事件, 消息
func irEvent(qq string, msgType, subMsgType int, msgFrom, tigObjAct, tigObjPas, msg, msgNum, msgId, rawMsg, json string, ptrNext int) int {
	return clvq.MT_CONTINUE
}
