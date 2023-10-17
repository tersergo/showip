package internal

// IPPacker ip地址包接口
type IPPacker interface {
	// GetIP 获取IP
	GetIP() string
	// GetIPArray 获取ip列表
	GetIPArray() []string
	// GetIPMap 获取ip列表map格式
	GetIPMap() map[string]string
	// String 字符串格式
	String() string
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
