package internal

import (
	"testing"
)

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

// TestToMergeArray test unit
func TestToMergeArray(t *testing.T) {
	arr0 := []string{"0"}
	arr1 := []string{"1", "2", "3", "4", "5", "6"}
	arrN := []string{"N"}

	newArr := MergeArray(arr0, arr1, arrN)

	total := len(arr0) + len(arr1) + len(arrN)
	if len(newArr) != total {
		t.Fail()
	}

	arrMap := map[int]string{
		0:         arr0[0],
		1:         arr1[0],
		total - 1: arrN[0],
	}
	for index, val := range arrMap {
		if newArr[index] != val {
			t.Fail()
		}
	}

}
