package internal

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

const (
	// DefaultForwardedIPKey  http header X-Forwarded-For
	DefaultForwardedIPKey = "X-Forwarded-For"
	// DefaultRealIPKey http header X-Real-Ip
	DefaultRealIPKey = "X-Real-Ip"
	// DefaultRemoteAddrKey RemoteAddr
	DefaultRemoteAddrKey = "RemoteAddress"
	// DefaultIPKey IP
	DefaultIPKey = "IP"
)

// ClientIP 客户端IP
type ClientIP struct {
	httpReq *http.Request     // http Request
	ipMap   map[string]string // ip map Set
	ipArray []string          // ip Array
}

// NewClientIP 客户端IP
func NewClientIP(req *http.Request) *ClientIP {
	return &ClientIP{httpReq: req}
}

// init argument
func (client *ClientIP) init() {
	if len(client.ipMap) > 0 {
		return
	}
	client.ipMap = make(map[string]string)
	client.ipArray = make([]string, 0)
	headers := []string{DefaultForwardedIPKey, DefaultRealIPKey, DefaultRemoteAddrKey}

	for _, key := range headers {
		val, has := client.TryGetHeader(key)
		if !has || len(val) == 0 {
			continue
		}

		valList := ToArray(val)
		client.ipArray = append(client.ipArray, valList...)
		client.ipMap[key] = val
	}

	if len(client.ipArray) > 0 {
		client.ipMap[DefaultIPKey] = client.ipArray[0]
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

// TryGetHeader 获取请求头部信息
func (client *ClientIP) TryGetHeader(name string) (val string, hasVal bool) {
	if len(name) == 0 || client.httpReq == nil {
		return
	}

	if name == DefaultRemoteAddrKey {
		val = client.GetRemoteIP()
	} else {
		val = client.httpReq.Header.Get(name)
	}

	hasVal = len(val) > 0 && val != "unknown"

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
func (client *ClientIP) GetQuery(key string) (val string) {
	if client.httpReq == nil || client.httpReq.URL == nil {
		return
	}
	query := client.httpReq.URL.Query()
	val = query.Get(key)

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

// GetXForwardedForIP 获取转发IP X-Forwarded-For
func (client *ClientIP) GetXForwardedForIP() (val string) {
	val, _ = client.TryGetHeader(DefaultForwardedIPKey)

	return
}

// GetXRealIP 获取真实IP X-Real-Ip
func (client *ClientIP) GetXRealIP() (val string) {
	val, _ = client.TryGetHeader(DefaultRealIPKey)

	return
}
