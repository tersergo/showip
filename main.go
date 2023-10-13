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
	configs := internal.GetConfigArg()

	flag.IntVar(&configs.ServerPort, "port", 80, "http services port config")
	flag.StringVar(&configs.ServerPath, "path", "/showip", "http services path config")
	flag.StringVar(&configs.Header, "header", "", "request header names")
	flag.StringVar(&configs.FormatArgName, "format", "format", "request output argument name")

	flag.Parse()
}

func main() {
	log.Println("start showip ", internal.NewServerIP().GetServerURL())
	serverPort := internal.GetConfigArg().ServerPort

	http.HandleFunc("/", webHandler)
	err := http.ListenAndServe(fmt.Sprintf(":%d", serverPort), nil)

	if err != nil {
		log.Fatalln("ListenAndServe err: ", err)
	}

}

func webHandler(rsp http.ResponseWriter, req *http.Request) {
	rspBody, rspCode := "", 200
	configs, client, server := internal.GetConfigArg(), internal.NewClientIP(req), internal.NewServerIP()
	var ipPack internal.IPPacker

	ipPack = client

	if req.URL.Path == configs.GetServerPath() { // 默认响应路径
		outType := internal.ToOutputType(client.GetQuery(configs.FormatArgName))
		switch outType {
		case internal.OutputArray:
			rspBody = internal.ToJson(ipPack.GetIPArray())
		case internal.OutputJSON:
			rspBody = internal.ToJson(ipPack.GetIPMap())
		case internal.OutputXML:
			rspBody = internal.ToXML(ipPack.GetIPMap(), configs.ModuleId)
		case internal.OutputHTML:
			rspBody = internal.ToHTML(ipPack.GetIPArray(), configs.ModuleId)
		// case OutputDefault:
		default:
			rspBody = fmt.Sprint("IP: ", ipPack.GetIP())
		}

		rsp.Header().Set(configs.XViaHeaderName, server.GetIP())
	} else {
		rspCode = 404
	}
	log.Println(rspCode, client.GetIP(), req.URL)

	rsp.WriteHeader(rspCode)

	if len(rspBody) == 0 {
		return
	}

	_, err := rsp.Write([]byte(rspBody))
	if err != nil {
		log.Println("err", err, "rspBody", rspBody)
	}
}
