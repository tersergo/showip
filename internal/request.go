package internal

import (
	"html"
	"net"
	"net/http"
	"strings"
)

// Request 客户端请求查询
type Request struct {
	httpReq *http.Request // http Request
}

// NewRequest 构造客户端请求实例
func NewRequest(req *http.Request) Request {
	return Request{httpReq: req}
}

// TryGetHeaderIP 获取请求头部ip信息
func (client Request) TryGetHeaderIP(name string) (val string, isValid bool) {
	if len(name) == 0 || client.httpReq == nil {
		return
	}

	if name == HeaderNameRemoteAddr {
		val = client.GetRemoteIP()
	} else {
		val = client.GetHeader(name)
	}
	// 有效ip ::1 or 127.0.0.1
	isValid = len(val) > 2 && !strings.EqualFold(val, "unknown")

	return
}

// GetHeader 获取请求头部信息
func (client Request) GetHeader(name string) (val string) {
	val = client.httpReq.Header.Get(name)

	if len(val) > 0 {
		val = strings.TrimSpace(val)
	}

	return
}

// GetRemoteIP  获取请求的网络地址IP(可能是代理服务器IP)
func (client Request) GetRemoteIP() (rip string) {
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
func (client Request) GetQuery(key string, defValue ...string) (val string) {
	if client.httpReq != nil && client.httpReq.URL != nil {
		val = client.httpReq.URL.Query().Get(key)
	}

	if len(val) > 0 {
		val = strings.TrimSpace(val)
		val = html.EscapeString(val) // 预防xss攻击

		return
	}

	if len(defValue) > 0 { // 有请求默认值
		return defValue[0]
	}

	return
}
