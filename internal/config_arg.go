package internal

import (
	"strings"
)

// configs 单一实例 环境配置参数
var configs = &configArg{
	ModuleName: "showip",

	ServerPort: 80,
	ServerPath: "/showip",

	FormatArgName:  "format",
	XViaHeaderName: "X-Via",
	ModeArgName:    "mode",
	ObjArgName:     "obj",
}

// configArg 配置参数
type configArg struct {
	// ModuleName 模块名称
	ModuleName string

	// ServerPort 服务响应端口
	ServerPort int
	// ServerPort 服务响应的路径
	ServerPath string

	// Header 追加优先获取用户指定头信息的ip参数
	Header string
	// XViaArgName X-Via响应头参数名称
	XViaHeaderName string

	// FormatArgName 用户请求返回格式参数名称
	FormatArgName string
	// ModeArgName 返回模式参数名称
	ModeArgName string
	// ModeArgName 返回对象参数名称
	ObjArgName string
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

// FormatIsValid 是否响应format输出参数
func (confArg *configArg) FormatIsValid() bool {

	return confArg.FormatArgName != "0"
}

// ModeIsValid 是否响应获取服务器IP的mode参数
func (confArg *configArg) ModeIsValid(host string) bool {

	return confArg.ModeArgName != "0" && strings.EqualFold(host, HostMode)
}

// XViaIsValid 是否输出包含服务器IP的X-Via响应头
func (confArg *configArg) XViaIsValid() bool {

	return confArg.XViaHeaderName != "0"
}
