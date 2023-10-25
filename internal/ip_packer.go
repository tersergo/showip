package internal

// IPPacker ip地址包接口
type IPPacker interface {
	// Default 获取默认IP
	Default() string
	// GetArray 获取ip列表
	GetArray() []string
	// GetMap 获取ip列表map格式
	GetMap() map[string]string
	// String 字符串格式
	String() string
	// TypeName 类名称
	TypeName() string
}

// DefaultHeaderList 默认ip相关头部列表
func DefaultHeaderList() []string {
	return []string{
		HeaderNameXFF,        // X-Forwarded-For
		HeaderNameXRIP,       // X-Real-IP
		HeaderNamePCIP,       // Proxy-Client-IP
		HeaderNameWLPCIP,     // WL-Proxy-Client-IP
		HeaderNameRemoteAddr, // RemoteAddr
	}
}
