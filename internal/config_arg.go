package internal

import (
	"strings"
)

// configs 单一实例 环境配置参数
var configs = &configArg{
	ServerPort:     80,
	ServerPath:     "/showip",
	ModuleId:       "showip",
	FormatArgName:  "format",
	XViaHeaderName: "X-Via",
}

// configArg 配置参数
type configArg struct {
	// ServerPort 服务响应端口
	ServerPort int
	// ServerPort 服务响应的路径
	ServerPath string
	// Header 追加优先获取用户指定头信息的ip参数
	Header string

	// ModuleId 模块名称: 用于html，xml展示时的节点名称
	ModuleId string
	// FormatArgName 用户请求返回格式参数名称
	FormatArgName string
	// XViaArgName X-Via响应参数名称
	XViaHeaderName string
}

// GetConfigArg 获取环境配置参数
func GetConfigArg() *configArg {
	return configs
}

// GetServerPath 获取服务响应的路径
func (confArg *configArg) GetServerPath() string {
	if strings.HasPrefix(confArg.ServerPath, "/") {
		return confArg.ServerPath
	}

	return "/" + confArg.ServerPath
}

// GetAddHeaders 优先获取用户请求header的指定名称的ip参数
func (confArg *configArg) GetAddHeaders() (header []string) {
	if len(confArg.Header) > 0 {
		header = ToArray(confArg.Header, ",")
	}

	return
}
