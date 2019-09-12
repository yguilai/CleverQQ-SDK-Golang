package clvq

import "C"
import "fmt"

var (
	//插件名称
	PluginName = "PluginName.IR"
	//插件版本
	PluginVer = "1.0.0"
	//插件作者
	PluginAuthor = "Author"
	//插件说明
	PluginDesc = "description"
	// 以上4者可自行修改
	//插件Skey
	PluginSkey = "8956RTEWDFG3216598WERDF3"
	//插件SDK
	PluginSDK = "S3"
	//以上两个变量请勿修改
	OnRecover = false
	//添加recover 以免程序panic直接导致主程序闪退
)

//export IR_Create
// IR_Create
// 由于export需要, 此函数名为大写字母开头, 因此导致此函数可导出, 请勿直接调用此函数
func IR_Create() *C.char {
	if IRCreate == nil {
		pluginInfo := fmt.Sprintf(
			"插件名称{%s}\r"+
				"插件版本{%s}\r"+
				"插件作者{%s}\r"+
				"插件说明{%s}\r"+
				"插件skey{%s}\r"+
				"插件sdk{%s}",
			PluginName, PluginVer, PluginAuthor, PluginDesc, PluginSkey, PluginSDK)
		return cString(pluginInfo)
	}
	return cString(IRCreate())
}

//export IR_Message
// IR_Message
// 此函数用于接收所有原始封包内容，
// 返回-1 已收到并拒绝处理
// 返回0 未收到并不处理
// 返回1 处理完且继续执行
// 返回2 处理完毕并不再让其他插件处理
// (Pro版可用)
// 由于export需要, 此函数名为大写字母开头, 因此导致此函数可导出, 请勿直接调用此函数
func IR_Message(QQ *C.char, MsgType C.int, RawMsg, Cookies, SessionKey, ClientKey *C.char) C.int {
	if IRMessage == nil {
		return cInt(1)
	}
	return cInt(IRMessage(goString(QQ), goInt(MsgType), goString(RawMsg), goString(Cookies), goString(SessionKey), goString(ClientKey)))
}

//export IR_Event
// IR_Event
// 此函数会分发IRC_机器人QQ接收到的所有：事件，消息；您可在此函数中自行调用所有参数
// 由于export需要, 此函数名为大写字母开头, 因此导致此函数可导出, 请勿直接调用此函数
func IR_Event(QQ *C.char, MsgType, SubMsgType C.int, MsgFrom, TigObjAct, TigObjPas, Msg, MsgNum, MsgId, RawMsg, Json *C.char, PtrNext C.int) C.int {
	if IREvent == nil {
		return MT_CONTINUE
	}
	if OnRecover {
		defer func() {
			if err := recover(); err != nil{
				IROutPutLog(fmt.Sprintf("%s", err))
			}
		}()
	}

	return cInt(IREvent(goString(QQ), goInt(MsgType), goInt(SubMsgType), goString(MsgFrom), goString(TigObjAct),
		goString(TigObjPas), goString(Msg), goString(MsgNum), goString(MsgId), goString(RawMsg), goString(Json), goInt(PtrNext)))
}

//export IR_SetUp
// IR_SetUp 当用户按下“设置”执行本函数
// 由于Golang对GUI编写目前并不友好, 目前不支持此函数, 后期会尝试完善
// 由于export需要, 此函数名为大写字母开头, 因此导致此函数可导出, 请勿直接调用此函数
func IR_SetUp() {
	if IRSetUp == nil {
		return
	}
	IRSetUp()
}

//export IR_DestroyPlugin
// 由于export需要, 此函数名为大写字母开头, 因此导致此函数可导出, 请勿直接调用此函数
func IR_DestroyPlugin() C.int {
	if IRDestroyPlugin == nil {
		return cInt(0)
	}
	return cInt(IRDestroyPlugin())
}

////要对以上4个函数自定义, 请实现以下对应函数

// IRCreate 插件创建
// 框架启动时将调用此函数读取插件信息
// 如需自定义IR_Create函数, 请实现实现此函数IRCreate
// 请务必保证返回如下格式的字符串:
// "插件名称{%s}\n插件版本{%s}\n插件作者{%s}\n插件说明{%s}\n插件skey{%s}\n插件sdk{%s}"
// %s替换为对应内容
// 否则可能导致插件无法正常导入, 或导致框架无法正常运行
var IRCreate func() string

// IRMessage
// 如需自定义IR_Message函数, 请实现此函数IRMessage (Pro版可用)
// @param qq string 响应QQ/机器人QQ 用于判定哪个QQ接收到该消息
// @param msgType int UDP收到的原始信息
// @param rawMsg string 经过Tea加密的原文
// @param cookies string 用于登录网页所需cookies，如不确定用途请忽略
// @param sessionKey string 通信包所用的加密秘钥
// @param clientKey string 登录网页服务用的秘钥
var IRMessage func(qq string, msgType int, rawMsg, cookies, sessionKey, clientKey string) int

// IREvent 处理事件, 消息
// 此函数会分发IRC_机器人QQ接收到的所有：事件，消息；您可在此函数中自行调用所有参数
// 如需自定义IR_Event函数, 请实现此函数IREvent
// @param qq string 响应QQ/机器人QQ 用于判定哪个QQ接收到该消息
// @param msgType int 接收到消息类型，该类型可在常量表中查询具体定义，此处仅列举： -1 未定义事件 0,在线状态临时会话 1,好友信息 2,群信息 3,讨论组信息 4,群临时会话 5,讨论组临时会话 6,财付通转账 7,好友验证回复会话
// @param subMsgType int 此参数在不同消息类型下，有不同的定义，暂定：接收财付通转账时 1为好友 4为群临时会话 5为讨论组临时会话    有人请求入群时，不良成员这里为1
// @param msgFrom string 此消息的来源，如：群号、讨论组ID、临时会话QQ、好友QQ等
// @param tigObjAct string 主动发送这条消息的QQ，踢人时为踢人管理员QQ
// @param tigObjPas string 被动触发的QQ，如某人被踢出群，则此参数为被踢出人QQ
// @param msg string 此参数有多重含义，常见为：对方发送的消息内容，但当IRC_消息类型为 某人申请入群，则为入群申请理由
// @param msgNum string 消息序号 此参数暂定用于消息撤回
// @param msgId string 消息id 此参数暂定用于消息撤回
// @param rawMsg string UDP收到的原始信息，特殊情况下会返回JSON结构（入群事件时，这里为该事件seq）
// @param json string JSON格式转账解析
// @param ptrNext int 此参数用于插件加载拒绝理由  用法：写到内存（“拒绝理由”，IRC_信息回传文本指针_Out）此注释从E语言中复制
// @return int 正常情况下请返回常量 MT_CONTINUE 即返回1
var IREvent func(qq string, msgType, subMsgType int, msgFrom, tigObjAct, tigObjPas, msg, msgNum, msgId, rawMsg, json string, ptrNext int) int

// IRSetUp 点击设置
// 如需自定义IR_SetUp函数, 请实现此函数IRSetUp
// 由于go语言对GUI编程不友好, 此函数暂未实现, 欢迎大佬参与完善
var IRSetUp func()

// IRDestroyPlugin 卸载插件
// 当插件被卸载时将会调用
// @return int  返回非0成功 亦可不返回或为空
var IRDestroyPlugin func() int
