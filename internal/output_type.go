package internal

import (
	"encoding/json"
	"fmt"
	"strings"
)

// OutputType 输出类型
type OutputType int

const (
	// OutputDefault 默认输出
	OutputDefault OutputType = iota
	// OutputHTML 输出html
	OutputHTML
	// OutputArray 输出array
	OutputArray
	// OutputJSON 输出json
	OutputJSON
	// OutputXML 输出xml
	OutputXML
)

// ToOutputType 装换为 OutputType
func ToOutputType(outArg string) (outType OutputType) {
	outType = OutputDefault

	if len(outArg) == 0 {
		return
	}

	outArg = strings.ToLower(strings.TrimSpace(outArg))
	switch outArg {
	case "array":
		outType = OutputArray
	case "html":
		outType = OutputHTML
	case "json":
		outType = OutputJSON
	case "xml":
		outType = OutputXML
	default:
		outType = OutputDefault
	}

	return
}

// ToJson 装换为JSON String
func ToJson(obj interface{}) (jsonText string) {
	jsonText = "{}"
	if obj == nil {
		return
	}

	jsonBytes, err := json.Marshal(obj)
	if err == nil && len(jsonBytes) > 0 {
		jsonText = string(jsonBytes)
	}

	return
}

// ToXML 装换为XML String
func ToXML(obj map[string]string, rootNode ...string) string {
	var builder strings.Builder

	builder.WriteString("<? version=\"1.0\" encoding=\"UTF-8\" ?>\n")
	if len(rootNode) == 0 {
		rootNode = []string{"root"}
	}

	builder.WriteString(fmt.Sprintf("<%s>\n", rootNode[0]))
	if len(obj) > 0 {
		for k, v := range obj {
			str := fmt.Sprintf("\t<%[1]s>%[2]s</%[1]s>\n", k, v)
			builder.WriteString(str)
		}
	}
	builder.WriteString(fmt.Sprintf("</%s>\n", rootNode[0]))

	return builder.String()
}

// ToHTML 装换为HTML(ul-li) String
func ToHTML(objArray []string, styleId ...string) (htmlText string) {
	if len(objArray) == 0 {
		return
	}

	var builder strings.Builder

	ulTag := "<ul>\n"
	if len(styleId) > 0 {
		ulTag = fmt.Sprintf("<ul class=\"%s\">\n", styleId[0])
	}
	builder.WriteString(ulTag)

	for _, v := range objArray {
		str := fmt.Sprintf("\t<li>%s</li>\n", v)
		builder.WriteString(str)
	}
	builder.WriteString("</ul>\n")

	return builder.String()
}

// ToArray 字符串拆分为 array(默认以半角逗号,拆分)
func ToArray(rawObj string, splitKeys ...string) (objs []string) {
	objs = []string{}
	if len(rawObj) == 0 {
		return
	}

	if len(splitKeys) == 0 {
		splitKeys = []string{","} // 默认以半角逗号,拆分
	}

	rawObj = strings.TrimSpace(rawObj)
	if !strings.Contains(rawObj, splitKeys[0]) {
		objs = append(objs, rawObj)
		return
	}

	rawList := strings.Split(rawObj, splitKeys[0])
	for _, rVal := range rawList {
		oVal := strings.TrimSpace(rVal)
		if len(oVal) > 0 {
			objs = append(objs, oVal)
		}
	}

	return
}
