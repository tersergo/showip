// package showip
package main

import (
	"flag"
	"fmt"
	"github.com/tersergo/showip/internal"
	"log"
	"net/http"
	"strings"
)

func init() {
	configs := internal.GetConfigArg()
	// services launch argument
	flag.IntVar(&configs.ServerPort, "port", 80, "http services port config")
	flag.StringVar(&configs.ServerPath, "path", "/showip", "http services path config")
	flag.StringVar(&configs.Header, "header", "", "append request ip header names")

	// response output header argument
	flag.StringVar(&configs.XViaHeaderName, "via", "X-Via", "response header via names")
	flag.StringVar(&configs.ObjId, "obj", "showip", "response output html styleId or xml nodeId")

	// request argument name
	flag.StringVar(&configs.FormatArgName, "format", "format", "request output argument name")
	flag.StringVar(&configs.ModeArgName, "mode", "mode", "request mode argument name")

	flag.Parse()
}

func main() {
	log.Println("launch showip services:", internal.NewServerIP().GetServerURL())
	envConf := internal.GetConfigArg()

	//http.HandleFunc("/", webHandler)
	http.HandleFunc(envConf.GetServerPath(), ipHandler) // 默认响应路径/showip
	err := http.ListenAndServe(fmt.Sprintf(":%d", envConf.ServerPort), nil)

	if err != nil {
		log.Fatalln("ListenAndServe err: ", err)
	}

}

func ipHandler(rsp http.ResponseWriter, req *http.Request) {
	rspCode, rspBody := 200, ""
	configs, client, server := internal.GetConfigArg(), internal.NewClientIP(req), internal.NewServerIP()

	log.Println(rspCode, client.GetIP(), req.URL)
	reqMode, reqFormat := client.GetQuery(configs.ModeArgName), client.GetQuery(configs.FormatArgName)

	var ipObj internal.IPPacker = client
	if configs.ModeIsValid() && strings.EqualFold(reqMode, "host") {
		ipObj = server
	}

	outType := internal.OutputDefault
	if configs.FormatIsValid() {
		outType = internal.ToOutputType(reqFormat)
	}

	switch outType {
	case internal.OutputArray:
		rspBody = internal.ToJson(ipObj.GetIPArray())
	case internal.OutputJSON:
		rspBody = internal.ToJson(ipObj.GetIPMap())
	case internal.OutputXML:
		rspBody = internal.ToXML(ipObj.GetIPMap(), configs.ObjId)
	case internal.OutputHTML:
		rspBody = internal.ToHTML(ipObj.GetIPArray(), configs.ObjId)
	// case OutputDefault:
	default:
		rspBody = fmt.Sprint("IP: ", ipObj.GetIP())
	}

	if configs.XViaIsValid() { //是否输出包含服务器IP的X-Via头信息
		rsp.Header().Set(configs.XViaHeaderName, server.GetIP())
	}
	rsp.WriteHeader(rspCode)

	_, err := rsp.Write([]byte(rspBody))
	if err != nil {
		log.Fatalln("err", err, "rspBody", rspBody)
	}
}
