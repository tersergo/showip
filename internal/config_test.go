package internal

import (
	"fmt"
	"strings"
	"testing"
)

// TestToConfig test unit
func TestToConfig(t *testing.T) {
	testHeader := []string{"a", "b", "c"}
	testCase := &Config{
		Port:      0, // port is invalid
		Path:      ModuleName,
		Header:    strings.Join(testHeader, ArraySplitKey),
		ViaArg:    "0", // via is invalid
		FormatArg: "0", // format is invalid
		ModeArg:   "0", // mode is invalid
		ObjArg:    "0",
	}

	answerPort := testCase.GetPort()
	if answerPort == testCase.Port {
		t.Errorf("GetPort err: %d => %d", testCase.Port, answerPort)
	}

	answerPath, wantPath := testCase.GetPath(), fmt.Sprintf("/%s", ModuleName)
	if answerPath != wantPath {
		t.Errorf("GetPath err: %s => %s (want: %s)", ModuleName, answerPath, wantPath)
	}

	if testCase.GetHeaders()[0] != testHeader[0] {
		t.Errorf("GetHeaders err: %s => %+v", testCase.Header, testCase.GetHeaders())
	}

	if testCase.FormatIsValid() {
		t.Errorf("FormatIsValid err: %s=%s ", FormatVarName, testCase.FormatArg)
	}

	if testCase.ViaIsValid() {
		t.Errorf("ViaIsValid err: %s=%s ", ViaVarName, testCase.ViaArg)
	}

	if testCase.ModeIsValid("") {
		t.Errorf("ModeIsValid err: %[1]s=%[2]s (ModeArg=%[2]s)", ModeVarName, testCase.ModeArg)
	}

}

// TestToConfigModeIsValid test unit
func TestToConfigModeIsValid(t *testing.T) {
	testCase := GetConfig() // mode default valid
	testMode := map[string]bool{
		"unknown":     false,
		ModeVarIsHost: true,
	}

	for input, want := range testMode {
		if testCase.ModeIsValid(input) != want {
			t.Errorf("ModeIsValid err: %s=%s (ModeArg=%s)", ModeVarName, input, testCase.ModeArg)
		}
	}
}

// TestToConfigGetVersion test unit
func TestToConfigGetVersion(t *testing.T) {
	vers := strings.Split(GetVersion(), " ")

	if len(vers) < 2 || len(vers[1]) < 4 {
		t.Errorf("GetVersion internal/Version file err: %s (e.g v1.0.0))", Version)
	}

}
