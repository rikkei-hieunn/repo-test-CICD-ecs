/*
Package converter implements logics convert data type.
*/
package converter

import (
	"bytes"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"reflect"
	"unsafe"
)

// StringToBytes high performance convert string to byte
func StringToBytes(s string) (b []byte) {
	/* #nosec G103 */
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	/* #nosec G103 */
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))

	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len

	return b
}

// BytesToString high performance convert byte to string
func BytesToString(b []byte) string {
	/* #nosec G103 */
	return *(*string)(unsafe.Pointer(&b))
}

func AppendString(strs ...string) *bytes.Buffer {
	var buff bytes.Buffer

	for i := range strs {
		buff.WriteString(strs[i])
	}

	return &buff
}

//ConvertInterfaceToListMapString convert interface to list map string interface
func ConvertInterfaceToListMapString(value interface{}) []map[string]interface{} {
	switch v := value.(type) {
	case []map[string]interface{}:
		return v
	default:
		return nil
	}
}

//ConvertInterfaceToPointer to interfacePointer
func ConvertInterfaceToPointer(value interface{}) *interface{} {
	switch v := value.(type) {
	case *interface{}:
		return v
	default:
		return nil
	}
}

// BytesToShiftJIS Convert an array of bytes (a valid UTF-8 string) to a ShiftJIS string
func BytesToShiftJIS(b []byte) ([]byte, error) {
	return transformEncoding(bytes.NewReader(b), japanese.ShiftJIS.NewEncoder())
}

// transformEncoding transform string encode
func transformEncoding(rawReader io.Reader, trans transform.Transformer) ([]byte, error) {
	ret, err := ioutil.ReadAll(transform.NewReader(rawReader, trans))
	if err == nil {
		return ret, nil
	}

	return nil, err
}
