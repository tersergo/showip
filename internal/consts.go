package internal

const (
	// ModuleName 系统模块名称
	ModuleName = "showip"

	// HeaderNameXFF   Squid,Nginx 代理服务ip头 X-Forwarded-For
	HeaderNameXFF = "X-Forwarded-For"
	// HeaderNameXRIP http代理服务ip头 X-Real-IP
	HeaderNameXRIP = "X-Real-IP"
	// HeaderNamePCIP Apache代理头 Proxy-Client-IP
	HeaderNamePCIP = "Proxy-Client-IP"
	// HeaderNameWLPCIP WebLogic代理服务ip头 WL-Proxy-Client-IP
	HeaderNameWLPCIP = "WL-Proxy-Client-IP"
	// HeaderNameRemoteAddr RemoteAddr头名称
	HeaderNameRemoteAddr = "RemoteAddress"

	// NodeNameIP IP
	NodeNameIP = "IP"
	// NodeNameIPV4 IPV4
	NodeNameIPV4 = "IPV4"
	// NodeNameIPV6 IPV6
	NodeNameIPV6 = "IPV6"

	// InvalidArg 无效参数
	InvalidArg = "0"
	// ObjVarName obj变量名称
	ObjVarName = "obj"
	// FormatVarName format变量名称
	FormatVarName = "format"
	// ViaVarName via变量默认
	ViaVarName = "X-Via"
	// ModeVarName mode变量名称
	ModeVarName = "mode"
	// ModeVarIsHost mode变量值
	ModeVarIsHost = "host"
)
