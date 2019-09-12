package main

import (
	"github.com/yguilai/CleverQQ-SDK-Golang/clvq"
	"io/ioutil"
	"net/http"
)

func init() {
	clvq.PluginName = "复读"
	clvq.PluginVer = "1.0.0"
	clvq.PluginDesc = "复读机一枚"
	clvq.PluginAuthor = "燕归来"
	clvq.IREvent = irEvent
	clvq.OnRecover = true
}

// main函数不写代码, 仅为编译需要
func main() {

}

func irEvent(qq string, msgType, subMsgType int, msgFrom, tigObjAct, tigObjPas, msg, msgNum, msgId, rawMsg, json string, ptrNext int) int {
	if msgType == clvq.MT_GROUP || msgType == clvq.MT_FRIEND {

		if msg == "图片" {
			res, err := http.Get("http://hbimg.b0.upaiyun.com/32f065b3afb3fb36b75a5cbc90051b1050e1e6b6e199-Ml6q9F_fw320")
			if err != nil {
				clvq.IROutPutLog(err.Error())
				return clvq.MT_CONTINUE
			}

			data, _ := ioutil.ReadAll(res.Body)

			guid := clvq.IRUploadPic(qq, msgType, msgFrom, data)

			clvq.IRSendMsg(qq, msgType, msgFrom, tigObjAct, `[IR:pic={`+ guid +`}.jpg]`, -1)
		}
	}
	return clvq.MT_CONTINUE
}
