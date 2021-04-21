package bytestrings

import (
	"bytes"
	"io"
	"io/ioutil"
)

func Buffer(rawString string) *bytes.Buffer {
	rawBytes := []byte(rawString)

	var b = new(bytes.Buffer)
	b.Write(rawBytes)

	b = bytes.NewBuffer(rawBytes)

	b = bytes.NewBufferString(rawString)

	return b
}

func toString(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	// ReadAll reads from r until an error or EOF and returns the data it read.
	// ReadAll 은 에러나 EOF까지 데이터 r을 읽는것
	if err != nil {
		return "", err
	} 
	return string(b), nil
}