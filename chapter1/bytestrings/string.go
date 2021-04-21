package bytestrings

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func SearchString() {
	s := "this is a test"
	fmt.Println(strings.Contains(s, "this"))
	fmt.Println(strings.HasPrefix(s, "this"))
	fmt.Println(strings.HasSuffix(s, "test"))
	fmt.Println("=============")
}

func ModifyString() {
	s := "simple string"
	fmt.Println(strings.Split(s, " "))
	fmt.Println(strings.Title(s))

	s = " simple string "
	fmt.Println(strings.TrimSpace(s))
	fmt.Println("=============")
}

func StringReader() {
	s := "simple stringn"
	r := strings.NewReader(s)
	io.Copy(os.Stdout, r)
}
