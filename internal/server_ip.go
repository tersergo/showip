package internal

import (
	"fmt"
	"net"
	"os"
	"reflect"
	"strings"
)

// ServerIP 服务器IP
type ServerIP struct {
	ipV4 []string
	ipV6 []string
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
			// if loopNet.IP.To4() != nil {
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
		ips[NodeNameIPV6] = strings.Join(server.ipV6, ",")
	}

	return
}

// String 字符串格式
func (server *ServerIP) String() string {
	server.init()

	return fmt.Sprint(NodeNameIP, ": ", server.Default())
}

// GetServerURL 获取服务器url
func (server *ServerIP) GetServerURL() string {
	host := server.Default()
	if len(host) == 0 {
		host, _ = os.Hostname()
	}

	if GetConfig().Port != 80 {
		host += fmt.Sprintf(":%d", GetConfig().Port)
	}

	return fmt.Sprint("http", "://", host, GetConfig().GetPath())
}

// TypeName  类名称
func (server *ServerIP) TypeName() string {
	return reflect.TypeOf(*server).Name()
}

// GetSimpleIP 默认服务器ip极简模式：ipv4返回后2位，ipv6返回后3位
func (server *ServerIP) GetSimpleIP() string {
	server.init()

	cip := server.Default()
	if len(cip) < 6 { // 非有效ip
		return cip
	}

	lastIndex, split := 0, ""
	if strings.Contains(cip, ".") { // ipv4
		lastIndex, split = 2, "."
	} else if strings.Contains(cip, ":") { // ipv6
		lastIndex, split = 3, ":"
	} else {
		return cip // 未知 IP 类型
	}

	ipSplits := ToArray(cip, split)
	maxLen := len(ipSplits)
	if maxLen <= lastIndex {
		return cip
	}

	var lastList []string
	for i := lastIndex; i > 0; i-- {
		lastList = append(lastList, ipSplits[maxLen-i])
	}

	return strings.Join(lastList, split)
}
