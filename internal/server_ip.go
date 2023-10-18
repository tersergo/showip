package internal

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
)

// ServerIP 服务器IP
type ServerIP struct {
	httpReq *http.Request // http Request
	ipV4    []string
	ipV6    []string
}

// NewServerIP 构造服务器IP实例
func NewServerIP() *ServerIP {
	return &ServerIP{}
}

// init ServerIP argument
func (server *ServerIP) init() {
	if len(server.ipV4) > 0 || len(server.ipV6) > 0 {
		return
	}

	ifList, err := net.InterfaceAddrs()
	if err != nil || len(ifList) == 0 {
		return
	}

	loopIPs := make([]*net.IPNet, 0)
	for _, addr := range ifList {
		ipNet, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}

		if ipNet.IP.IsLoopback() {
			loopIPs = append(loopIPs, ipNet)
			continue
		}

		if !ipNet.IP.IsPrivate() {
			continue
		}

		if ipNet.IP.To4() != nil {
			server.ipV4 = append(server.ipV4, ipNet.IP.String())
		} else {
			server.ipV6 = append(server.ipV6, ipNet.IP.String())
		}

	}
	// 本机没有有效ip时，补充本机环回ip
	if len(server.ipV4) == 0 && len(server.ipV6) == 0 && len(loopIPs) > 0 {
		for _, loopNet := range loopIPs {
			if len(loopNet.IP) == net.IPv4len {
				server.ipV4 = []string{loopNet.IP.String()}
			} else {
				server.ipV6 = []string{loopNet.IP.String()}
			}
		}
	}
}

// Default 获取默认IP
func (server *ServerIP) Default() (cip string) {
	server.init()

	if len(server.ipV4) > 0 {
		return server.ipV4[0]
	}
	if len(server.ipV6) > 0 {
		return server.ipV6[0]
	}

	return

}

// GetArray 获取服务器ip列表
func (server *ServerIP) GetArray() (ips []string) {
	server.init()

	return MergeArray(server.ipV4, server.ipV6)
}

// GetMap 获取服务器ip列表map格式
func (server *ServerIP) GetMap() (ips map[string]string) {
	server.init()

	defaultIP := server.Default()
	if len(defaultIP) == 0 {
		return
	}

	ips = map[string]string{
		NodeNameIP: defaultIP,
	}

	if len(server.ipV4) > 0 {
		ips[NodeNameIPV4] = strings.Join(server.ipV4, ",")
	}
	if len(server.ipV6) > 0 {
		ips[NodeNameIPV6] = strings.Join(server.ipV4, ",")
	}

	return
}

// String 字符串格式
func (server *ServerIP) String() string {
	server.init()
	var builder strings.Builder

	ipMap := server.GetMap()
	for k, v := range ipMap {
		str := fmt.Sprintf("%s: %s\n", k, v)
		builder.WriteString(str)
	}

	return builder.String()
}

// GetServerURL 获取服务器url
func (server *ServerIP) GetServerURL() string {
	host := server.Default()
	if len(host) == 0 {
		host, _ = os.Hostname()
	}

	if GetConfigs().Port != 80 {
		host += fmt.Sprintf(":%d", GetConfigs().Port)
	}

	return fmt.Sprint("http", "://", host, GetConfigs().GetServerPath())
}
