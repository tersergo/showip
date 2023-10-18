package internal

import "testing"

// TestToOutputType test unit
func TestToOutputType(t *testing.T) {
	outputTestMap := map[string]OutputType{
		"undefined": OutputText,
		"HtMl":      OutputHTML,
		"Array":     OutputArray,
		"Json":      OutputJSON,
		"xml":       OutputXML,
	}

	for arg, outType := range outputTestMap {
		argType := ToOutputType(arg)
		if argType != outType {
			t.Fail()
		}
	}
}
