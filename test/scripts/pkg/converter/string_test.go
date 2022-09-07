package converter

import (
	"bytes"
	"reflect"
	"repo-test-CICD-S3/pkg/converter"
	"testing"
)

// StringToBytes
func TestStringToBytes(t *testing.T) {
	inputs := []string{
		"        ",
		"ABV\tRCMFXV\tNOPV\tTKTIM\tTKZXD",
	}
	outputsExpect := [][]byte{
		[]byte("        "),
		[]byte("ABV\tRCMFXV\tNOPV\tTKTIM\tTKZXD"),
	}
	for idx := range inputs {
		if bytes.Equal(converter.StringToBytes(inputs[idx]), outputsExpect[idx]) {
			t.Logf("CASE %d: SUCCESS", idx)
		} else {
			t.Logf("CASE %d: FAILED", idx)
		}
	}
}

// BytesToString
func TestBytesToString(t *testing.T) {
	inputs := [][]byte{
		[]byte("        "),
		[]byte("ABV\tRCMFXV\tNOPV\tTKTIM\tTKZXD"),
	}
	outputsExpect := []string{
		"        ",
		"ABV\tRCMFXV\tNOPV\tTKTIM\tTKZXD",
	}
	for idx := range inputs {
		if converter.BytesToString(inputs[idx]) == outputsExpect[idx] {
			t.Logf("CASE %d: SUCCESS", idx)
		} else {
			t.Logf("CASE %d: FAILED", idx)
		}
	}
}

// ConvertInterfaceToListMapString
func TestConvertInterfaceToListMapString(t *testing.T) {
	inputs := []interface{}{
		[]map[string]interface{}{{"money_flow_data": "20220607\t18:00\t1\t2700\r\n"}},
		[]map[string]interface{}{{"best_quote_data": "20220607\t18:00\t1\t2700\r\n"}},
		[]map[int]interface{}{{1: "20220607\t18:00\t1\t2700\n"}},
	}
	outputsExpect := [][]map[string]interface{}{
		{{"money_flow_data": "20220607\t18:00\t1\t2700\r\n"}},
		{{"best_quote_data": "20220607\t18:00\t1\t2700\r\n"}},
		nil,
	}
	for idx := range inputs {
		if reflect.DeepEqual(converter.ConvertInterfaceToListMapString(inputs[idx]), outputsExpect[idx]) {
			t.Logf("CASE %d: SUCCESS", idx)
		} else {
			t.Errorf("CASE %d: FAILED", idx)
		}
	}
}

// ConvertInterfaceToPointer
func TestConvertInterfaceToPointer(t *testing.T) {
	var nV interface{} = 2503
	nP := &nV
	inputs := []interface{}{
		nP, nV,
	}
	var nilInterfacePointer *interface{} = nil
	outputsExpect := []interface{}{nP, nilInterfacePointer}
	for idx := range inputs {
		if reflect.DeepEqual(converter.ConvertInterfaceToPointer(inputs[idx]), outputsExpect[idx]) {
			t.Logf("CASE %d: SUCCESS", idx)
		} else {
			t.Errorf("CASE %d: FAILED", idx)
		}
	}
}
