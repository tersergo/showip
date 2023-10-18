// package showip
package main

import (
	"flag"
	"fmt"
	"github.com/tersergo/showip/internal"
	"log"
	"net/http"
)

func init() {
	configs := internal.GetConfigs()
	// services launch argument
	flag.IntVar(&configs.Port, "port", 80, "http services port config")
	flag.StringVar(&configs.Path, "path", "/"+internal.ModuleName, "http services path config")
	flag.StringVar(&configs.Header, "header", "", "append request ip header names")

	// response output header argument
	flag.StringVar(&configs.ViaArg, "via", internal.ViaVarName, "response header X-Via names")

	// request argument name
	flag.StringVar(&configs.FormatArg, "format", internal.FormatVarName, "request output argument name")
	flag.StringVar(&configs.ModeArg, "mode", internal.ModeVarName, "request mode argument name")

	flag.Parse()
}

func main() {
	envConf := internal.GetConfigs()
	log.Println("launch showip services:", internal.NewServerIP().GetServerURL())
	log.Println("load environment config", internal.ToJson(envConf))

	//http.HandleFunc("/", webHandler)
	http.HandleFunc(envConf.GetServerPath(), ipHandler) // 默认响应路径/showip
	err := http.ListenAndServe(fmt.Sprintf(":%d", envConf.Port), nil)

	if err != nil {
		log.Fatalln("ListenAndServe err: ", err)
	}

}

func ipHandler(rsp http.ResponseWriter, req *http.Request) {
	rspCode, rspBody := 200, ""
	configs, client, server := internal.GetConfigs(), internal.NewClientIP(req), internal.NewServerIP()

	log.Println(rspCode, client.Default(), req.URL)
	reqMode, reqFormat := client.GetQuery(configs.ModeArg), client.GetQuery(configs.FormatArg)
	reqObjId := client.GetQuery(configs.ObjArg, internal.ModuleName)

	var ipObj internal.IPPacker = client
	if configs.ModeIsValid(reqMode) { //是否响应mode参数，返回服务器ip信息
		ipObj = server
	}

	outType := internal.OutputText
	if configs.FormatIsValid() { //是否响应format参数
		outType = internal.ToOutputType(reqFormat)
	}

	switch outType {
	case internal.OutputArray:
		rspBody = internal.ToJson(ipObj.GetArray())
	case internal.OutputJSON:
		rspBody = internal.ToJson(ipObj.GetMap())
	case internal.OutputXML:
		rspBody = internal.ToXML(ipObj.GetMap(), reqObjId)
	case internal.OutputHTML:
		rspBody = internal.ToHTML(ipObj.GetArray(), reqObjId)
	default: // case OutputText:
		rspBody = fmt.Sprint(internal.NodeNameIP, ": ", ipObj.Default())
	}

	if configs.ViaIsValid() { // 是否输出包含服务器IP的X-Via头信息
		rsp.Header().Set(configs.ViaArg, server.Default())
	}
	rsp.WriteHeader(rspCode)

	_, err := rsp.Write([]byte(rspBody))
	if err != nil {
		log.Fatalln("err", err, "rspBody", rspBody)
	}
}
