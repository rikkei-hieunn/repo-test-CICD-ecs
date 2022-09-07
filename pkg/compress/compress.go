/*
Package compress implements logics compress data.
*/
package compress

import (
	"bytes"
	"compress/zlib"
)

// Compress compress data
func Compress(input []byte) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	if len(input) == 0 {
		return &buf, nil
	}
	writer := zlib.NewWriter(&buf)
	defer func() {
		_ = writer.Close()
	}()
	_, err := writer.Write(input)
	if err != nil {
		return nil, err
	}

	return &buf, nil
}
