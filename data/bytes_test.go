package data

import (
	"bytes"
	"testing"
)

func TestBytes(t *testing.T) {

	resultBytes := make([]byte, 10, 10)
	resultBF := bytes.NewBuffer(resultBytes)
	resultBF.Truncate(0)
	resultBF.WriteString("abcde")
	t.Log(resultBytes)
	t.Log(resultBF.Bytes())
}
