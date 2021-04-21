package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("Geeks")

	var buffer1, buffer2 bytes.Buffer
	writer := io.MultiWriter(&buffer1, &buffer2)

	n, err := io.Copy(writer, reader)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Number of bytes in the buffer: %v\n", n)
	fmt.Printf("Buffer1: %v\n", buffer1.String())
	fmt.Printf("Buffer2: %v\n", buffer2.String())
}