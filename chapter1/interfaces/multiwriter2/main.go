package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("GeeksforGeeks\nis\na\nCS-Portal!")
	fmt.Printf("%v\n", reader)

	var buf1, buf2 bytes.Buffer
	writer := io.MultiWriter(&buf1, &buf2)

	n, err := io.Copy(writer, reader)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Number of bytes in the buffer: %v\n", n)
	fmt.Printf("Buffer1: %v\n", buf1.String())
	fmt.Printf("Buffer2: %v\n", buf2.String())
}