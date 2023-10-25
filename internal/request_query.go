package internal

import (
	"html"
	"net"
	"net/http"
	"strings"
)

// RequestQuery 客户端请求查询
type RequestQuery struct {
	httpReq *http.Request // http Request
}

// NewQuery 构造客户端请求实例
func NewQuery(req *http.Request) RequestQuery {
	return RequestQuery{httpReq: req}
}

// TryGetHeaderIP 获取请求头部ip信息
func (client RequestQuery) TryGetHeaderIP(name string) (val string, isValid bool) {
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
func (client RequestQuery) GetHeader(name string) (val string) {
	val = client.httpReq.Header.Get(name)

	if len(val) > 0 {
		val = strings.TrimSpace(val)
	}

	return
}

// GetRemoteIP  获取请求的网络地址IP(可能是代理服务器IP)
func (client RequestQuery) GetRemoteIP() (rip string) {
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
func (client RequestQuery) GetQuery(key string, defaultVal ...string) (val string) {
	if client.httpReq != nil && client.httpReq.URL != nil {
		val = client.httpReq.URL.Query().Get(key)
	}

	if len(val) > 0 {
		val = strings.TrimSpace(val)
		val = html.EscapeString(val) //预防xss攻击

		return
	}

	if len(defaultVal) > 0 { // 有配置请求默认值
		return defaultVal[0]
	}

	return
}
