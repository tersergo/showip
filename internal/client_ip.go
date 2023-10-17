package internal

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

// ClientIP 客户端IP
type ClientIP struct {
	httpReq *http.Request     // http Request
	ipMap   map[string]string // ip map Set
	ipArray []string          // ip Array
}

// NewClientIP 构造客户端IP实例
func NewClientIP(req *http.Request) *ClientIP {
	return &ClientIP{httpReq: req}
}

// init ClientIP argument
func (client *ClientIP) init() {
	if len(client.ipMap) > 0 {
		return
	}
	client.ipMap, client.ipArray = make(map[string]string), make([]string, 0)
	headers, appendHeader := DefaultHeaderList(), GetConfigArg().GetAddHeaders()

	if len(appendHeader) > 0 { // 配置中包含指定header名称
		headers = MergeArray(appendHeader, headers)
	}

	for _, key := range headers {
		if _, ok := client.ipMap[key]; ok { // 过滤名称重复key
			continue
		}

		val, has := client.TryGetHeaderIP(key)
		if !has || len(val) == 0 {
			continue
		}

		valList := ToArray(val)
		client.ipArray = append(client.ipArray, valList...)
		client.ipMap[key] = val
	}

	if len(client.ipArray) > 0 {
		client.ipMap[NodeNameIP] = client.ipArray[0]
	}
}

// GetIP 获取IP
func (client *ClientIP) GetIP() (cip string) {
	client.init()
	if len(client.ipArray) > 0 {
		return client.ipArray[0]
	}

	return

}

// String 字符串格式
func (client *ClientIP) String() string {
	client.init()
	var builder strings.Builder

	for k, v := range client.ipMap {
		str := fmt.Sprintf("%s: %s\n", k, v)
		builder.WriteString(str)
	}

	return builder.String()
}

// TryGetHeaderIP 获取请求头部ip信息
func (client *ClientIP) TryGetHeaderIP(name string) (val string, isValid bool) {
	if len(name) == 0 || client.httpReq == nil {
		return
	}

	if name == HeaderNameRemoteAddr {
		val = client.GetRemoteIP()
	} else {
		val = strings.TrimSpace(client.httpReq.Header.Get(name))
	}
	// 有效ip ::1
	isValid = len(val) > 2 && !strings.EqualFold(val, "unknown")

	return
}

// GetRemoteIP  获取远程IP
func (client *ClientIP) GetRemoteIP() (rip string) {
	if client.httpReq == nil {
		return
	}

	remoteAddr := client.httpReq.RemoteAddr
	if len(remoteAddr) == 0 {
		return
	}

	host, _, err := net.SplitHostPort(strings.TrimSpace(remoteAddr))
	if err == nil {
		return host
	}

	return
}

// GetQuery  获取Get请求参数
func (client *ClientIP) GetQuery(key string, defaultVal ...string) (val string) {
	if client.httpReq == nil || client.httpReq.URL == nil {
		return
	}
	query := client.httpReq.URL.Query()
	val = query.Get(key)

	if len(val) == 0 && len(defaultVal) > 0 {
		val = defaultVal[0]
	}

	return
}

// GetIPArray 获取客户端ip列表
func (client *ClientIP) GetIPArray() (ips []string) {
	client.init()
	if len(client.ipArray) > 0 {
		return client.ipArray
	}

	return []string{}
}

// GetIPMap 获取客户端ip列表map格式
func (client *ClientIP) GetIPMap() (ips map[string]string) {
	client.init()
	if len(client.ipMap) > 0 {
		return client.ipMap
	}

	return map[string]string{}
}
