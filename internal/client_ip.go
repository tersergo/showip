package internal

import (
	"fmt"
	"net/http"
	"reflect"
)

// ClientIP 客户端IP
type ClientIP struct {
	ipMap   map[string]string // ip map Set
	ipArray []string          // ip Array

	Request
}

// NewClientIP 构造客户端IP实例
func NewClientIP(req *http.Request) *ClientIP {
	return &ClientIP{Request: NewRequest(req)}
}

// init ClientIP argument
func (client *ClientIP) init() {
	if len(client.ipMap) > 0 {
		return
	}
	client.ipMap, client.ipArray = make(map[string]string), make([]string, 0)
	headers, appendHeader := DefaultHeaderList(), GetConfig().GetHeaders()

	if len(appendHeader) > 0 { // 配置中包含指定header名称
		headers = MergeArray(appendHeader, headers)
	}

	ipKeys := make(map[string]bool)
	for _, key := range headers {
		if _, ok := client.ipMap[key]; ok { // 过滤名称重复key
			continue
		}

		val, has := client.TryGetHeaderIP(key)
		if !has || len(val) == 0 {
			continue
		}

		client.ipMap[key] = val
		valList := ToArray(val)
		for _, ip := range valList { // 过滤重复ip
			if _, ok := ipKeys[ip]; !ok {
				client.ipArray = append(client.ipArray, ip)
				ipKeys[ip] = true
			}
		}
	}

	if len(client.ipArray) > 0 {
		client.ipMap[NodeNameIP] = client.ipArray[0]
	}
}

// Default 获取默认IP
func (client *ClientIP) Default() (cip string) {
	client.init()
	if len(client.ipArray) > 0 {
		return client.ipArray[0]
	}

	return

}

// String 字符串格式
func (client *ClientIP) String() string {
	client.init()

	return fmt.Sprint(NodeNameIP, ": ", client.Default())
}

// GetArray 获取客户端ip列表
func (client *ClientIP) GetArray() (ips []string) {
	client.init()
	if len(client.ipArray) > 0 {
		return client.ipArray
	}

	return []string{}
}

// GetMap 获取客户端ip列表map格式
func (client *ClientIP) GetMap() (ips map[string]string) {
	client.init()
	if len(client.ipMap) > 0 {
		return client.ipMap
	}

	return map[string]string{}
}

// TypeName  类名称
func (client *ClientIP) TypeName() string {
	return reflect.TypeOf(*client).Name()
}
