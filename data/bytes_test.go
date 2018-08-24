package data

import (
	"bytes"
	"testing"
)

func TestBytes(t *testing.T) {

	resultBytes := make([]byte, 50, 50)
	resultBF := bytes.NewBuffer(resultBytes)
	resultBF.Truncate(10)
	resultBF.WriteString("abcdefg")
	t.Log(resultBytes)
	t.Log(resultBF.Bytes())
}
