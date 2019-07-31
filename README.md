# CleverQQ-SDK-Golang
本项目为CleverQQ(原IRQQ)插件的Golang SDK, 方便Gopher们开发插件. 

使用过程中有任何问题或建议, 请大胆提Issue或PR.

## 注意
* 本项目使用了cgo, 而框架提供的`IRapi.dll`为32位, 故请确保你的系统中有安装32位GCC编译器
> 附我的[GCC编译器](https://www.lanzous.com/i58etri?_blank)
* 编译时, 请将输出的.dll文件名修改为`插件名称.IR.dll`, 与你代码中的插件名称保持一致, 否则会导致框架无法正常启动(闪退)
* 请勿在项目的其他包内导入clvq包, 因为这会导致`IR_Create IR_Message`等4个函数重复export, 最后导致框架调用到你的插件时会闪退

## 获取
```bash
go get github.com/yguilai/CleverQQ-SDK-Golang
```

## 使用
0.导入本项目
```go
package main

import "github.com/yguilai/CleverQQ-SDK-Golang/clvq"

func init()  {
    clvq.PluginName = "插件名称"
    clvq.PluginVer = "1.0.0"    //插件版本
    clvq.PluginDesc = "插件说明"
    clvq.PluginAuthor = "作者"
    clvq.IREvent = irEvent
}

// main函数不写代码, 仅为编译需要
func main() {}

//通过实现clvq包内的IREvent方式来处理机器人分发下来的所有消息, 事件
//如果你有e语言开发插件的经验,相信你会很快上手这个SDK
func irEvent(qq string, msgType, subMsgType int, msgFrom, tigObjAct, tigObjPas, msg, msgNum, msgId, rawMsg, json string, ptrNext int) int {
	return clvq.MT_CONTINUE
}
```

1.修改几个GO的系统环境变量
```bash
SET CGO_LDFLAGS=-Wl,--kill-at
SET CGO_ENABLED=1
SET GOOS=windows
SET GOARCH=386
```

2.编译
```bash
go build -ldflags "-s -w" -buildmode=c-shared -o 插件名称.IR.dll
```
> 建议使用bat脚本, 这会让步骤1中的修改变为临时的. 如本项目demo中的build.bat,

## 快速编译工具
### 使用

将tool目录下的`build.exe`和`build.bat`文件复制到与你导入本项目的.go文件同级目录下, 双击执行build.exe即可. 

> 注意: 你的.go文件必须包含你自定义的插件名, 例如`clvq.PluginName = "demo"`

## TODO List

* [ ] 测试当前所有Api
* [ ] 完善文档
* [x] 开发快速编译dll工具
* [ ] ...

## License

[MIT](https://github.com/yguilai/CleverQQ-SDK-Golang/blob/master/LICENSE)