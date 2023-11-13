package internal

import (
	"encoding/json"
	"fmt"
	"strings"
)

// OutputType 输出类型
type OutputType int

const (
	// OutputText 默认输出
	OutputText OutputType = iota
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
	outType = OutputText

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
	default: // case "text":
		outType = OutputText
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
func ToXML(obj IPPacker, rootNode string) string {
	var builder strings.Builder
	builder.WriteString("<? version=\"1.0\" encoding=\"UTF-8\" ?>\n")
	if obj == nil {
		return builder.String()
	}

	if len(rootNode) == 0 {
		rootNode = obj.TypeName()
	}

	builder.WriteString(fmt.Sprintf("<%s>\n", rootNode))
	objList := obj.GetMap()
	if len(objList) > 0 {
		for k, v := range objList {
			elem := fmt.Sprintf("\t<%[1]s>%[2]s</%[1]s>\n", k, v)
			builder.WriteString(elem)
		}
	}
	builder.WriteString(fmt.Sprintf("</%s>\n", rootNode))

	return builder.String()
}

// ToHTML 装换为HTML(ul-li) String
func ToHTML(objArray []string, objId string) (htmlText string) {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("<ul class=\"%s\"", ModuleName))
	if len(objId) > 0 {
		builder.WriteString(fmt.Sprintf(" id=\"%s\"", objId))
	}
	builder.WriteString(" >\n")

	for _, v := range objArray {
		liStr := fmt.Sprintf("\t<li>%s</li>\n", v)
		builder.WriteString(liStr)
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
		splitKeys = []string{ArraySplitKey} // 默认以半角逗号,拆分
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

// MergeArray 合并多个数组
func MergeArray(srcList ...[]string) (newArr []string) {
	if len(srcList) == 0 {
		return
	}

	total := 0
	for _, src := range srcList {
		total += len(src)
	}
	newArr = make([]string, total)
	startIndex := 0
	for _, src := range srcList {
		length := len(src)
		if length == 0 {
			return
		}
		copy(newArr[startIndex:], src)

		startIndex += length
	}

	return
}
