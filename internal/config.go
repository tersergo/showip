package internal

import (
	"fmt"
	"strings"
)

// Config 单一实例 环境配置参数
var _configs = &Config{
	Port:      80,
	Path:      "/" + ModuleName,
	FormatArg: FormatVarName,
	ViaArg:    ViaVarName,
	ModeArg:   ModeVarName,
	ObjArg:    ObjVarName,
}

// Config 配置参数
type Config struct {
	// Port 服务响应端口
	Port int
	// Port 服务响应的路径
	Path string

	// Header 追加或优先获取IP协议头参数
	Header string
	// ViaArg X-Via响应头参数名称
	ViaArg string

	// FormatArg 用户请求返回格式参数名称
	FormatArg string
	// ModeArg 返回模式参数名称
	ModeArg string
	// ModeArg 返回对象参数名称
	ObjArg string
}

// GetConfig 获取环境配置参数
func GetConfig() *Config {
	return _configs
}

// GetPath 获取服务响应的路径
func (conf *Config) GetPath() string {
	if strings.HasPrefix(conf.Path, "/") {
		return conf.Path
	}
	return "/" + conf.Path
}

// GetPort 获取服务响应的端口
func (conf *Config) GetPort() int {
	if conf.Port < 1 {
		return 80
	}

	return conf.Port
}

// GetHeaders 优先获取用户请求header的指定名称的ip参数
func (conf *Config) GetHeaders() (header []string) {
	if len(conf.Header) > 0 {
		header = ToArray(conf.Header, ArraySplitKey)
	}

	return
}

// FormatIsValid 是否响应format输出参数
func (conf *Config) FormatIsValid() bool {
	return conf.FormatArg != InvalidArg
}

// ModeIsValid 是否响应获取服务器IP的mode参数
func (conf *Config) ModeIsValid(host string) bool {
	return conf.ModeArg != InvalidArg && strings.EqualFold(host, ModeVarIsHost)
}

// ViaIsValid 是否输出包含服务器IP的X-Via响应头
func (conf *Config) ViaIsValid() bool {
	return conf.ViaArg != InvalidArg
}

// GetVersion 获取当前版本
func GetVersion() string {
	return fmt.Sprint(ModuleName, " ", Version)
}
