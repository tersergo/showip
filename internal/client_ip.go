package internal

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

// ClientIP 客户端IP
type ClientIP struct {
	ipMap   map[string]string // ip map Set
	ipArray []string          // ip Array

	RequestQuery
}

// NewClientIP 构造客户端IP实例
func NewClientIP(req *http.Request) *ClientIP {
	return &ClientIP{RequestQuery: NewQuery(req)}
}

// init ClientIP argument
func (client *ClientIP) init() {
	if len(client.ipMap) > 0 {
		return
	}
	client.ipMap, client.ipArray = make(map[string]string), make([]string, 0)
	headers, appendHeader := DefaultHeaderList(), GetConfigs().GetHeaders()

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
	var builder strings.Builder

	for k, v := range client.ipMap {
		str := fmt.Sprintf("%s: %s\n", k, v)
		builder.WriteString(str)
	}

	return builder.String()
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
