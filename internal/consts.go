package internal

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

	ObjVarName = "obj"
	HostMode   = "host"
)
