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

var (
	// serverPort 服务响应端口
	serverPort int
	// serverPath 服务响应的路径
	serverPath string
	// modName 模块名称
	modName string = "showip"
)

func init() {
	flag.IntVar(&serverPort, "port", 80, "http services port config")
	flag.StringVar(&serverPath, "path", "/showip", "http services path config")
	flag.Parse()
}

func main() {
	if !strings.HasPrefix(serverPath, "/") {
		serverPath = "/" + serverPath
	}

	log.Printf("start showip http services port(%d) path(%s)", serverPort, serverPath)

	http.HandleFunc("/", webHandler)
	err := http.ListenAndServe(fmt.Sprintf(":%d", serverPort), nil)

	if err != nil {
		log.Fatalf("fatal err: %v", err)
	}
}

func webHandler(rsp http.ResponseWriter, req *http.Request) {
	path, rspBody, rspCode := req.URL.Path, "", 200

	client := internal.NewClientIP(req)

	if path == serverPath {
		outType := internal.ToOutputType(client.GetQuery("format"))
		switch outType {
		case internal.OutputArray:
			rspBody = internal.ToJson(client.GetIPArray())
		case internal.OutputJSON:
			rspBody = internal.ToJson(client.GetIPMap())
		case internal.OutputXML:
			rspBody = internal.ToXML(client.GetIPMap(), modName)
		case internal.OutputHTML:
			rspBody = internal.ToHTML(client.GetIPArray(), modName)
		// case OutputDefault:
		default:
			rspBody = fmt.Sprint("IP: ", client.GetIP())
		}
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
