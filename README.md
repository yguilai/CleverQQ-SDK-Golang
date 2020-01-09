# ~~CleverQQ-SDK-Golang~~
# CleverQQ框架 官方已停止服务

[![deprecated](http://badges.github.io/stability-badges/dist/deprecated.svg)](https://github.com/yguilai/CleverQQ-SDK-Golang/)

本项目为CleverQQ(原IRQQ)插件的Golang SDK, 方便Gopher们开发插件. 

使用过程中有任何问题或建议, 请大胆提Issue或PR.

## 更新日志
* 2019-09-12: 修复IRUploadPic无法正常上传图片的问题, 新增`file2byte()`函数用以读取文件返回`[]byte`, 新增TODO List任务.
* 2019-08-04: 给IR_Event事件添加recover, 避免插件panic直接导致框架主程序闪退, 但该功能默认情况下关闭, 需要`clvq.OnRecover=true`开启
* 2019-08-01: 完善文档, 提出获取本项目可能出现的异常及其解决措施
* 2019-07-31: 增加快速编译dll工具
* 2019-07-30: 修复IRIfFriend函数的Bug
* 2019-07-30: 过滤一些无用文件
* 2019-07-30: 上传本项目

## 注意
* 本项目使用了cgo, 而框架提供的`IRapi.dll`为32位, 故请确保你的系统中有安装32位GCC编译器
> 附我的[GCC编译器](https://www.lanzous.com/i58etri?_blank)
* 编译时, 请将输出的.dll文件名修改为`插件名称.IR.dll`, 与你代码中的插件名称保持一致, 否则会导致框架无法正常启动(闪退)
* 请勿在项目的其他包内导入clvq包, 因为这会导致`IR_Create IR_Message`等4个函数重复export, 最后导致框架调用到你的插件时会闪退

## 获取
```bash
go get github.com/yguilai/CleverQQ-SDK-Golang
```

### 获取本项目可能出现的异常
* 由于地区原因, 无法成功获取到`golang.org`的包
```bash
https fetch failed: Get https://golang.org/x/text/encoding/simplifiedchinese?go-get=1: dial tcp 216.239.37.1:443: 
// connectex: A connection attempt failed because the connected party did not properly respond after a period of time,
or established connection failed because connected host has failed to respond.
```

请尝试修改GO的系统环境变量后, 重新`go get`本项目. 如下:
```bash
//打开cmd,执行
SET GO111MODULE=on
SET GOPROXY=https://goproxy.io
```
通过此设置, 当你`go get`通过`go mod`方式管理的项目时, 会将`https://goproxy.io`作为你的代理
> 此功能目前对非go mod方式管理的go语言项目无效

* 本地GCC编译器非32位, 或无32位GCC编译器, `go get`后得到如下异常: 
```bash
# github.com/yguilai/CleverQQ-SDK-Golang/clvq
cc1.exe: sorry, unimplemented: 64-bit mode not compiled in
```
此异常可以忽略, 你只需遵循[使用](#使用)章节中的规则编译插件即可.

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
> ~~建议使用bat脚本, 这会让步骤1中的修改变为临时的. 如本项目demo中的build.bat~~.
> 建议按照规则使用快速编译工具.

## 快速编译工具
### 使用

将tool目录下的`build.exe`和`build.bat`文件复制到与你导入本项目的.go文件同级目录下, 双击执行build.exe即可. 

> 注意: 你的.go文件必须包含你自定义的插件名, 例如`clvq.PluginName = "demo"`


## TODO List

* [ ] 测试当前所有Api
* [ ] 完善文档
* [x] 开发快速编译dll工具
* [ ] 修复点击卸载插件会导致程序闪退问题
* [ ] 修复插件输出日志来源不正常的显示
* [ ] ...

## License

[MIT](https://github.com/yguilai/CleverQQ-SDK-Golang/blob/master/LICENSE)
