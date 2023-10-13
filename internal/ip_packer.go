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

const (
	// HeaderNameXFF   Squid,Nginx 代理服务ip头 X-Forwarded-For
	HeaderNameXFF = "X-Forwarded-For"
	// HeaderNameXRIP http代理服务ip头 X-Real-IP
	HeaderNameXRIP = "X-Real-IP"
	// HeaderNamePCIP Apache代理头 Proxy-Client-IP
	HeaderNamePCIP = "Proxy-Client-IP"
	// HeaderNameWLPCIP WebLogic代理服务ip头 WL-Proxy-Client-IP
	HeaderNameWLPCIP = "WL-Proxy-Client-IP"
	// HeaderNameRemoteAddr RemoteAddr头
	HeaderNameRemoteAddr = "RemoteAddress"

	// NodeNameIP IP
	NodeNameIP = "IP"
	// NodeNameIPV4 IPV4
	NodeNameIPV4 = "IPV4"
	// NodeNameIPV6 IPV6
	NodeNameIPV6 = "IPV6"
)

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
