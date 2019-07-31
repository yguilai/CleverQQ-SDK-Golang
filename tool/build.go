package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	for _, v := range getGoFile() {
		pluginName := getPlugin_Name(v)
		if pluginName != "" {
			build(v, pluginName)
		}
	}
}

func getGoFile() []string {
	pwd, _ := os.Getwd()
	fileInfoList, err := ioutil.ReadDir(pwd)
	if err != nil {
		log.Fatal(err)
	}
	gofile := make([]string, 0)
	for i := range fileInfoList {
		name := fileInfoList[i].Name()
		if strings.Index(name, ".go") != -1 {
			gofile = append(gofile, name)
		}
	}
	return gofile
}

func getPlugin_Name(name string) string {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	flag := false
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}
		if strings.Index(line, "package main") != -1 {
			flag = true
		}
		if strings.Index(line, "PluginName") != -1 && flag {
			index_p := strings.Index(line, "\"")
			index_s := strings.LastIndex(line, "\"")
			return line[index_p+1 : index_s]
		}
	}
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()
	return ""
}

func build(goName string, pluginName string) {
	cmd := exec.Command("build.bat", pluginName)
	stdout, err := cmd.StdoutPipe()
	errpanic(err)
	defer stdout.Close()
	if err := cmd.Start(); err != nil { // 运行命令
		panic(err)
	}
	if opBytes, err := ioutil.ReadAll(stdout); err != nil { // 读取输出结果
		panic(err)
	} else {
		log.Println(string(opBytes))
	}
}

func errpanic(err error) {
	if err != nil {
		panic(err)
	}
}
