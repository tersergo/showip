// package showip
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tersergo/showip/internal"
)

// isVerMode 是否展示版本
var isVerMode bool

func init() {
	configs := internal.GetConfig()
	// services launch argument
	flag.IntVar(&configs.Port, "port", 80, "http services port config")
	flag.StringVar(&configs.Path, "path", "/"+internal.ModuleName, "http services path config")
	flag.StringVar(&configs.Header, "header", "", "append request ip header names")
	// response output header argument
	flag.StringVar(&configs.ViaArg, "via", internal.ViaVarName, "response header X-Via names")
	// request argument name
	flag.StringVar(&configs.FormatArg, "format", internal.FormatVarName, "request output argument name")
	flag.StringVar(&configs.ModeArg, "mode", internal.ModeVarName, "request mode argument name")
	flag.BoolVar(&isVerMode, "version", false, "print package name and version")

	flag.Parse()
}

func main() {
	if isVerMode {
		fmt.Println(internal.GetVersion())
		os.Exit(0)
	}

	envConf := internal.GetConfig()
	log.Println("launch showip services:", internal.NewServerIP().GetServerURL())
	log.Println("load environment config", internal.ToJson(envConf))

	// http.HandleFunc("/", webHandler)
	http.HandleFunc(envConf.GetPath(), ipHandler) // 默认响应路径/showip
	err := http.ListenAndServe(fmt.Sprintf(":%d", envConf.GetPort()), nil)

	if err != nil {
		log.Fatalln("ListenAndServe err: ", err)
	}

}

func ipHandler(rsp http.ResponseWriter, req *http.Request) {
	configs, client, server := internal.GetConfig(), internal.NewClientIP(req), internal.NewServerIP()

	log.Println(client.Default(), req.URL)

	var ipObj internal.IPPacker = client
	reqMode, reqFormat, reqObjId := client.GetQuery(configs.ModeArg), client.GetQuery(configs.FormatArg),
		client.GetQuery(configs.ObjArg)

	if configs.ModeIsValid(reqMode) { // 是否响应mode参数，返回服务器ip信息
		ipObj = server
	}

	outType, rspBody := internal.OutputText, ""
	if configs.FormatIsValid() { // 是否响应format参数
		outType = internal.ToOutputType(reqFormat)
	}

	switch outType {
	case internal.OutputArray:
		rspBody = internal.ToJson(ipObj.GetArray())
	case internal.OutputJSON:
		rspBody = internal.ToJson(ipObj.GetMap())
	case internal.OutputXML:
		rspBody = internal.ToXML(ipObj, reqObjId)
	case internal.OutputHTML:
		rspBody = internal.ToHTML(ipObj.GetArray(), reqObjId)
	default: // case OutputText:
		rspBody = ipObj.String()
	}

	if configs.ViaIsValid() { // 是否输出包含服务器IP的X-Via头信息
		rsp.Header().Set(configs.ViaArg, server.GetSimpleIP())
	}
	rsp.WriteHeader(http.StatusOK)

	_, err := rsp.Write([]byte(rspBody + "\n"))
	if err != nil {
		log.Fatalln("err", err, "response", rspBody)
	}
}
